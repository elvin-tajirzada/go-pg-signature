package main

import (
	"fmt"
	signature "github.com/elvin-tacirzade/go-pg-signature"
	"github.com/jmoiron/sqlx"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
	"log"
	"os"
)

/*

Here are table, procedure and function

--Table
CREATE TABLE IF NOT EXISTS public.users
(
    user_id serial PRIMARY KEY,
    name VARCHAR NOT NULL,
    email VARCHAR NOT NULL,
    created_at TIMESTAMP WITHOUT TIME ZONE DEFAULT NOW()
);

--Procedure
CREATE OR REPLACE PROCEDURE public.create_user(
    _name INOUT VARCHAR,
    _email INOUT VARCHAR,
    _user_id INOUT INTEGER DEFAULT NULL
)
    LANGUAGE plpgsql
AS
$$

BEGIN

    INSERT INTO public.users (name, email)
    VALUES (_name, _email)
    RETURNING user_id, name, email INTO _user_id, _name, _email;
    COMMIT;

END;
$$;

--Function
CREATE OR REPLACE FUNCTION public.get_user(id INTEGER)
    RETURNS TABLE
            (
                _user_id    INTEGER,
                _name       VARCHAR,
                _email      VARCHAR
            )
AS
$$

SELECT user_id, name, email
FROM public.users
WHERE user_id = id

$$ LANGUAGE sql;
*/

type User struct {
	ID    int    `db:"_user_id"`
	Name  string `db:"_name"`
	Email string `db:"_email"`
}

func main() {
	var user User

	loadEnvErr := godotenv.Load()
	if loadEnvErr != nil {
		log.Fatalf("failed to load .env file: %v", loadEnvErr)
	}

	dsn := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=%s",
		os.Getenv("POSTGRES_HOST"), os.Getenv("POSTGRES_PORT"), os.Getenv("POSTGRES_USER"),
		os.Getenv("POSTGRES_PASSWORD"), os.Getenv("POSTGRES_DB_NAME"), os.Getenv("POSTGRES_SSL_MODE"))

	db, dbErr := sqlx.Connect("postgres", dsn)
	if dbErr != nil {
		log.Fatalf("failed to connect postgres: %v", dbErr)
	}

	sign := signature.New(db)

	/*
		rows, rowsErr := sign.RunFunction("public", "get_user", map[string]interface{}{ "id": 1 })
	*/

	rows, rowsErr := sign.RunProcedure("public", "create_user", map[string]interface{}{
		"_name":  "John Doe",
		"_email": "john.doe@gmail.com",
	})
	if rowsErr != nil {
		log.Fatalf("failed to run procedure: %v", rowsErr)
	}

	for rows.Next() {
		rowsScanErr := rows.StructScan(&user)
		if rowsScanErr != nil {
			log.Fatalf("failed to scan row: %v", rowsScanErr)
		}
	}

	log.Println(user)
}

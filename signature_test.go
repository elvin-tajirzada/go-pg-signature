package signature

import (
	"github.com/jmoiron/sqlx"
	"github.com/stretchr/testify/assert"
	"gopkg.in/DATA-DOG/go-sqlmock.v1"
	"reflect"
	"regexp"
	"testing"
)

type User struct {
	ID    int    `db:"_user_id"`
	Name  string `db:"_name"`
	Email string `db:"_email"`
}

const (
	queryProcedure    = "CALL public.create_user(_name => ?, _email => ?);"
	queryFunction     = "SELECT * FROM public.get_user(id => ?);"
	expectedQueryMake = "public.create_user(_name => :_name, _email => :_email)"
)

var expectedUser = User{ID: 1, Name: "John Doe", Email: "john.doe@gmil.com"}

func TestSignature_RunProcedure(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("failed to create a new sqlmock: %v", err)
	}

	defer db.Close()

	dbx := sqlx.NewDb(db, "sqlmock")

	rows := sqlmock.NewRows([]string{"_user_id", "_name", "_email"}).AddRow(expectedUser.ID, expectedUser.Name, expectedUser.Email)

	mock.ExpectQuery(regexp.QuoteMeta(queryProcedure)).WillReturnRows(rows)

	sign := NewSignature(dbx)

	procedureRows, procedureRowsErr := sign.RunProcedure("public", "create_user", map[string]interface{}{
		"_name":  "John Doe",
		"_email": "john.doe@gmail.com",
	})
	if procedureRowsErr != nil {
		t.Fatalf("failed to run procedure: %v", procedureRowsErr)
	}

	var actualUser User
	for procedureRows.Next() {
		scanErr := procedureRows.StructScan(&actualUser)
		if scanErr != nil {
			t.Fatalf("failed to scan rows: %v", scanErr)
		}
	}

	assert.True(t, reflect.DeepEqual(expectedUser, actualUser))
}

func TestSignature_RunFunction(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("failed to create a new sqlmock: %v", err)
	}

	defer db.Close()

	dbx := sqlx.NewDb(db, "sqlmock")

	rows := sqlmock.NewRows([]string{"_user_id", "_name", "_email"}).AddRow(expectedUser.ID, expectedUser.Name, expectedUser.Email)

	mock.ExpectQuery(regexp.QuoteMeta(queryFunction)).WillReturnRows(rows)

	sign := NewSignature(dbx)

	functionRows, functionRowsErr := sign.RunFunction("public", "get_user", map[string]interface{}{
		"id": 1,
	})
	if functionRowsErr != nil {
		t.Fatalf("failed to run function: %v", functionRowsErr)
	}

	var actualUser User
	for functionRows.Next() {
		scanErr := functionRows.StructScan(&actualUser)
		if scanErr != nil {
			t.Fatalf("failed to scan rows: %v", scanErr)
		}
	}

	assert.True(t, reflect.DeepEqual(expectedUser, actualUser))
}

func TestMakeQuery(t *testing.T) {
	schemaName := "public"
	signatureName := "create_user"
	parameters := map[string]interface{}{
		"_name":  "John Doe",
		"_email": "john.doe@gmail.com",
	}

	actualQuery := makeQuery(schemaName, signatureName, parameters)

	assert.Equal(t, expectedQueryMake, actualQuery)
}

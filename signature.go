// Package signature provides to run procedure and function for postgresql
package signature

import (
	"github.com/jmoiron/sqlx"
	"strings"
)

// Signature contains sqlx database and has RunProcedure and RunFunction functions
type Signature struct {
	DB *sqlx.DB
}

// New creates a new Signature
func New(db *sqlx.DB) *Signature {
	return &Signature{DB: db}
}

// RunProcedure runs procedure that return *sqlx.Row and error
func (s *Signature) RunProcedure(schemaName, procedureName string, params map[string]interface{}) (*sqlx.Rows, error) {
	query := "CALL " + makeQuery(schemaName, procedureName, params) + ";"
	return s.DB.NamedQuery(query, params)
}

// RunFunction runs function that return *sqlx.Row and error
func (s *Signature) RunFunction(schemaName, functionName string, params map[string]interface{}) (*sqlx.Rows, error) {
	query := "SELECT * FROM " + makeQuery(schemaName, functionName, params) + ";"
	return s.DB.NamedQuery(query, params)
}

// makeQuery makes query for a signature
func makeQuery(schemaName, signatureName string, parameters map[string]interface{}) string {
	var params []string

	if len(parameters) > 0 {
		for key := range parameters {
			params = append(params, key+" => :"+key)
		}
	}

	paramsToString := strings.Join(params, ", ")

	return schemaName + "." + signatureName + "(" + paramsToString + ")"
}

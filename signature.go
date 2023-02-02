// Package signature provides to run procedure and function for postgresql
package signature

import (
	"github.com/jmoiron/sqlx"
	"strings"
)

type (
	// ISignature interface includes RunProcedure and RunFunction functions
	ISignature interface {
		RunProcedure(schemaName, procedureName string, args map[string]interface{}) (*sqlx.Rows, error)
		RunFunction(schemaName, functionName string, params map[string]interface{}) (*sqlx.Rows, error)
	}

	// signature struct includes all dependency that uses packages
	signature struct {
		DB *sqlx.DB
	}
)

// NewSignature function creates a new signature struct
func NewSignature(db *sqlx.DB) ISignature {
	return &signature{DB: db}
}

// RunProcedure function runs procedure that return *sqlx.Row and error
func (s *signature) RunProcedure(schemaName, procedureName string, params map[string]interface{}) (*sqlx.Rows, error) {
	query := "CALL " + makeQuery(schemaName, procedureName, params) + ";"
	return s.DB.NamedQuery(query, params)
}

// RunFunction function runs function that return *sqlx.Row and error
func (s *signature) RunFunction(schemaName, functionName string, params map[string]interface{}) (*sqlx.Rows, error) {
	query := "SELECT * FROM " + makeQuery(schemaName, functionName, params) + ";"
	return s.DB.NamedQuery(query, params)
}

// makeQuery function makes query for a signature
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

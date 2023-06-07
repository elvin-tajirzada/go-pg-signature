# Go Postgres Signature Library
[![Go Reference](https://pkg.go.dev/badge/github.com/elvin-tacirzade/go-pg-signature.svg)](https://pkg.go.dev/github.com/elvin-tacirzade/go-pg-signature)

This library provides to run procedure and function for postgresql
## Installation
```
go get -u github.com/elvin-tacirzade/go-pg-signature
```
## Usage
We call New function. This function takes *sqlx.DB parameter and return Signature interface and error.

Signature interface includes two functions:
1. RunProcedure()
2. RunFunction()

See the [example](https://github.com/elvin-tacirzade/go-pg-signature/tree/main/example) subdirectory for more information.

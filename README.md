# Go Postgres Signature Library
[![Go Reference](https://pkg.go.dev/badge/github.com/elvin-tacirzade/go-pg-signature.svg)](https://pkg.go.dev/github.com/elvin-tacirzade/go-pg-signature)

This library provides to run procedure and function for postgresql
## Installation
```
go get -u github.com/elvin-tacirzade/go-pg-signature
```
## Usage
We call NewSignature function. This function takes a *sqlx.DB parameter and return ISignature interface and error.

ISignature interface includes two functions:
1. RunProcedure()
2. RunFunction()

See the [example](https://github.com/elvin-tacirzade/go-pg-signature/tree/main/example) subdirectory for more information.

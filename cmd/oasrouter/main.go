package main

import (
	"context"
	"fmt"
	"net/http"

	"github.com/getkin/kin-openapi/openapi3"
	"github.com/gorilla/mux"
)

func main() {
	var (
		r   = setupRouter()
		ctx = context.Background()
	)

	// Load the OpenAPI spec into memory
	doc, err := openapi3.NewLoader().LoadFromFile("openapi/petstore.yaml")
	if err != nil {
		panic(fmt.Sprintf("failed to load spec: %v", err))
	}

	// See here for complete list of validation options:
	// https://github.com/getkin/kin-openapi/blob/master/openapi3/validation_options.go
	opts := []openapi3.ValidationOption{
		openapi3.EnableExamplesValidation(),
		openapi3.EnableSchemaDefaultsValidation(),
		openapi3.EnableSchemaPatternValidation(),
		openapi3.EnableSchemaFormatValidation(),
	}

	// Ensure document is a valid, parsable OpenAPI spec
	err = doc.Validate(ctx, opts...)
	if err != nil {
		panic(fmt.Sprintf("failed to validate spec: %v", err))
	}

	http.ListenAndServe(":8080", r)
}

func setupRouter() *mux.Router {
	r := mux.NewRouter()
	r.HandleFunc("/", testHandler)
	return r
}

func testHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Test func"))
}

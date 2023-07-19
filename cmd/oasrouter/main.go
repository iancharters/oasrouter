package main

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/getkin/kin-openapi/openapi3"
	"github.com/getkin/kin-openapi/routers/gorillamux"
	"github.com/gorilla/mux"
	"github.com/iancharters/oasrouter/server"
)

func main() {
	ctx := context.Background()

	loader := &openapi3.Loader{
		Context:               ctx,
		IsExternalRefsAllowed: true,
	}

	// Load the OpenAPI spec into memory
	doc, err := loader.LoadFromFile("openapi/petstore.yaml")
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

	// Create new router for OpenAPI spec.  This router is not used with the server
	// itself, but is used to validate incoming requests against the OpenAPI spec.
	oasRouter, err := gorillamux.NewRouter(doc)
	if err != nil {
		panic(fmt.Sprintf("Failed to create OAS router: %v", err))
	}

	// Create new gorilla/mux router
	r := mux.NewRouter()

	// Register routes
	r.HandleFunc("/pet", createPetHandler).Methods(http.MethodPost)
	r.HandleFunc("/pet/{petId}", getPetHandler).Methods(http.MethodGet)

	// Register validation middleware
	r.Use(server.ValidateRequestMiddleware(oasRouter, nil))
	r.Use(server.ValidateResponseMiddleware(oasRouter, nil))

	http.ListenAndServe(":8080", r)
}

func createPetHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("firing createPetHandler")
	b, _ := json.Marshal(map[string]string{"hi": "2u"})
	w.Header().Set("Content-Type", "application/json")
	w.Write(b)
}

func getPetHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("firing getPetHandler")
	w.Header().Set("Content-Type", "application/json")
	w.Write([]byte("pet: " + mux.Vars(r)["petId"]))
}

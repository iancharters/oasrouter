package main

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/getkin/kin-openapi/openapi3"
	"github.com/getkin/kin-openapi/openapi3filter"
	"github.com/getkin/kin-openapi/routers"
	"github.com/getkin/kin-openapi/routers/gorillamux"
	"github.com/gorilla/mux"
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
	r.Use(validateRequestMiddleware(oasRouter, nil))
	r.Use(validateResponseMiddleware(oasRouter, nil))

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

func validateRequestMiddleware(oasRouter routers.Router, opts *openapi3filter.Options) func(next http.Handler) http.Handler {
	if opts == nil {
		opts = &openapi3filter.Options{
			MultiError:         true,
			AuthenticationFunc: openapi3filter.NoopAuthenticationFunc,
		}
	}

	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			ctx := r.Context()

			route, params, err := oasRouter.FindRoute(r)
			if err != nil {
				// do something if this happens
				fmt.Printf("Error finding route: %v\n", err)
			}

			req := &openapi3filter.RequestValidationInput{
				Request:     r,
				PathParams:  params,
				Route:       route,
				Options:     opts,
				QueryParams: r.URL.Query(),
			}

			err = openapi3filter.ValidateRequest(ctx, req)
			if err != nil {
				// do something if this happens
				fmt.Printf("validate request error: %v\n", err)
			}

			next.ServeHTTP(w, r)
		})
	}
}

type responseWriterInterceptor struct {
	http.ResponseWriter
	buffer    bytes.Buffer
	status    int
	headerMap http.Header
}

func (rwi *responseWriterInterceptor) Write(p []byte) (int, error) {
	rwi.buffer.Write(p)
	return rwi.ResponseWriter.Write(p)
}

func (rwi *responseWriterInterceptor) WriteHeader(statusCode int) {
	rwi.status = statusCode
	rwi.ResponseWriter.WriteHeader(statusCode)
}

func (rwi *responseWriterInterceptor) Header() http.Header {
	rwi.headerMap = rwi.ResponseWriter.Header().Clone()
	return rwi.ResponseWriter.Header()
}

func validateResponseMiddleware(oasRouter routers.Router, opts *openapi3filter.Options) func(next http.Handler) http.Handler {
	if opts == nil {
		opts = &openapi3filter.Options{
			MultiError:         true,
			AuthenticationFunc: openapi3filter.NoopAuthenticationFunc,
		}
	}

	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			ctx := r.Context()
			rwi := &responseWriterInterceptor{ResponseWriter: w}

			next.ServeHTTP(rwi, r)

			route, params, err := oasRouter.FindRoute(r)
			if err != nil {
				// do something if this happens
				fmt.Printf("Error finding route: %v\n", err)
			}

			// Validate response
			responseValidationInput := &openapi3filter.ResponseValidationInput{
				RequestValidationInput: &openapi3filter.RequestValidationInput{
					Request:    r,
					PathParams: params,
					Route:      route,
					Options:    opts,
				},
				Status: rwi.status,
				Header: rwi.headerMap,
			}

			fmt.Println(rwi.buffer.String())
			responseValidationInput.SetBodyBytes(rwi.buffer.Bytes())

			err = openapi3filter.ValidateResponse(ctx, responseValidationInput)
			if err != nil {
				// do something if this happens
				fmt.Printf("validate response error: %v\n", err)
			}

		})
	}
}

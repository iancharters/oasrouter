package server

import (
	"bytes"
	"fmt"
	"net/http"

	"github.com/getkin/kin-openapi/openapi3filter"
	"github.com/getkin/kin-openapi/routers"
)

func ValidateRequestMiddleware(oasRouter routers.Router, opts *openapi3filter.Options) func(next http.Handler) http.Handler {
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
				fmt.Printf("Error finding route: %v\n", err)

				return
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

func ValidateResponseMiddleware(oasRouter routers.Router, opts *openapi3filter.Options) func(next http.Handler) http.Handler {
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
				fmt.Printf("Error finding route: %v\n", err)

				return
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

			responseValidationInput.SetBodyBytes(rwi.buffer.Bytes())

			err = openapi3filter.ValidateResponse(ctx, responseValidationInput)
			if err != nil {
				// do something if this happens
				fmt.Printf("validate response error: %v\n", err)
			}

		})
	}
}

/*
 * Swagger Petstore - OpenAPI 3.0
 *
 * This is a sample Pet Store Server based on the OpenAPI 3.0 specification.  You can find out more about Swagger at [http://swagger.io](http://swagger.io). In the third iteration of the pet store, we've switched to the design first approach! You can now help us improve the API whether it's by making changes to the definition itself or to the code. That way, with time, we can improve the API in general, and expose some of the new features in OAS3.  Some useful links: - [The Pet Store repository](https://github.com/swagger-api/swagger-petstore) - [The source API definition for the Pet Store](https://github.com/swagger-api/swagger-petstore/blob/master/src/main/resources/openapi.yaml)
 *
 * API version: 1.0.17
 * Contact: apiteam@swagger.io
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi

import (
	"encoding/json"
	"net/http"
	"strings"

	"github.com/gorilla/mux"
)

// StoreApiController binds http requests to an api service and writes the service results to the http response
type StoreApiController struct {
	service StoreApiServicer
	errorHandler ErrorHandler
}

// StoreApiOption for how the controller is set up.
type StoreApiOption func(*StoreApiController)

// WithStoreApiErrorHandler inject ErrorHandler into controller
func WithStoreApiErrorHandler(h ErrorHandler) StoreApiOption {
	return func(c *StoreApiController) {
		c.errorHandler = h
	}
}

// NewStoreApiController creates a default api controller
func NewStoreApiController(s StoreApiServicer, opts ...StoreApiOption) Router {
	controller := &StoreApiController{
		service:      s,
		errorHandler: DefaultErrorHandler,
	}

	for _, opt := range opts {
		opt(controller)
	}

	return controller
}

// Routes returns all the api routes for the StoreApiController
func (c *StoreApiController) Routes() Routes {
	return Routes{ 
		{
			"DeleteOrder",
			strings.ToUpper("Delete"),
			"/store/order/{orderId}",
			c.DeleteOrder,
		},
		{
			"GetInventory",
			strings.ToUpper("Get"),
			"/store/inventory",
			c.GetInventory,
		},
		{
			"GetOrderById",
			strings.ToUpper("Get"),
			"/store/order/{orderId}",
			c.GetOrderById,
		},
		{
			"PlaceOrder",
			strings.ToUpper("Post"),
			"/store/order",
			c.PlaceOrder,
		},
	}
}

// DeleteOrder - Delete purchase order by ID
func (c *StoreApiController) DeleteOrder(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	orderIdParam, err := parseInt64Parameter(params["orderId"], true)
	if err != nil {
		c.errorHandler(w, r, &ParsingError{Err: err}, nil)
		return
	}
	result, err := c.service.DeleteOrder(r.Context(), orderIdParam)
	// If an error occurred, encode the error with the status code
	if err != nil {
		c.errorHandler(w, r, err, &result)
		return
	}
	// If no error, encode the body and the result code
	EncodeJSONResponse(result.Body, &result.Code, w)

}

// GetInventory - Returns pet inventories by status
func (c *StoreApiController) GetInventory(w http.ResponseWriter, r *http.Request) {
	result, err := c.service.GetInventory(r.Context())
	// If an error occurred, encode the error with the status code
	if err != nil {
		c.errorHandler(w, r, err, &result)
		return
	}
	// If no error, encode the body and the result code
	EncodeJSONResponse(result.Body, &result.Code, w)

}

// GetOrderById - Find purchase order by ID
func (c *StoreApiController) GetOrderById(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	orderIdParam, err := parseInt64Parameter(params["orderId"], true)
	if err != nil {
		c.errorHandler(w, r, &ParsingError{Err: err}, nil)
		return
	}
	result, err := c.service.GetOrderById(r.Context(), orderIdParam)
	// If an error occurred, encode the error with the status code
	if err != nil {
		c.errorHandler(w, r, err, &result)
		return
	}
	// If no error, encode the body and the result code
	EncodeJSONResponse(result.Body, &result.Code, w)

}

// PlaceOrder - Place an order for a pet
func (c *StoreApiController) PlaceOrder(w http.ResponseWriter, r *http.Request) {
	orderParam := Order{}
	d := json.NewDecoder(r.Body)
	d.DisallowUnknownFields()
	if err := d.Decode(&orderParam); err != nil {
		c.errorHandler(w, r, &ParsingError{Err: err}, nil)
		return
	}
	if err := AssertOrderRequired(orderParam); err != nil {
		c.errorHandler(w, r, err, nil)
		return
	}
	result, err := c.service.PlaceOrder(r.Context(), orderParam)
	// If an error occurred, encode the error with the status code
	if err != nil {
		c.errorHandler(w, r, err, &result)
		return
	}
	// If no error, encode the body and the result code
	EncodeJSONResponse(result.Body, &result.Code, w)

}
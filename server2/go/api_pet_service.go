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
	"context"
	"net/http"
	"errors"
	"os"
)

// PetApiService is a service that implements the logic for the PetApiServicer
// This service should implement the business logic for every endpoint for the PetApi API.
// Include any external packages or services that will be required by this service.
type PetApiService struct {
}

// NewPetApiService creates a default api service
func NewPetApiService() PetApiServicer {
	return &PetApiService{}
}

// AddPet - Add a new pet to the store
func (s *PetApiService) AddPet(ctx context.Context, pet Pet) (ImplResponse, error) {
	// TODO - update AddPet with the required logic for this service method.
	// Add api_pet_service.go to the .openapi-generator-ignore to avoid overwriting this service implementation when updating open api generation.

	//TODO: Uncomment the next line to return response Response(200, Pet{}) or use other options such as http.Ok ...
	//return Response(200, Pet{}), nil

	//TODO: Uncomment the next line to return response Response(405, {}) or use other options such as http.Ok ...
	//return Response(405, nil),nil

	return Response(http.StatusNotImplemented, nil), errors.New("AddPet method not implemented")
}

// DeletePet - Deletes a pet
func (s *PetApiService) DeletePet(ctx context.Context, petId int64, apiKey string) (ImplResponse, error) {
	// TODO - update DeletePet with the required logic for this service method.
	// Add api_pet_service.go to the .openapi-generator-ignore to avoid overwriting this service implementation when updating open api generation.

	//TODO: Uncomment the next line to return response Response(400, {}) or use other options such as http.Ok ...
	//return Response(400, nil),nil

	return Response(http.StatusNotImplemented, nil), errors.New("DeletePet method not implemented")
}

// FindPetsByStatus - Finds Pets by status
func (s *PetApiService) FindPetsByStatus(ctx context.Context, status string) (ImplResponse, error) {
	// TODO - update FindPetsByStatus with the required logic for this service method.
	// Add api_pet_service.go to the .openapi-generator-ignore to avoid overwriting this service implementation when updating open api generation.

	//TODO: Uncomment the next line to return response Response(200, []Pet{}) or use other options such as http.Ok ...
	//return Response(200, []Pet{}), nil

	//TODO: Uncomment the next line to return response Response(400, {}) or use other options such as http.Ok ...
	//return Response(400, nil),nil

	return Response(http.StatusNotImplemented, nil), errors.New("FindPetsByStatus method not implemented")
}

// FindPetsByTags - Finds Pets by tags
func (s *PetApiService) FindPetsByTags(ctx context.Context, tags []string) (ImplResponse, error) {
	// TODO - update FindPetsByTags with the required logic for this service method.
	// Add api_pet_service.go to the .openapi-generator-ignore to avoid overwriting this service implementation when updating open api generation.

	//TODO: Uncomment the next line to return response Response(200, []Pet{}) or use other options such as http.Ok ...
	//return Response(200, []Pet{}), nil

	//TODO: Uncomment the next line to return response Response(400, {}) or use other options such as http.Ok ...
	//return Response(400, nil),nil

	return Response(http.StatusNotImplemented, nil), errors.New("FindPetsByTags method not implemented")
}

// GetPetById - Find pet by ID
func (s *PetApiService) GetPetById(ctx context.Context, petId int64) (ImplResponse, error) {
	// TODO - update GetPetById with the required logic for this service method.
	// Add api_pet_service.go to the .openapi-generator-ignore to avoid overwriting this service implementation when updating open api generation.

	//TODO: Uncomment the next line to return response Response(200, Pet{}) or use other options such as http.Ok ...
	//return Response(200, Pet{}), nil

	//TODO: Uncomment the next line to return response Response(400, {}) or use other options such as http.Ok ...
	//return Response(400, nil),nil

	//TODO: Uncomment the next line to return response Response(404, {}) or use other options such as http.Ok ...
	//return Response(404, nil),nil

	return Response(http.StatusNotImplemented, nil), errors.New("GetPetById method not implemented")
}

// UpdatePet - Update an existing pet
func (s *PetApiService) UpdatePet(ctx context.Context, pet Pet) (ImplResponse, error) {
	// TODO - update UpdatePet with the required logic for this service method.
	// Add api_pet_service.go to the .openapi-generator-ignore to avoid overwriting this service implementation when updating open api generation.

	//TODO: Uncomment the next line to return response Response(200, Pet{}) or use other options such as http.Ok ...
	//return Response(200, Pet{}), nil

	//TODO: Uncomment the next line to return response Response(400, {}) or use other options such as http.Ok ...
	//return Response(400, nil),nil

	//TODO: Uncomment the next line to return response Response(404, {}) or use other options such as http.Ok ...
	//return Response(404, nil),nil

	//TODO: Uncomment the next line to return response Response(405, {}) or use other options such as http.Ok ...
	//return Response(405, nil),nil

	return Response(http.StatusNotImplemented, nil), errors.New("UpdatePet method not implemented")
}

// UpdatePetWithForm - Updates a pet in the store with form data
func (s *PetApiService) UpdatePetWithForm(ctx context.Context, petId int64, name string, status string) (ImplResponse, error) {
	// TODO - update UpdatePetWithForm with the required logic for this service method.
	// Add api_pet_service.go to the .openapi-generator-ignore to avoid overwriting this service implementation when updating open api generation.

	//TODO: Uncomment the next line to return response Response(405, {}) or use other options such as http.Ok ...
	//return Response(405, nil),nil

	return Response(http.StatusNotImplemented, nil), errors.New("UpdatePetWithForm method not implemented")
}

// UploadFile - uploads an image
func (s *PetApiService) UploadFile(ctx context.Context, petId int64, additionalMetadata string, body *os.File) (ImplResponse, error) {
	// TODO - update UploadFile with the required logic for this service method.
	// Add api_pet_service.go to the .openapi-generator-ignore to avoid overwriting this service implementation when updating open api generation.

	//TODO: Uncomment the next line to return response Response(200, ApiResponse{}) or use other options such as http.Ok ...
	//return Response(200, ApiResponse{}), nil

	return Response(http.StatusNotImplemented, nil), errors.New("UploadFile method not implemented")
}

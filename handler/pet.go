package handlers

import (
	"net/http"
)

func CreatePetHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// oasutil.Body(r)

		w.Write([]byte("CreatePetHandler"))
	}
}

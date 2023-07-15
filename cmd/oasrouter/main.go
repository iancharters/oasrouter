package main

import (
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/", TestHandler)
	http.ListenAndServe(":8080", r)
}

func TestHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Test func"))
}

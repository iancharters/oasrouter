package main

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/iancharters/oasrouter/oasutil"
)

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/", TestHandler)
	http.ListenAndServe(":8080", r)
}

func TestHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte(oasutil.TestFunc()))
}

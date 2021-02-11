package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	route := mux.NewRouter()

	route.HandleFunc("/api/books", service.getBooks).Methods("GET")

	log.Fatal(http.ListenAndServe(":8080", r))
}

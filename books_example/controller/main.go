package main

import (
	"log"
	"net/http"

	service "books_example/service"

	"github.com/gorilla/mux"
)

func main() {
	route := mux.NewRouter()

	route.HandleFunc("/api/books", service.GetBooks).Methods("GET")

	log.Fatal(http.ListenAndServe(":8080", route))
}

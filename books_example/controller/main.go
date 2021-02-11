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
	route.HandleFunc("/api/books/{id}", service.GetBook).Methods("GET")
	route.HandleFunc("/api/books", service.CreateBook).Methods("POST")
	route.HandleFunc("/api/books/{id}", service.UpdateBook).Methods("PUT")
	route.HandleFunc("/api/books/{id}", service.DeleteBook).Methods("DELETE")

	log.Fatal(http.ListenAndServe(":8080", route))
}

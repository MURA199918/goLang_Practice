package service

import (
	"context"
	"encoding/json"
	"log"
	"net/http"

	"books_example/model"
	database "books_example/repository"

	"go.mongodb.org/mongo-driver/bson"
)

var collection = database.ConnectDatabase()

//GetBooks function
func GetBooks(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var books []model.Book

	cur, err := collection.Find(context.TODO(), bson.M{})
	if err != nil {
		log.Fatal(err)
	}
	defer cur.Close(context.TODO())

	for cur.Next(context.TODO()) {
		var book model.Book

		err := cur.Decode(&book)
		if err != nil {
			log.Fatal(err)
		}
		books = append(books, book)
	}

	if err := cur.Err(); err != nil {
		log.Fatal(err)
	}
	json.NewEncoder(w).Encode(books)
}

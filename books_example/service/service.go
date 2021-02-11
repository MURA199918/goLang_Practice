package service

import (
	"context"
	"encoding/json"
	"log"
	"net/http"

	"books_example/model"
	database "books_example/respository"

	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

var collection = database.ConnectDatabase()

//GetBooks function
func GetBooks(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var books []model.Book

	cur, err := collection.Find(context.TODO(), bson.M{}) //bson.M{} to get all data from database
	if err != nil {
		log.Fatal(err)
	}
	defer cur.Close(context.TODO())

	for cur.Next(context.TODO()) {
		var book model.Book

		err := cur.Decode(&book) //deserialize, & returns memory address
		if err != nil {
			log.Fatal(err)
		}
		books = append(books, book)
	}

	if err := cur.Err(); err != nil {
		log.Fatal(err)
	}
	json.NewEncoder(w).Encode(books) //encode to display in browser
}

//GetBook function
func GetBook(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json") //set header
	var book model.Book
	var params = mux.Vars(r) //get all parameters by mux

	id, _ := primitive.ObjectIDFromHex(params["id"]) //finding mongodb document by it bson id

	filter := bson.M{"_id": id}                                     //filter id with bson id
	err := collection.FindOne(context.TODO(), filter).Decode(&book) // use filter to search and findone from database

	if err != nil {
		log.Fatal(err)
	}
	json.NewEncoder(w).Encode(book)
}

//CreateBook function
func CreateBook(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var book model.Book
	_ = json.NewDecoder(r.Body).Decode(&book) //decoding our body request

	result, err := collection.InsertOne(context.TODO(), book)

	if err != nil {
		log.Fatal(err)
	}
	json.NewEncoder(w).Encode(result) //encoding our request and display in browser
}

//UpdateBook function
func UpdateBook(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var params = mux.Vars(r)

	id, _ := primitive.ObjectIDFromHex(params["id"])
	var book model.Book
	filter := bson.M{"_id": id}

	_ = json.NewDecoder(r.Body).Decode(&book)

	update := bson.D{
		{"$set", bson.D{
			{"title", book.Title},
			{"author", bson.D{
				{"firstname", book.Author.FirstName},
				{"lastname", book.Author.LastName},
			}},
		}},
	}

	err := collection.FindOneAndUpdate(context.TODO(), filter, update).Decode(&book)

	if err != nil {
		log.Fatal(err)
	}
	book.ID = id
	json.NewEncoder(w).Encode(book)
}

//DeleteBook function
func DeleteBook(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var params = mux.Vars(r)

	id, err := primitive.ObjectIDFromHex(params["id"])
	filter := bson.M{"_id": id}
	deleteResult, err := collection.DeleteOne(context.TODO(), filter)

	if err != nil {
		log.Fatal(err)
	}
	json.NewEncoder(w).Encode(deleteResult)
}

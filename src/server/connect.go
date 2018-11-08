package main

import (
	"fmt"
	"io"
	"log"
	"net/http"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"gopkg.in/mgo.v2/bson"
)

type Person struct {
	ID   bson.ObjectId `bson:"_id"`
	Name string        `bson:"name"`
	Age  int           `bson:"age"`
}

func test(w http.ResponseWriter, r *http.Request) {
}

func PostImageBlob(w http.ResponseWriter, r *http.Request) {
	fmt.Print("post detect")
	reader, err := r.MultipartReader()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	for {
		tar, err := reader.NextPart()
		if err == io.EOF {
			break
		}
		if tar.FileName() == "" {
			continue
		}
		fmt.Printf("%+v", tar)
	}
}

func main() {
	router := mux.NewRouter()
	allowedOrigins := handlers.AllowedOrigins([]string{"http://localhost:8080"})
	allowedMethods := handlers.AllowedMethods([]string{"GET", "POST", "DELETE", "PUT", "OPTIONS"})
	allowedHeaders := handlers.AllowedHeaders([]string{"X-Requested-With", "Content-Type", "Authorization"})
	router.HandleFunc("/images/post", PostImageBlob).Methods("POST")
	router.HandleFunc("/images/post", test).Methods("GET")
	log.Fatal(http.ListenAndServe(":5000", handlers.CORS(allowedOrigins, allowedMethods, allowedHeaders)(router)))
}

package main

import (
	"fmt"
	"io"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"gopkg.in/mgo.v2/bson"
)

type Person struct {
	ID   bson.ObjectId `bson:"_id"`
	Name string        `bson:"name"`
	Age  int           `bson:"age"`
}

func PostImageBlob(w http.ResponseWriter, r *http.Request) {
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
	router.HandleFunc("/images/{id}", PostImageBlob).Methods("POST")
	log.Fatal(http.ListenAndServe(":5000", router))
}

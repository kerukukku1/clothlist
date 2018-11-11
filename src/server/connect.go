package main

import (
	"fmt"
	"io"
	"log"
	"math/rand"
	"net/http"
	"os"
	"time"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	mgo "gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

type Image struct {
	ID    bson.ObjectId `bson:"_id"`
	Title string        `bson:"title"`
	Path  string        `bson:"path"`
	link  string        `bson:"link"`
}

func init() {
	rand.Seed(time.Now().UnixNano())
}

var rs1Letters = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")

func RandString1(n int) string {
	b := make([]rune, n)
	for i := range b {
		b[i] = rs1Letters[rand.Intn(len(rs1Letters))]
	}
	return string(b)
}

func test(w http.ResponseWriter, r *http.Request) {
}

func PostImageBlob(w http.ResponseWriter, r *http.Request) {
	err := r.ParseMultipartForm(32 << 20) // maxMemory
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	file, _, err := r.FormFile("imagefile")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer file.Close()
	title := r.FormValue("title")
	if len(title) == 0 {
		return
	}
	link := RandString1(20)
	p, _ := os.Getwd()
	path := p + "/images/" + link + ".jpg"
	f, err := os.Create(path)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	newImage := &Image{
		ID:    bson.NewObjectId(),
		Path:  path,
		Title: title,
		link:  link,
	}
	mongoSaveImage(*newImage)

	defer f.Close()
	io.Copy(f, file)
}

func mongoSaveImage(newImage Image) {
	session, _ := mgo.Dial("mongodb://localhost/test")
	//この関数が終わる時にsessionをcloseする
	defer session.Close()
	db := session.DB("test")
	column := db.C("images")
	if err := column.Insert(&newImage); err != nil {
		log.Fatalln(err)
		fmt.Println(err)
	}
}

func main() {
	router := mux.NewRouter()
	os.Mkdir("images", 0777)
	allowedOrigins := handlers.AllowedOrigins([]string{"http://localhost:8080", "http://192.168.11.8:8080"})
	allowedMethods := handlers.AllowedMethods([]string{"GET", "POST", "DELETE", "PUT", "OPTIONS"})
	allowedHeaders := handlers.AllowedHeaders([]string{"X-Requested-With", "Content-Type", "Authorization"})
	router.HandleFunc("/images/post", PostImageBlob).Methods("POST")
	router.HandleFunc("/images/post", test).Methods("GET")
	log.Fatal(http.ListenAndServe(":5000", handlers.CORS(allowedOrigins, allowedMethods, allowedHeaders)(router)))
}

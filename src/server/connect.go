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
)

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
	fmt.Print("post detect")
	err := r.ParseMultipartForm(32 << 20) // maxMemory
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	fmt.Print("memory ok")
	file, _, err := r.FormFile("imagefile")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	fmt.Print("file get ok")
	defer file.Close()
	title := r.FormValue("title")
	if len(title) == 0 {
		return
	}
	fmt.Print("title get ok")
	f, err := os.Create("./images/" + RandString1(20) + ".jpg")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	fmt.Print("image save ok")
	defer f.Close()
	io.Copy(f, file)
}

func main() {
	router := mux.NewRouter()
	os.Mkdir("images", 0777)
	allowedOrigins := handlers.AllowedOrigins([]string{"http://localhost:8080"})
	allowedMethods := handlers.AllowedMethods([]string{"GET", "POST", "DELETE", "PUT", "OPTIONS"})
	allowedHeaders := handlers.AllowedHeaders([]string{"X-Requested-With", "Content-Type", "Authorization"})
	router.HandleFunc("/images/post", PostImageBlob).Methods("POST")
	router.HandleFunc("/images/post", test).Methods("GET")
	log.Fatal(http.ListenAndServe(":5000", handlers.CORS(allowedOrigins, allowedMethods, allowedHeaders)(router)))
}

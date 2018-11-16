package main

import (
	"encoding/json"
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
}

type Detail struct {
	ID        bson.ObjectId   `bson:"_id"`
	ParentID  bson.ObjectId   `bson:"parentObjectId"`
	Comment   string          `bson:"comment"`
	Withcloth []bson.ObjectId `bson:"withCloth"`
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
	// p, _ := os.Getwd()
	// path := p + "/images/" + RandString1(20) + ".jpg"
	path := "images/" + RandString1(20) + ".jpg"
	f, err := os.Create(path)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	newImage := &Image{
		ID:    bson.NewObjectId(),
		Path:  path,
		Title: title,
	}
	mongoSaveImage(*newImage)

	defer f.Close()
	io.Copy(f, file)
}

func convertJSON(image Image) string {
	img, err := json.Marshal(image)
	if err != nil {
		log.Fatalln(err)
	}
	return string(img)
}
func findObjID(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	session, _ := mgo.Dial("mongodb://localhost/test")
	defer session.Close()
	db := session.DB("test")
	idStr := params["objID"]
	if !bson.IsObjectIdHex(idStr) {
		fmt.Println("not objectId")
		return
	}
	id := bson.ObjectIdHex(idStr)

	var image Image
	if err := db.C(params["column"]).FindId(id).One(&image); err != nil {
		fmt.Println(err)
		return
	}
	w.Write([]byte(convertJSON(image)))
}

func findTagImage(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	session, _ := mgo.Dial("mongodb://localhost/test")
	defer session.Close()
	var images []Image
	db := session.DB("test")
	fmt.Println(params["column"])
	if err := db.C(params["column"]).Find(bson.M{}).All(&images); err != nil {
		log.Fatal(err)
	}
	ret := ""
	ret += "["
	for _, image := range images {
		ret += convertJSON(image)
		ret += ","
	}
	sc := []rune(ret)
	ret = string(sc[:(len(sc) - 1)])
	ret += "]"
	w.Write([]byte(ret))
}

func mongoSaveImage(newImage Image) {
	session, _ := mgo.Dial("mongodb://localhost/test")
	//この関数が終わる時にsessionをcloseする
	defer session.Close()
	db := session.DB("test")
	img_column := db.C("images")
	if err := img_column.Insert(&newImage); err != nil {
		log.Fatalln(err)
		fmt.Println(err)
	}
	var withClothes []bson.ObjectId
	fmt.Println(withClothes)
	newDetail := &Detail{
		ID:        bson.NewObjectId(),
		ParentID:  newImage.ID,
		Comment:   "# Markdown document is not setting yet.",
		Withcloth: withClothes,
	}
	println("mongo save")
	detail_column := db.C("detail")
	if err := detail_column.Insert(&newDetail); err != nil {
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
	router.HandleFunc("/api/{column}", findTagImage).Methods("GET")
	router.HandleFunc("/api/{column}/{objID}", findObjID).Methods("GET")
	log.Fatal(http.ListenAndServe(":5000", handlers.CORS(allowedOrigins, allowedMethods, allowedHeaders)(router)))
}

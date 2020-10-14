package main

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"log"
	"math/rand"
	"net/http"
	"strconv"
)
type Picture struct {
	ID string `json:"id"`
	Title string `json:"title"`
	Body string `json:"body"`
}

var pictures []Picture

func main() {

	router := mux.NewRouter()

	//pictures = append(pictures, Picture{ID: "1", Title: "My first post", Body: "This is the content of my first post"})

	router.HandleFunc("/pictures", createPicture).Methods("POST")
	router.HandleFunc("/pictures", getPictures).Methods("GET")
	//router.HandleFunc("/pictures/{id}", getPicture).Methods("GET")

	log.Fatal(http.ListenAndServe(":8080", router))
}

func createPicture (w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var picture Picture
	_ = json.NewDecoder(r.Body).Decode(&picture)
	picture.ID = strconv.Itoa(rand.Intn(1000000))
	pictures = append(pictures, picture)
	json.NewEncoder(w).Encode(&picture)
}

func getPictures(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	for _, item := range pictures {
		json.NewEncoder(w).Encode(&item)
	}
}

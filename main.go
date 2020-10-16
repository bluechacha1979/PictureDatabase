package main

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"github.com/jinzhu/gorm"
	_ "github.com/mattn/go-sqlite3"
	"io/ioutil"
	"net/http"
)

//Database Model
type Picture struct {
	ID string `json:"id"`
	Title string `json:"title"`
	Body string `json:"body"`
}

var database *gorm.DB
var err error

//Initial Database Migration
func InitialMigration() {
	database, err = gorm.Open("sqlite3", "test.db")
	if err != nil {
		panic("Failed to connect to database!")
	}
	defer database.Close()
	database.AutoMigrate(&Picture{})
}

func main() {

	InitialMigration()

	database, err = gorm.Open("sqlite3", "test.db")
	if err != nil {
		panic("Failed to connect to database!")
	}

	handleRequests()

	defer database.Close()
}


//Routes
func handleRequests() {
	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/pictures", createPicture).Methods("POST")
	router.HandleFunc("/pictures", getPictures).Methods("GET")
	router.HandleFunc("/pictures/{id}", getPicture).Methods("GET")
	http.ListenAndServe(":8080", router)
}

//Handlers
func createPicture (w http.ResponseWriter, r *http.Request) {
	reqBody, _ := ioutil.ReadAll(r.Body)
	var picture Picture
	json.Unmarshal(reqBody, &picture)
	database.Create(&picture)
	fmt.Println("Creating New Picture")
	json.NewEncoder(w).Encode(picture)
}

func getPictures(w http.ResponseWriter, r *http.Request) {
	pictures := []Picture{}
	database.Find(&pictures)
	fmt.Println("Getting All the Pictures")
	json.NewEncoder(w).Encode(pictures)
}

func getPicture(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	pictures := []Picture{}
	database.Find(&pictures)

	for _, item := range pictures {
		if item.ID == vars["id"] {
			json.NewEncoder(w).Encode(item)
			fmt.Println("Get Single Picture" )
			return
		}
	}
	fmt.Println("Did not find the Picture")
}

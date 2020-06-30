package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/atra2396/hyperlink/auth"
	model "github.com/atra2396/hyperlink/models"
	"github.com/atra2396/hyperlink/routing"
	"github.com/gorilla/mux"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

func main() {
	db := initDb()
	defer db.Close()

	routing.InitDbConnection(db)

	r := mux.NewRouter()

	r.HandleFunc("/", greet).Methods("GET")
	r.HandleFunc("/login/{username}/{password}", auth.Login).Methods("GET")

	secure := r.PathPrefix("/auth").Subrouter()
	secure.Use(auth.AuthenticationMiddleware)

	secure.HandleFunc("/api/v1/text/{id}", routing.GetText).Methods("GET")
	secure.HandleFunc("/api/v1/text", routing.CreateText).Methods("POST")
	secure.HandleFunc("/", greet).Methods("GET")

	fmt.Println("Starting server...")
	log.Fatal(http.ListenAndServe(":8080", r))
}

func greet(w http.ResponseWriter, r *http.Request) {
	welcome := "Welcome to the GO server!"
	json.NewEncoder(w).Encode(welcome)
}

func initDb() gorm.DB {
	db, err := gorm.Open("sqlite3", "file::memory:?cache=shared")
	if err != nil {
		log.Panicln("Error: Could not open database connection")
	}

	db.AutoMigrate(&model.User{}, &model.Text{}, &model.Fragment{}, &model.Link{})
	return *db
}

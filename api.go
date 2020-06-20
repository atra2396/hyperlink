package main

import (
	"log"
	"net/http"

	"github.com/atra2396/hyperlink/routing"
	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/api/v1/text/{id}", routing.GetText).Methods("GET")
	r.HandleFunc("/api/v1/text", routing.CreateText).Methods("POST")

	log.Fatal(http.ListenAndServe(":8080", r))
}

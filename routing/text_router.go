package routing

import (
	"encoding/json"
	"net/http"
	"strconv"

	model "github.com/atra2396/hyperlink/models"
	"github.com/gorilla/mux"
	"github.com/jinzhu/gorm"
)

var db gorm.DB

func InitDbConnection(database gorm.DB) {
	db = database
}

func CreateText(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var newText model.Text
	_ = json.NewDecoder(r.Body).Decode(&newText)

	db.Create(&newText)

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(newText)
}

func GetText(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	givenId, err := strconv.Atoi(params["id"])
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("invalid id given"))
		return
	}

	castedId := uint(givenId)

	var text model.Text
	db.Where("id = ?", castedId).First(&text)

	json.NewEncoder(w).Encode(text)
}

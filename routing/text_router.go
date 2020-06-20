package routing

import (
	"encoding/json"
	"net/http"
	"strconv"

	model "github.com/atra2396/hyperlink/models"
	"github.com/gorilla/mux"
)

var texts []model.Text

func CreateText(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var newText model.Text
	_ = json.NewDecoder(r.Body).Decode(&newText)
	newText.ID = uint(len(texts))
	texts = append(texts, newText)
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

	for _, val := range texts {
		if val.ID == castedId {
			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode(val)
			return
		}
	}

	w.WriteHeader(http.StatusNotFound)
	w.Write([]byte("id could not be found"))
	return
}

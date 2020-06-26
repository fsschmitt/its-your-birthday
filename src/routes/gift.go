package routes

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/fsschmitt/its-your-birthday/src/models"
)

func CreateGift(w http.ResponseWriter, r *http.Request) {
	var gift models.Gift

	err := json.NewDecoder(r.Body).Decode(&gift)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	fmt.Fprintf(w, "Gift: %+v", gift)

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
}

func GetGift(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
}

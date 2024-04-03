package handlers

import (
	"encoding/json"
	"go-restapi-app/src/models"
	"net/http"
)

var items = []models.Item {
	{ID: "0001", Name: "Chocolate"},
	{ID: "0002", Name: "Biscuits"},
}

func GetItems(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(items)
}
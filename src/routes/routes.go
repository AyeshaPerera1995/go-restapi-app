package routes

import (
	"go-restapi-app/src/handlers"
	"github.com/gorilla/mux"
)

func NewRouter() *mux.Router {

	r := mux.NewRouter()
	r.HandleFunc("/items", handlers.GetItems).Methods("GET")

	return r

}
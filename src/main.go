package main

import (
	"go-restapi-app/src/routes"
	"log"
	"net/http"
)

func main() {
	r := routes.NewRouter()

	log.Fatal(http.ListenAndServe(":5000", r))

}
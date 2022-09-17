package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/SpasZahariev/go-play-around-with-it/db-book-management-system/pkg/routes"
	"github.com/gorilla/mux"
)

func main() {
	fmt.Println("Starting Book Management app!")
	var router *mux.Router = mux.NewRouter()
	routes.RegisterBookStoreRoutes(router)

	http.Handle("/", router)
	log.Fatal(http.ListenAndServe("localhost:9010", router))
}
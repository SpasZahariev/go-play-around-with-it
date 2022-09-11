package routes

import (
	"github.com/gorilla/mux"
	"github.com/SpasZahariev/go-play-around-with-it/db-book-management-system/pkg/controllers"
)

var RegisterBookStoreRoutes = func(router *mux.Router) {
	// todo maybe I dont need this to be /book/
	router.HandleFunc("/book/", controllers.CreateBook)
}
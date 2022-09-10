package main

import (
	"fmt"
	"net/http"
)

func catHandler(writer http.ResponseWriter, request *http.Request) {

	if request.Method == "GET" {
		fmt.Fprint(writer, "Oh yes we have all the Catsies")
	}
}

func favouriteHandler(writer http.ResponseWriter, request *http.Request) {

	fmt.Println("We just Received a POST request? maybe: " + request.Method)

	theName := request.FormValue("myName")
	theAge := request.FormValue("myAge")
	theCat := request.FormValue("myFavouriteCat")

	fmt.Fprint(writer, "Heyo! You just sent us a form with your details. Name: " + theName + " Age: " + theAge + " YOUR MOFO CAT: " + theCat)
}

func main() {
	fmt.Println("Hello from Dusi's app")
	fileServer := http.FileServer(http.Dir("./static"))

	http.Handle("/", fileServer)
	http.HandleFunc("/cats", catHandler)
	http.HandleFunc("/favourite-cats", favouriteHandler)


	fmt.Println("Server has started on port 6969")

	http.ListenAndServe(":6969", nil)

}

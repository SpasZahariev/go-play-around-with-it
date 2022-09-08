package main

import (
	"fmt"
	"log"
	"net/http"
)

func formHandler(writer http.ResponseWriter, request *http.Request) {
	if err := request.ParseForm(); err != nil {
		fmt.Fprint(writer, "Parse Error", err)
		return
	}

	fmt.Fprint(writer, "POST request was successful \n")

	name := request.FormValue("name")
	cuteness := request.FormValue("cuteness")

	fmt.Fprintf(writer, "Name = %s  \n", name)
	fmt.Fprintf(writer, "Your Cuteness = %s  \n", cuteness)
}

func helloHandler(writer http.ResponseWriter, request *http.Request) {
	if request.URL.Path != "/hello" {
		http.Error(writer, "404 i didn't find it", http.StatusNotFound)
		return
	}

	if request.Method != "GET" {
		http.Error(writer, "method is not supported", http.StatusNotFound)

	}
	fmt.Fprint(writer, "hello friend!")

}


func main() {
	fileServer := http.FileServer(http.Dir("./static"))
	http.Handle("/", fileServer)
	http.HandleFunc("/form", formHandler)
	http.HandleFunc("/hello", helloHandler)

	fmt.Println("Starting server at port 8080")

	error := http.ListenAndServe(":8080", nil)
	if error != nil {
		log.Fatal(error)
	}

}

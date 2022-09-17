package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/SpasZahariev/go-play-around-with-it/db-book-management-system/pkg/models"
	"github.com/SpasZahariev/go-play-around-with-it/db-book-management-system/pkg/utils"
	"github.com/gorilla/mux"
)

var Newbook models.Book

func GetBook(writer http.ResponseWriter, request *http.Request) {
	newBooks := models.GetAllBooks()

	res, _ := json.Marshal(newBooks)
	writer.Header().Set("Content-Type", "application/json")
	writer.WriteHeader(http.StatusOK)
	writer.Write(res)
}

func GetBookById(writer http.ResponseWriter, request *http.Request) {
	params := mux.Vars(request)
	bookId := params["bookId"]

	id, err := strconv.ParseInt(bookId, 0, 0)

	if err != nil {
		fmt.Println("Error while parsing book id")
	}

	book, _ := models.GetBookById(id)

	res, _ := json.Marshal(book)
	writer.Header().Set("Content-Type", "application/json")
	writer.WriteHeader(http.StatusOK)
	writer.Write(res)

}

func CreateBook(writer http.ResponseWriter, request *http.Request) {
	newBook := models.Book{} //pointer to new empty Book

	utils.ParseBody(request, &newBook) // write into the Book

	b := newBook.CreateBook() // persisting in DB

	res, _ := json.Marshal(b)


	writer.Header().Set("Content-Type", "application/json")
	writer.WriteHeader(http.StatusOK)
	writer.Write(res)
}

func DeleteBook(writer http.ResponseWriter, request *http.Request) {
	params := mux.Vars(request)
	bookId := params["bookId"]

	id, err := strconv.ParseInt(bookId, 0, 0)

	if err != nil {
		fmt.Println("Error while parsing book id")
	}

	book := models.DeleteBook(id)

	res, _ := json.Marshal(book)
	writer.Header().Set("Content-Type", "application/json")
	writer.WriteHeader(http.StatusOK)
	writer.Write(res)

}

func UpdateBook(writer http.ResponseWriter, request *http.Request) {
	var updateBook = models.Book{}
	utils.ParseBody(request, &updateBook)
	vars := mux.Vars(request)
	bookId := vars["bookId"]
	id, err := strconv.ParseInt(bookId, 0, 0)

	if err != nil {
		fmt.Println("Problem while parsing bookId")
	}
	bookDetails, db := models.GetBookById(id)

	if updateBook.Name != "" {
		bookDetails.Name = updateBook.Name
	}

	if updateBook.Author != "" {
		bookDetails.Author = updateBook.Author
	}

	if updateBook.Publication != "" {
		bookDetails.Publication = updateBook.Publication
	}

	db.Save(&bookDetails)
	res, _ := json.Marshal(bookDetails)
	writer.Header().Set("Content-Type", "application/json")
	writer.WriteHeader(http.StatusOK)
	writer.Write(res)
}

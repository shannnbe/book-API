package handlers

import (
	"book_store/models"
	// "book_store/pkg/mocks"
	"encoding/json"
	"fmt"
	"io"
	"log"
	// "math/rand"
	"net/http"
)

// var bookIDCounter int = 1 // This will store the current ID for the next book

func (h handler) CreateBook(w http.ResponseWriter, r *http.Request){
	// read request body (title, desc, qty)
	defer r.Body.Close()
	body, err := io.ReadAll(r.Body)

	if err != nil {
		log.Fatalln(err)
	}

	var book models.BookModel
	json.Unmarshal(body, &book)

	if result := h.DB.Create(&book); result.Error != nil {
		fmt.Println(result.Error)
	}

	// Increment the book ID counter and assign it to the new book
	// book.Id = bookIDCounter
	// bookIDCounter++

	// // append request to mocks
	// // book.Id = rand.Intn(100) // make random ID
	// mocks.Books = append(mocks.Books, book)

	// send response
	var str = fmt.Sprint(
		"Book is created.\n","\n",
		"Book ID\t\t: ", book.Id, "\n",
		"Title\t\t: ", book.Title, "\n",
		"Description\t: ", book.Description, "\n",
		"Quantty\t\t: ", book.Qty, "\n",
		)

	w.WriteHeader(http.StatusCreated)
	w.Header().Add("Content-Type", "application/json")
	io.WriteString(w, str)
}
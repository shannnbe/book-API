package handlers

import (
	"book_store/models"
	// "book_store/pkg/mocks"
	"encoding/json"
	"io"
	"log"
	"net/http"
	"strconv"
	"fmt"
	"github.com/gorilla/mux"
)

func (h handler) UpdateBook(w http.ResponseWriter, r *http.Request) {
	// read id
	vars := mux.Vars(r)
	id, _ := strconv.Atoi(vars["id"])
	found := false

	// read request body
	defer r.Body.Close()
	body, err := io.ReadAll(r.Body)

	if err != nil {
		log.Fatalln(err)
	}

	var updatedBook models.BookModel
	json.Unmarshal(body, &updatedBook)

	var book models.BookModel

	// cari updatedBooks via id di mocks yg udh ada
	if result := h.DB.First(&book, id); result.Error != nil {
		fmt.Println(result.Error)
	} else {
		if err := h.DB.Save(&book); err.Error != nil {
			w.WriteHeader(http.StatusNotFound)
			json.NewEncoder(w).Encode(found)
			io.WriteString(w, "Update is not successful. Book ID not found.")
			return
		}
	}
	
	book.Title = updatedBook.Title
	book.Description = updatedBook.Description
	book.Qty = updatedBook.Qty

	var str = fmt.Sprint(
		"Book updated! Here's the summary.\n","\n",
		"Book ID\t\t: ", book.Id, "\n",
		"Title\t\t: ", book.Title, "\n",
		"Description\t: ", book.Description, "\n",
		"Quantty\t\t: ", book.Qty, "\n",
		)

	// send response
	w.WriteHeader(http.StatusOK)
	w.Header().Add("Content-Type", "application/json")
	json.NewEncoder(w).Encode(!found)
	io.WriteString(w, str)
}

		



package handlers

import (
	"book_store/models"
	"fmt"
	// "book_store/pkg/mocks"
	"encoding/json"
	"io"
	"net/http"
	"strconv"
	"github.com/gorilla/mux"
)

func (h handler) DeleteBook(w http.ResponseWriter, r *http.Request) {
	// read id
	vars := mux.Vars(r)
	id, _ := strconv.Atoi(vars["id"])
	found := false

	// search id from the mocks
	var book models.BookModel

	if result := h.DB.Find(&book, id); result.Error != nil {
		fmt.Println(result.Error)
	} else {
		if err := h.DB.Delete(&book).Error; err != nil {
			// Error occurred while deleting the book
			w.WriteHeader(http.StatusInternalServerError)
			json.NewEncoder(w).Encode(found)
			io.WriteString(w, "Failed to delete the book. Book ID is not found")
			return
		}
	}
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(!found)
	io.WriteString(w, "Delete succesful.")
}
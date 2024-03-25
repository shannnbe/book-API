package handlers

import (
	"book_store/models"
	// "book_store/pkg/mocks"
	"encoding/json"
	"fmt"
	"net/http"
)

func (h handler) GetBooks(w http.ResponseWriter, r *http.Request) {

	var book []models.BookModel

	if result := h.DB.Find(&book); result.Error != nil {
		fmt.Println(result.Error)
	}

	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(book)
}
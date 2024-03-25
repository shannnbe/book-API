package handlers

import (
	// "book_store/models"
	"book_store/models"
	"book_store/pkg/mocks"
	"fmt"
	"io"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	// "golang.org/x/text/cases"
)

func (h handler) GetBookById(w http.ResponseWriter, r *http.Request){
	// read by id
	vars := mux.Vars(r)
	id, _ := strconv.Atoi(vars["id"])
	found := false


	// search id from the mocks
	outerLoop:
	for _, book := range mocks.Books{
		switch {
		case book.Id == id:

			var book models.BookModel
			if result := h.DB.First(&book, id); result.Error != nil {
				fmt.Println(result.Error)
			}

			found = true
			var str = fmt.Sprint(
				"Book ID found!\n", "\n",
				"Book ID\t\t: ", book.Id, "\n",
				"Title\t\t: ", book.Title, "\n",
				"Description\t: ", book.Description, "\n",
				"Quantty\t\t: ", book.Qty, "\n",
				)
				
			w.WriteHeader(http.StatusOK)
			w.Header().Add("Content-Type", "application/json")
			io.WriteString(w, str)
			break outerLoop
		}
	}
	if !found {
		w.WriteHeader(http.StatusNotFound)
		defer io.WriteString(w, `Can't search by ID. Book ID not found.`)
	}
}
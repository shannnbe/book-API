package main

import (
	// "book_store/pkg/db"
	"book_store/pkg/handlers"
	"book_store/pkg/handlers/db"
	"net/http"

	"github.com/gorilla/mux"
	log "github.com/sirupsen/logrus"
)

func main() {
	DB := db.Init()
	h := handlers.New(DB)
	router := mux.NewRouter()

	log.Println("Starting API")
	router.HandleFunc("/books", h.GetBooks).Methods("GET")
	router.HandleFunc("/book", h.CreateBook).Methods("POST")
	router.HandleFunc("/book/{id}", h.UpdateBook).Methods("PUT")
	router.HandleFunc("/book/{id}", h.DeleteBook).Methods("DELETE")

	// coba-coba
	router.HandleFunc("/books/{id}", h.GetBookById).Methods("GET")
	http.ListenAndServe(":9000", router)
}

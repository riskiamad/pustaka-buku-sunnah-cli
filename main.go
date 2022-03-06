package main

import (
	"log"
	"net/http"
	"pustaka-buku-sunnah-cli/handler"

	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/", handler.HomeHandler).Methods("GET")
	r.HandleFunc("/books/add", handler.AddBookHandler).Methods("GET")
	r.HandleFunc("/books/add", handler.AddBookProcessHandler).Methods("POST")
	r.HandleFunc("/books/{id}", handler.BookDetailsHandler).Methods("GET")
	r.HandleFunc("/books/edit/{id}", handler.EditBookHandler).Methods("GET")
	r.HandleFunc("/books/edit/{id}", handler.EditBookProcessHandler).Methods("POST")
	r.HandleFunc("/books/delete/{id}", handler.DeleteBookHandler).Methods("GET")

	log.Fatal(http.ListenAndServe(":80", r))

}

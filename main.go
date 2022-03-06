package main

import (
	"log"
	"net/http"
	"os"
	"pustaka-buku-sunnah-cli/handler"

	"github.com/gorilla/context"
	"github.com/gorilla/mux"
	"gopkg.in/paytm/grace.v1"
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
	http.Handle("/", r)

	port := os.Getenv("PORT")
	if port == "" {
		port = "8000"
	}
	err := grace.Serve(":"+port, context.ClearHandler(http.DefaultServeMux))
	if err != nil {
		log.Fatalln(err)
	}
}

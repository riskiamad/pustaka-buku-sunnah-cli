package handler

import (
	"bytes"
	"embed"
	"encoding/json"
	"html/template"
	"log"

	"net/http"
	"pustaka-buku-sunnah-cli/entity"
	"strconv"

	"github.com/gorilla/mux"
)

var Views embed.FS

func HomeHandler(w http.ResponseWriter, r *http.Request) {
	resp, err := http.Get("https://api.iamrisk.my.id/v1/books/")
	if err != nil {
		log.Fatalln(err)
	}

	defer resp.Body.Close()

	var result map[string]interface{}
	json.NewDecoder(resp.Body).Decode(&result)
	tmpl := template.Must(template.ParseFS(Views, "views/*.html"))
	err = tmpl.ExecuteTemplate(w, "index", result)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func AddBookHandler(w http.ResponseWriter, r *http.Request) {
	tmpl := template.Must(template.ParseFS(Views, "views/*.html"))
	err := tmpl.ExecuteTemplate(w, "add", nil)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

}

func BookDetailsHandler(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["id"]
	url := "https://api.iamrisk.my.id/v1/books/" + id
	resp, err := http.Get(url)
	if err != nil {
		log.Fatalln(err)
	}

	defer resp.Body.Close()

	var result map[string]interface{}
	json.NewDecoder(resp.Body).Decode(&result)
	tmpl := template.Must(template.ParseFS(Views, "views/*.html"))
	err = tmpl.ExecuteTemplate(w, "details", result)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

}

func EditBookHandler(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["id"]
	url := "https://api.iamrisk.my.id/v1/books/" + id
	resp, err := http.Get(url)
	if err != nil {
		log.Fatalln(err)
	}

	defer resp.Body.Close()

	var result map[string]interface{}
	json.NewDecoder(resp.Body).Decode(&result)
	tmpl := template.Must(template.ParseFS(Views, "views/*.html"))
	err = tmpl.ExecuteTemplate(w, "update", result)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

}

func AddBookProcessHandler(w http.ResponseWriter, r *http.Request) {
	title := r.FormValue("title")
	description := r.FormValue("description")
	rating, _ := strconv.Atoi(r.FormValue("rating"))
	price, _ := strconv.Atoi(r.FormValue("price"))
	discount, _ := strconv.Atoi(r.FormValue("discount"))
	var book = entity.BookRequest{
		Title:       title,
		Description: description,
		Price:       price,
		Discount:    discount,
		Rating:      rating,
	}
	jsonReq, err := json.Marshal(book)
	if err != nil {
		log.Fatalln(err)
	}
	resp, err := http.Post("https://api.iamrisk.my.id/v1/books/", "application/json; charset=utf-8", bytes.NewBuffer(jsonReq))
	if err != nil {
		log.Fatalln(err)
	}

	defer resp.Body.Close()
	http.Redirect(w, r, "/", http.StatusFound)

}

func EditBookProcessHandler(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["id"]
	url := "https://api.iamrisk.my.id/v1/books/" + id
	title := r.FormValue("title")
	description := r.FormValue("description")
	rating, _ := strconv.Atoi(r.FormValue("rating"))
	price, _ := strconv.Atoi(r.FormValue("price"))
	discount, _ := strconv.Atoi(r.FormValue("discount"))
	var book = entity.BookRequest{
		Title:       title,
		Description: description,
		Price:       price,
		Discount:    discount,
		Rating:      rating,
	}
	jsonReq, err := json.Marshal(book)
	if err != nil {
		log.Fatalln(err)
	}
	req, err := http.NewRequest(http.MethodPut, url, bytes.NewBuffer(jsonReq))
	if err != nil {
		log.Fatalln(err)
	}
	req.Header.Set("Content-Type", "application/json; charset=utf-8")
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		log.Fatalln(err)
	}
	defer resp.Body.Close()
	http.Redirect(w, r, "/", http.StatusFound)
}

func DeleteBookHandler(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["id"]
	url := "https://api.iamrisk.my.id/v1/books/" + id
	req, err := http.NewRequest(http.MethodDelete, url, nil)
	if err != nil {
		log.Fatalln(err)
	}
	req.Header.Set("Content-Type", "application/json; charset=utf-8")
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		log.Fatalln(err)
	}

	defer resp.Body.Close()
	http.Redirect(w, r, "/", http.StatusFound)
}

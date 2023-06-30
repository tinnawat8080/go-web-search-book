package main

import (
	"encoding/json"
	"html/template"
	"net/http"
	"os"

	"github.com/joho/godotenv"
)

var tmpl = template.Must(template.ParseGlob("templates/*"))

func main() {
	godotenv.Load()
	port := os.Getenv("PORT")

	http.HandleFunc("/", indexHandler)
	http.HandleFunc("/search", searchHandler)
	http.ListenAndServe(":"+port, nil)
}

func indexHandler(w http.ResponseWriter, r *http.Request) {
	var books []Book = getRecomendBook()
	tmpl.ExecuteTemplate(w, "index", books)
}

func searchHandler(w http.ResponseWriter, r *http.Request) {
	var bookName string
	json.NewDecoder(r.Body).Decode(&bookName)
	var books []Book = searchBook(bookName)
	response, _ := json.Marshal(books)
	w.Write(response)
}

func getRecomendBook() []Book {
	var books []Book
	b1 := Book{}
	b1.Name = "Core JAVA"
	b1.Author = "Cay S. Horstman"
	b1.ImageUrl = "https://horstmann.com/corejava/cj12v2.jpg"
	b1.Price = 369.50
	books = append(books, b1)

	b2 := Book{}
	b2.Name = "Design Pattern"
	b2.Author = "Jame R. Trott"
	b2.ImageUrl = "https://m.media-amazon.com/images/I/81xYhfUxw8L._AC_UF1000,1000_QL80_.jpg"
	b2.Price = 355.50
	books = append(books, b2)

	return books
}

func searchBook(bookName string) []Book {
	var books []Book
	if bookName == "python" {
		b1 := Book{}
		b1.Name = "Python"
		b1.Author = "Mark Luiz"
		b1.ImageUrl = "https://m.media-amazon.com/images/I/519YsN69yEL._AC_UF1000,1000_QL80_.jpg"
		b1.Price = 369.50
		books = append(books, b1)
	} else if bookName == "java" {
		b1 := Book{}
		b1.Name = "Core JAVA"
		b1.Author = "Cay S. Horstman"
		b1.ImageUrl = "https://horstmann.com/corejava/cj12v2.jpg"
		b1.Price = 369.50
		books = append(books, b1)
	}
	return books
}

type Book struct {
	Name     string
	Author   string
	Price    float32
	ImageUrl string
}

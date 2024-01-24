package main

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"log"
	"math/rand"
	"net/http"
	"strconv"
)

// Структура Книги
type Book struct {
	ID     string  `json:"id"`
	Isbn   string  `json:"isbn"`
	Title  string  `json:"title"`
	Author *Author `json:"author"`
}

// Структура Автор
type Author struct {
	ID      string `json:"id"`
	Name    string `json:"name"`
	Surname string `json:"surname"`
}

var books []Book
var author []Author

func main() {
	//Инциализация маршутизатора
	r := mux.NewRouter()

	author = append(author, Author{ID: "2", Name: "Николай", Surname: "Гоголь"})
	author = append(author, Author{ID: "1", Name: "Александр", Surname: "Пушкин"})

	books = append(books, Book{ID: "1", Isbn: "4487324", Title: "Капитанская Дочка", Author: &author[1]})
	books = append(books, Book{ID: "2", Isbn: "4484365", Title: "Шинель ", Author: &author[0]})

	//Обработчик маршрутизатора
	r.HandleFunc("/api/books", getBooks).Methods("GET")
	r.HandleFunc("/api/books/{id}", getBook).Methods("GET")
	r.HandleFunc("/api/books/{id}", createBook).Methods("POST")
	r.HandleFunc("/api/books/{id}/{idauthor}", updateBooks).Methods("PUT")
	r.HandleFunc("/api/books/{id}", deleteBook).Methods("DELETE")
	r.HandleFunc("/api/authors", getAuthors).Methods("GET")
	r.HandleFunc("/api/authors/{id}", getAuthor).Methods("GET")
	r.HandleFunc("/api/authors", createAuthor).Methods("POST")
	r.HandleFunc("/api/authors/{id}", updateAuthor).Methods("PUT")
	r.HandleFunc("/api/authors/{id}", deleteAuthor).Methods("DELETE")

	log.Fatal(http.ListenAndServe(":8000", r))
}

func getBook(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	parms := mux.Vars(r)
	for _, item := range books {
		if item.ID == parms["id"] {
			json.NewEncoder(w).Encode(item)
			return
		}
	}
	json.NewEncoder(w).Encode(Book{})
}

func getBooks(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(books)
}

func createBook(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	parms := mux.Vars(r)
	var book Book
	_ = json.NewDecoder(r.Body).Decode(&book)
	book.ID = strconv.Itoa(rand.Intn(1000000))
	for i, item := range author {
		if item.ID == parms["id"] {
			fmt.Print(&author[i])
			book.Author = &author[i]
			break
		}
	}
	if book.Author != nil {
		books = append(books, book)

		json.NewEncoder(w).Encode(book)
	} else {
		w.WriteHeader(http.StatusNotFound)
	}

}

func deleteBook(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	parms := mux.Vars(r)
	for index, item := range books {
		if item.ID == parms["id"] {
			books = append(books[:index], books[index+1:]...)
			break
		}
	}
	json.NewEncoder(w).Encode(books)
}

func updateBooks(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	parms := mux.Vars(r)

	for i, itemB := range author {
		if itemB.ID == parms["idauthor"] {
			for index, item := range books {
				if item.ID == parms["id"] {
					books = append(books[:index], books[index+1:]...)
					var book Book
					_ = json.NewDecoder(r.Body).Decode(&book)
					book.ID = parms["id"]
					book.Author = &author[i]
					books = append(books, book)
					json.NewEncoder(w).Encode(book)
					break
				}

			}
		} else {
			w.WriteHeader(http.StatusNotFound)
			break
		}
	}

}

func getAuthors(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(author)
}
func getAuthor(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	parms := mux.Vars(r)
	for _, item := range author {
		if item.ID == parms["id"] {
			json.NewEncoder(w).Encode(item)
			return
		}
	}
	json.NewEncoder(w).Encode(Book{})
}

func createAuthor(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var authorUP Author
	_ = json.NewDecoder(r.Body).Decode(&authorUP)
	authorUP.ID = strconv.Itoa(rand.Intn(1000000))
	author = append(author, authorUP)
	json.NewEncoder(w).Encode(author)

}
func deleteAuthor(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	parms := mux.Vars(r)
	for index, item := range author {
		if item.ID == parms["id"] {
			author = append(author[:index], author[index+1:]...)
			break
		}
	}
	json.NewEncoder(w).Encode(books)
}

func updateAuthor(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	parms := mux.Vars(r)

	for index, item := range author {
		if item.ID == parms["id"] {
			author = append(author[:index], author[index+1:]...)
			var auth Author
			_ = json.NewDecoder(r.Body).Decode(&auth)
			auth.ID = parms["id"]
			author = append(author, auth)
			json.NewEncoder(w).Encode(author)
			return
		}
	}

}

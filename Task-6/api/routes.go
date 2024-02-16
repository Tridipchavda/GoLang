package api

import (
	"crud_postgres/api/bookapi"
	"crud_postgres/dbconn"
	"crud_postgres/models"
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

// Api struct with App and Book which contains Bookservice Interface
type API struct {
	app  *dbconn.App
	Book bookapi.Book
}

// Function to create new API
func NewApi(app *dbconn.App) *API {
	return &API{
		app:  app,
		Book: bookapi.NewService(app),
	}
}

// Function to read Book Data from database
func (a *API) GetAllBook(w http.ResponseWriter, r *http.Request) {
	books, err := a.Book.GetAllBook()

	// Handling Error in Operation as Internal Server Error
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}

	// Send all the Data to the Client
	for _, v := range books {
		year := strconv.Itoa(v.Year)
		w.Write([]byte(v.Author + " " + v.Genre + " " + v.Title + " " + v.Publisher + " " + v.ISBN + " " + year + "\n"))
	}
	w.Write([]byte("Successfully retrieved all Books"))
}

// Function to read particular Book Data with ISBN number from database
func (a *API) GetOneBook(w http.ResponseWriter, r *http.Request) {

	// Get ISBN Number from URL
	isbn := mux.Vars(r)["isbn"]
	fmt.Println(isbn)
	book, err := a.Book.GetOneBook(isbn)

	// Handling Error in Operation as Internal Server Error
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}

	// Send all the Data to the Client
	year := strconv.Itoa(book.Year)
	w.Write([]byte(book.Author + " " + book.Genre + " " + book.Title + " " + book.Publisher + " " + book.ISBN + " " + year + "\n"))

	w.Write([]byte("Successfully retrieved all Books"))
}

func (a *API) CreateBook(w http.ResponseWriter, r *http.Request) {
	// Parse the Form to get the Fields
	err := r.ParseForm()
	if err != nil {
		http.Error(w, "Unable to parse form", http.StatusBadRequest)
		return
	}

	// Converting year to Int
	year, err := strconv.Atoi(r.Form.Get("year"))
	if err != nil {
		http.Error(w, "Unable to parse form", http.StatusBadRequest)
		return
	}

	// Enter all Daata in Book and Pass it to CreateBook Function of Book
	book := &models.Book{
		Title:     r.Form.Get("title"),
		Author:    r.Form.Get("author"),
		ISBN:      r.Form.Get("isbn"),
		Publisher: r.Form.Get("publisher"),
		Year:      year,
		Genre:     r.Form.Get("genre"),
	}
	b, err := a.Book.CreateBook(*book)

	// Handling Error in Operation as Internal Server Error
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}
	log.Println(b)
	w.Write([]byte("Successfully created Book"))
}

func (a *API) DeleteOneBook(w http.ResponseWriter, r *http.Request) {
	// Parse the Form to get Fields
	isbn := mux.Vars(r)["isbn"]

	// Handling Error in Operation as Internal Server Error
	err := a.Book.DeleteOneBook(isbn)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}
	w.Write([]byte("Successfully Deleted Book"))
}

func (a *API) UpdateOneBook(w http.ResponseWriter, r *http.Request) {
	// Parse the Form to get Fields
	err := r.ParseForm()
	if err != nil {
		http.Error(w, "Unable to parse form", http.StatusBadRequest)
		return
	}
	// Convert Year into int
	year, err := strconv.Atoi(r.Form.Get("year"))
	if err != nil {
		http.Error(w, "Unable to parse form", http.StatusBadRequest)
		return
	}
	// Passing the book to UpdateOneBook function and handling the error
	book := &models.Book{
		Title:     r.Form.Get("title"),
		Author:    r.Form.Get("author"),
		ISBN:      r.Form.Get("isbn"),
		Publisher: r.Form.Get("publisher"),
		Year:      year,
		Genre:     r.Form.Get("genre"),
	}

	err = a.Book.UpdateOneBook(*book, r.Form.Get("isbn"))
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}
	w.Write([]byte("Successfully Updated Book"))
}

func (api *API) InitializeRouters(router *mux.Router) {
	book := NewApi(api.app)
	router.HandleFunc("/", book.GetAllBook).Methods("GET")
	router.HandleFunc("/{isbn}", book.GetOneBook).Methods("GET")
	router.HandleFunc("/", book.CreateBook).Methods("POST")
	router.HandleFunc("/", book.UpdateOneBook).Methods("PUT")
	router.HandleFunc("/{isbn}", book.DeleteOneBook).Methods("DELETE")
}

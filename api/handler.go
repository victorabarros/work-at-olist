package api

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/victorabarros/challenge-olist/business"
)

// SetUpRoutes set api routes
func SetUpRoutes(r *mux.Router, a *business.Author, b business.Book) {
	r.HandleFunc("/authors", listAuthors(a)).Methods(http.MethodGet)
	r.HandleFunc("/authors/{id:[0-9]+}", getAuthorByID(a)).Methods(http.MethodGet)
	r.HandleFunc("/books", createBook(b)).Methods(http.MethodPost)
	r.HandleFunc("/books", listBooks(b)).Methods(http.MethodGet)
	r.HandleFunc("/books/{id:[0-9]+}", getBook(b)).Methods(http.MethodGet)
	r.HandleFunc("/books/{id:[0-9]+}", putBook(b)).Methods(http.MethodPut)
	r.HandleFunc("/books/{id:[0-9]+}", patchBook(b)).Methods(http.MethodPatch)
	r.HandleFunc("/books/{id:[0-9]+}", deleteBook(b)).Methods(http.MethodDelete)
	// TODO add liveness and probeness
}

// TODO move to utils
type response struct {
	Message string `json:"message"`
}

// TODO move to utils
func writeServiceError(w http.ResponseWriter) {
	// TODO log error
	w.WriteHeader(http.StatusServiceUnavailable)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response{"Fail connection on DB."})
}

package api

import (
	"net/http"
)

// Book type with name and title isbn
type Book struct {
	// implement here
	Title  string
	Author string
	ISBN   int
}

// marshall stuff
func (b Book) ToJSON() []byte {
	return nil
}

// unmarshall stuff
func FromJSON(data []byte) Book {
	return Book{}
}

// http handle for serializing books
func BooksHandleFunc(w http.ResponseWriter, r *http.Request) {

}

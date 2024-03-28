// This is a simple RESTful API for managing books
// Author: Samuel Parente

package main

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"
	"sync"
)

// Book represents a book entity
type Book struct {
	ID     int    `json:"id"`
	Title  string `json:"title"`
	Author string `json:"author"`
}

// MemoryStore stores books in memory
type MemoryStore struct {
	sync.RWMutex
	books map[int]Book
}

// Global memory store
var store MemoryStore

// Create a new book
func createBook(w http.ResponseWriter, r *http.Request) {
	var book Book
	err := json.NewDecoder(r.Body).Decode(&book)
	if err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	store.Lock()
	defer store.Unlock()
	store.books[book.ID] = book

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(book)
}

// Get all books
func getBooks(w http.ResponseWriter, r *http.Request) {
	store.RLock()
	defer store.RUnlock()

	books := make([]Book, 0, len(store.books))
	for _, book := range store.books {
		books = append(books, book)
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(books)
}

// Get a book by ID
func getBookByID(w http.ResponseWriter, r *http.Request) {
	params := r.URL.Query()
	idStr := params.Get("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "Invalid book ID", http.StatusBadRequest)
		return
	}

	store.RLock()
	defer store.RUnlock()

	book, ok := store.books[id]
	if !ok {
		http.Error(w, "Book not found", http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(book)
}

// Update a book by ID
func updateBookByID(w http.ResponseWriter, r *http.Request) {
	params := r.URL.Query()
	idStr := params.Get("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "Invalid book ID", http.StatusBadRequest)
		return
	}

	var updatedBook Book
	err = json.NewDecoder(r.Body).Decode(&updatedBook)
	if err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	store.Lock()
	defer store.Unlock()

	_, ok := store.books[id]
	if !ok {
		http.Error(w, "Book not found", http.StatusNotFound)
		return
	}

	updatedBook.ID = id
	store.books[id] = updatedBook

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(updatedBook)
}

// Delete a book by ID
func deleteBookByID(w http.ResponseWriter, r *http.Request) {
	params := r.URL.Query()
	idStr := params.Get("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "Invalid book ID", http.StatusBadRequest)
		return
	}

	store.Lock()
	defer store.Unlock()

	_, ok := store.books[id]
	if !ok {
		http.Error(w, "Book not found", http.StatusNotFound)
		return
	}

	delete(store.books, id)

	w.WriteHeader(http.StatusNoContent)
}

func main() {
	store.books = make(map[int]Book)

	http.HandleFunc("/books", getBooks)
	http.HandleFunc("/books/create", createBook)
	http.HandleFunc("/books/get", getBookByID)
	http.HandleFunc("/books/update", updateBookByID)
	http.HandleFunc("/books/delete", deleteBookByID)

	log.Fatal(http.ListenAndServe(":8080", nil))
}

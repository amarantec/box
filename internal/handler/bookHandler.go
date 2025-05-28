package handler

import (
	"context"
	"encoding/json"
	"net/http"
	"strconv"
	"time"

	"github.com/amarantec/box/internal"
	"github.com/amarantec/box/internal/book"
)

type BookHandler struct {
	Service book.IBookService
}

func NewBookHandler(service book.IBookService) *BookHandler {
	return &BookHandler{Service: service}
}

func (h *BookHandler) RegisterBook(w http.ResponseWriter, r *http.Request) {
	ctxTimeout, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	var book internal.Book

	if err :=
		json.NewDecoder(r.Body).Decode(&book); err != nil {
		http.Error(w,
			"Could not decode this request. Error: "+err.Error(),
			http.StatusBadRequest)
		return
	}

	response, err := h.Service.RegisterBook(ctxTimeout, book)
	if err != nil {
		http.Error(w,
			"Could not register this book. Error: "+err.Error(),
			http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	if err := json.NewEncoder(w).Encode(map[string]any{
		"response": response,
	}); err != nil {
		http.Error(w,
			"Could not encode this response. Error: "+err.Error(),
			http.StatusInternalServerError)
		return
	}
}

func (h *BookHandler) ListBooks(w http.ResponseWriter, r *http.Request) {
	ctxTimeout, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	response, err := h.Service.ListBooks(ctxTimeout)
	if err != nil {
		http.Error(w,
			"Could not list books from repository. Error: "+err.Error(),
			http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	if err := json.NewEncoder(w).Encode(map[string]any{
		"response": response,
	}); err != nil {
		http.Error(w,
			"Could not encode this response. Error: "+err.Error(),
			http.StatusInternalServerError)
		return
	}
}

func (h *BookHandler) GetBookById(w http.ResponseWriter, r *http.Request) {
	ctxTimeout, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	bookId, err := strconv.ParseInt(r.PathValue("bookId"), 10, 64)
	if err != nil {
		http.Error(w,
			"Invalid parameter. Error: "+err.Error(),
			http.StatusBadRequest)
		return
	}

	response, err := h.Service.GetBookById(ctxTimeout, bookId)
	if err != nil {
		http.Error(w,
			"Could not get this book from repository. Error: "+err.Error(),
			http.StatusInternalServerError)
		return

	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(map[string]any{
		"response": response,
	}); err != nil {
		http.Error(w,
			"Could not encode this response. Error: "+err.Error(),
			http.StatusInternalServerError)
		return

	}
}

func (h *BookHandler) UpdateBook(w http.ResponseWriter, r *http.Request) {
	ctxTimeout, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	var book internal.Book

	if err :=
		json.NewDecoder(r.Body).Decode(&book); err != nil {
		http.Error(w,
			"Could not decode this request. Error: "+err.Error(),
			http.StatusBadRequest)
		return
	}

	response, err := h.Service.UpdateBook(ctxTimeout, book)
	if err != nil {
		http.Error(w,
			"Could not update this book. Error: "+err.Error(),
			http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusNoContent)
	if err := json.NewEncoder(w).Encode(map[string]any{
		"response": response,
	}); err != nil {
		http.Error(w,
			"Could not encode this response. Error: "+err.Error(),
			http.StatusInternalServerError)
		return
	}
}

func (h *BookHandler) DeleteBook(w http.ResponseWriter, r *http.Request) {
	ctxTimeout, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	bookId, err := strconv.ParseInt(r.PathValue("bookId"), 10, 64)
	if err != nil {
		http.Error(w,
			"Invalid parameter. Error: "+err.Error(),
			http.StatusBadRequest)
		return
	}

	response, err := h.Service.DeleteBook(ctxTimeout, bookId)
	if err != nil {
		http.Error(w,
			"Could not delete this book from repository. Error: "+err.Error(),
			http.StatusInternalServerError)
		return

	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusNoContent)
	if err := json.NewEncoder(w).Encode(map[string]any{
		"response": response,
	}); err != nil {
		http.Error(w,
			"Could not encode this response. Error: "+err.Error(),
			http.StatusInternalServerError)
		return

	}
}

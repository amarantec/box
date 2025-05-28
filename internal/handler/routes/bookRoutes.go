package routes

import (
	"net/http"

	"github.com/amarantec/box/internal/handler"
)

func bookRoutes(handler *handler.BookHandler) *http.ServeMux {
	bookMux := http.NewServeMux()

	bookMux.HandleFunc("/register-book", handler.RegisterBook)
	bookMux.HandleFunc("/list-books", handler.ListBooks)
	bookMux.HandleFunc("/get-book/{bookId}", handler.GetBookById)
	bookMux.HandleFunc("/update-book", handler.UpdateBook)
	bookMux.HandleFunc("/delete-book/{bookId}", handler.DeleteBook)

	return bookMux
}

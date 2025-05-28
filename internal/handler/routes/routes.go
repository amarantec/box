package routes

import (
	"net/http"

	"github.com/amarantec/box/internal/book"
	"github.com/amarantec/box/internal/handler"
	"github.com/jackc/pgx/v5/pgxpool"
)

func Router(conn *pgxpool.Pool) *http.ServeMux {
	mux := http.NewServeMux()

	bookRepository := book.NewBookRepository(conn)
	bookService := book.NewBookService(bookRepository)
	bookHandler := handler.NewBookHandler(bookService)

	mux.Handle("/books/", http.StripPrefix("/books", bookRoutes(bookHandler)))

	return mux
}

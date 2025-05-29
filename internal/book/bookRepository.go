package book

import (
	"context"
	"log"
	"time"

	"github.com/amarantec/box/internal"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
)

type IBookRepository interface {
	RegisterBook(ctx context.Context, b internal.Book) (int64, error)
	ListBooks(ctx context.Context) ([]internal.Book, error)
	GetBookById(ctx context.Context, bookId int64) (internal.Book, error)
	UpdateBook(ctx context.Context, book internal.Book) (bool, error)
	DeleteBook(ctx context.Context, bookId int64) (bool, error)
	ListBooksByGenre(ctx context.Context, genre string) ([]internal.Book, error)
	ListBooksByAuthor(ctx context.Context, author string) ([]internal.Book, error)
}

type bookRepository struct {
	Conn *pgxpool.Pool
}

func NewBookRepository(conn *pgxpool.Pool) IBookRepository {
	return &bookRepository{Conn: conn}
}

func (r *bookRepository) RegisterBook(ctx context.Context, b internal.Book) (int64, error) {
	err :=
		r.Conn.QueryRow(
			ctx,
			`INSERT INTO books (title, description, genre, author, publish_date, publisher, pages) VALUES ($1, $2, $3, $4, $5, $6, $7) RETURNING id;`, b.Title, b.Description, b.Genre, b.Author, b.PublishDate, b.Publisher, b.Pages).Scan(&b.ID)

	if err != nil {
		return internal.ZERO, err
	}

	return b.ID, nil
}

func (r *bookRepository) ListBooks(ctx context.Context) ([]internal.Book, error) {
	rows, err :=
		r.Conn.Query(
			ctx,
			`SELECT id, title, description, genre, author, publish_date, publisher, pages 
                FROM books WHERE deleted_at IS NULL;`)

	if err != nil {
		return []internal.Book{}, err
	}

	defer rows.Close()

	var books []internal.Book
	for rows.Next() {
		b := internal.Book{}
		if err := rows.Scan(
			&b.ID,
			&b.Title,
			&b.Description,
			&b.Genre,
			&b.Author,
			&b.PublishDate,
			&b.Publisher,
			&b.Pages,
		); err != nil {
			return []internal.Book{}, err
		}
		books = append(books, b)
	}

	return books, nil
}

func (r *bookRepository) GetBookById(ctx context.Context, bookId int64) (internal.Book, error) {
	var b internal.Book
	if err :=
		r.Conn.QueryRow(
			ctx,
			`SELECT title, description, genre, author, publish_date, publisher, pages 
                FROM books WHERE id = $1 AND deleted_at IS NULL;`, bookId).Scan(&b.Title, &b.Description, &b.Genre, &b.Author, &b.PublishDate,
			&b.Publisher, &b.Pages); err != nil {
		if err == pgx.ErrNoRows {
			return internal.Book{}, internal.ErrBookNotFound
		}
		return internal.Book{}, err
	}

	return b, nil
}

func (r *bookRepository) UpdateBook(ctx context.Context, b internal.Book) (bool, error) {
	result, err :=
		r.Conn.Exec(
			ctx,
			`UPDATE books SET title = $2, description = $3, genre = $4, author = $5, publish_date = $6, publisher = $7, pages = $8, updated_at = $9 WHERE id = $1 AND deleted_at IS NULL;`, b.ID, b.Title, b.Description, b.Genre, b.Author, b.PublishDate, b.Publisher, b.Pages, time.Now(),
		)

	if err != nil {
		return false, err
	}

	if result.RowsAffected() == internal.ZERO {
		log.Printf("Book not found, %d rows affected.\n", result.RowsAffected())
		return false, internal.ErrBookNotFound
	} else {
		log.Printf("Book with ID %d updated.\n", b.ID)
		return true, nil
	}
}

func (r *bookRepository) DeleteBook(ctx context.Context, bookId int64) (bool, error) {
	result, err :=
		r.Conn.Exec(
			ctx,
			"UPDATE books SET deleted_at = $2 WHERE id = $1;", bookId, time.Now())

	if err != nil {
		return false, err
	}

	if result.RowsAffected() == internal.ZERO {
		log.Printf("Book not found, %d rows affected.\n", result.RowsAffected())
		return false, internal.ErrBookNotFound
	} else {
		log.Printf("Book with ID %d deleted.\n", bookId)
		return true, nil
	}
}

func (r *bookRepository) ListBooksByGenre(ctx context.Context, genre string) ([]internal.Book, error) {
	rows, err :=
		r.Conn.Query(
			ctx,
			`SELECT id, title, description, author, publish_date, publisher, pages
            FROM books WHERE genre = $1 AND deleted_at IS NULL;`, genre)

	if err != nil {
		return []internal.Book{}, err
	}

	defer rows.Close()

	var books []internal.Book
	for rows.Next() {
		b := internal.Book{}
		if err := rows.Scan(
			&b.ID,
			&b.Title,
			&b.Description,
			&b.Author,
			&b.PublishDate,
			&b.Publisher,
			&b.Pages); err != nil {
			return []internal.Book{}, err
		}
		books = append(books, b)
	}

	return books, nil
}

func (r *bookRepository) ListBooksByAuthor(ctx context.Context, author string) ([]internal.Book, error) {
	rows, err :=
		r.Conn.Query(
			ctx,
			`SELECT id, title, description, genre, publish_date, publisher, pages
            FROM books WHERE author = $1 AND deleted_at IS NULL;`, author)

	if err != nil {
		return []internal.Book{}, err
	}

	defer rows.Close()

	var books []internal.Book
	for rows.Next() {
		b := internal.Book{}
		if err := rows.Scan(
			&b.ID,
			&b.Title,
			&b.Description,
			&b.Genre,
			&b.PublishDate,
			&b.Publisher,
			&b.Pages); err != nil {
			return []internal.Book{}, err
		}
		books = append(books, b)
	}

	return books, nil
}

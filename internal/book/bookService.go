package book

import (
	"context"

	"github.com/amarantec/box/internal"
)

type IBookService interface {
	RegisterBook(ctx context.Context, b internal.Book) (internal.Response[int64], error)
	ListBooks(ctx context.Context) (internal.Response[[]internal.Book], error)
	GetBookById(ctx context.Context, bookId int64) (internal.Response[internal.Book], error)
	UpdateBook(ctx context.Context, book internal.Book) (internal.Response[bool], error)
	DeleteBook(ctx context.Context, bookId int64) (internal.Response[bool], error)
}

type bookService struct {
	bookRepo IBookRepository
}

func NewBookService(repository IBookRepository) IBookService {
	return &bookService{bookRepo: repository}
}

func (s *bookService) RegisterBook(ctx context.Context, b internal.Book) (internal.Response[int64], error) {
	var response internal.Response[int64]
	data, err := s.bookRepo.RegisterBook(ctx, b)
	if err != nil {
		response.Data = internal.ZERO
		response.Success = false
		return response, err
	}

	response.Data = data
	response.Success = true
	response.Message = "Book Registered Successfuly."
	return response, nil
}

func (s *bookService) ListBooks(ctx context.Context) (internal.Response[[]internal.Book], error) {
	var response internal.Response[[]internal.Book]

	data, err := s.bookRepo.ListBooks(ctx)
	if err != nil {
		response.Data = []internal.Book{}
		response.Success = false
		return response, err
	}
	response.Data = data
	response.Success = true
	response.Message = "All books registered in the system."
	return response, nil
}

func (s *bookService) GetBookById(ctx context.Context, id int64) (internal.Response[internal.Book], error) {
	var response internal.Response[internal.Book]

	data, err := s.bookRepo.GetBookById(ctx, id)
	if err != nil {
		response.Data = internal.Book{}
		response.Success = false
		return response, err
	}
	response.Data = data
	response.Success = true
	response.Message = "Book found successfully."
	return response, nil
}

func (s *bookService) UpdateBook(ctx context.Context, book internal.Book) (internal.Response[bool], error) {
	var response internal.Response[bool]

	data, err := s.bookRepo.UpdateBook(ctx, book)
	if err != nil {
		response.Data = false
		response.Success = false
		return response, err
	}

	response.Data = data
	response.Success = true
	response.Message = "Book updated successfully."
	return response, nil
}

func (s *bookService) DeleteBook(ctx context.Context, bookId int64) (internal.Response[bool], error) {
	var response internal.Response[bool]

	data, err := s.bookRepo.DeleteBook(ctx, bookId)
	if err != nil {
		response.Data = false
		response.Success = false
		return response, err
	}

	response.Data = data
	response.Success = true
	response.Message = "Book deleted successfully."
	return response, nil
}

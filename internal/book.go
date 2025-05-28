package internal

import "time"

type Book struct {
	ID          int64
	Title       string
	Description string
	Genre       []string
	Author      []string
	PublishDate time.Time
	Publisher   string
	Pages       int
	CreatedAt   time.Time
	UpdatedAt   *time.Time
	DeletedAt   *time.Time
}

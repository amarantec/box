package internal

type Response[T any] struct {
	Data    T
	Success bool
	Message string
}

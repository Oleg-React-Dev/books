package repository

import "bookApi/internal/domain"

type BookRepository interface {
	GetAll() []domain.Book
	GetByID(id int) (domain.Book, bool)
	Create(book domain.Book) domain.Book
	Update(id int, book domain.Book) bool
	Delete(id int) bool
}

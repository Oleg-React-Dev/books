package repository

import (
	"bookApi/internal/domain"
	"sync"
)

type BookMemoryRepo struct {
	books []domain.Book
	mu    sync.RWMutex
}

func NewBookMemoryRepo() *BookMemoryRepo {
	return &BookMemoryRepo{books: []domain.Book{}}
}

func (r *BookMemoryRepo) GetAll() []domain.Book {
	r.mu.RLock()
	defer r.mu.RUnlock()
	return r.books
}

func (r *BookMemoryRepo) GetByID(id int) (domain.Book, bool) {
	r.mu.RLock()
	defer r.mu.RUnlock()
	for _, book := range r.books {
		if book.ID == id {
			return book, true
		}
	}
	return domain.Book{}, false
}

func (r *BookMemoryRepo) Create(book domain.Book) domain.Book {
	r.mu.Lock()
	defer r.mu.Unlock()
	book.ID = len(r.books) + 1
	r.books = append(r.books, book)
	return book
}

func (r *BookMemoryRepo) Update(id int, updatedBook domain.Book) bool {
	r.mu.Lock()
	defer r.mu.Unlock()
	for i, book := range r.books {
		if book.ID == id {
			r.books[i] = updatedBook
			r.books[i].ID = id
			return true
		}
	}
	return false
}

func (r *BookMemoryRepo) Delete(id int) bool {
	r.mu.Lock()
	defer r.mu.Unlock()
	for i, book := range r.books {
		if book.ID == id {
			r.books = append(r.books[:i], r.books[i+1:]...)
			return true
		}
	}
	return false
}

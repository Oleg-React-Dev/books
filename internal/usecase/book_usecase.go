package usecase

import (
	"bookApi/internal/domain"
	"bookApi/internal/repository"
)

type BookUseCase struct {
	repo repository.BookRepository
}

func NewBookUseCase(repo repository.BookRepository) *BookUseCase {
	return &BookUseCase{repo: repo}
}

func (uc *BookUseCase) GetAllBooks() []domain.Book {
	return uc.repo.GetAll()
}

func (uc *BookUseCase) GetBookByID(id int) (domain.Book, bool) {
	return uc.repo.GetByID(id)
}

func (uc *BookUseCase) CreateBook(book domain.Book) domain.Book {
	return uc.repo.Create(book)
}

func (uc *BookUseCase) UpdateBook(id int, book domain.Book) bool {
	return uc.repo.Update(id, book)
}

func (uc *BookUseCase) DeleteBook(id int) bool {
	return uc.repo.Delete(id)
}

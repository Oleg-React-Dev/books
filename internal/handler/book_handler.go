package handler

import (
	"bookApi/internal/models"
	"bookApi/internal/repository"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

type BookHandler struct {
	repo repository.BookRepository
}

func NewBookHandler(repo repository.BookRepository) *BookHandler {
	return &BookHandler{repo}
}

// @Summary Получить список книг
// @Description Возвращает все книги из базы данных
// @Tags Books
// @Accept json
// @Produce json
// @Success 200 {array} models.Book "Список книг"
// @Failure 500 {string} string "Ошибка сервера"
// @Router /books [get]
func (h *BookHandler) GetBooks(c echo.Context) error {
	books, err := h.repo.GetAll()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusOK, books)
}

// @Summary Получить книгу по ID
// @Description Возвращает книгу по её уникальному идентификатору
// @Tags Books
// @Accept json
// @Produce json
// @Param id path int true "ID книги"
// @Success 200 {object} models.Book "Найденная книга"
// @Failure 404 {string} string "Книга не найдена"
// @Failure 500 {string} string "Ошибка сервера"
// @Router /books/{id} [get]
func (h *BookHandler) GetBook(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	book, err := h.repo.GetByID(uint(id))
	if err != nil {
		return c.JSON(http.StatusNotFound, "Book not found")
	}
	return c.JSON(http.StatusOK, book)
}

// @Summary Создать новую книгу
// @Description Добавляет новую книгу в базу данных
// @Tags Books
// @Accept json
// @Produce json
// @Param book body models.Book true "Данные книги"
// @Success 201 {object} models.Book "Созданная книга"
// @Failure 400 {string} string "Ошибка валидации данных"
// @Failure 500 {string} string "Ошибка сервера"
// @Router /books [post]
func (h *BookHandler) CreateBook(c echo.Context) error {
	var book models.Book
	if err := c.Bind(&book); err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}
	if err := h.repo.Create(&book); err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusCreated, book)
}

// @Summary Обновить книгу
// @Description Обновляет данные существующей книги по её ID
// @Tags Books
// @Accept json
// @Produce json
// @Param id path int true "ID книги"
// @Param book body models.Book true "Обновлённые данные книги"
// @Success 200 {object} models.Book "Обновлённая книга"
// @Failure 400 {string} string "Ошибка валидации данных"
// @Failure 500 {string} string "Ошибка сервера"
// @Router /books/{id} [put]
func (h *BookHandler) UpdateBook(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))

	var book models.Book
	if err := c.Bind(&book); err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	book.ID = uint(id)
	if err := h.repo.Update(&book); err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, book)
}

// @Summary Удалить книгу
// @Description Удаляет книгу по её уникальному идентификатору
// @Tags Books
// @Accept json
// @Produce json
// @Param id path int true "ID книги"
// @Success 204 "Книга успешно удалена"
// @Failure 500 {string} string "Ошибка сервера"
// @Router /books/{id} [delete]
func (h *BookHandler) DeleteBook(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	if err := h.repo.Delete(uint(id)); err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	return c.NoContent(http.StatusNoContent)
}

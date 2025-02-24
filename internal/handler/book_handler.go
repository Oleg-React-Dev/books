package handler

import (
	"bookApi/internal/models"
	"bookApi/internal/repository"
	"net/http"
	"strconv"

	"github.com/labstack/echo"
)

type BookHandler struct {
	repo repository.BookRepository
}

func NewBookHandler(repo repository.BookRepository) *BookHandler {
	return &BookHandler{repo}
}

func (h *BookHandler) GetBooks(c echo.Context) error {
	books, err := h.repo.GetAll()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusOK, books)
}

func (h *BookHandler) GetBook(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	book, err := h.repo.GetByID(uint(id))
	if err != nil {
		return c.JSON(http.StatusNotFound, "Book not found")
	}
	return c.JSON(http.StatusOK, book)
}

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

func (h *BookHandler) DeleteBook(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	if err := h.repo.Delete(uint(id)); err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	return c.NoContent(http.StatusNoContent)
}

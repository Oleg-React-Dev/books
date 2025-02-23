package delivery

import (
	"bookApi/internal/domain"
	"bookApi/internal/usecase"
	"net/http"
	"strconv"

	"github.com/labstack/echo"
)

type BookHandler struct {
	useCase *usecase.BookUseCase
}

func NewBookHandler(e *echo.Echo, useCase *usecase.BookUseCase) {
	handler := &BookHandler{useCase: useCase}

	e.GET("/books", handler.GetAllBooks)
	e.GET("/books/:id", handler.GetBookByID)
	e.POST("/books", handler.CreateBook)
	e.PUT("/books/:id", handler.UpdateBook)
	e.DELETE("/books/:id", handler.DeleteBook)
}

func (h *BookHandler) GetAllBooks(c echo.Context) error {
	books := h.useCase.GetAllBooks()
	return c.JSON(http.StatusOK, books)
}

func (h *BookHandler) GetBookByID(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "invalid book ID"})
	}

	book, found := h.useCase.GetBookByID(id)
	if !found {
		return c.JSON(http.StatusNotFound, map[string]string{"error": "book not found"})
	}

	return c.JSON(http.StatusOK, book)
}

func (h *BookHandler) CreateBook(c echo.Context) error {
	var book domain.Book
	if err := c.Bind(&book); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "invalid input"})
	}

	createdBook := h.useCase.CreateBook(book)
	return c.JSON(http.StatusCreated, createdBook)
}

func (h *BookHandler) UpdateBook(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "invalid book ID"})
	}

	var book domain.Book
	if err := c.Bind(&book); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "invalid input"})
	}

	if !h.useCase.UpdateBook(id, book) {
		return c.JSON(http.StatusNotFound, map[string]string{"error": "book not found"})
	}

	return c.JSON(http.StatusOK, book)
}

func (h *BookHandler) DeleteBook(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "invalid book ID"})
	}

	if !h.useCase.DeleteBook(id) {
		return c.JSON(http.StatusNotFound, map[string]string{"error": "book not found"})
	}

	return c.JSON(http.StatusNoContent, nil)
}

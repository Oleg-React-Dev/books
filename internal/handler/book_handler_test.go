package handler

import (
	"bookApi/internal/models"
	"bytes"
	"encoding/json"
	"errors"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/labstack/echo"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"gorm.io/gorm"
)

type MockBookRepository struct {
	mock.Mock
}

func (m *MockBookRepository) GetAll() ([]models.Book, error) {
	args := m.Called()
	return args.Get(0).([]models.Book), args.Error(1)
}

func (m *MockBookRepository) GetByID(id uint) (*models.Book, error) {
	args := m.Called(id)
	if args.Get(0) != nil {
		return args.Get(0).(*models.Book), args.Error(1)
	}
	return nil, args.Error(1)
}

func (m *MockBookRepository) Create(book *models.Book) error {
	args := m.Called(book)
	return args.Error(0)
}

func (m *MockBookRepository) Update(book *models.Book) error {
	args := m.Called(book)
	return args.Error(0)
}

func (m *MockBookRepository) Delete(id uint) error {
	args := m.Called(id)
	return args.Error(0)
}

func TestGetBooks(t *testing.T) {

	e := echo.New()
	req := httptest.NewRequest(http.MethodGet, "/books", nil)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	mockRepo := new(MockBookRepository)
	handler := NewBookHandler(mockRepo)

	expectedBooks := []models.Book{
		{Model: gorm.Model{ID: 1}, Title: "Book 1", Author: "Author 1"},
		{Model: gorm.Model{ID: 2}, Title: "Book 2", Author: "Author 2"},
	}

	mockRepo.On("GetAll").Return(expectedBooks, nil)

	err := handler.GetBooks(c)
	assert.NoError(t, err)
	assert.Equal(t, http.StatusOK, rec.Code)

	var books []models.Book
	err = json.Unmarshal(rec.Body.Bytes(), &books)
	assert.NoError(t, err)
	assert.Equal(t, expectedBooks, books)
	mockRepo.AssertExpectations(t)
}

func TestGetBook(t *testing.T) {
	e := echo.New()
	req := httptest.NewRequest(http.MethodGet, "/books/1", nil)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	c.SetParamNames("id")
	c.SetParamValues("1")

	mockRepo := new(MockBookRepository)
	handler := NewBookHandler(mockRepo)

	expectedBook := models.Book{Model: gorm.Model{ID: 1}, Title: "Book 1", Author: "Author 1"}
	mockRepo.On("GetByID", uint(1)).Return(&expectedBook, nil)

	err := handler.GetBook(c)
	assert.NoError(t, err)
	assert.Equal(t, http.StatusOK, rec.Code)

	var book models.Book
	err = json.Unmarshal(rec.Body.Bytes(), &book)
	assert.NoError(t, err)
	assert.Equal(t, expectedBook, book)
	mockRepo.AssertExpectations(t)
}

func TestGetBook_NotFound(t *testing.T) {
	e := echo.New()
	req := httptest.NewRequest(http.MethodGet, "/books/99", nil)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	c.SetParamNames("id")
	c.SetParamValues("99")

	mockRepo := new(MockBookRepository)
	handler := NewBookHandler(mockRepo)

	mockRepo.On("GetByID", uint(99)).Return(&models.Book{}, errors.New("not found"))

	err := handler.GetBook(c)
	assert.NoError(t, err)
	assert.Equal(t, http.StatusNotFound, rec.Code)
	assert.Equal(t, `"Book not found"
`, rec.Body.String())
	mockRepo.AssertExpectations(t)
}

func TestCreateBook(t *testing.T) {
	e := echo.New()
	book := models.Book{Title: "New Book", Author: "New Author"}
	bookJSON, _ := json.Marshal(book)

	req := httptest.NewRequest(http.MethodPost, "/books", bytes.NewBuffer(bookJSON))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	mockRepo := new(MockBookRepository)
	handler := NewBookHandler(mockRepo)

	mockRepo.On("Create", mock.AnythingOfType("*models.Book")).Return(nil)

	err := handler.CreateBook(c)
	assert.NoError(t, err)
	assert.Equal(t, http.StatusCreated, rec.Code)

	var createdBook models.Book
	err = json.Unmarshal(rec.Body.Bytes(), &createdBook)
	assert.NoError(t, err)
	assert.Equal(t, book.Title, createdBook.Title)
	mockRepo.AssertExpectations(t)
}

func TestUpdateBook(t *testing.T) {
	e := echo.New()
	book := models.Book{Model: gorm.Model{ID: 1}, Title: "Updated Title", Author: "Updated Author"}
	bookJSON, _ := json.Marshal(book)

	req := httptest.NewRequest(http.MethodPut, "/books/1", bytes.NewBuffer(bookJSON))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	c.SetParamNames("id")
	c.SetParamValues("1")

	mockRepo := new(MockBookRepository)
	handler := NewBookHandler(mockRepo)

	mockRepo.On("Update", mock.AnythingOfType("*models.Book")).Return(nil)

	err := handler.UpdateBook(c)
	assert.NoError(t, err)
	assert.Equal(t, http.StatusOK, rec.Code)

	var updatedBook models.Book
	err = json.Unmarshal(rec.Body.Bytes(), &updatedBook)
	assert.NoError(t, err)
	assert.Equal(t, book.Title, updatedBook.Title)
	mockRepo.AssertExpectations(t)
}

func TestDeleteBook(t *testing.T) {
	e := echo.New()
	req := httptest.NewRequest(http.MethodDelete, "/books/1", nil)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	c.SetParamNames("id")
	c.SetParamValues("1")

	mockRepo := new(MockBookRepository)
	handler := NewBookHandler(mockRepo)

	mockRepo.On("Delete", uint(1)).Return(nil)

	err := handler.DeleteBook(c)
	assert.NoError(t, err)
	assert.Equal(t, http.StatusNoContent, rec.Code)
	mockRepo.AssertExpectations(t)
}

func TestCreateBook_EmptyPayload(t *testing.T) {
	e := echo.New()
	mockRepo := new(MockBookRepository)
	h := NewBookHandler(mockRepo)
	req := httptest.NewRequest(http.MethodPost, "/books", strings.NewReader(""))
	req.Header.Set("Content-Type", "application/json")
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	err := h.CreateBook(c)

	assert.NoError(t, err)
	assert.Equal(t, http.StatusBadRequest, rec.Code)
}

func TestUpdateBook_InvalidID(t *testing.T) {
	e := echo.New()
	mockRepo := new(MockBookRepository)
	h := NewBookHandler(mockRepo)
	req := httptest.NewRequest(http.MethodPut, "/books/abc", nil)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	c.SetPath("/books/:id")
	c.SetParamNames("id")
	c.SetParamValues("abc")

	err := h.UpdateBook(c)

	assert.NoError(t, err)
	assert.Equal(t, http.StatusBadRequest, rec.Code)
}

func TestDeleteBook_NotFound(t *testing.T) {
	e := echo.New()
	mockRepo := new(MockBookRepository)
	mockRepo.On("Delete", uint(999)).Return(errors.New("not found"))

	h := NewBookHandler(mockRepo)
	req := httptest.NewRequest(http.MethodDelete, "/books/999", nil)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	c.SetPath("/books/:id")
	c.SetParamNames("id")
	c.SetParamValues("999")

	err := h.DeleteBook(c)

	assert.NoError(t, err)
	assert.Equal(t, http.StatusInternalServerError, rec.Code)
}

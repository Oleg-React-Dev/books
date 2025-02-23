package main

import (
	"bookApi/internal/database"
	"bookApi/internal/handler"
	"bookApi/internal/repository"

	"github.com/labstack/echo"
)

func main() {
	db := database.InitDB()
	database.ApplyMigrations(db)

	repo := repository.NewBookRepository(db)
	bookHandler := handler.NewBookHandler(repo)

	e := echo.New()
	e.GET("/books", bookHandler.GetBooks)
	e.GET("/books/:id", bookHandler.GetBook)
	e.POST("/books", bookHandler.CreateBook)
	e.PUT("/books/:id", bookHandler.UpdateBook)
	e.DELETE("/books/:id", bookHandler.DeleteBook)

	e.Start(":8080")
}

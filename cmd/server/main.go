package main

import (
	"bookApi/internal/database"
	"bookApi/internal/handler"
	"bookApi/internal/repository"

	_ "bookApi/docs"

	"github.com/labstack/echo/v4"
	echoSwagger "github.com/swaggo/echo-swagger"
)

func main() {
	db := database.InitDB()
	database.ApplyMigrations(db)

	repo := repository.NewBookRepository(db)
	bookHandler := handler.NewBookHandler(repo)

	e := echo.New()
	e.GET("/swagger/*", echoSwagger.WrapHandler)
	e.GET("/books", bookHandler.GetBooks)
	e.GET("/books/:id", bookHandler.GetBook)
	e.POST("/books", bookHandler.CreateBook)
	e.PUT("/books/:id", bookHandler.UpdateBook)
	e.DELETE("/books/:id", bookHandler.DeleteBook)

	e.Start(":8080")
}

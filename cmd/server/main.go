package main

import (
	delivery "bookApi/internal/delivery/http"
	"bookApi/internal/repository"
	"bookApi/internal/usecase"

	"github.com/labstack/echo"
)

func main() {
	e := echo.New()

    repo := repository.NewBookMemoryRepo()
    useCase := usecase.NewBookUseCase(repo)
    delivery.NewBookHandler(e, useCase)

    e.Logger.Fatal(e.Start(":8080"))
}

package main

import (
	"CustomerApi/Customer/src/handlers"
	"CustomerApi/Customer/src/repositories"
	"CustomerApi/Customer/src/services"
	//	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
)

func main() {
	e := echo.New()
	//	v := validator.New()

	repository := repositories.NewRepository()
	service := services.NewService(repository)
	handler := handlers.NewHandler(service, e) //, v)

	handler.InitEndpoints()

	e.Logger.Fatal(e.Start(":8080"))
}

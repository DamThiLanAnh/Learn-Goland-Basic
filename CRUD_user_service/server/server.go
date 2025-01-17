package main

import (
	"github.com/labstack/echo/v4"
	"golang-basic/CRUD_user_service/handlers"
)

func main() {
	e := echo.New()

	//Routes
	e.POST("/users", handlers.CreateUser)
	e.GET("/users/:id", handlers.GetUser)
	e.PUT("/users/:id", handlers.UpdateUser)
	e.DELETE("/users/:id", handlers.DeleteUser)

	// Start server at localhost:8080
	e.Logger.Fatal(e.Start(":8080"))
}

package main

import (
	"fmt"

	"github.com/j-yw/portfolio/handlers"
	"github.com/labstack/echo/v4"
)

func main() {
	fmt.Println("I'm working")
	app := echo.New()
	app.Static("/assets", "assets")
	pageHandler := handlers.PageHandler{}
	app.GET("/", pageHandler.HandlePage)
	app.Start(":3000")
	fmt.Println("server running on port 3000")
}

package hello

import (
	"github.com/labstack/echo"
)

// Routes defines the routes for the Hello resource
func Routes(e *echo.Echo) {

	// Routes
	e.GET("/hello", GetHello)

}

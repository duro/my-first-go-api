package hello

import (
	"fmt"
	"net/http"

	"github.com/labstack/echo"
)

// GetHello returns a hello resource
func GetHello(c echo.Context) error {
	db := DB

	fmt.Println(db)

	return c.String(http.StatusOK, "Hello, Fucking World!")
}

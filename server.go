package main

import (
	"fmt"
	"log"

	"github.com/duro/my-first-go-api/api/hello"
	"github.com/duro/my-first-go-api/models"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"github.com/labstack/echo"
	"github.com/labstack/echo/engine/standard"
	"github.com/labstack/echo/middleware"
)

// DB is the global DB connection for app
var DB *gorm.DB

func main() {
	// Create new server
	e := echo.New()

	// Connect to DB
	db, err := gorm.Open("postgres", "host=docker.internal user=postgres dbname=first_go_api sslmode=disable password=notsosecretpassword")
	if err != nil {
		log.Println("DB Connection Failure")
		log.Fatal(err)
	} else {
		fmt.Println("DB Connection Success")
		DB = db
	}

	// Run Auto Migrations
	models.DoMigrate(DB)

	// Middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	// Routes
	hello.Routes(e)

	fmt.Println("Server Running at http://localhost:1323")

	// Start Server
	e.Run(standard.New(":1323"))
}

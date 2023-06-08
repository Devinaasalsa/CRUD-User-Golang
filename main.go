package main

import (
	"myapp-me/config"
	"myapp-me/middleware"
	"myapp-me/models"

	// validators "myapp-me/helpers/validator"
	"myapp-me/routes"
	_ "net/http"

	"time"

	"github.com/labstack/echo/v4"
	_ "github.com/labstack/echo/v4/middleware"
	// "github.com/go-playground/validator/v10"
)

func main() {
	app := echo.New()

	//Database init
	config.DatabaseInit()
	gorm := config.DB()

	// validators.Init()

	//Initialize database (auto migrate)
	config.DB().AutoMigrate(
		&models.User{},
	)

	dbGorm, err := gorm.DB()
	if err != nil {
		panic("Couldn't connect database, please check your database connection")
	}

	dbGorm.Ping()

	//Routing registration
	routes.RoutesInit(app)

	//Middleware registration
	app.Use(middleware.SetCORS())
	app.Use(middleware.LoggingRequest)
	app.HTTPErrorHandler = middleware.ErrorHandler

	lock := make(chan error)
	go func(lock chan error) {
		lock <- app.Start(":8080")
	}(lock)

	time.Sleep(1 * time.Millisecond)
	middleware.MakeLogEntry(nil).Warning("Application started without SSL/TTL enabled")

	err = <-lock
	if err != nil {
		middleware.MakeLogEntry(nil).Panic("Failed to start application!")
	}
}

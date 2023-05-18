package main

import (
	"context"
	"log"

	"github.com/dev-hyunsang/daily-todo/database"
	"github.com/dev-hyunsang/daily-todo/middleware"
	"github.com/fatih/color"
	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New()

	middleware.Middleware(app)

	client, err := database.ConnectionDB()
	if err != nil {
		log.Println(color.RedString("[ERROR]"), "Failed Connection DataBase")
		panic(err)
	}

	defer client.Close()

	if err := client.Schema.Create(context.Background()); err != nil {
		panic(err)
	}

	if err := app.Listen(":3000"); err != nil {
		panic(err)
	}
}

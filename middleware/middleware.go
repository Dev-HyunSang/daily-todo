package middleware

import (
	"github.com/dev-hyunsang/daily-todo/users"
	"github.com/gofiber/fiber/v2"
)

func Middleware(app *fiber.App) {
	api := app.Group("/api")

	user := api.Group("/users")
	user.Post("/join", users.JoinUserHandler)
	user.Post("/login", users.LoginUserHandler)
}

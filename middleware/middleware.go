package middleware

import (
	ToDo "github.com/dev-hyunsang/daily-todo/todos"
	"github.com/dev-hyunsang/daily-todo/users"
	"github.com/gofiber/fiber/v2"
)

func Middleware(app *fiber.App) {
	api := app.Group("/api")

	user := api.Group("/users")
	user.Post("/join", users.JoinUserHandler)
	user.Post("/login", users.LoginUserHandler)
	user.Post("/logout", users.LogoutUserHandler)

	todos := api.Group("/todos")
	todos.Post("/create", ToDo.CreateToDoHandler)
	todos.Post("/list", ToDo.AllListToDoHandler)
}

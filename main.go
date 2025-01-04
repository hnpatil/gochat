package main

import (
	"github.com/hnpatil/gochat/handlers"
	"github.com/hnpatil/gochat/handlers/message"
	"github.com/hnpatil/gochat/handlers/room"
	"github.com/hnpatil/gochat/handlers/user"
	"github.com/hnpatil/gochat/migrations"
	"github.com/hnpatil/lofr"
	"gofr.dev/pkg/gofr"
)

func main() {
	app := gofr.New()

	app.Migrate(migrations.All())
	registerRoutes(app, user.New(), room.New(), message.New())

	app.Run()
}

func registerRoutes(app *gofr.App, user handlers.User, room handlers.Room, message handlers.Message) {
	app.POST("/users", lofr.Handler(user.Create))
	app.GET("/users", lofr.Handler(user.List))
	app.PATCH("/users/{id}", lofr.Handler(user.Update))
	app.GET("/users/{id}", lofr.Handler(user.Get))
	app.DELETE("/users/{id}", lofr.Handler(user.Delete))

	app.POST("/rooms", lofr.Handler(room.Create))
	app.GET("/rooms", lofr.Handler(room.List))
	app.PATCH("/rooms/{id}", lofr.Handler(room.Update))
	app.GET("/rooms/{id}", lofr.Handler(room.Get))
	app.DELETE("/rooms/{id}", lofr.Handler(room.Delete))

	app.POST("/rooms/{roomID}/messages", lofr.Handler(message.Create))
	app.GET("/rooms/{roomID}/messages", lofr.Handler(message.List))
	app.PATCH("/rooms/{roomID}/messages/{messageID}", lofr.Handler(message.Update))
	app.GET("/rooms/{roomID}/messages/{messageID}", lofr.Handler(message.Get))
	app.DELETE("/rooms/{roomID}/messages/{messageID}", lofr.Handler(message.Delete))
}

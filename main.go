package main

import (
	"github.com/hnpatil/gochat/handlers"
	"github.com/hnpatil/gochat/handlers/message"
	"github.com/hnpatil/gochat/handlers/room"
	"github.com/hnpatil/gochat/handlers/user"
	"github.com/hnpatil/gochat/migrations"
	messagerepo "github.com/hnpatil/gochat/repos/message"
	roomrepo "github.com/hnpatil/gochat/repos/room"
	userrepo "github.com/hnpatil/gochat/repos/user"
	messagesvc "github.com/hnpatil/gochat/services/message"
	roomsvc "github.com/hnpatil/gochat/services/room"
	usersvc "github.com/hnpatil/gochat/services/user"
	"github.com/hnpatil/lofr"
	"github.com/huandu/go-sqlbuilder"
	"gofr.dev/pkg/gofr"
	"strings"
)

func main() {
	app := gofr.New()

	app.Migrate(migrations.All())
	apiKeys := app.Config.Get("API_KEYS")

	app.EnableAPIKeyAuth(strings.Split(apiKeys, ",")...)
	app.UseMiddleware(lofr.Middleware)

	sqlbuilder.DefaultFlavor = sqlbuilder.PostgreSQL

	roomServices := roomsvc.New(roomrepo.New())

	userHandler := user.New(usersvc.New(userrepo.New()))
	roomHandler := room.New(roomServices)
	messageHandler := message.New(messagesvc.New(messagerepo.New(), roomServices))

	registerRoutes(app, userHandler, roomHandler, messageHandler)

	app.Run()
}

func registerRoutes(app *gofr.App, user handlers.User, room handlers.Room, message handlers.Message) {
	app.POST("/v1/users", lofr.Handler(user.Create))
	app.GET("/v1/users", lofr.Handler(user.List))
	app.PATCH("/v1/users", lofr.Handler(user.Update))
	app.DELETE("/v1/users", lofr.Handler(user.Delete))

	app.POST("/v1/rooms", lofr.Handler(room.Create))
	app.GET("/v1/rooms", lofr.Handler(room.List))
	app.PATCH("/v1/rooms/{id}", lofr.Handler(room.Update))
	app.GET("/v1/rooms/{id}", lofr.Handler(room.Get))
	app.DELETE("/v1/rooms/{id}", lofr.Handler(room.Delete))

	app.POST("/v1/rooms/{roomID}/messages", lofr.Handler(message.Create))
	app.GET("/v1/rooms/{roomID}/messages", lofr.Handler(message.List))
	app.PATCH("/v1/rooms/{roomID}/messages/{messageID}", lofr.Handler(message.Update))
	app.GET("/v1/rooms/{roomID}/messages/{messageID}", lofr.Handler(message.Get))
	app.DELETE("/v1/rooms/{roomID}/messages/{messageID}", lofr.Handler(message.Delete))
}

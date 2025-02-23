package main

import (
	"fmt"
	"github.com/gocql/gocql"
	"github.com/hnpatil/gochat/handlers"
	"github.com/hnpatil/gochat/handlers/message"
	"github.com/hnpatil/gochat/handlers/space"
	"github.com/hnpatil/gochat/migrations"
	messagerepo "github.com/hnpatil/gochat/repos/message"
	spacerepo "github.com/hnpatil/gochat/repos/space"
	messagesvc "github.com/hnpatil/gochat/services/message"
	spacesvc "github.com/hnpatil/gochat/services/space"
	"github.com/hnpatil/lofr"
	"github.com/huandu/go-sqlbuilder"
	"gofr.dev/pkg/gofr"
	"strconv"
	"strings"
)

// @SecurityDefinitions.apikey ApiKeyAuth
// @In header
// @Name X-API-KEY

func main() {
	app := gofr.New()

	cass, err := getCassandra(app)
	if err != nil {
		panic(err)
	}

	err = migrations.Run(app, cass)
	if err != nil {
		panic(err)
	}

	app.Migrate(migrations.All())

	apiKeys := app.Config.Get("API_KEYS")

	app.EnableAPIKeyAuth(strings.Split(apiKeys, ",")...)
	app.UseMiddleware(lofr.Middleware)

	sqlbuilder.DefaultFlavor = sqlbuilder.CQL

	spaceRepo := spacerepo.New(cass)
	messageRepo := messagerepo.New(cass)

	registerRoutes(app, space.New(spacesvc.New(spaceRepo)), message.New(messagesvc.New(messageRepo, spaceRepo)))

	app.Run()
}

func registerRoutes(app *gofr.App, space handlers.UserSpace, message handlers.Message) {
	app.POST("/v1/messages", lofr.Handler(message.Create))
	app.GET("/v1/messages", lofr.Handler(message.List))
	app.GET("/v1/spaces", lofr.Handler(space.List))
}

func getCassandra(app *gofr.App) (*gocql.Session, error) {
	cluster := gocql.NewCluster(app.Config.Get("CASS_HOST"))
	cluster.Port, _ = strconv.Atoi(app.Config.GetOrDefault("CASS_PORT", "9042"))
	cluster.Keyspace = app.Config.Get("CASS_KEYSPACE")
	cluster.AuthProvider = func(h *gocql.HostInfo) (gocql.Authenticator, error) {
		return gocql.PasswordAuthenticator{
			Username: app.Config.Get("CASS_USERNAME"),
			Password: app.Config.Get("CASS_PASSWORD"),
		}, nil
	}

	session, err := cluster.CreateSession()
	if err != nil {
		return nil, fmt.Errorf("failed to connect to Cassandra: %v", err)
	}

	return session, nil
}

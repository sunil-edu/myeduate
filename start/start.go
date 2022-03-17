package main

import (
	"context"
	"myeduate/ent"
	"myeduate/ent/migrate"
	"myeduate/logger"
	"myeduate/routes"
	"myeduate/utils"
	"net/http"
	"time"

	_ "myeduate/ent/runtime"

	"entgo.io/ent/dialect"
	"entgo.io/ent/dialect/sql/schema"

	gohandlers "github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"

	_ "github.com/lib/pq"
	"go.uber.org/zap"
)

var log *zap.SugaredLogger

func main() {

	ctx := context.Background()

	config, err := utils.LoadConfig(".")

	log = logger.InitLogger(config)

	if err != nil {
		log.Fatalf("cannot load config: %v", err)
	}

	if err != nil {
		log.Fatalf("failed opening connection to firebase: %v", err)
	}

	// Create an ent.Client with Postgres database.
	client, err := ent.Open(dialect.Postgres,
		config.DbSource,
		ent.UtilsConfig(&config),
	)
	if err != nil {
		log.Fatalf("failed opening connection to postgress: %v", err)
	}
	defer client.Close()

	if err := client.Schema.Create(ctx,
		schema.WithAtlas(true),
		migrate.WithDropColumn(true),
		migrate.WithDropIndex(true),
		migrate.WithGlobalUniqueID(true)); err != nil {
		log.Fatalf("opening ent client", err)
	}

	// Add a global hook that runs on all types and all operations.
	client.Use(func(next ent.Mutator) ent.Mutator {
		return ent.MutateFunc(func(ctx context.Context, m ent.Mutation) (ent.Value, error) {
			start := time.Now()
			defer func() {
				log.Debugf("Operation = %s \t Table-Name = %s \t Time = %s \t ConcreteType = %T",
					m.Op(), m.Type(), time.Since(start), m)
			}()
			return next.Mutate(ctx, m)
		})
	})

	app := echo.New()

	app.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{config.AllowOrigins},
		AllowHeaders: []string{echo.HeaderOrigin, echo.HeaderContentType, echo.HeaderAccept},
	}))

	routes.Init(app, client)

	app.Logger.Fatal(app.Start(":8081"))
}

func cors(allowedOrigins string) mux.MiddlewareFunc {
	return gohandlers.CORS(
		gohandlers.AllowedOrigins([]string{allowedOrigins}),
		gohandlers.AllowedHeaders([]string{"Access-Control-Allow-Origin", "Access-Control-Allow-Headers", "Origin", "X-Requested-With", "Content-Type", "Accept"}),
		gohandlers.AllowedMethods([]string{http.MethodGet, http.MethodHead, http.MethodPost, http.MethodPut, http.MethodOptions}),
		gohandlers.AllowCredentials(),
	)
}

package main

import (
	"context"
	"fmt"
	"myeduate/ent"
	"myeduate/ent/migrate"
	"myeduate/logger"
	"myeduate/resolvers"
	"myeduate/utils"
	"net/http"
	"time"

	_ "myeduate/ent/runtime"

	"entgo.io/contrib/entgql"
	"entgo.io/ent/dialect"
	"entgo.io/ent/dialect/sql/schema"
	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/handler/transport"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/go-chi/chi"
	"github.com/gorilla/websocket"
	_ "github.com/lib/pq"
	"github.com/rs/cors"
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

	// create a new serve mux and register the handlers
	router := chi.NewRouter()

	// Add CORS middleware around every request
	// See https://github.com/rs/cors for full option listing
	router.Use(cors.New(cors.Options{
		AllowedOrigins:   []string{config.AllowOrigins},
		AllowCredentials: true,
		Debug:            true,
	}).Handler)

	// Configure the server and start listening on :8081.
	srv := handler.NewDefaultServer(resolvers.NewSchema(client))
	srv.Use(entgql.Transactioner{TxOpener: client})
	srv.AddTransport(&transport.Websocket{
		Upgrader: websocket.Upgrader{
			CheckOrigin: func(r *http.Request) bool {
				// Check against your desired domains here
				return r.Host == "example.org"
			},
			ReadBufferSize:  1024,
			WriteBufferSize: 1024,
		},
	})

	router.Handle("/", playground.Handler("myeduate", "/query"))
	router.Handle("/query", srv)
	fmt.Printf("listening on : %v \n", config.ServerAddr)
	err = http.ListenAndServe(config.ServerAddr, router)
	if err != nil {
		panic(err)
	}
}

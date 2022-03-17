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
	"github.com/99designs/gqlgen/graphql/handler/extension"
	"github.com/99designs/gqlgen/graphql/handler/lru"
	"github.com/99designs/gqlgen/graphql/handler/transport"
	"github.com/99designs/gqlgen/graphql/playground"

	gohandlers "github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"github.com/gorilla/websocket"

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

	// create a new serve mux and register the handlers
	route := mux.NewRouter()

	// CORS
	ch := cors(config.AllowOrigins)

	// create a new server
	s := http.Server{
		Addr:    config.ServerAddr, // configure the bind address
		Handler: ch(route),         // set the default handler
		//		ErrorLog:     l.StandardLogger(&hclog.StandardLoggerOptions{}), // set the logger for the server
		ReadTimeout:  5 * time.Second,   // max time to read request from the client
		WriteTimeout: 10 * time.Second,  // max time to write response to the client
		IdleTimeout:  120 * time.Second, // max time for connections using TCP Keep-Alive
	}
	// Configure the server and start listening on :8081.
	srv := handler.New(resolvers.NewSchema(client))

	srv.AddTransport(&transport.Websocket{
		Upgrader: websocket.Upgrader{
			CheckOrigin: func(r *http.Request) bool {
				return true
			},
			ReadBufferSize:  1024,
			WriteBufferSize: 1024,
		},
		KeepAlivePingInterval: 10 * time.Second,
	})

	srv.AddTransport(transport.Options{})
	srv.AddTransport(transport.GET{})
	srv.AddTransport(transport.POST{})
	srv.AddTransport(transport.MultipartForm{})

	srv.SetQueryCache(lru.New(1000))

	srv.Use(extension.Introspection{})
	srv.Use(extension.AutomaticPersistedQuery{
		Cache: lru.New(100),
	})

	srv.Use(entgql.Transactioner{TxOpener: client})

	route.Handle("/", playground.Handler("myeduate", "/query"))
	route.Handle("/query", srv)
	fmt.Printf("listening on : %v \n", config.ServerAddr)
	if err := s.ListenAndServe(); err != nil {
		log.Fatal("http server terminated", err)
	}
}

func cors(allowedOrigins string) mux.MiddlewareFunc {
	return gohandlers.CORS(
		gohandlers.AllowedOrigins([]string{allowedOrigins}),
		gohandlers.AllowedHeaders([]string{"Access-Control-Allow-Origin", "Access-Control-Allow-Headers", "Origin", "X-Requested-With", "Content-Type", "Accept"}),
		gohandlers.AllowedMethods([]string{http.MethodGet, http.MethodHead, http.MethodPost, http.MethodPut, http.MethodOptions}),
		gohandlers.AllowCredentials(),
	)
}

package gql

import (
	"myeduate/ent"
	"myeduate/resolvers"
	"net/http"
	"time"

	"entgo.io/contrib/entgql"
	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/handler/extension"
	"github.com/99designs/gqlgen/graphql/handler/transport"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/gorilla/websocket"
	"github.com/labstack/echo/v4"
)

func Init(app *echo.Echo, client *ent.Client) {

	// do not forget the middleware

	app.GET("/", PlaygroundHandler())
	app.POST("/query", GraphqlHandler(client))
	app.Any("/subscriptions", GraphqlWsHandler(client))

	app.GET("/ws", PlaygroundWsHandler())
}

func PlaygroundHandler() echo.HandlerFunc {
	h := playground.Handler("GraphQL", "/query")

	return func(c echo.Context) error {
		h.ServeHTTP(c.Response(), c.Request())
		return nil
	}
}

func PlaygroundWsHandler() echo.HandlerFunc {
	h := playground.Handler("GraphQL WS", "/subscription")
	return func(c echo.Context) error {
		h.ServeHTTP(c.Response(), c.Request())
		return nil
	}
}

func GraphqlHandler(client *ent.Client) echo.HandlerFunc {
	h := handler.NewDefaultServer(resolvers.NewSchema(client))
	h.Use(entgql.Transactioner{TxOpener: client})

	return func(c echo.Context) error {
		h.ServeHTTP(c.Response(), c.Request())
		return nil
	}
}

func GraphqlWsHandler(client *ent.Client) echo.HandlerFunc {
	h := handler.New(resolvers.NewSchema(client))
	h.AddTransport(transport.POST{})
	h.AddTransport(&transport.Websocket{
		KeepAlivePingInterval: 10 * time.Second,
		Upgrader: websocket.Upgrader{
			CheckOrigin: func(r *http.Request) bool {
				return true
			},
		},
	})
	h.Use(extension.Introspection{})
	return func(c echo.Context) error {
		h.ServeHTTP(c.Response(), c.Request())
		return nil
	}
}



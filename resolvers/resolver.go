package resolvers

import (
	"myeduate"
	"myeduate/ent"
	"myeduate/logger"

	"go.uber.org/zap"

	"github.com/99designs/gqlgen/graphql"
)

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.

// Resolver is the resolver root.
type Resolver struct {
	client *ent.Client
	log    *zap.SugaredLogger
}

// NewSchema creates a graphql executable schema.
func NewSchema(client *ent.Client) graphql.ExecutableSchema {
	return myeduate.NewExecutableSchema(myeduate.Config{
		Resolvers: &Resolver{client, logger.GetLogger()},
	})
}

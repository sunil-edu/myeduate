package resolvers

import (
	"myeduate"
	"myeduate/ent"
	"myeduate/logger"
	"sync"

	"go.uber.org/zap"

	"github.com/99designs/gqlgen/graphql"
)

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.

var schema *graphql.ExecutableSchema

// Resolver is the resolver root.
type Resolver struct {
	client       *ent.Client
	log          *zap.SugaredLogger
	ChatMessages []*ent.MsgChannelMessage
	// All active subscriptions
	ChatObservers map[string]chan []*ent.MsgChannelMessage
	mutex         sync.Mutex
}

// NewSchema creates a graphql executable schema.
func NewSchema(client *ent.Client) graphql.ExecutableSchema {

	if schema == nil {
		schema = new(graphql.ExecutableSchema)
		*schema = myeduate.NewExecutableSchema(myeduate.Config{
			Resolvers: &Resolver{
				client:        client,
				log:           logger.GetLogger(),
				ChatMessages:  []*ent.MsgChannelMessage{},
				ChatObservers: map[string]chan []*ent.MsgChannelMessage{},
			},
		})
	}
	return *schema
}

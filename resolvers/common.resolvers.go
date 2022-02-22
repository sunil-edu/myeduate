package resolvers

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"myeduate"
	"myeduate/ent"
	"myeduate/ent/schema/pulid"
)

func (r *queryResolver) Node(ctx context.Context, token string, id pulid.ID) (ent.Noder, error) {
	return r.client.Noder(ctx, id, ent.WithNodeType(ent.IDToType))
}

func (r *queryResolver) Nodes(ctx context.Context, token string, ids []pulid.ID) ([]ent.Noder, error) {
	return r.client.Noders(ctx, ids, ent.WithNodeType(ent.IDToType))
}

// Query returns myeduate.QueryResolver implementation.
func (r *Resolver) Query() myeduate.QueryResolver { return &queryResolver{r} }

type queryResolver struct{ *Resolver }

package resolvers

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"myeduate"
	"myeduate/ent"
)

func (r *queryResolver) Node(ctx context.Context, token string, id int) (ent.Noder, error) {
	return r.client.Noder(ctx, id)
}

func (r *queryResolver) Nodes(ctx context.Context, token string, ids []int) ([]ent.Noder, error) {
	return r.client.Noders(ctx, ids)
}

// Query returns myeduate.QueryResolver implementation.
func (r *Resolver) Query() myeduate.QueryResolver { return &queryResolver{r} }

type queryResolver struct{ *Resolver }

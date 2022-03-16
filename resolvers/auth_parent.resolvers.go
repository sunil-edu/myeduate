package resolvers

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"fmt"
	"myeduate"
	"myeduate/ent"
	"myeduate/ent/authparent"
)

func (r *mutationResolver) AddAuthParentUser(ctx context.Context, token string, input ent.CreateAuthParentInput) (*ent.AuthParent, error) {
	client := ent.FromContext(ctx)
	return client.AuthParent.Create().SetInput(input).Save(ctx)
}

func (r *queryResolver) GetAuthParentUserNamesByIds(ctx context.Context, token string, id []int) ([]*myeduate.UserNamesByID, error) {
	var parents []*myeduate.UserNamesByID

	if len(id) == 0 {
		return nil, fmt.Errorf("no input")
	}

	for _, i := range id {
		r.client.AuthParent.Query().Where(authparent.IDEQ(i)).Select(authparent.FieldParentFirstName, authparent.FieldParentMiddleName, authparent.FieldParentLastName, authparent.FieldID).ScanX(ctx, &parents)

	}

	return parents, nil
}

// Mutation returns myeduate.MutationResolver implementation.
func (r *Resolver) Mutation() myeduate.MutationResolver { return &mutationResolver{r} }

// Query returns myeduate.QueryResolver implementation.
func (r *Resolver) Query() myeduate.QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }

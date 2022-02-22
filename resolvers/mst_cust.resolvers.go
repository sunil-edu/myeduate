package resolvers

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"fmt"
	"myeduate"
	"myeduate/ent"
	"myeduate/ent/schema/pulid"
)

func (r *mutationResolver) AddCustomer(ctx context.Context, token string, input ent.CreateMstCustomerInput) (*ent.MstCustomer, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) UpdateCustomer(ctx context.Context, token string, id pulid.ID, input ent.UpdateMstCustomerInput) (*ent.MstCustomer, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) ListCustomers(ctx context.Context, token string, after *ent.Cursor, first *int, before *ent.Cursor, last *int, orderBy *ent.MstCustomerOrder) (*ent.MstCustomerConnection, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) GetCustomerByID(ctx context.Context, token string, id pulid.ID) (*ent.MstCustomer, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) GetCustomerIdsByNames(ctx context.Context, token string, custNames []string) ([]*myeduate.CustData, error) {
	panic(fmt.Errorf("not implemented"))
}

// Mutation returns myeduate.MutationResolver implementation.
func (r *Resolver) Mutation() myeduate.MutationResolver { return &mutationResolver{r} }

type mutationResolver struct{ *Resolver }

package resolvers

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"fmt"
	"myeduate"
	"myeduate/ent"
)

func (r *mutationResolver) AddCustomer(ctx context.Context, token string, input ent.CreateMstCustomerInput) (*ent.MstCustomer, error) {
	client := ent.FromContext(ctx)
	return client.MstCustomer.Create().SetInput(input).Save(ctx)
}

func (r *mutationResolver) UpdateCustomer(ctx context.Context, token string, id int, input ent.UpdateMstCustomerInput) (*ent.MstCustomer, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) ListCustomers(ctx context.Context, token string, after *ent.Cursor, first *int, before *ent.Cursor, last *int, orderBy *ent.MstCustomerOrder) (*ent.MstCustomerConnection, error) {
	return r.client.MstCustomer.Query().
		Paginate(ctx, after, first, before, last,
			ent.WithMstCustomerOrder(orderBy))
}

func (r *queryResolver) GetCustomerByID(ctx context.Context, token string, id int) (*ent.MstCustomer, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) GetCustomerIdsByNames(ctx context.Context, token string, custNames []string) ([]*myeduate.CustData, error) {
	panic(fmt.Errorf("not implemented"))
}

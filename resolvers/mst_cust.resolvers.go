package resolvers

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"fmt"
	"myeduate"
	"myeduate/ent"
	"myeduate/ent/mstcustomer"

	"github.com/google/uuid"
)

func (r *mutationResolver) AddCustomer(ctx context.Context, token string, input ent.CreateMstCustomerInput) (*ent.MstCustomer, error) {
	client := ent.FromContext(ctx)

	r.log.Debugf("received request to add new customer %s", input.CustName)
	return client.MstCustomer.Create().SetInput(input).Save(ctx)
}

func (r *mutationResolver) UpdateCustomer(ctx context.Context, token string, id uuid.UUID, input ent.UpdateMstCustomerInput) (*ent.MstCustomer, error) {
	client := ent.FromContext(ctx)

	r.log.Debugf("received request to udate customer %s", input.CustName)
	return client.MstCustomer.UpdateOneID(id).SetInput(input).Save(ctx)
}

func (r *queryResolver) ListCustomers(ctx context.Context, token string, after *ent.Cursor, first *int, before *ent.Cursor, last *int, orderBy *ent.MstCustomerOrder, where *ent.MstCustomerWhereInput) (*ent.MstCustomerConnection, error) {
	return r.client.MstCustomer.Query().
		Paginate(ctx, after, first, before, last,
			ent.WithMstCustomerOrder(orderBy),
			ent.WithMstCustomerFilter(where.Filter))
}

func (r *queryResolver) GetCustomerByID(ctx context.Context, token string, id uuid.UUID) (*ent.MstCustomer, error) {
	return r.client.MstCustomer.Query().Where(mstcustomer.IDEQ(id)).Only(ctx)
}

func (r *queryResolver) GetCustomerIdsByNames(ctx context.Context, token string, custNames []string) ([]*myeduate.CustData, error) {
	custData := make([]*myeduate.CustData, len(custNames))

	if len(custNames) == 0 {
		return nil, fmt.Errorf("no input")
	}

	for i, name := range custNames {
		id, err := r.client.MstCustomer.Query().Where(mstcustomer.CustNameEQ(name)).Select(mstcustomer.FieldID).OnlyID(ctx)

		if err != nil {

		} else {

			custData[i] = &myeduate.CustData{
				CustomerName: name,
				ID:           id,
			}
		}
	}

	return custData, nil
}

// Mutation returns myeduate.MutationResolver implementation.
func (r *Resolver) Mutation() myeduate.MutationResolver { return &mutationResolver{r} }

type mutationResolver struct{ *Resolver }

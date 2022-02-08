package resolvers

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"fmt"
	"myeduate"
	"myeduate/ent"

	"github.com/google/uuid"
)

func (r *mutationResolver) AddInst(ctx context.Context, token string, input ent.CreateMstInstInput) (*ent.MstInst, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) UpdateInst(ctx context.Context, token string, id uuid.UUID, input ent.UpdateMstInstInput) (*ent.MstInst, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) ListInsts(ctx context.Context, token string, after *ent.Cursor, first *int, before *ent.Cursor, last *int, orderBy *ent.MstInstOrder) (*ent.MstInstConnection, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) ListInstsByCustID(ctx context.Context, token string, customerID uuid.UUID, after *ent.Cursor, first *int, before *ent.Cursor, last *int, orderBy *ent.MstInstOrder) (*ent.MstInstConnection, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) GetInstByID(ctx context.Context, token string, id uuid.UUID) (*ent.MstInst, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) GetInstIdsByNames(ctx context.Context, token string, instNames []string) ([]*myeduate.InstData, error) {
	panic(fmt.Errorf("not implemented"))
}

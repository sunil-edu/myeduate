package resolvers

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"fmt"
	"myeduate"
	"myeduate/ent"
	"myeduate/ent/authstaff"
)

func (r *mutationResolver) AddAuthStaffUser(ctx context.Context, token string, input ent.CreateAuthStaffInput) (*ent.AuthStaff, error) {
	client := ent.FromContext(ctx)

	return client.AuthStaff.Create().SetInput(input).Save(ctx)
}

func (r *queryResolver) GetAuthStaffUserIds(ctx context.Context, token string, id []int) ([]*myeduate.UserNamesByID, error) {
	var staff []*myeduate.UserNamesByID

	if len(id) == 0 {
		return nil, fmt.Errorf("no input")
	}

	for _, i := range id {
		r.client.AuthStaff.Query().Where(authstaff.IDEQ(i)).Select(authstaff.FieldStaffFirstName, authstaff.FieldStaffMiddleName, authstaff.FieldStaffLastName, authstaff.FieldID).ScanX(ctx, &staff)

	}

	return staff, nil
}

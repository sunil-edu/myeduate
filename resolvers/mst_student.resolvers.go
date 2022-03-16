package resolvers

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"fmt"
	"myeduate"
	"myeduate/ent"
	"myeduate/ent/mststudent"
)

func (r *mutationResolver) AddStudent(ctx context.Context, token string, input ent.CreateMstStudentInput) (*ent.MstStudent, error) {
	client := ent.FromContext(ctx)
	return client.MstStudent.Create().SetInput(input).Save(ctx)
}

func (r *queryResolver) GetStudentNamesByIds(ctx context.Context, token string, id []int) ([]*myeduate.UserNamesByID, error) {
	var students []*myeduate.UserNamesByID

	if len(id) == 0 {
		return nil, fmt.Errorf("no input")
	}

	for _, i := range id {
		r.client.MstStudent.Query().Where(mststudent.IDEQ(i)).Select(mststudent.FieldStdFirstName, mststudent.FieldStdMiddleName, mststudent.FieldStdLastName, mststudent.FieldID).ScanX(ctx, &students)

	}

	return students, nil
}

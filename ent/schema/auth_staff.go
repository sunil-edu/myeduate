package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
)

// AuthStaff holds the schema definition for the AuthStaff entity.
type AuthStaff struct {
	ent.Schema
}

func (AuthStaff) Mixin() []ent.Mixin {
	return []ent.Mixin{
		TimeMixin{},
	}
}

// Fields of the AuthStaff.
func (AuthStaff) Fields() []ent.Field {
	return []ent.Field{

		field.String("staff_first_name").
			NotEmpty().
			StructTag(`gqlgen:"first_name"`),

		field.String("staff_middle_name").
			StructTag(`gqlgen:"middle_name"`),

		field.String("staff_last_name").
			StructTag(`gqlgen:"last_name"`),


	}
}

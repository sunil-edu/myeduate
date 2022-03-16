package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
)

// AuthParent holds the schema definition for the AuthParent entity.
type AuthParent struct {
	ent.Schema
}

func (AuthParent) Mixin() []ent.Mixin {
	return []ent.Mixin{
		TimeMixin{},
	}
}

// Fields of the AuthParent.
func (AuthParent) Fields() []ent.Field {
	return []ent.Field{

		field.String("parent_first_name").
			NotEmpty().
			StructTag(`gqlgen:"first_name"`),

		field.String("parent_middle_name").
			StructTag(`gqlgen:"middle_name"`),

		field.String("parent_last_name").
			StructTag(`gqlgen:"last_name"`),

		field.String("parent_address"),
		field.String("parent_place"),
		field.String("parent_state"),
		field.String("parent_pin"),

		field.String("parent_email").Unique().NotEmpty(),
		field.String("parent_mobile").Unique().NotEmpty(),
	}
}

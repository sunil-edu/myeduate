package schema

import (
	"myeduate/ent/customtypes"

	"entgo.io/ent"

	"entgo.io/ent/schema/field"
)

// MstStudent holds the schema definition for the MstStudent entity.
type MstStudent struct {
	ent.Schema
}

func (MstStudent) Mixin() []ent.Mixin {
	return []ent.Mixin{
		TimeMixin{},
	}
}

// Fields of the MstStudent.
func (MstStudent) Fields() []ent.Field {
	return []ent.Field{

		field.String("std_first_name").
			NotEmpty().
			StructTag(`gqlgen:"first_name"`),

		field.String("std_middle_name").
			StructTag(`gqlgen:"middle_name"`),

		field.String("std_last_name").
			StructTag(`gqlgen:"last_name"`),

		field.Bool("std_studying").Default(true),

		field.Enum("std_status").
			GoType(customtypes.StdStatus("")).
			Default("CUR"),

		field.Enum("std_sex").
			GoType(customtypes.Sex("")),

		field.String("std_reg_no"),
	}
}

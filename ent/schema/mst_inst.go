package schema

import (
	"myeduate/ent/customtypes"

	"entgo.io/contrib/entgql"
	"entgo.io/ent"
	"entgo.io/ent/dialect"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/index"
	"github.com/google/uuid"
)

// MstInst holds the schema definition for the MstInst entity.
type MstInst struct {
	ent.Schema
}

func (MstInst) Mixin() []ent.Mixin {
	return []ent.Mixin{
		TimeMixin{},
	}
}

// Fields of the MstInst.
func (MstInst) Fields() []ent.Field {
	return []ent.Field{

		field.UUID("id", uuid.UUID{}).Default(uuid.New),

		field.String("inst_code").
			Annotations(entgql.OrderField("INST_CODE")),

		field.String("inst_name").
			NotEmpty().
			Annotations(entgql.OrderField("INST_NAME")),

		field.String("inst_short_name"),

		field.String("inst_address"),

		field.String("inst_place").
			Annotations(entgql.OrderField("INST_PLACE")),

		field.String("inst_state"),

		field.String("inst_pin"),

		field.String("inst_contact_person").
			Annotations(entgql.OrderField("INST_CONTACT_PERSON")),

		field.String("inst_phone").
			Annotations(entgql.OrderField("INST_PHONE")),

		field.String("inst_email"),

		field.String("inst_mobile"),

		field.String("inst_url"),

		field.String("inst_banner_1"),

		field.String("inst_banner_2"),

		field.String("inst_logo_url"),

		field.Enum("inst_is_active").GoType(customtypes.IsActive("")).Default("ACTIVE"),

		field.String("inst_status").
			Annotations(entgql.OrderField("INST_STATUS")),

		field.Time("inst_time_zone").
			SchemaType(map[string]string{
				dialect.Postgres: "DATE",
			}),

		field.UUID("customer_id", uuid.UUID{}).Annotations(
			entgql.OrderField("CUSTOMER_ID")),
	}
}

// Edges of the MstInst.
func (MstInst) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("InstfromCust", MstCustomer.Type).
			Ref("Cust2Inst").
			Unique().
			Required().
			Field("customer_id").
			Annotations(entgql.Bind()),
	}
}

func (MstInst) Indexes() []ent.Index {
	return []ent.Index{

		index.Fields("inst_name", "customer_id").
			Unique(),
	}
}

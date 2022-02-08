package schema

import (
	"myeduate/ent/customtypes"

	"entgo.io/contrib/entgql"
	"entgo.io/ent"
	"entgo.io/ent/dialect"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"github.com/google/uuid"
)

// MstCustomer holds the schema definition for the MstCustomer entity.
type MstCustomer struct {
	ent.Schema
}

func (MstCustomer) Mixin() []ent.Mixin {
	return []ent.Mixin{
		TimeMixin{},
	}
}

// Fields of the MstCustomer.
func (MstCustomer) Fields() []ent.Field {
	return []ent.Field{

		field.UUID("id", uuid.UUID{}).Default(uuid.New),

		field.String("cust_code").
			Unique().
			NotEmpty().
			Annotations(entgql.OrderField("CUST_CODE")),

		field.String("cust_name").
			Unique().
			NotEmpty().
			Annotations(entgql.OrderField("CUST_NAME")),

		field.String("cust_address"),

		field.String("cust_place").
			Annotations(entgql.OrderField("CUST_PLACE")),

		field.String("cust_state"),

		field.String("cust_pin"),

		field.String("cust_contact_person").
			Annotations(entgql.OrderField("CUST_CONTACT_PERSON")),

		field.String("cust_phone").
			Annotations(entgql.OrderField("CUST_PHONE")),

		field.String("cust_email"),

		field.String("cust_mobile"),

		field.String("cust_url"),

		field.String("cust_banner_1"),

		field.String("cust_banner_2"),

		field.String("cust_logo_url"),

		field.Enum("cust_is_active").GoType(customtypes.IsActive("")).Default("ACTIVE"),

		field.String("cust_status").
			Annotations(entgql.OrderField("STATUS")),

		field.Int("cust_num_inst").
			Default(0),

		field.Time("cust_time_zone").
			SchemaType(map[string]string{
				dialect.Postgres: "DATE",
			}),
	}
}

// Edges of the MstCustomer.
func (MstCustomer) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("Cust2Inst", MstInst.Type).Annotations(entgql.Bind()),
	}
}

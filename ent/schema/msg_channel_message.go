package schema

import (
	"time"

	"entgo.io/contrib/entgql"
	"entgo.io/ent"
	"entgo.io/ent/dialect"
	"entgo.io/ent/schema/field"
)

// MsgChannelMessage holds the schema definition for the MsgChannelMessage entity.
type MsgChannelMessage struct {
	ent.Schema
}

func (MsgChannelMessage) Mixin() []ent.Mixin {
	return []ent.Mixin{
		TimeMixin{},
	}
}

// Fields of the MsgChannelMessage.
func (MsgChannelMessage) Fields() []ent.Field {
	return []ent.Field{

		field.Time("msg_date").Default(func() time.Time { return time.UnixMilli(0) }).Optional().Nillable().
			SchemaType(map[string]string{dialect.Postgres: "DATE"}).Annotations(entgql.OrderField("MSG_DATE")),

		field.Bool("msg_is_expiry").Default(false),

		field.Time("msg_expiry_date").Default(func() time.Time { return time.UnixMilli(0) }).Optional().Nillable().
			SchemaType(map[string]string{dialect.Postgres: "DATE"}),

		field.Bool("msg_is_text").Default(true),

		field.String("msg_content"),

		field.String("msg_media_type"),

		field.String("msg_media_content"),

		field.Bool("msg_active").Default(false),

		field.Bool("msg_is_individual"),

		field.String("msg_recv_or_sent"),
	}
}

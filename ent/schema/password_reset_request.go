package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

type PasswordResetRequest struct {
	ent.Schema
}

func (PasswordResetRequest) Fields() []ent.Field {
	return []ent.Field{
		field.Int("user_id").Optional(),
		field.String("token_value"),
		field.Time("expiration_time"),
	}
}

func (PasswordResetRequest) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("user", User.Type).Ref("password_reset_request").Unique().Field("user_id"),
	}
}

func (PasswordResetRequest) Annotations() []schema.Annotation {
	return []schema.Annotation{entsql.Annotation{Table: "password_reset_request"}}

}

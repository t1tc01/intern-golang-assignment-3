package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

type Token struct {
	ent.Schema
}

func (Token) Fields() []ent.Field {
	return []ent.Field{
		field.Int("user_id").Optional(),
		field.String("token_value"),
		field.Time("expiration_time"),
	}
}

func (Token) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("user", User.Type).Ref("token").Unique().Field("user_id"),
	}
}

func (Token) Annotations() []schema.Annotation {
	return []schema.Annotation{entsql.Annotation{Table: "token"}}
}

package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

type Credentials struct {
	ent.Schema
}

func (Credentials) Fields() []ent.Field {
	return []ent.Field{
		field.Int("user_id").Optional(),
		field.String("hashed_password"),
		field.Time("last_login"),
	}
}

func (Credentials) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("user", User.Type).Ref("credentials").Unique().Field("user_id")}
}

func (Credentials) Annotations() []schema.Annotation {
	return []schema.Annotation{entsql.Annotation{Table: "credentials"}}

}

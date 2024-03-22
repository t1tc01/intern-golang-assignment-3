package schema

import (
	"entgo.io/contrib/entgql"
	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

type User struct {
	ent.Schema
}

func (User) Fields() []ent.Field {
	return []ent.Field{
		field.Int("id"),
		field.String("username"),
		field.String("email"),
		field.String("password"),
		field.Int("role"),
		field.Time("created_at"),
		field.Time("updated_at"),
		field.Time("deleted_at").Optional(),
	}
}

func (User) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("session", Session.Type),
		edge.To("password_reset_request", PasswordResetRequest.Type),
		edge.To("credentials", Credentials.Type),
		edge.To("token", Token.Type),
	}
}

func (User) Annotations() []schema.Annotation {
	return []schema.Annotation{entsql.Annotation{Table: "user"},
		entgql.QueryField(),
		entgql.Mutations(entgql.MutationCreate())}
}

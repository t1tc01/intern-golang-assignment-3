package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

type Location struct {
	ent.Schema
}

func (Location) Fields() []ent.Field {
	return []ent.Field{
		field.Int("id"),
		field.String("place").Optional(),
		field.String("address").Optional(),
		field.Time("created_at"),
		field.Time("updated_at"),
		field.Time("deleted_at").Optional()}
}
func (Location) Edges() []ent.Edge {
	return []ent.Edge{edge.To("geometries", Geometry.Type)}
}
func (Location) Annotations() []schema.Annotation {
	return []schema.Annotation{entsql.Annotation{Table: "location"}}
}

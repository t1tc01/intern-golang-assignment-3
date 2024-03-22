package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

type Report struct {
	ent.Schema
}

func (Report) Fields() []ent.Field {
	return []ent.Field{
		field.Int("id"),
		field.Int32("felt").Optional(),
		field.Float("cdi").Optional(),
		field.Float("mmi").Optional(),
		field.String("alert").Optional(),
		field.Time("created_at"),
		field.Time("updated_at"),
		field.Time("deleted_at").Optional()}
}
func (Report) Edges() []ent.Edge {
	return []ent.Edge{edge.To("earthquakes", Earthquake.Type)}
}
func (Report) Annotations() []schema.Annotation {
	return []schema.Annotation{entsql.Annotation{Table: "report"}}
}

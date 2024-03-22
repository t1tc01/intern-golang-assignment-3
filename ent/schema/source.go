package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

type Source struct {
	ent.Schema
}

func (Source) Fields() []ent.Field {
	return []ent.Field{field.Int("id"),
		field.String("name").Optional(),
		field.Time("created_at"),
		field.Time("updated_at"),
		field.Time("deleted_at").Optional()}
}
func (Source) Edges() []ent.Edge {
	return []ent.Edge{edge.To("source_earthquakes", SourceEarthquake.Type)}
}
func (Source) Annotations() []schema.Annotation {
	return []schema.Annotation{entsql.Annotation{Table: "source"}}
}

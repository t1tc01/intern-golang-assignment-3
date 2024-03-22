package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

type Geometry struct {
	ent.Schema
}

func (Geometry) Fields() []ent.Field {
	return []ent.Field{field.Int("id"), field.Int("location_id").Optional(), field.Float("longitude"), field.Float("latitude"), field.Float("depth"), field.Time("created_at"), field.Time("updated_at"), field.Time("deleted_at").Optional()}
}
func (Geometry) Edges() []ent.Edge {
	return []ent.Edge{edge.To("earthquakes", Earthquake.Type),
		edge.From("location", Location.Type).Ref("geometries").Unique().Field("location_id")}
}
func (Geometry) Annotations() []schema.Annotation {
	return []schema.Annotation{entsql.Annotation{Table: "geometry"}}
}

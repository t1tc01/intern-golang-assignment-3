package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

type SourceEarthquake struct {
	ent.Schema
}

func (SourceEarthquake) Fields() []ent.Field {
	return []ent.Field{field.Int("id"), field.Int("s_id").Optional(), field.Int("eq_id").Optional(), field.Time("created_at"), field.Time("updated_at"), field.Time("deleted_at").Optional()}
}
func (SourceEarthquake) Edges() []ent.Edge {
	return []ent.Edge{edge.From("earthquake", Earthquake.Type).Ref("source_earthquakes").Unique().Field("eq_id"), edge.From("source", Source.Type).Ref("source_earthquakes").Unique().Field("s_id")}
}
func (SourceEarthquake) Annotations() []schema.Annotation {
	return []schema.Annotation{entsql.Annotation{Table: "source_earthquake"}}
}

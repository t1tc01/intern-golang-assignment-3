package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

type FeatureType struct {
	ent.Schema
}

func (FeatureType) Fields() []ent.Field {
	return []ent.Field{field.Int("id"), field.String("feat_type").Optional(), field.Time("created_at"), field.Time("updated_at"), field.Time("deleted_at").Optional()}
}
func (FeatureType) Edges() []ent.Edge {
	return []ent.Edge{edge.To("ftype_earthquakes", FtypeEarthquake.Type)}
}
func (FeatureType) Annotations() []schema.Annotation {
	return []schema.Annotation{entsql.Annotation{Table: "feature_type"}}
}

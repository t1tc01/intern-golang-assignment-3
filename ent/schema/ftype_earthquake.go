package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

type FtypeEarthquake struct {
	ent.Schema
}

func (FtypeEarthquake) Fields() []ent.Field {
	return []ent.Field{field.Int("id"), field.Int("ft_id").Optional(), field.Int("eq_id").Optional(), field.Time("created_at"), field.Time("updated_at"), field.Time("deleted_at").Optional()}
}
func (FtypeEarthquake) Edges() []ent.Edge {
	return []ent.Edge{edge.From("earthquake", Earthquake.Type).Ref("ftype_earthquakes").Unique().Field("eq_id"), edge.From("feature_type", FeatureType.Type).Ref("ftype_earthquakes").Unique().Field("ft_id")}
}
func (FtypeEarthquake) Annotations() []schema.Annotation {
	return []schema.Annotation{entsql.Annotation{Table: "ftype_earthquake"}}
}

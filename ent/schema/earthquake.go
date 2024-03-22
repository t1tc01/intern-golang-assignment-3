package schema

import (
	"entgo.io/contrib/entgql"
	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

type Earthquake struct {
	ent.Schema
}

func (Earthquake) Fields() []ent.Field {
	return []ent.Field{
		field.Int("geo_id").Optional(),
		field.Int("report_id").Optional(),
		field.Float("mag"),
		field.Time("time"),
		field.Time("updated_time").Optional(),
		field.Int32("tz").Optional(),
		field.String("url").Optional(),
		field.String("detail").Optional(),
		field.String("status").Optional(),
		field.Int32("tsunami").Optional(),
		field.Int32("sig").Optional(),
		field.String("net").Optional(),
		field.String("code").Optional(),
		field.Int32("nst").Optional(),
		field.Float("dmin").Optional(),
		field.Float("rms").Optional(),
		field.Float("gap").Optional(),
		field.String("mag_type").Optional(),
		field.String("eq_type").Optional(),
		field.Time("created_at"),
		field.Time("updated_at"),
		field.Time("deleted_at").Optional(),
	}
}

func (Earthquake) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("geometry", Geometry.Type).Ref("earthquakes").Unique().Field("geo_id"),
		edge.From("report", Report.Type).Ref("earthquakes").Unique().Field("report_id"),
		edge.To("ftype_earthquakes", FtypeEarthquake.Type),
		edge.To("source_earthquakes", SourceEarthquake.Type)}
}
func (Earthquake) Annotations() []schema.Annotation {
	return []schema.Annotation{entsql.Annotation{Table: "earthquake"},
		entgql.QueryField(),
		entgql.Mutations(entgql.MutationCreate())}
}

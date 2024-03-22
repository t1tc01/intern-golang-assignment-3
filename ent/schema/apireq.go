package schema

import (
	"entgo.io/contrib/entgql"
	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/field"
)

type Apireq struct {
	ent.Schema
}

func (Apireq) Fields() []ent.Field {
	return []ent.Field{field.Int("id"),
		field.Time("req_time"),
		field.JSON("req_param", map[string]interface{}{}).Optional(),
		field.JSON("req_body", map[string]interface{}{}).Optional(),
		field.JSON("req_headers", map[string]interface{}{}).Optional(),
		field.JSON("req_metadata", map[string]interface{}{}).Optional(),
		field.Time("resp_time"),
		field.Int("resp_status").Optional(),
		field.JSON("resp_body", map[string]interface{}{}).Optional(),
		field.JSON("resp_headers", map[string]interface{}{}).Optional(),
		field.JSON("resp_metadata", map[string]interface{}{}).Optional(),
		field.Time("created_at"),
		field.Time("updated_at"),
		field.Time("deleted_at").Optional()}
}

func (Apireq) Edges() []ent.Edge {
	return nil
}

func (Apireq) Annotations() []schema.Annotation {
	return []schema.Annotation{entsql.Annotation{Table: "apireq"},
		entgql.QueryField(),
		entgql.Mutations(entgql.MutationCreate())}
}

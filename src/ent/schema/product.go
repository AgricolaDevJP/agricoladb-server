package schema

import (
	"entgo.io/contrib/entgql"
	"entgo.io/ent"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

type Product struct {
	ent.Schema
}

func (Product) Fields() []ent.Field {
	return []ent.Field{
		field.Int("id"),
		field.Int("revision_id").Immutable(),
		field.Bool("is_official_ja"),
		field.String("name_ja").Optional(),
		field.String("name_en").Optional(),
		field.Int("published_year").Optional(),
	}
}

func (Product) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("cards", Card.Type).Annotations(entgql.RelayConnection()),
		edge.From("revision", Revision.Type).Ref("products").Unique().Field("revision_id").Immutable().Required(),
	}
}

func (Product) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entgql.QueryField(),
	}
}

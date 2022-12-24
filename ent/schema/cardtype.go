package schema

import (
	"entgo.io/contrib/entgql"
	"entgo.io/ent"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

type CardType struct {
	ent.Schema
}

func (CardType) Fields() []ent.Field {
	return []ent.Field{
		field.Int("id"),
		field.String("key").Immutable().NotEmpty().Unique(),
		field.String("name_ja").Optional(),
		field.String("name_en").Optional(),
	}
}

func (CardType) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("cards", Card.Type).Annotations(entgql.RelayConnection()),
	}
}

func (CardType) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entgql.QueryField(),
	}
}

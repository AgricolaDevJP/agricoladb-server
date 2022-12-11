package schema

import (
	"entgo.io/contrib/entgql"
	"entgo.io/ent"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

type Revision struct {
	ent.Schema
}

func (Revision) Fields() []ent.Field {
	return []ent.Field{
		field.String("key").Immutable().NotEmpty().Unique(),
		field.String("name_ja").Optional(),
		field.String("name_en").Optional(),
	}
}

func (Revision) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("cards", Card.Type).Annotations(entgql.RelayConnection()),
		edge.To("products", Product.Type),
		edge.To("decks", Deck.Type),
	}
}

func (Revision) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entgql.QueryField(),
	}
}

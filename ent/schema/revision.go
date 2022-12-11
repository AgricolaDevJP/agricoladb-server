package schema

import (
	"entgo.io/ent"
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
		edge.To("cards", Card.Type),
		edge.To("products", Product.Type),
		edge.To("decks", Deck.Type),
	}
}

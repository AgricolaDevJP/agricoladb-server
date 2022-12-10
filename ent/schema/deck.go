package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

type Deck struct {
	ent.Schema
}

func (Deck) Fields() []ent.Field {
	return []ent.Field{
		field.String("key").Immutable().NotEmpty().Unique(),
		field.Int("revision_id").Immutable(),
		field.String("name_ja").Optional(),
		field.String("name_en").Optional(),
	}
}

func (Deck) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("cards", Card.Type),
		edge.From("revision", Revision.Type).Ref("decks").Unique().Field("revision_id").Immutable().Required(),
	}
}

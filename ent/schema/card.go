package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/index"
)

type Card struct {
	ent.Schema
}

func (Card) Fields() []ent.Field {
	return []ent.Field{
		field.String("literal_id").MaxLen(32).NotEmpty().Immutable(),
		field.Int("revision_id").Immutable(),
		field.String("printed_id").Optional(),
		field.String("play_agricola_card_id").Optional(),
		field.Int("deck_id").Optional(),
		field.Int("card_type_id"),
		field.Int("card_special_color_id").Optional(),
		field.String("name_ja").Optional(),
		field.String("name_en").Optional(),
		field.Int("min_players_number").Optional(),
		field.String("prerequisite").Optional(),
		field.String("cost").Optional(),
		field.Text("function_text").Optional(),
		field.Bool("is_official_ja"),
		field.Int("victory_point").Optional(),
		field.Bool("is_variable_victory_point"),
		field.Bool("has_arrrow"),
		field.Bool("has_bonus_point_icon"),
		field.Bool("has_negative_bonus_point_icon"),
		field.Bool("has_pan_icon"),
		field.Bool("has_bread_icon"),
		field.Bool("has_farm_planner_icon"),
		field.Bool("has_actions_booster_icon"),
		field.Bool("has_points_provider_icon"),
		field.Bool("has_goods_provider_icon"),
		field.Bool("has_food_provider_icon"),
		field.Bool("has_crop_provider_icon"),
		field.Bool("has_building_resource_provider_icon"),
		field.Bool("has_livestock_provider_icon"),
		field.Bool("has_cut_peat_icon"),
		field.Bool("has_fell_trees_icon"),
		field.Bool("has_slash_and_burn_icon"),
		field.Bool("has_hiring_fare_icon"),
	}
}

func (Card) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("revision", Revision.Type).Ref("cards").Unique().Field("revision_id").Immutable().Required(),
		edge.From("products", Product.Type).Ref("cards"),
		edge.From("deck", Deck.Type).Ref("cards").Unique().Field("deck_id"),
		edge.From("card_type", CardType.Type).Ref("cards").Unique().Field("card_type_id").Required(),
		edge.From("card_special_color", CardSpecialColor.Type).Ref("cards").Unique().Field("card_special_color_id"),
		edge.To("ancestors", Card.Type).From("children"),
	}
}

func (Card) Indexes() []ent.Index {
	return []ent.Index{
		index.Fields("revision_id", "literal_id").Unique(),
	}
}

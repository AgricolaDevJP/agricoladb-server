// Code generated by ent, DO NOT EDIT.

package migrate

import (
	"entgo.io/ent/dialect/sql/schema"
	"entgo.io/ent/schema/field"
)

var (
	// CardsColumns holds the columns for the "cards" table.
	CardsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "literal_id", Type: field.TypeString, Size: 32},
		{Name: "printed_id", Type: field.TypeString, Nullable: true},
		{Name: "play_agricola_card_id", Type: field.TypeString, Nullable: true},
		{Name: "name_ja", Type: field.TypeString, Nullable: true},
		{Name: "name_en", Type: field.TypeString, Nullable: true},
		{Name: "min_players_number", Type: field.TypeInt, Nullable: true},
		{Name: "prerequisite", Type: field.TypeString, Nullable: true},
		{Name: "cost", Type: field.TypeString, Nullable: true},
		{Name: "description", Type: field.TypeString, Nullable: true, Size: 2147483647},
		{Name: "note", Type: field.TypeString, Nullable: true, Size: 2147483647},
		{Name: "is_official_ja", Type: field.TypeBool},
		{Name: "victory_point", Type: field.TypeInt, Nullable: true},
		{Name: "special_victory_point", Type: field.TypeString, Nullable: true},
		{Name: "has_arrow", Type: field.TypeBool},
		{Name: "has_bonus_point_icon", Type: field.TypeBool},
		{Name: "has_negative_bonus_point_icon", Type: field.TypeBool},
		{Name: "has_pan_icon", Type: field.TypeBool},
		{Name: "has_bread_icon", Type: field.TypeBool},
		{Name: "has_farm_planner_icon", Type: field.TypeBool},
		{Name: "has_actions_booster_icon", Type: field.TypeBool},
		{Name: "has_points_provider_icon", Type: field.TypeBool},
		{Name: "has_goods_provider_icon", Type: field.TypeBool},
		{Name: "has_food_provider_icon", Type: field.TypeBool},
		{Name: "has_crop_provider_icon", Type: field.TypeBool},
		{Name: "has_building_resource_provider_icon", Type: field.TypeBool},
		{Name: "has_livestock_provider_icon", Type: field.TypeBool},
		{Name: "has_cut_peat_icon", Type: field.TypeBool},
		{Name: "has_fell_trees_icon", Type: field.TypeBool},
		{Name: "has_slash_and_burn_icon", Type: field.TypeBool},
		{Name: "has_hiring_fare_icon", Type: field.TypeBool},
		{Name: "card_special_color_id", Type: field.TypeInt, Nullable: true},
		{Name: "card_type_id", Type: field.TypeInt},
		{Name: "deck_id", Type: field.TypeInt, Nullable: true},
		{Name: "revision_id", Type: field.TypeInt},
	}
	// CardsTable holds the schema information for the "cards" table.
	CardsTable = &schema.Table{
		Name:       "cards",
		Columns:    CardsColumns,
		PrimaryKey: []*schema.Column{CardsColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "cards_card_special_colors_cards",
				Columns:    []*schema.Column{CardsColumns[31]},
				RefColumns: []*schema.Column{CardSpecialColorsColumns[0]},
				OnDelete:   schema.SetNull,
			},
			{
				Symbol:     "cards_card_types_cards",
				Columns:    []*schema.Column{CardsColumns[32]},
				RefColumns: []*schema.Column{CardTypesColumns[0]},
				OnDelete:   schema.NoAction,
			},
			{
				Symbol:     "cards_decks_cards",
				Columns:    []*schema.Column{CardsColumns[33]},
				RefColumns: []*schema.Column{DecksColumns[0]},
				OnDelete:   schema.SetNull,
			},
			{
				Symbol:     "cards_revisions_cards",
				Columns:    []*schema.Column{CardsColumns[34]},
				RefColumns: []*schema.Column{RevisionsColumns[0]},
				OnDelete:   schema.NoAction,
			},
		},
		Indexes: []*schema.Index{
			{
				Name:    "card_revision_id_literal_id",
				Unique:  true,
				Columns: []*schema.Column{CardsColumns[34], CardsColumns[1]},
			},
		},
	}
	// CardSpecialColorsColumns holds the columns for the "card_special_colors" table.
	CardSpecialColorsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "key", Type: field.TypeString, Unique: true},
		{Name: "name_ja", Type: field.TypeString, Nullable: true},
		{Name: "name_en", Type: field.TypeString, Nullable: true},
	}
	// CardSpecialColorsTable holds the schema information for the "card_special_colors" table.
	CardSpecialColorsTable = &schema.Table{
		Name:       "card_special_colors",
		Columns:    CardSpecialColorsColumns,
		PrimaryKey: []*schema.Column{CardSpecialColorsColumns[0]},
	}
	// CardTypesColumns holds the columns for the "card_types" table.
	CardTypesColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "key", Type: field.TypeString, Unique: true},
		{Name: "name_ja", Type: field.TypeString, Nullable: true},
		{Name: "name_en", Type: field.TypeString, Nullable: true},
	}
	// CardTypesTable holds the schema information for the "card_types" table.
	CardTypesTable = &schema.Table{
		Name:       "card_types",
		Columns:    CardTypesColumns,
		PrimaryKey: []*schema.Column{CardTypesColumns[0]},
	}
	// DecksColumns holds the columns for the "decks" table.
	DecksColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "key", Type: field.TypeString, Unique: true},
		{Name: "name_ja", Type: field.TypeString, Nullable: true},
		{Name: "name_en", Type: field.TypeString, Nullable: true},
		{Name: "revision_id", Type: field.TypeInt},
	}
	// DecksTable holds the schema information for the "decks" table.
	DecksTable = &schema.Table{
		Name:       "decks",
		Columns:    DecksColumns,
		PrimaryKey: []*schema.Column{DecksColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "decks_revisions_decks",
				Columns:    []*schema.Column{DecksColumns[4]},
				RefColumns: []*schema.Column{RevisionsColumns[0]},
				OnDelete:   schema.NoAction,
			},
		},
	}
	// ProductsColumns holds the columns for the "products" table.
	ProductsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "is_official_ja", Type: field.TypeBool},
		{Name: "name_ja", Type: field.TypeString, Nullable: true},
		{Name: "name_en", Type: field.TypeString, Nullable: true},
		{Name: "published_year", Type: field.TypeInt, Nullable: true},
		{Name: "revision_id", Type: field.TypeInt},
	}
	// ProductsTable holds the schema information for the "products" table.
	ProductsTable = &schema.Table{
		Name:       "products",
		Columns:    ProductsColumns,
		PrimaryKey: []*schema.Column{ProductsColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "products_revisions_products",
				Columns:    []*schema.Column{ProductsColumns[5]},
				RefColumns: []*schema.Column{RevisionsColumns[0]},
				OnDelete:   schema.NoAction,
			},
		},
	}
	// RevisionsColumns holds the columns for the "revisions" table.
	RevisionsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "key", Type: field.TypeString, Unique: true},
		{Name: "name_ja", Type: field.TypeString, Nullable: true},
		{Name: "name_en", Type: field.TypeString, Nullable: true},
	}
	// RevisionsTable holds the schema information for the "revisions" table.
	RevisionsTable = &schema.Table{
		Name:       "revisions",
		Columns:    RevisionsColumns,
		PrimaryKey: []*schema.Column{RevisionsColumns[0]},
	}
	// CardAncestorsColumns holds the columns for the "card_ancestors" table.
	CardAncestorsColumns = []*schema.Column{
		{Name: "card_id", Type: field.TypeInt},
		{Name: "child_id", Type: field.TypeInt},
	}
	// CardAncestorsTable holds the schema information for the "card_ancestors" table.
	CardAncestorsTable = &schema.Table{
		Name:       "card_ancestors",
		Columns:    CardAncestorsColumns,
		PrimaryKey: []*schema.Column{CardAncestorsColumns[0], CardAncestorsColumns[1]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "card_ancestors_card_id",
				Columns:    []*schema.Column{CardAncestorsColumns[0]},
				RefColumns: []*schema.Column{CardsColumns[0]},
				OnDelete:   schema.Cascade,
			},
			{
				Symbol:     "card_ancestors_child_id",
				Columns:    []*schema.Column{CardAncestorsColumns[1]},
				RefColumns: []*schema.Column{CardsColumns[0]},
				OnDelete:   schema.Cascade,
			},
		},
	}
	// ProductCardsColumns holds the columns for the "product_cards" table.
	ProductCardsColumns = []*schema.Column{
		{Name: "product_id", Type: field.TypeInt},
		{Name: "card_id", Type: field.TypeInt},
	}
	// ProductCardsTable holds the schema information for the "product_cards" table.
	ProductCardsTable = &schema.Table{
		Name:       "product_cards",
		Columns:    ProductCardsColumns,
		PrimaryKey: []*schema.Column{ProductCardsColumns[0], ProductCardsColumns[1]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "product_cards_product_id",
				Columns:    []*schema.Column{ProductCardsColumns[0]},
				RefColumns: []*schema.Column{ProductsColumns[0]},
				OnDelete:   schema.Cascade,
			},
			{
				Symbol:     "product_cards_card_id",
				Columns:    []*schema.Column{ProductCardsColumns[1]},
				RefColumns: []*schema.Column{CardsColumns[0]},
				OnDelete:   schema.Cascade,
			},
		},
	}
	// Tables holds all the tables in the schema.
	Tables = []*schema.Table{
		CardsTable,
		CardSpecialColorsTable,
		CardTypesTable,
		DecksTable,
		ProductsTable,
		RevisionsTable,
		CardAncestorsTable,
		ProductCardsTable,
	}
)

func init() {
	CardsTable.ForeignKeys[0].RefTable = CardSpecialColorsTable
	CardsTable.ForeignKeys[1].RefTable = CardTypesTable
	CardsTable.ForeignKeys[2].RefTable = DecksTable
	CardsTable.ForeignKeys[3].RefTable = RevisionsTable
	DecksTable.ForeignKeys[0].RefTable = RevisionsTable
	ProductsTable.ForeignKeys[0].RefTable = RevisionsTable
	CardAncestorsTable.ForeignKeys[0].RefTable = CardsTable
	CardAncestorsTable.ForeignKeys[1].RefTable = CardsTable
	ProductCardsTable.ForeignKeys[0].RefTable = ProductsTable
	ProductCardsTable.ForeignKeys[1].RefTable = CardsTable
}

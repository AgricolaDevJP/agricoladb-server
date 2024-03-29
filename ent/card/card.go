// Code generated by ent, DO NOT EDIT.

package card

import (
	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
)

const (
	// Label holds the string label denoting the card type in the database.
	Label = "card"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldLiteralID holds the string denoting the literal_id field in the database.
	FieldLiteralID = "literal_id"
	// FieldRevisionID holds the string denoting the revision_id field in the database.
	FieldRevisionID = "revision_id"
	// FieldPrintedID holds the string denoting the printed_id field in the database.
	FieldPrintedID = "printed_id"
	// FieldPlayAgricolaCardID holds the string denoting the play_agricola_card_id field in the database.
	FieldPlayAgricolaCardID = "play_agricola_card_id"
	// FieldDeckID holds the string denoting the deck_id field in the database.
	FieldDeckID = "deck_id"
	// FieldCardTypeID holds the string denoting the card_type_id field in the database.
	FieldCardTypeID = "card_type_id"
	// FieldCardSpecialColorID holds the string denoting the card_special_color_id field in the database.
	FieldCardSpecialColorID = "card_special_color_id"
	// FieldNameJa holds the string denoting the name_ja field in the database.
	FieldNameJa = "name_ja"
	// FieldNameEn holds the string denoting the name_en field in the database.
	FieldNameEn = "name_en"
	// FieldMinPlayersNumber holds the string denoting the min_players_number field in the database.
	FieldMinPlayersNumber = "min_players_number"
	// FieldPrerequisite holds the string denoting the prerequisite field in the database.
	FieldPrerequisite = "prerequisite"
	// FieldCost holds the string denoting the cost field in the database.
	FieldCost = "cost"
	// FieldDescription holds the string denoting the description field in the database.
	FieldDescription = "description"
	// FieldNote holds the string denoting the note field in the database.
	FieldNote = "note"
	// FieldIsOfficialJa holds the string denoting the is_official_ja field in the database.
	FieldIsOfficialJa = "is_official_ja"
	// FieldVictoryPoint holds the string denoting the victory_point field in the database.
	FieldVictoryPoint = "victory_point"
	// FieldSpecialVictoryPoint holds the string denoting the special_victory_point field in the database.
	FieldSpecialVictoryPoint = "special_victory_point"
	// FieldHasArrow holds the string denoting the has_arrow field in the database.
	FieldHasArrow = "has_arrow"
	// FieldHasBonusPointIcon holds the string denoting the has_bonus_point_icon field in the database.
	FieldHasBonusPointIcon = "has_bonus_point_icon"
	// FieldHasNegativeBonusPointIcon holds the string denoting the has_negative_bonus_point_icon field in the database.
	FieldHasNegativeBonusPointIcon = "has_negative_bonus_point_icon"
	// FieldHasPanIcon holds the string denoting the has_pan_icon field in the database.
	FieldHasPanIcon = "has_pan_icon"
	// FieldHasBreadIcon holds the string denoting the has_bread_icon field in the database.
	FieldHasBreadIcon = "has_bread_icon"
	// FieldHasFarmPlannerIcon holds the string denoting the has_farm_planner_icon field in the database.
	FieldHasFarmPlannerIcon = "has_farm_planner_icon"
	// FieldHasActionsBoosterIcon holds the string denoting the has_actions_booster_icon field in the database.
	FieldHasActionsBoosterIcon = "has_actions_booster_icon"
	// FieldHasPointsProviderIcon holds the string denoting the has_points_provider_icon field in the database.
	FieldHasPointsProviderIcon = "has_points_provider_icon"
	// FieldHasGoodsProviderIcon holds the string denoting the has_goods_provider_icon field in the database.
	FieldHasGoodsProviderIcon = "has_goods_provider_icon"
	// FieldHasFoodProviderIcon holds the string denoting the has_food_provider_icon field in the database.
	FieldHasFoodProviderIcon = "has_food_provider_icon"
	// FieldHasCropProviderIcon holds the string denoting the has_crop_provider_icon field in the database.
	FieldHasCropProviderIcon = "has_crop_provider_icon"
	// FieldHasBuildingResourceProviderIcon holds the string denoting the has_building_resource_provider_icon field in the database.
	FieldHasBuildingResourceProviderIcon = "has_building_resource_provider_icon"
	// FieldHasLivestockProviderIcon holds the string denoting the has_livestock_provider_icon field in the database.
	FieldHasLivestockProviderIcon = "has_livestock_provider_icon"
	// FieldHasCutPeatIcon holds the string denoting the has_cut_peat_icon field in the database.
	FieldHasCutPeatIcon = "has_cut_peat_icon"
	// FieldHasFellTreesIcon holds the string denoting the has_fell_trees_icon field in the database.
	FieldHasFellTreesIcon = "has_fell_trees_icon"
	// FieldHasSlashAndBurnIcon holds the string denoting the has_slash_and_burn_icon field in the database.
	FieldHasSlashAndBurnIcon = "has_slash_and_burn_icon"
	// FieldHasHiringFareIcon holds the string denoting the has_hiring_fare_icon field in the database.
	FieldHasHiringFareIcon = "has_hiring_fare_icon"
	// EdgeRevision holds the string denoting the revision edge name in mutations.
	EdgeRevision = "revision"
	// EdgeProducts holds the string denoting the products edge name in mutations.
	EdgeProducts = "products"
	// EdgeDeck holds the string denoting the deck edge name in mutations.
	EdgeDeck = "deck"
	// EdgeCardType holds the string denoting the card_type edge name in mutations.
	EdgeCardType = "card_type"
	// EdgeCardSpecialColor holds the string denoting the card_special_color edge name in mutations.
	EdgeCardSpecialColor = "card_special_color"
	// EdgeChildren holds the string denoting the children edge name in mutations.
	EdgeChildren = "children"
	// EdgeAncestors holds the string denoting the ancestors edge name in mutations.
	EdgeAncestors = "ancestors"
	// Table holds the table name of the card in the database.
	Table = "cards"
	// RevisionTable is the table that holds the revision relation/edge.
	RevisionTable = "cards"
	// RevisionInverseTable is the table name for the Revision entity.
	// It exists in this package in order to avoid circular dependency with the "revision" package.
	RevisionInverseTable = "revisions"
	// RevisionColumn is the table column denoting the revision relation/edge.
	RevisionColumn = "revision_id"
	// ProductsTable is the table that holds the products relation/edge. The primary key declared below.
	ProductsTable = "product_cards"
	// ProductsInverseTable is the table name for the Product entity.
	// It exists in this package in order to avoid circular dependency with the "product" package.
	ProductsInverseTable = "products"
	// DeckTable is the table that holds the deck relation/edge.
	DeckTable = "cards"
	// DeckInverseTable is the table name for the Deck entity.
	// It exists in this package in order to avoid circular dependency with the "deck" package.
	DeckInverseTable = "decks"
	// DeckColumn is the table column denoting the deck relation/edge.
	DeckColumn = "deck_id"
	// CardTypeTable is the table that holds the card_type relation/edge.
	CardTypeTable = "cards"
	// CardTypeInverseTable is the table name for the CardType entity.
	// It exists in this package in order to avoid circular dependency with the "cardtype" package.
	CardTypeInverseTable = "card_types"
	// CardTypeColumn is the table column denoting the card_type relation/edge.
	CardTypeColumn = "card_type_id"
	// CardSpecialColorTable is the table that holds the card_special_color relation/edge.
	CardSpecialColorTable = "cards"
	// CardSpecialColorInverseTable is the table name for the CardSpecialColor entity.
	// It exists in this package in order to avoid circular dependency with the "cardspecialcolor" package.
	CardSpecialColorInverseTable = "card_special_colors"
	// CardSpecialColorColumn is the table column denoting the card_special_color relation/edge.
	CardSpecialColorColumn = "card_special_color_id"
	// ChildrenTable is the table that holds the children relation/edge. The primary key declared below.
	ChildrenTable = "card_ancestors"
	// AncestorsTable is the table that holds the ancestors relation/edge. The primary key declared below.
	AncestorsTable = "card_ancestors"
)

// Columns holds all SQL columns for card fields.
var Columns = []string{
	FieldID,
	FieldLiteralID,
	FieldRevisionID,
	FieldPrintedID,
	FieldPlayAgricolaCardID,
	FieldDeckID,
	FieldCardTypeID,
	FieldCardSpecialColorID,
	FieldNameJa,
	FieldNameEn,
	FieldMinPlayersNumber,
	FieldPrerequisite,
	FieldCost,
	FieldDescription,
	FieldNote,
	FieldIsOfficialJa,
	FieldVictoryPoint,
	FieldSpecialVictoryPoint,
	FieldHasArrow,
	FieldHasBonusPointIcon,
	FieldHasNegativeBonusPointIcon,
	FieldHasPanIcon,
	FieldHasBreadIcon,
	FieldHasFarmPlannerIcon,
	FieldHasActionsBoosterIcon,
	FieldHasPointsProviderIcon,
	FieldHasGoodsProviderIcon,
	FieldHasFoodProviderIcon,
	FieldHasCropProviderIcon,
	FieldHasBuildingResourceProviderIcon,
	FieldHasLivestockProviderIcon,
	FieldHasCutPeatIcon,
	FieldHasFellTreesIcon,
	FieldHasSlashAndBurnIcon,
	FieldHasHiringFareIcon,
}

var (
	// ProductsPrimaryKey and ProductsColumn2 are the table columns denoting the
	// primary key for the products relation (M2M).
	ProductsPrimaryKey = []string{"product_id", "card_id"}
	// ChildrenPrimaryKey and ChildrenColumn2 are the table columns denoting the
	// primary key for the children relation (M2M).
	ChildrenPrimaryKey = []string{"card_id", "child_id"}
	// AncestorsPrimaryKey and AncestorsColumn2 are the table columns denoting the
	// primary key for the ancestors relation (M2M).
	AncestorsPrimaryKey = []string{"card_id", "child_id"}
)

// ValidColumn reports if the column name is valid (part of the table columns).
func ValidColumn(column string) bool {
	for i := range Columns {
		if column == Columns[i] {
			return true
		}
	}
	return false
}

var (
	// LiteralIDValidator is a validator for the "literal_id" field. It is called by the builders before save.
	LiteralIDValidator func(string) error
)

// OrderOption defines the ordering options for the Card queries.
type OrderOption func(*sql.Selector)

// ByID orders the results by the id field.
func ByID(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldID, opts...).ToFunc()
}

// ByLiteralID orders the results by the literal_id field.
func ByLiteralID(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldLiteralID, opts...).ToFunc()
}

// ByRevisionID orders the results by the revision_id field.
func ByRevisionID(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldRevisionID, opts...).ToFunc()
}

// ByPrintedID orders the results by the printed_id field.
func ByPrintedID(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldPrintedID, opts...).ToFunc()
}

// ByPlayAgricolaCardID orders the results by the play_agricola_card_id field.
func ByPlayAgricolaCardID(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldPlayAgricolaCardID, opts...).ToFunc()
}

// ByDeckID orders the results by the deck_id field.
func ByDeckID(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldDeckID, opts...).ToFunc()
}

// ByCardTypeID orders the results by the card_type_id field.
func ByCardTypeID(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldCardTypeID, opts...).ToFunc()
}

// ByCardSpecialColorID orders the results by the card_special_color_id field.
func ByCardSpecialColorID(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldCardSpecialColorID, opts...).ToFunc()
}

// ByNameJa orders the results by the name_ja field.
func ByNameJa(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldNameJa, opts...).ToFunc()
}

// ByNameEn orders the results by the name_en field.
func ByNameEn(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldNameEn, opts...).ToFunc()
}

// ByMinPlayersNumber orders the results by the min_players_number field.
func ByMinPlayersNumber(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldMinPlayersNumber, opts...).ToFunc()
}

// ByPrerequisite orders the results by the prerequisite field.
func ByPrerequisite(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldPrerequisite, opts...).ToFunc()
}

// ByCost orders the results by the cost field.
func ByCost(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldCost, opts...).ToFunc()
}

// ByDescription orders the results by the description field.
func ByDescription(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldDescription, opts...).ToFunc()
}

// ByNote orders the results by the note field.
func ByNote(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldNote, opts...).ToFunc()
}

// ByIsOfficialJa orders the results by the is_official_ja field.
func ByIsOfficialJa(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldIsOfficialJa, opts...).ToFunc()
}

// ByVictoryPoint orders the results by the victory_point field.
func ByVictoryPoint(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldVictoryPoint, opts...).ToFunc()
}

// BySpecialVictoryPoint orders the results by the special_victory_point field.
func BySpecialVictoryPoint(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldSpecialVictoryPoint, opts...).ToFunc()
}

// ByHasArrow orders the results by the has_arrow field.
func ByHasArrow(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldHasArrow, opts...).ToFunc()
}

// ByHasBonusPointIcon orders the results by the has_bonus_point_icon field.
func ByHasBonusPointIcon(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldHasBonusPointIcon, opts...).ToFunc()
}

// ByHasNegativeBonusPointIcon orders the results by the has_negative_bonus_point_icon field.
func ByHasNegativeBonusPointIcon(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldHasNegativeBonusPointIcon, opts...).ToFunc()
}

// ByHasPanIcon orders the results by the has_pan_icon field.
func ByHasPanIcon(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldHasPanIcon, opts...).ToFunc()
}

// ByHasBreadIcon orders the results by the has_bread_icon field.
func ByHasBreadIcon(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldHasBreadIcon, opts...).ToFunc()
}

// ByHasFarmPlannerIcon orders the results by the has_farm_planner_icon field.
func ByHasFarmPlannerIcon(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldHasFarmPlannerIcon, opts...).ToFunc()
}

// ByHasActionsBoosterIcon orders the results by the has_actions_booster_icon field.
func ByHasActionsBoosterIcon(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldHasActionsBoosterIcon, opts...).ToFunc()
}

// ByHasPointsProviderIcon orders the results by the has_points_provider_icon field.
func ByHasPointsProviderIcon(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldHasPointsProviderIcon, opts...).ToFunc()
}

// ByHasGoodsProviderIcon orders the results by the has_goods_provider_icon field.
func ByHasGoodsProviderIcon(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldHasGoodsProviderIcon, opts...).ToFunc()
}

// ByHasFoodProviderIcon orders the results by the has_food_provider_icon field.
func ByHasFoodProviderIcon(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldHasFoodProviderIcon, opts...).ToFunc()
}

// ByHasCropProviderIcon orders the results by the has_crop_provider_icon field.
func ByHasCropProviderIcon(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldHasCropProviderIcon, opts...).ToFunc()
}

// ByHasBuildingResourceProviderIcon orders the results by the has_building_resource_provider_icon field.
func ByHasBuildingResourceProviderIcon(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldHasBuildingResourceProviderIcon, opts...).ToFunc()
}

// ByHasLivestockProviderIcon orders the results by the has_livestock_provider_icon field.
func ByHasLivestockProviderIcon(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldHasLivestockProviderIcon, opts...).ToFunc()
}

// ByHasCutPeatIcon orders the results by the has_cut_peat_icon field.
func ByHasCutPeatIcon(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldHasCutPeatIcon, opts...).ToFunc()
}

// ByHasFellTreesIcon orders the results by the has_fell_trees_icon field.
func ByHasFellTreesIcon(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldHasFellTreesIcon, opts...).ToFunc()
}

// ByHasSlashAndBurnIcon orders the results by the has_slash_and_burn_icon field.
func ByHasSlashAndBurnIcon(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldHasSlashAndBurnIcon, opts...).ToFunc()
}

// ByHasHiringFareIcon orders the results by the has_hiring_fare_icon field.
func ByHasHiringFareIcon(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldHasHiringFareIcon, opts...).ToFunc()
}

// ByRevisionField orders the results by revision field.
func ByRevisionField(field string, opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newRevisionStep(), sql.OrderByField(field, opts...))
	}
}

// ByProductsCount orders the results by products count.
func ByProductsCount(opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborsCount(s, newProductsStep(), opts...)
	}
}

// ByProducts orders the results by products terms.
func ByProducts(term sql.OrderTerm, terms ...sql.OrderTerm) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newProductsStep(), append([]sql.OrderTerm{term}, terms...)...)
	}
}

// ByDeckField orders the results by deck field.
func ByDeckField(field string, opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newDeckStep(), sql.OrderByField(field, opts...))
	}
}

// ByCardTypeField orders the results by card_type field.
func ByCardTypeField(field string, opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newCardTypeStep(), sql.OrderByField(field, opts...))
	}
}

// ByCardSpecialColorField orders the results by card_special_color field.
func ByCardSpecialColorField(field string, opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newCardSpecialColorStep(), sql.OrderByField(field, opts...))
	}
}

// ByChildrenCount orders the results by children count.
func ByChildrenCount(opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborsCount(s, newChildrenStep(), opts...)
	}
}

// ByChildren orders the results by children terms.
func ByChildren(term sql.OrderTerm, terms ...sql.OrderTerm) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newChildrenStep(), append([]sql.OrderTerm{term}, terms...)...)
	}
}

// ByAncestorsCount orders the results by ancestors count.
func ByAncestorsCount(opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborsCount(s, newAncestorsStep(), opts...)
	}
}

// ByAncestors orders the results by ancestors terms.
func ByAncestors(term sql.OrderTerm, terms ...sql.OrderTerm) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newAncestorsStep(), append([]sql.OrderTerm{term}, terms...)...)
	}
}
func newRevisionStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(RevisionInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.M2O, true, RevisionTable, RevisionColumn),
	)
}
func newProductsStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(ProductsInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.M2M, true, ProductsTable, ProductsPrimaryKey...),
	)
}
func newDeckStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(DeckInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.M2O, true, DeckTable, DeckColumn),
	)
}
func newCardTypeStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(CardTypeInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.M2O, true, CardTypeTable, CardTypeColumn),
	)
}
func newCardSpecialColorStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(CardSpecialColorInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.M2O, true, CardSpecialColorTable, CardSpecialColorColumn),
	)
}
func newChildrenStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(Table, FieldID),
		sqlgraph.Edge(sqlgraph.M2M, true, ChildrenTable, ChildrenPrimaryKey...),
	)
}
func newAncestorsStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(Table, FieldID),
		sqlgraph.Edge(sqlgraph.M2M, false, AncestorsTable, AncestorsPrimaryKey...),
	)
}

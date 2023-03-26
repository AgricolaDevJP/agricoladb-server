// Code generated by ent, DO NOT EDIT.

package revision

const (
	// Label holds the string label denoting the revision type in the database.
	Label = "revision"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldKey holds the string denoting the key field in the database.
	FieldKey = "key"
	// FieldNameJa holds the string denoting the name_ja field in the database.
	FieldNameJa = "name_ja"
	// FieldNameEn holds the string denoting the name_en field in the database.
	FieldNameEn = "name_en"
	// EdgeCards holds the string denoting the cards edge name in mutations.
	EdgeCards = "cards"
	// EdgeProducts holds the string denoting the products edge name in mutations.
	EdgeProducts = "products"
	// EdgeDecks holds the string denoting the decks edge name in mutations.
	EdgeDecks = "decks"
	// Table holds the table name of the revision in the database.
	Table = "revisions"
	// CardsTable is the table that holds the cards relation/edge.
	CardsTable = "cards"
	// CardsInverseTable is the table name for the Card entity.
	// It exists in this package in order to avoid circular dependency with the "card" package.
	CardsInverseTable = "cards"
	// CardsColumn is the table column denoting the cards relation/edge.
	CardsColumn = "revision_id"
	// ProductsTable is the table that holds the products relation/edge.
	ProductsTable = "products"
	// ProductsInverseTable is the table name for the Product entity.
	// It exists in this package in order to avoid circular dependency with the "product" package.
	ProductsInverseTable = "products"
	// ProductsColumn is the table column denoting the products relation/edge.
	ProductsColumn = "revision_id"
	// DecksTable is the table that holds the decks relation/edge.
	DecksTable = "decks"
	// DecksInverseTable is the table name for the Deck entity.
	// It exists in this package in order to avoid circular dependency with the "deck" package.
	DecksInverseTable = "decks"
	// DecksColumn is the table column denoting the decks relation/edge.
	DecksColumn = "revision_id"
)

// Columns holds all SQL columns for revision fields.
var Columns = []string{
	FieldID,
	FieldKey,
	FieldNameJa,
	FieldNameEn,
}

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
	// KeyValidator is a validator for the "key" field. It is called by the builders before save.
	KeyValidator func(string) error
)
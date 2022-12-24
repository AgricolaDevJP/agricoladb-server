// Code generated by ent, DO NOT EDIT.

package product

const (
	// Label holds the string label denoting the product type in the database.
	Label = "product"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldRevisionID holds the string denoting the revision_id field in the database.
	FieldRevisionID = "revision_id"
	// FieldIsOfficialJa holds the string denoting the is_official_ja field in the database.
	FieldIsOfficialJa = "is_official_ja"
	// FieldNameJa holds the string denoting the name_ja field in the database.
	FieldNameJa = "name_ja"
	// FieldNameEn holds the string denoting the name_en field in the database.
	FieldNameEn = "name_en"
	// EdgeCards holds the string denoting the cards edge name in mutations.
	EdgeCards = "cards"
	// EdgeRevision holds the string denoting the revision edge name in mutations.
	EdgeRevision = "revision"
	// Table holds the table name of the product in the database.
	Table = "products"
	// CardsTable is the table that holds the cards relation/edge. The primary key declared below.
	CardsTable = "product_cards"
	// CardsInverseTable is the table name for the Card entity.
	// It exists in this package in order to avoid circular dependency with the "card" package.
	CardsInverseTable = "cards"
	// RevisionTable is the table that holds the revision relation/edge.
	RevisionTable = "products"
	// RevisionInverseTable is the table name for the Revision entity.
	// It exists in this package in order to avoid circular dependency with the "revision" package.
	RevisionInverseTable = "revisions"
	// RevisionColumn is the table column denoting the revision relation/edge.
	RevisionColumn = "revision_id"
)

// Columns holds all SQL columns for product fields.
var Columns = []string{
	FieldID,
	FieldRevisionID,
	FieldIsOfficialJa,
	FieldNameJa,
	FieldNameEn,
}

var (
	// CardsPrimaryKey and CardsColumn2 are the table columns denoting the
	// primary key for the cards relation (M2M).
	CardsPrimaryKey = []string{"product_id", "card_id"}
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

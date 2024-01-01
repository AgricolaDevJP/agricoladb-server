// Code generated by ent, DO NOT EDIT.

package product

import (
	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"github.com/AgricolaDevJP/agricoladb-server/ent/predicate"
)

// ID filters vertices based on their ID field.
func ID(id int) predicate.Product {
	return predicate.Product(sql.FieldEQ(FieldID, id))
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id int) predicate.Product {
	return predicate.Product(sql.FieldEQ(FieldID, id))
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id int) predicate.Product {
	return predicate.Product(sql.FieldNEQ(FieldID, id))
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...int) predicate.Product {
	return predicate.Product(sql.FieldIn(FieldID, ids...))
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...int) predicate.Product {
	return predicate.Product(sql.FieldNotIn(FieldID, ids...))
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id int) predicate.Product {
	return predicate.Product(sql.FieldGT(FieldID, id))
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id int) predicate.Product {
	return predicate.Product(sql.FieldGTE(FieldID, id))
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id int) predicate.Product {
	return predicate.Product(sql.FieldLT(FieldID, id))
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id int) predicate.Product {
	return predicate.Product(sql.FieldLTE(FieldID, id))
}

// RevisionID applies equality check predicate on the "revision_id" field. It's identical to RevisionIDEQ.
func RevisionID(v int) predicate.Product {
	return predicate.Product(sql.FieldEQ(FieldRevisionID, v))
}

// IsOfficialJa applies equality check predicate on the "is_official_ja" field. It's identical to IsOfficialJaEQ.
func IsOfficialJa(v bool) predicate.Product {
	return predicate.Product(sql.FieldEQ(FieldIsOfficialJa, v))
}

// NameJa applies equality check predicate on the "name_ja" field. It's identical to NameJaEQ.
func NameJa(v string) predicate.Product {
	return predicate.Product(sql.FieldEQ(FieldNameJa, v))
}

// NameEn applies equality check predicate on the "name_en" field. It's identical to NameEnEQ.
func NameEn(v string) predicate.Product {
	return predicate.Product(sql.FieldEQ(FieldNameEn, v))
}

// PublishedYear applies equality check predicate on the "published_year" field. It's identical to PublishedYearEQ.
func PublishedYear(v int) predicate.Product {
	return predicate.Product(sql.FieldEQ(FieldPublishedYear, v))
}

// RevisionIDEQ applies the EQ predicate on the "revision_id" field.
func RevisionIDEQ(v int) predicate.Product {
	return predicate.Product(sql.FieldEQ(FieldRevisionID, v))
}

// RevisionIDNEQ applies the NEQ predicate on the "revision_id" field.
func RevisionIDNEQ(v int) predicate.Product {
	return predicate.Product(sql.FieldNEQ(FieldRevisionID, v))
}

// RevisionIDIn applies the In predicate on the "revision_id" field.
func RevisionIDIn(vs ...int) predicate.Product {
	return predicate.Product(sql.FieldIn(FieldRevisionID, vs...))
}

// RevisionIDNotIn applies the NotIn predicate on the "revision_id" field.
func RevisionIDNotIn(vs ...int) predicate.Product {
	return predicate.Product(sql.FieldNotIn(FieldRevisionID, vs...))
}

// IsOfficialJaEQ applies the EQ predicate on the "is_official_ja" field.
func IsOfficialJaEQ(v bool) predicate.Product {
	return predicate.Product(sql.FieldEQ(FieldIsOfficialJa, v))
}

// IsOfficialJaNEQ applies the NEQ predicate on the "is_official_ja" field.
func IsOfficialJaNEQ(v bool) predicate.Product {
	return predicate.Product(sql.FieldNEQ(FieldIsOfficialJa, v))
}

// NameJaEQ applies the EQ predicate on the "name_ja" field.
func NameJaEQ(v string) predicate.Product {
	return predicate.Product(sql.FieldEQ(FieldNameJa, v))
}

// NameJaNEQ applies the NEQ predicate on the "name_ja" field.
func NameJaNEQ(v string) predicate.Product {
	return predicate.Product(sql.FieldNEQ(FieldNameJa, v))
}

// NameJaIn applies the In predicate on the "name_ja" field.
func NameJaIn(vs ...string) predicate.Product {
	return predicate.Product(sql.FieldIn(FieldNameJa, vs...))
}

// NameJaNotIn applies the NotIn predicate on the "name_ja" field.
func NameJaNotIn(vs ...string) predicate.Product {
	return predicate.Product(sql.FieldNotIn(FieldNameJa, vs...))
}

// NameJaGT applies the GT predicate on the "name_ja" field.
func NameJaGT(v string) predicate.Product {
	return predicate.Product(sql.FieldGT(FieldNameJa, v))
}

// NameJaGTE applies the GTE predicate on the "name_ja" field.
func NameJaGTE(v string) predicate.Product {
	return predicate.Product(sql.FieldGTE(FieldNameJa, v))
}

// NameJaLT applies the LT predicate on the "name_ja" field.
func NameJaLT(v string) predicate.Product {
	return predicate.Product(sql.FieldLT(FieldNameJa, v))
}

// NameJaLTE applies the LTE predicate on the "name_ja" field.
func NameJaLTE(v string) predicate.Product {
	return predicate.Product(sql.FieldLTE(FieldNameJa, v))
}

// NameJaContains applies the Contains predicate on the "name_ja" field.
func NameJaContains(v string) predicate.Product {
	return predicate.Product(sql.FieldContains(FieldNameJa, v))
}

// NameJaHasPrefix applies the HasPrefix predicate on the "name_ja" field.
func NameJaHasPrefix(v string) predicate.Product {
	return predicate.Product(sql.FieldHasPrefix(FieldNameJa, v))
}

// NameJaHasSuffix applies the HasSuffix predicate on the "name_ja" field.
func NameJaHasSuffix(v string) predicate.Product {
	return predicate.Product(sql.FieldHasSuffix(FieldNameJa, v))
}

// NameJaIsNil applies the IsNil predicate on the "name_ja" field.
func NameJaIsNil() predicate.Product {
	return predicate.Product(sql.FieldIsNull(FieldNameJa))
}

// NameJaNotNil applies the NotNil predicate on the "name_ja" field.
func NameJaNotNil() predicate.Product {
	return predicate.Product(sql.FieldNotNull(FieldNameJa))
}

// NameJaEqualFold applies the EqualFold predicate on the "name_ja" field.
func NameJaEqualFold(v string) predicate.Product {
	return predicate.Product(sql.FieldEqualFold(FieldNameJa, v))
}

// NameJaContainsFold applies the ContainsFold predicate on the "name_ja" field.
func NameJaContainsFold(v string) predicate.Product {
	return predicate.Product(sql.FieldContainsFold(FieldNameJa, v))
}

// NameEnEQ applies the EQ predicate on the "name_en" field.
func NameEnEQ(v string) predicate.Product {
	return predicate.Product(sql.FieldEQ(FieldNameEn, v))
}

// NameEnNEQ applies the NEQ predicate on the "name_en" field.
func NameEnNEQ(v string) predicate.Product {
	return predicate.Product(sql.FieldNEQ(FieldNameEn, v))
}

// NameEnIn applies the In predicate on the "name_en" field.
func NameEnIn(vs ...string) predicate.Product {
	return predicate.Product(sql.FieldIn(FieldNameEn, vs...))
}

// NameEnNotIn applies the NotIn predicate on the "name_en" field.
func NameEnNotIn(vs ...string) predicate.Product {
	return predicate.Product(sql.FieldNotIn(FieldNameEn, vs...))
}

// NameEnGT applies the GT predicate on the "name_en" field.
func NameEnGT(v string) predicate.Product {
	return predicate.Product(sql.FieldGT(FieldNameEn, v))
}

// NameEnGTE applies the GTE predicate on the "name_en" field.
func NameEnGTE(v string) predicate.Product {
	return predicate.Product(sql.FieldGTE(FieldNameEn, v))
}

// NameEnLT applies the LT predicate on the "name_en" field.
func NameEnLT(v string) predicate.Product {
	return predicate.Product(sql.FieldLT(FieldNameEn, v))
}

// NameEnLTE applies the LTE predicate on the "name_en" field.
func NameEnLTE(v string) predicate.Product {
	return predicate.Product(sql.FieldLTE(FieldNameEn, v))
}

// NameEnContains applies the Contains predicate on the "name_en" field.
func NameEnContains(v string) predicate.Product {
	return predicate.Product(sql.FieldContains(FieldNameEn, v))
}

// NameEnHasPrefix applies the HasPrefix predicate on the "name_en" field.
func NameEnHasPrefix(v string) predicate.Product {
	return predicate.Product(sql.FieldHasPrefix(FieldNameEn, v))
}

// NameEnHasSuffix applies the HasSuffix predicate on the "name_en" field.
func NameEnHasSuffix(v string) predicate.Product {
	return predicate.Product(sql.FieldHasSuffix(FieldNameEn, v))
}

// NameEnIsNil applies the IsNil predicate on the "name_en" field.
func NameEnIsNil() predicate.Product {
	return predicate.Product(sql.FieldIsNull(FieldNameEn))
}

// NameEnNotNil applies the NotNil predicate on the "name_en" field.
func NameEnNotNil() predicate.Product {
	return predicate.Product(sql.FieldNotNull(FieldNameEn))
}

// NameEnEqualFold applies the EqualFold predicate on the "name_en" field.
func NameEnEqualFold(v string) predicate.Product {
	return predicate.Product(sql.FieldEqualFold(FieldNameEn, v))
}

// NameEnContainsFold applies the ContainsFold predicate on the "name_en" field.
func NameEnContainsFold(v string) predicate.Product {
	return predicate.Product(sql.FieldContainsFold(FieldNameEn, v))
}

// PublishedYearEQ applies the EQ predicate on the "published_year" field.
func PublishedYearEQ(v int) predicate.Product {
	return predicate.Product(sql.FieldEQ(FieldPublishedYear, v))
}

// PublishedYearNEQ applies the NEQ predicate on the "published_year" field.
func PublishedYearNEQ(v int) predicate.Product {
	return predicate.Product(sql.FieldNEQ(FieldPublishedYear, v))
}

// PublishedYearIn applies the In predicate on the "published_year" field.
func PublishedYearIn(vs ...int) predicate.Product {
	return predicate.Product(sql.FieldIn(FieldPublishedYear, vs...))
}

// PublishedYearNotIn applies the NotIn predicate on the "published_year" field.
func PublishedYearNotIn(vs ...int) predicate.Product {
	return predicate.Product(sql.FieldNotIn(FieldPublishedYear, vs...))
}

// PublishedYearGT applies the GT predicate on the "published_year" field.
func PublishedYearGT(v int) predicate.Product {
	return predicate.Product(sql.FieldGT(FieldPublishedYear, v))
}

// PublishedYearGTE applies the GTE predicate on the "published_year" field.
func PublishedYearGTE(v int) predicate.Product {
	return predicate.Product(sql.FieldGTE(FieldPublishedYear, v))
}

// PublishedYearLT applies the LT predicate on the "published_year" field.
func PublishedYearLT(v int) predicate.Product {
	return predicate.Product(sql.FieldLT(FieldPublishedYear, v))
}

// PublishedYearLTE applies the LTE predicate on the "published_year" field.
func PublishedYearLTE(v int) predicate.Product {
	return predicate.Product(sql.FieldLTE(FieldPublishedYear, v))
}

// PublishedYearIsNil applies the IsNil predicate on the "published_year" field.
func PublishedYearIsNil() predicate.Product {
	return predicate.Product(sql.FieldIsNull(FieldPublishedYear))
}

// PublishedYearNotNil applies the NotNil predicate on the "published_year" field.
func PublishedYearNotNil() predicate.Product {
	return predicate.Product(sql.FieldNotNull(FieldPublishedYear))
}

// HasCards applies the HasEdge predicate on the "cards" edge.
func HasCards() predicate.Product {
	return predicate.Product(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.M2M, false, CardsTable, CardsPrimaryKey...),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasCardsWith applies the HasEdge predicate on the "cards" edge with a given conditions (other predicates).
func HasCardsWith(preds ...predicate.Card) predicate.Product {
	return predicate.Product(func(s *sql.Selector) {
		step := newCardsStep()
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// HasRevision applies the HasEdge predicate on the "revision" edge.
func HasRevision() predicate.Product {
	return predicate.Product(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, RevisionTable, RevisionColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasRevisionWith applies the HasEdge predicate on the "revision" edge with a given conditions (other predicates).
func HasRevisionWith(preds ...predicate.Revision) predicate.Product {
	return predicate.Product(func(s *sql.Selector) {
		step := newRevisionStep()
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.Product) predicate.Product {
	return predicate.Product(sql.AndPredicates(predicates...))
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.Product) predicate.Product {
	return predicate.Product(sql.OrPredicates(predicates...))
}

// Not applies the not operator on the given predicate.
func Not(p predicate.Product) predicate.Product {
	return predicate.Product(sql.NotPredicates(p))
}

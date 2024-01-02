// Code generated by ent, DO NOT EDIT.

package cardtype

import (
	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"github.com/AgricolaDevJP/agricoladb-server/ent/predicate"
)

// ID filters vertices based on their ID field.
func ID(id int) predicate.CardType {
	return predicate.CardType(sql.FieldEQ(FieldID, id))
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id int) predicate.CardType {
	return predicate.CardType(sql.FieldEQ(FieldID, id))
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id int) predicate.CardType {
	return predicate.CardType(sql.FieldNEQ(FieldID, id))
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...int) predicate.CardType {
	return predicate.CardType(sql.FieldIn(FieldID, ids...))
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...int) predicate.CardType {
	return predicate.CardType(sql.FieldNotIn(FieldID, ids...))
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id int) predicate.CardType {
	return predicate.CardType(sql.FieldGT(FieldID, id))
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id int) predicate.CardType {
	return predicate.CardType(sql.FieldGTE(FieldID, id))
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id int) predicate.CardType {
	return predicate.CardType(sql.FieldLT(FieldID, id))
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id int) predicate.CardType {
	return predicate.CardType(sql.FieldLTE(FieldID, id))
}

// Key applies equality check predicate on the "key" field. It's identical to KeyEQ.
func Key(v string) predicate.CardType {
	return predicate.CardType(sql.FieldEQ(FieldKey, v))
}

// NameJa applies equality check predicate on the "name_ja" field. It's identical to NameJaEQ.
func NameJa(v string) predicate.CardType {
	return predicate.CardType(sql.FieldEQ(FieldNameJa, v))
}

// NameEn applies equality check predicate on the "name_en" field. It's identical to NameEnEQ.
func NameEn(v string) predicate.CardType {
	return predicate.CardType(sql.FieldEQ(FieldNameEn, v))
}

// KeyEQ applies the EQ predicate on the "key" field.
func KeyEQ(v string) predicate.CardType {
	return predicate.CardType(sql.FieldEQ(FieldKey, v))
}

// KeyNEQ applies the NEQ predicate on the "key" field.
func KeyNEQ(v string) predicate.CardType {
	return predicate.CardType(sql.FieldNEQ(FieldKey, v))
}

// KeyIn applies the In predicate on the "key" field.
func KeyIn(vs ...string) predicate.CardType {
	return predicate.CardType(sql.FieldIn(FieldKey, vs...))
}

// KeyNotIn applies the NotIn predicate on the "key" field.
func KeyNotIn(vs ...string) predicate.CardType {
	return predicate.CardType(sql.FieldNotIn(FieldKey, vs...))
}

// KeyGT applies the GT predicate on the "key" field.
func KeyGT(v string) predicate.CardType {
	return predicate.CardType(sql.FieldGT(FieldKey, v))
}

// KeyGTE applies the GTE predicate on the "key" field.
func KeyGTE(v string) predicate.CardType {
	return predicate.CardType(sql.FieldGTE(FieldKey, v))
}

// KeyLT applies the LT predicate on the "key" field.
func KeyLT(v string) predicate.CardType {
	return predicate.CardType(sql.FieldLT(FieldKey, v))
}

// KeyLTE applies the LTE predicate on the "key" field.
func KeyLTE(v string) predicate.CardType {
	return predicate.CardType(sql.FieldLTE(FieldKey, v))
}

// KeyContains applies the Contains predicate on the "key" field.
func KeyContains(v string) predicate.CardType {
	return predicate.CardType(sql.FieldContains(FieldKey, v))
}

// KeyHasPrefix applies the HasPrefix predicate on the "key" field.
func KeyHasPrefix(v string) predicate.CardType {
	return predicate.CardType(sql.FieldHasPrefix(FieldKey, v))
}

// KeyHasSuffix applies the HasSuffix predicate on the "key" field.
func KeyHasSuffix(v string) predicate.CardType {
	return predicate.CardType(sql.FieldHasSuffix(FieldKey, v))
}

// KeyEqualFold applies the EqualFold predicate on the "key" field.
func KeyEqualFold(v string) predicate.CardType {
	return predicate.CardType(sql.FieldEqualFold(FieldKey, v))
}

// KeyContainsFold applies the ContainsFold predicate on the "key" field.
func KeyContainsFold(v string) predicate.CardType {
	return predicate.CardType(sql.FieldContainsFold(FieldKey, v))
}

// NameJaEQ applies the EQ predicate on the "name_ja" field.
func NameJaEQ(v string) predicate.CardType {
	return predicate.CardType(sql.FieldEQ(FieldNameJa, v))
}

// NameJaNEQ applies the NEQ predicate on the "name_ja" field.
func NameJaNEQ(v string) predicate.CardType {
	return predicate.CardType(sql.FieldNEQ(FieldNameJa, v))
}

// NameJaIn applies the In predicate on the "name_ja" field.
func NameJaIn(vs ...string) predicate.CardType {
	return predicate.CardType(sql.FieldIn(FieldNameJa, vs...))
}

// NameJaNotIn applies the NotIn predicate on the "name_ja" field.
func NameJaNotIn(vs ...string) predicate.CardType {
	return predicate.CardType(sql.FieldNotIn(FieldNameJa, vs...))
}

// NameJaGT applies the GT predicate on the "name_ja" field.
func NameJaGT(v string) predicate.CardType {
	return predicate.CardType(sql.FieldGT(FieldNameJa, v))
}

// NameJaGTE applies the GTE predicate on the "name_ja" field.
func NameJaGTE(v string) predicate.CardType {
	return predicate.CardType(sql.FieldGTE(FieldNameJa, v))
}

// NameJaLT applies the LT predicate on the "name_ja" field.
func NameJaLT(v string) predicate.CardType {
	return predicate.CardType(sql.FieldLT(FieldNameJa, v))
}

// NameJaLTE applies the LTE predicate on the "name_ja" field.
func NameJaLTE(v string) predicate.CardType {
	return predicate.CardType(sql.FieldLTE(FieldNameJa, v))
}

// NameJaContains applies the Contains predicate on the "name_ja" field.
func NameJaContains(v string) predicate.CardType {
	return predicate.CardType(sql.FieldContains(FieldNameJa, v))
}

// NameJaHasPrefix applies the HasPrefix predicate on the "name_ja" field.
func NameJaHasPrefix(v string) predicate.CardType {
	return predicate.CardType(sql.FieldHasPrefix(FieldNameJa, v))
}

// NameJaHasSuffix applies the HasSuffix predicate on the "name_ja" field.
func NameJaHasSuffix(v string) predicate.CardType {
	return predicate.CardType(sql.FieldHasSuffix(FieldNameJa, v))
}

// NameJaIsNil applies the IsNil predicate on the "name_ja" field.
func NameJaIsNil() predicate.CardType {
	return predicate.CardType(sql.FieldIsNull(FieldNameJa))
}

// NameJaNotNil applies the NotNil predicate on the "name_ja" field.
func NameJaNotNil() predicate.CardType {
	return predicate.CardType(sql.FieldNotNull(FieldNameJa))
}

// NameJaEqualFold applies the EqualFold predicate on the "name_ja" field.
func NameJaEqualFold(v string) predicate.CardType {
	return predicate.CardType(sql.FieldEqualFold(FieldNameJa, v))
}

// NameJaContainsFold applies the ContainsFold predicate on the "name_ja" field.
func NameJaContainsFold(v string) predicate.CardType {
	return predicate.CardType(sql.FieldContainsFold(FieldNameJa, v))
}

// NameEnEQ applies the EQ predicate on the "name_en" field.
func NameEnEQ(v string) predicate.CardType {
	return predicate.CardType(sql.FieldEQ(FieldNameEn, v))
}

// NameEnNEQ applies the NEQ predicate on the "name_en" field.
func NameEnNEQ(v string) predicate.CardType {
	return predicate.CardType(sql.FieldNEQ(FieldNameEn, v))
}

// NameEnIn applies the In predicate on the "name_en" field.
func NameEnIn(vs ...string) predicate.CardType {
	return predicate.CardType(sql.FieldIn(FieldNameEn, vs...))
}

// NameEnNotIn applies the NotIn predicate on the "name_en" field.
func NameEnNotIn(vs ...string) predicate.CardType {
	return predicate.CardType(sql.FieldNotIn(FieldNameEn, vs...))
}

// NameEnGT applies the GT predicate on the "name_en" field.
func NameEnGT(v string) predicate.CardType {
	return predicate.CardType(sql.FieldGT(FieldNameEn, v))
}

// NameEnGTE applies the GTE predicate on the "name_en" field.
func NameEnGTE(v string) predicate.CardType {
	return predicate.CardType(sql.FieldGTE(FieldNameEn, v))
}

// NameEnLT applies the LT predicate on the "name_en" field.
func NameEnLT(v string) predicate.CardType {
	return predicate.CardType(sql.FieldLT(FieldNameEn, v))
}

// NameEnLTE applies the LTE predicate on the "name_en" field.
func NameEnLTE(v string) predicate.CardType {
	return predicate.CardType(sql.FieldLTE(FieldNameEn, v))
}

// NameEnContains applies the Contains predicate on the "name_en" field.
func NameEnContains(v string) predicate.CardType {
	return predicate.CardType(sql.FieldContains(FieldNameEn, v))
}

// NameEnHasPrefix applies the HasPrefix predicate on the "name_en" field.
func NameEnHasPrefix(v string) predicate.CardType {
	return predicate.CardType(sql.FieldHasPrefix(FieldNameEn, v))
}

// NameEnHasSuffix applies the HasSuffix predicate on the "name_en" field.
func NameEnHasSuffix(v string) predicate.CardType {
	return predicate.CardType(sql.FieldHasSuffix(FieldNameEn, v))
}

// NameEnIsNil applies the IsNil predicate on the "name_en" field.
func NameEnIsNil() predicate.CardType {
	return predicate.CardType(sql.FieldIsNull(FieldNameEn))
}

// NameEnNotNil applies the NotNil predicate on the "name_en" field.
func NameEnNotNil() predicate.CardType {
	return predicate.CardType(sql.FieldNotNull(FieldNameEn))
}

// NameEnEqualFold applies the EqualFold predicate on the "name_en" field.
func NameEnEqualFold(v string) predicate.CardType {
	return predicate.CardType(sql.FieldEqualFold(FieldNameEn, v))
}

// NameEnContainsFold applies the ContainsFold predicate on the "name_en" field.
func NameEnContainsFold(v string) predicate.CardType {
	return predicate.CardType(sql.FieldContainsFold(FieldNameEn, v))
}

// HasCards applies the HasEdge predicate on the "cards" edge.
func HasCards() predicate.CardType {
	return predicate.CardType(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, CardsTable, CardsColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasCardsWith applies the HasEdge predicate on the "cards" edge with a given conditions (other predicates).
func HasCardsWith(preds ...predicate.Card) predicate.CardType {
	return predicate.CardType(func(s *sql.Selector) {
		step := newCardsStep()
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.CardType) predicate.CardType {
	return predicate.CardType(sql.AndPredicates(predicates...))
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.CardType) predicate.CardType {
	return predicate.CardType(sql.OrPredicates(predicates...))
}

// Not applies the not operator on the given predicate.
func Not(p predicate.CardType) predicate.CardType {
	return predicate.CardType(sql.NotPredicates(p))
}
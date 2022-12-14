// Code generated by ent, DO NOT EDIT.

package cardspecialcolor

import (
	"agricoladb/ent/predicate"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
)

// ID filters vertices based on their ID field.
func ID(id int) predicate.CardSpecialColor {
	return predicate.CardSpecialColor(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldID), id))
	})
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id int) predicate.CardSpecialColor {
	return predicate.CardSpecialColor(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldID), id))
	})
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id int) predicate.CardSpecialColor {
	return predicate.CardSpecialColor(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldID), id))
	})
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...int) predicate.CardSpecialColor {
	return predicate.CardSpecialColor(func(s *sql.Selector) {
		v := make([]any, len(ids))
		for i := range v {
			v[i] = ids[i]
		}
		s.Where(sql.In(s.C(FieldID), v...))
	})
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...int) predicate.CardSpecialColor {
	return predicate.CardSpecialColor(func(s *sql.Selector) {
		v := make([]any, len(ids))
		for i := range v {
			v[i] = ids[i]
		}
		s.Where(sql.NotIn(s.C(FieldID), v...))
	})
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id int) predicate.CardSpecialColor {
	return predicate.CardSpecialColor(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldID), id))
	})
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id int) predicate.CardSpecialColor {
	return predicate.CardSpecialColor(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldID), id))
	})
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id int) predicate.CardSpecialColor {
	return predicate.CardSpecialColor(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldID), id))
	})
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id int) predicate.CardSpecialColor {
	return predicate.CardSpecialColor(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldID), id))
	})
}

// Key applies equality check predicate on the "key" field. It's identical to KeyEQ.
func Key(v string) predicate.CardSpecialColor {
	return predicate.CardSpecialColor(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldKey), v))
	})
}

// NameJa applies equality check predicate on the "name_ja" field. It's identical to NameJaEQ.
func NameJa(v string) predicate.CardSpecialColor {
	return predicate.CardSpecialColor(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldNameJa), v))
	})
}

// NameEn applies equality check predicate on the "name_en" field. It's identical to NameEnEQ.
func NameEn(v string) predicate.CardSpecialColor {
	return predicate.CardSpecialColor(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldNameEn), v))
	})
}

// KeyEQ applies the EQ predicate on the "key" field.
func KeyEQ(v string) predicate.CardSpecialColor {
	return predicate.CardSpecialColor(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldKey), v))
	})
}

// KeyNEQ applies the NEQ predicate on the "key" field.
func KeyNEQ(v string) predicate.CardSpecialColor {
	return predicate.CardSpecialColor(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldKey), v))
	})
}

// KeyIn applies the In predicate on the "key" field.
func KeyIn(vs ...string) predicate.CardSpecialColor {
	v := make([]any, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.CardSpecialColor(func(s *sql.Selector) {
		s.Where(sql.In(s.C(FieldKey), v...))
	})
}

// KeyNotIn applies the NotIn predicate on the "key" field.
func KeyNotIn(vs ...string) predicate.CardSpecialColor {
	v := make([]any, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.CardSpecialColor(func(s *sql.Selector) {
		s.Where(sql.NotIn(s.C(FieldKey), v...))
	})
}

// KeyGT applies the GT predicate on the "key" field.
func KeyGT(v string) predicate.CardSpecialColor {
	return predicate.CardSpecialColor(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldKey), v))
	})
}

// KeyGTE applies the GTE predicate on the "key" field.
func KeyGTE(v string) predicate.CardSpecialColor {
	return predicate.CardSpecialColor(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldKey), v))
	})
}

// KeyLT applies the LT predicate on the "key" field.
func KeyLT(v string) predicate.CardSpecialColor {
	return predicate.CardSpecialColor(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldKey), v))
	})
}

// KeyLTE applies the LTE predicate on the "key" field.
func KeyLTE(v string) predicate.CardSpecialColor {
	return predicate.CardSpecialColor(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldKey), v))
	})
}

// KeyContains applies the Contains predicate on the "key" field.
func KeyContains(v string) predicate.CardSpecialColor {
	return predicate.CardSpecialColor(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldKey), v))
	})
}

// KeyHasPrefix applies the HasPrefix predicate on the "key" field.
func KeyHasPrefix(v string) predicate.CardSpecialColor {
	return predicate.CardSpecialColor(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldKey), v))
	})
}

// KeyHasSuffix applies the HasSuffix predicate on the "key" field.
func KeyHasSuffix(v string) predicate.CardSpecialColor {
	return predicate.CardSpecialColor(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldKey), v))
	})
}

// KeyEqualFold applies the EqualFold predicate on the "key" field.
func KeyEqualFold(v string) predicate.CardSpecialColor {
	return predicate.CardSpecialColor(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldKey), v))
	})
}

// KeyContainsFold applies the ContainsFold predicate on the "key" field.
func KeyContainsFold(v string) predicate.CardSpecialColor {
	return predicate.CardSpecialColor(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldKey), v))
	})
}

// NameJaEQ applies the EQ predicate on the "name_ja" field.
func NameJaEQ(v string) predicate.CardSpecialColor {
	return predicate.CardSpecialColor(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldNameJa), v))
	})
}

// NameJaNEQ applies the NEQ predicate on the "name_ja" field.
func NameJaNEQ(v string) predicate.CardSpecialColor {
	return predicate.CardSpecialColor(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldNameJa), v))
	})
}

// NameJaIn applies the In predicate on the "name_ja" field.
func NameJaIn(vs ...string) predicate.CardSpecialColor {
	v := make([]any, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.CardSpecialColor(func(s *sql.Selector) {
		s.Where(sql.In(s.C(FieldNameJa), v...))
	})
}

// NameJaNotIn applies the NotIn predicate on the "name_ja" field.
func NameJaNotIn(vs ...string) predicate.CardSpecialColor {
	v := make([]any, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.CardSpecialColor(func(s *sql.Selector) {
		s.Where(sql.NotIn(s.C(FieldNameJa), v...))
	})
}

// NameJaGT applies the GT predicate on the "name_ja" field.
func NameJaGT(v string) predicate.CardSpecialColor {
	return predicate.CardSpecialColor(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldNameJa), v))
	})
}

// NameJaGTE applies the GTE predicate on the "name_ja" field.
func NameJaGTE(v string) predicate.CardSpecialColor {
	return predicate.CardSpecialColor(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldNameJa), v))
	})
}

// NameJaLT applies the LT predicate on the "name_ja" field.
func NameJaLT(v string) predicate.CardSpecialColor {
	return predicate.CardSpecialColor(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldNameJa), v))
	})
}

// NameJaLTE applies the LTE predicate on the "name_ja" field.
func NameJaLTE(v string) predicate.CardSpecialColor {
	return predicate.CardSpecialColor(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldNameJa), v))
	})
}

// NameJaContains applies the Contains predicate on the "name_ja" field.
func NameJaContains(v string) predicate.CardSpecialColor {
	return predicate.CardSpecialColor(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldNameJa), v))
	})
}

// NameJaHasPrefix applies the HasPrefix predicate on the "name_ja" field.
func NameJaHasPrefix(v string) predicate.CardSpecialColor {
	return predicate.CardSpecialColor(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldNameJa), v))
	})
}

// NameJaHasSuffix applies the HasSuffix predicate on the "name_ja" field.
func NameJaHasSuffix(v string) predicate.CardSpecialColor {
	return predicate.CardSpecialColor(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldNameJa), v))
	})
}

// NameJaIsNil applies the IsNil predicate on the "name_ja" field.
func NameJaIsNil() predicate.CardSpecialColor {
	return predicate.CardSpecialColor(func(s *sql.Selector) {
		s.Where(sql.IsNull(s.C(FieldNameJa)))
	})
}

// NameJaNotNil applies the NotNil predicate on the "name_ja" field.
func NameJaNotNil() predicate.CardSpecialColor {
	return predicate.CardSpecialColor(func(s *sql.Selector) {
		s.Where(sql.NotNull(s.C(FieldNameJa)))
	})
}

// NameJaEqualFold applies the EqualFold predicate on the "name_ja" field.
func NameJaEqualFold(v string) predicate.CardSpecialColor {
	return predicate.CardSpecialColor(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldNameJa), v))
	})
}

// NameJaContainsFold applies the ContainsFold predicate on the "name_ja" field.
func NameJaContainsFold(v string) predicate.CardSpecialColor {
	return predicate.CardSpecialColor(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldNameJa), v))
	})
}

// NameEnEQ applies the EQ predicate on the "name_en" field.
func NameEnEQ(v string) predicate.CardSpecialColor {
	return predicate.CardSpecialColor(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldNameEn), v))
	})
}

// NameEnNEQ applies the NEQ predicate on the "name_en" field.
func NameEnNEQ(v string) predicate.CardSpecialColor {
	return predicate.CardSpecialColor(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldNameEn), v))
	})
}

// NameEnIn applies the In predicate on the "name_en" field.
func NameEnIn(vs ...string) predicate.CardSpecialColor {
	v := make([]any, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.CardSpecialColor(func(s *sql.Selector) {
		s.Where(sql.In(s.C(FieldNameEn), v...))
	})
}

// NameEnNotIn applies the NotIn predicate on the "name_en" field.
func NameEnNotIn(vs ...string) predicate.CardSpecialColor {
	v := make([]any, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.CardSpecialColor(func(s *sql.Selector) {
		s.Where(sql.NotIn(s.C(FieldNameEn), v...))
	})
}

// NameEnGT applies the GT predicate on the "name_en" field.
func NameEnGT(v string) predicate.CardSpecialColor {
	return predicate.CardSpecialColor(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldNameEn), v))
	})
}

// NameEnGTE applies the GTE predicate on the "name_en" field.
func NameEnGTE(v string) predicate.CardSpecialColor {
	return predicate.CardSpecialColor(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldNameEn), v))
	})
}

// NameEnLT applies the LT predicate on the "name_en" field.
func NameEnLT(v string) predicate.CardSpecialColor {
	return predicate.CardSpecialColor(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldNameEn), v))
	})
}

// NameEnLTE applies the LTE predicate on the "name_en" field.
func NameEnLTE(v string) predicate.CardSpecialColor {
	return predicate.CardSpecialColor(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldNameEn), v))
	})
}

// NameEnContains applies the Contains predicate on the "name_en" field.
func NameEnContains(v string) predicate.CardSpecialColor {
	return predicate.CardSpecialColor(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldNameEn), v))
	})
}

// NameEnHasPrefix applies the HasPrefix predicate on the "name_en" field.
func NameEnHasPrefix(v string) predicate.CardSpecialColor {
	return predicate.CardSpecialColor(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldNameEn), v))
	})
}

// NameEnHasSuffix applies the HasSuffix predicate on the "name_en" field.
func NameEnHasSuffix(v string) predicate.CardSpecialColor {
	return predicate.CardSpecialColor(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldNameEn), v))
	})
}

// NameEnIsNil applies the IsNil predicate on the "name_en" field.
func NameEnIsNil() predicate.CardSpecialColor {
	return predicate.CardSpecialColor(func(s *sql.Selector) {
		s.Where(sql.IsNull(s.C(FieldNameEn)))
	})
}

// NameEnNotNil applies the NotNil predicate on the "name_en" field.
func NameEnNotNil() predicate.CardSpecialColor {
	return predicate.CardSpecialColor(func(s *sql.Selector) {
		s.Where(sql.NotNull(s.C(FieldNameEn)))
	})
}

// NameEnEqualFold applies the EqualFold predicate on the "name_en" field.
func NameEnEqualFold(v string) predicate.CardSpecialColor {
	return predicate.CardSpecialColor(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldNameEn), v))
	})
}

// NameEnContainsFold applies the ContainsFold predicate on the "name_en" field.
func NameEnContainsFold(v string) predicate.CardSpecialColor {
	return predicate.CardSpecialColor(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldNameEn), v))
	})
}

// HasCards applies the HasEdge predicate on the "cards" edge.
func HasCards() predicate.CardSpecialColor {
	return predicate.CardSpecialColor(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(CardsTable, FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, CardsTable, CardsColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasCardsWith applies the HasEdge predicate on the "cards" edge with a given conditions (other predicates).
func HasCardsWith(preds ...predicate.Card) predicate.CardSpecialColor {
	return predicate.CardSpecialColor(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(CardsInverseTable, FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, CardsTable, CardsColumn),
		)
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.CardSpecialColor) predicate.CardSpecialColor {
	return predicate.CardSpecialColor(func(s *sql.Selector) {
		s1 := s.Clone().SetP(nil)
		for _, p := range predicates {
			p(s1)
		}
		s.Where(s1.P())
	})
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.CardSpecialColor) predicate.CardSpecialColor {
	return predicate.CardSpecialColor(func(s *sql.Selector) {
		s1 := s.Clone().SetP(nil)
		for i, p := range predicates {
			if i > 0 {
				s1.Or()
			}
			p(s1)
		}
		s.Where(s1.P())
	})
}

// Not applies the not operator on the given predicate.
func Not(p predicate.CardSpecialColor) predicate.CardSpecialColor {
	return predicate.CardSpecialColor(func(s *sql.Selector) {
		p(s.Not())
	})
}

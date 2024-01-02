// Code generated by ent, DO NOT EDIT.

package ent

import (
	"fmt"
	"strings"

	"entgo.io/ent"
	"entgo.io/ent/dialect/sql"
	"github.com/AgricolaDevJP/agricoladb-server/ent/cardspecialcolor"
)

// CardSpecialColor is the model entity for the CardSpecialColor schema.
type CardSpecialColor struct {
	config `json:"-"`
	// ID of the ent.
	ID int `json:"id,omitempty"`
	// Key holds the value of the "key" field.
	Key string `json:"key,omitempty"`
	// NameJa holds the value of the "name_ja" field.
	NameJa string `json:"name_ja,omitempty"`
	// NameEn holds the value of the "name_en" field.
	NameEn string `json:"name_en,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the CardSpecialColorQuery when eager-loading is set.
	Edges        CardSpecialColorEdges `json:"edges"`
	selectValues sql.SelectValues
}

// CardSpecialColorEdges holds the relations/edges for other nodes in the graph.
type CardSpecialColorEdges struct {
	// Cards holds the value of the cards edge.
	Cards []*Card `json:"cards,omitempty"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [1]bool
	// totalCount holds the count of the edges above.
	totalCount [1]map[string]int

	namedCards map[string][]*Card
}

// CardsOrErr returns the Cards value or an error if the edge
// was not loaded in eager-loading.
func (e CardSpecialColorEdges) CardsOrErr() ([]*Card, error) {
	if e.loadedTypes[0] {
		return e.Cards, nil
	}
	return nil, &NotLoadedError{edge: "cards"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*CardSpecialColor) scanValues(columns []string) ([]any, error) {
	values := make([]any, len(columns))
	for i := range columns {
		switch columns[i] {
		case cardspecialcolor.FieldID:
			values[i] = new(sql.NullInt64)
		case cardspecialcolor.FieldKey, cardspecialcolor.FieldNameJa, cardspecialcolor.FieldNameEn:
			values[i] = new(sql.NullString)
		default:
			values[i] = new(sql.UnknownType)
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the CardSpecialColor fields.
func (csc *CardSpecialColor) assignValues(columns []string, values []any) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case cardspecialcolor.FieldID:
			value, ok := values[i].(*sql.NullInt64)
			if !ok {
				return fmt.Errorf("unexpected type %T for field id", value)
			}
			csc.ID = int(value.Int64)
		case cardspecialcolor.FieldKey:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field key", values[i])
			} else if value.Valid {
				csc.Key = value.String
			}
		case cardspecialcolor.FieldNameJa:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field name_ja", values[i])
			} else if value.Valid {
				csc.NameJa = value.String
			}
		case cardspecialcolor.FieldNameEn:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field name_en", values[i])
			} else if value.Valid {
				csc.NameEn = value.String
			}
		default:
			csc.selectValues.Set(columns[i], values[i])
		}
	}
	return nil
}

// Value returns the ent.Value that was dynamically selected and assigned to the CardSpecialColor.
// This includes values selected through modifiers, order, etc.
func (csc *CardSpecialColor) Value(name string) (ent.Value, error) {
	return csc.selectValues.Get(name)
}

// QueryCards queries the "cards" edge of the CardSpecialColor entity.
func (csc *CardSpecialColor) QueryCards() *CardQuery {
	return NewCardSpecialColorClient(csc.config).QueryCards(csc)
}

// Update returns a builder for updating this CardSpecialColor.
// Note that you need to call CardSpecialColor.Unwrap() before calling this method if this CardSpecialColor
// was returned from a transaction, and the transaction was committed or rolled back.
func (csc *CardSpecialColor) Update() *CardSpecialColorUpdateOne {
	return NewCardSpecialColorClient(csc.config).UpdateOne(csc)
}

// Unwrap unwraps the CardSpecialColor entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (csc *CardSpecialColor) Unwrap() *CardSpecialColor {
	_tx, ok := csc.config.driver.(*txDriver)
	if !ok {
		panic("ent: CardSpecialColor is not a transactional entity")
	}
	csc.config.driver = _tx.drv
	return csc
}

// String implements the fmt.Stringer.
func (csc *CardSpecialColor) String() string {
	var builder strings.Builder
	builder.WriteString("CardSpecialColor(")
	builder.WriteString(fmt.Sprintf("id=%v, ", csc.ID))
	builder.WriteString("key=")
	builder.WriteString(csc.Key)
	builder.WriteString(", ")
	builder.WriteString("name_ja=")
	builder.WriteString(csc.NameJa)
	builder.WriteString(", ")
	builder.WriteString("name_en=")
	builder.WriteString(csc.NameEn)
	builder.WriteByte(')')
	return builder.String()
}

// NamedCards returns the Cards named value or an error if the edge was not
// loaded in eager-loading with this name.
func (csc *CardSpecialColor) NamedCards(name string) ([]*Card, error) {
	if csc.Edges.namedCards == nil {
		return nil, &NotLoadedError{edge: name}
	}
	nodes, ok := csc.Edges.namedCards[name]
	if !ok {
		return nil, &NotLoadedError{edge: name}
	}
	return nodes, nil
}

func (csc *CardSpecialColor) appendNamedCards(name string, edges ...*Card) {
	if csc.Edges.namedCards == nil {
		csc.Edges.namedCards = make(map[string][]*Card)
	}
	if len(edges) == 0 {
		csc.Edges.namedCards[name] = []*Card{}
	} else {
		csc.Edges.namedCards[name] = append(csc.Edges.namedCards[name], edges...)
	}
}

// CardSpecialColors is a parsable slice of CardSpecialColor.
type CardSpecialColors []*CardSpecialColor
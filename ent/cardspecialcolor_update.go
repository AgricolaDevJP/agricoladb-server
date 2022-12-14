// Code generated by ent, DO NOT EDIT.

package ent

import (
	"agricoladb/ent/card"
	"agricoladb/ent/cardspecialcolor"
	"agricoladb/ent/predicate"
	"context"
	"errors"
	"fmt"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// CardSpecialColorUpdate is the builder for updating CardSpecialColor entities.
type CardSpecialColorUpdate struct {
	config
	hooks    []Hook
	mutation *CardSpecialColorMutation
}

// Where appends a list predicates to the CardSpecialColorUpdate builder.
func (cscu *CardSpecialColorUpdate) Where(ps ...predicate.CardSpecialColor) *CardSpecialColorUpdate {
	cscu.mutation.Where(ps...)
	return cscu
}

// SetNameJa sets the "name_ja" field.
func (cscu *CardSpecialColorUpdate) SetNameJa(s string) *CardSpecialColorUpdate {
	cscu.mutation.SetNameJa(s)
	return cscu
}

// SetNillableNameJa sets the "name_ja" field if the given value is not nil.
func (cscu *CardSpecialColorUpdate) SetNillableNameJa(s *string) *CardSpecialColorUpdate {
	if s != nil {
		cscu.SetNameJa(*s)
	}
	return cscu
}

// ClearNameJa clears the value of the "name_ja" field.
func (cscu *CardSpecialColorUpdate) ClearNameJa() *CardSpecialColorUpdate {
	cscu.mutation.ClearNameJa()
	return cscu
}

// SetNameEn sets the "name_en" field.
func (cscu *CardSpecialColorUpdate) SetNameEn(s string) *CardSpecialColorUpdate {
	cscu.mutation.SetNameEn(s)
	return cscu
}

// SetNillableNameEn sets the "name_en" field if the given value is not nil.
func (cscu *CardSpecialColorUpdate) SetNillableNameEn(s *string) *CardSpecialColorUpdate {
	if s != nil {
		cscu.SetNameEn(*s)
	}
	return cscu
}

// ClearNameEn clears the value of the "name_en" field.
func (cscu *CardSpecialColorUpdate) ClearNameEn() *CardSpecialColorUpdate {
	cscu.mutation.ClearNameEn()
	return cscu
}

// AddCardIDs adds the "cards" edge to the Card entity by IDs.
func (cscu *CardSpecialColorUpdate) AddCardIDs(ids ...int) *CardSpecialColorUpdate {
	cscu.mutation.AddCardIDs(ids...)
	return cscu
}

// AddCards adds the "cards" edges to the Card entity.
func (cscu *CardSpecialColorUpdate) AddCards(c ...*Card) *CardSpecialColorUpdate {
	ids := make([]int, len(c))
	for i := range c {
		ids[i] = c[i].ID
	}
	return cscu.AddCardIDs(ids...)
}

// Mutation returns the CardSpecialColorMutation object of the builder.
func (cscu *CardSpecialColorUpdate) Mutation() *CardSpecialColorMutation {
	return cscu.mutation
}

// ClearCards clears all "cards" edges to the Card entity.
func (cscu *CardSpecialColorUpdate) ClearCards() *CardSpecialColorUpdate {
	cscu.mutation.ClearCards()
	return cscu
}

// RemoveCardIDs removes the "cards" edge to Card entities by IDs.
func (cscu *CardSpecialColorUpdate) RemoveCardIDs(ids ...int) *CardSpecialColorUpdate {
	cscu.mutation.RemoveCardIDs(ids...)
	return cscu
}

// RemoveCards removes "cards" edges to Card entities.
func (cscu *CardSpecialColorUpdate) RemoveCards(c ...*Card) *CardSpecialColorUpdate {
	ids := make([]int, len(c))
	for i := range c {
		ids[i] = c[i].ID
	}
	return cscu.RemoveCardIDs(ids...)
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (cscu *CardSpecialColorUpdate) Save(ctx context.Context) (int, error) {
	var (
		err      error
		affected int
	)
	if len(cscu.hooks) == 0 {
		affected, err = cscu.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*CardSpecialColorMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			cscu.mutation = mutation
			affected, err = cscu.sqlSave(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(cscu.hooks) - 1; i >= 0; i-- {
			if cscu.hooks[i] == nil {
				return 0, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = cscu.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, cscu.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// SaveX is like Save, but panics if an error occurs.
func (cscu *CardSpecialColorUpdate) SaveX(ctx context.Context) int {
	affected, err := cscu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (cscu *CardSpecialColorUpdate) Exec(ctx context.Context) error {
	_, err := cscu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (cscu *CardSpecialColorUpdate) ExecX(ctx context.Context) {
	if err := cscu.Exec(ctx); err != nil {
		panic(err)
	}
}

func (cscu *CardSpecialColorUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   cardspecialcolor.Table,
			Columns: cardspecialcolor.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: cardspecialcolor.FieldID,
			},
		},
	}
	if ps := cscu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := cscu.mutation.NameJa(); ok {
		_spec.SetField(cardspecialcolor.FieldNameJa, field.TypeString, value)
	}
	if cscu.mutation.NameJaCleared() {
		_spec.ClearField(cardspecialcolor.FieldNameJa, field.TypeString)
	}
	if value, ok := cscu.mutation.NameEn(); ok {
		_spec.SetField(cardspecialcolor.FieldNameEn, field.TypeString, value)
	}
	if cscu.mutation.NameEnCleared() {
		_spec.ClearField(cardspecialcolor.FieldNameEn, field.TypeString)
	}
	if cscu.mutation.CardsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   cardspecialcolor.CardsTable,
			Columns: []string{cardspecialcolor.CardsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: card.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := cscu.mutation.RemovedCardsIDs(); len(nodes) > 0 && !cscu.mutation.CardsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   cardspecialcolor.CardsTable,
			Columns: []string{cardspecialcolor.CardsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: card.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := cscu.mutation.CardsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   cardspecialcolor.CardsTable,
			Columns: []string{cardspecialcolor.CardsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: card.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, cscu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{cardspecialcolor.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	return n, nil
}

// CardSpecialColorUpdateOne is the builder for updating a single CardSpecialColor entity.
type CardSpecialColorUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *CardSpecialColorMutation
}

// SetNameJa sets the "name_ja" field.
func (cscuo *CardSpecialColorUpdateOne) SetNameJa(s string) *CardSpecialColorUpdateOne {
	cscuo.mutation.SetNameJa(s)
	return cscuo
}

// SetNillableNameJa sets the "name_ja" field if the given value is not nil.
func (cscuo *CardSpecialColorUpdateOne) SetNillableNameJa(s *string) *CardSpecialColorUpdateOne {
	if s != nil {
		cscuo.SetNameJa(*s)
	}
	return cscuo
}

// ClearNameJa clears the value of the "name_ja" field.
func (cscuo *CardSpecialColorUpdateOne) ClearNameJa() *CardSpecialColorUpdateOne {
	cscuo.mutation.ClearNameJa()
	return cscuo
}

// SetNameEn sets the "name_en" field.
func (cscuo *CardSpecialColorUpdateOne) SetNameEn(s string) *CardSpecialColorUpdateOne {
	cscuo.mutation.SetNameEn(s)
	return cscuo
}

// SetNillableNameEn sets the "name_en" field if the given value is not nil.
func (cscuo *CardSpecialColorUpdateOne) SetNillableNameEn(s *string) *CardSpecialColorUpdateOne {
	if s != nil {
		cscuo.SetNameEn(*s)
	}
	return cscuo
}

// ClearNameEn clears the value of the "name_en" field.
func (cscuo *CardSpecialColorUpdateOne) ClearNameEn() *CardSpecialColorUpdateOne {
	cscuo.mutation.ClearNameEn()
	return cscuo
}

// AddCardIDs adds the "cards" edge to the Card entity by IDs.
func (cscuo *CardSpecialColorUpdateOne) AddCardIDs(ids ...int) *CardSpecialColorUpdateOne {
	cscuo.mutation.AddCardIDs(ids...)
	return cscuo
}

// AddCards adds the "cards" edges to the Card entity.
func (cscuo *CardSpecialColorUpdateOne) AddCards(c ...*Card) *CardSpecialColorUpdateOne {
	ids := make([]int, len(c))
	for i := range c {
		ids[i] = c[i].ID
	}
	return cscuo.AddCardIDs(ids...)
}

// Mutation returns the CardSpecialColorMutation object of the builder.
func (cscuo *CardSpecialColorUpdateOne) Mutation() *CardSpecialColorMutation {
	return cscuo.mutation
}

// ClearCards clears all "cards" edges to the Card entity.
func (cscuo *CardSpecialColorUpdateOne) ClearCards() *CardSpecialColorUpdateOne {
	cscuo.mutation.ClearCards()
	return cscuo
}

// RemoveCardIDs removes the "cards" edge to Card entities by IDs.
func (cscuo *CardSpecialColorUpdateOne) RemoveCardIDs(ids ...int) *CardSpecialColorUpdateOne {
	cscuo.mutation.RemoveCardIDs(ids...)
	return cscuo
}

// RemoveCards removes "cards" edges to Card entities.
func (cscuo *CardSpecialColorUpdateOne) RemoveCards(c ...*Card) *CardSpecialColorUpdateOne {
	ids := make([]int, len(c))
	for i := range c {
		ids[i] = c[i].ID
	}
	return cscuo.RemoveCardIDs(ids...)
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (cscuo *CardSpecialColorUpdateOne) Select(field string, fields ...string) *CardSpecialColorUpdateOne {
	cscuo.fields = append([]string{field}, fields...)
	return cscuo
}

// Save executes the query and returns the updated CardSpecialColor entity.
func (cscuo *CardSpecialColorUpdateOne) Save(ctx context.Context) (*CardSpecialColor, error) {
	var (
		err  error
		node *CardSpecialColor
	)
	if len(cscuo.hooks) == 0 {
		node, err = cscuo.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*CardSpecialColorMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			cscuo.mutation = mutation
			node, err = cscuo.sqlSave(ctx)
			mutation.done = true
			return node, err
		})
		for i := len(cscuo.hooks) - 1; i >= 0; i-- {
			if cscuo.hooks[i] == nil {
				return nil, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = cscuo.hooks[i](mut)
		}
		v, err := mut.Mutate(ctx, cscuo.mutation)
		if err != nil {
			return nil, err
		}
		nv, ok := v.(*CardSpecialColor)
		if !ok {
			return nil, fmt.Errorf("unexpected node type %T returned from CardSpecialColorMutation", v)
		}
		node = nv
	}
	return node, err
}

// SaveX is like Save, but panics if an error occurs.
func (cscuo *CardSpecialColorUpdateOne) SaveX(ctx context.Context) *CardSpecialColor {
	node, err := cscuo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (cscuo *CardSpecialColorUpdateOne) Exec(ctx context.Context) error {
	_, err := cscuo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (cscuo *CardSpecialColorUpdateOne) ExecX(ctx context.Context) {
	if err := cscuo.Exec(ctx); err != nil {
		panic(err)
	}
}

func (cscuo *CardSpecialColorUpdateOne) sqlSave(ctx context.Context) (_node *CardSpecialColor, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   cardspecialcolor.Table,
			Columns: cardspecialcolor.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: cardspecialcolor.FieldID,
			},
		},
	}
	id, ok := cscuo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "CardSpecialColor.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := cscuo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, cardspecialcolor.FieldID)
		for _, f := range fields {
			if !cardspecialcolor.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != cardspecialcolor.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := cscuo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := cscuo.mutation.NameJa(); ok {
		_spec.SetField(cardspecialcolor.FieldNameJa, field.TypeString, value)
	}
	if cscuo.mutation.NameJaCleared() {
		_spec.ClearField(cardspecialcolor.FieldNameJa, field.TypeString)
	}
	if value, ok := cscuo.mutation.NameEn(); ok {
		_spec.SetField(cardspecialcolor.FieldNameEn, field.TypeString, value)
	}
	if cscuo.mutation.NameEnCleared() {
		_spec.ClearField(cardspecialcolor.FieldNameEn, field.TypeString)
	}
	if cscuo.mutation.CardsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   cardspecialcolor.CardsTable,
			Columns: []string{cardspecialcolor.CardsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: card.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := cscuo.mutation.RemovedCardsIDs(); len(nodes) > 0 && !cscuo.mutation.CardsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   cardspecialcolor.CardsTable,
			Columns: []string{cardspecialcolor.CardsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: card.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := cscuo.mutation.CardsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   cardspecialcolor.CardsTable,
			Columns: []string{cardspecialcolor.CardsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: card.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	_node = &CardSpecialColor{config: cscuo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, cscuo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{cardspecialcolor.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	return _node, nil
}

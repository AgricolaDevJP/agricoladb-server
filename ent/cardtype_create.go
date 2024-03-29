// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/AgricolaDevJP/agricoladb-server/ent/card"
	"github.com/AgricolaDevJP/agricoladb-server/ent/cardtype"
)

// CardTypeCreate is the builder for creating a CardType entity.
type CardTypeCreate struct {
	config
	mutation *CardTypeMutation
	hooks    []Hook
	conflict []sql.ConflictOption
}

// SetKey sets the "key" field.
func (ctc *CardTypeCreate) SetKey(s string) *CardTypeCreate {
	ctc.mutation.SetKey(s)
	return ctc
}

// SetNameJa sets the "name_ja" field.
func (ctc *CardTypeCreate) SetNameJa(s string) *CardTypeCreate {
	ctc.mutation.SetNameJa(s)
	return ctc
}

// SetNillableNameJa sets the "name_ja" field if the given value is not nil.
func (ctc *CardTypeCreate) SetNillableNameJa(s *string) *CardTypeCreate {
	if s != nil {
		ctc.SetNameJa(*s)
	}
	return ctc
}

// SetNameEn sets the "name_en" field.
func (ctc *CardTypeCreate) SetNameEn(s string) *CardTypeCreate {
	ctc.mutation.SetNameEn(s)
	return ctc
}

// SetNillableNameEn sets the "name_en" field if the given value is not nil.
func (ctc *CardTypeCreate) SetNillableNameEn(s *string) *CardTypeCreate {
	if s != nil {
		ctc.SetNameEn(*s)
	}
	return ctc
}

// SetID sets the "id" field.
func (ctc *CardTypeCreate) SetID(i int) *CardTypeCreate {
	ctc.mutation.SetID(i)
	return ctc
}

// AddCardIDs adds the "cards" edge to the Card entity by IDs.
func (ctc *CardTypeCreate) AddCardIDs(ids ...int) *CardTypeCreate {
	ctc.mutation.AddCardIDs(ids...)
	return ctc
}

// AddCards adds the "cards" edges to the Card entity.
func (ctc *CardTypeCreate) AddCards(c ...*Card) *CardTypeCreate {
	ids := make([]int, len(c))
	for i := range c {
		ids[i] = c[i].ID
	}
	return ctc.AddCardIDs(ids...)
}

// Mutation returns the CardTypeMutation object of the builder.
func (ctc *CardTypeCreate) Mutation() *CardTypeMutation {
	return ctc.mutation
}

// Save creates the CardType in the database.
func (ctc *CardTypeCreate) Save(ctx context.Context) (*CardType, error) {
	return withHooks(ctx, ctc.sqlSave, ctc.mutation, ctc.hooks)
}

// SaveX calls Save and panics if Save returns an error.
func (ctc *CardTypeCreate) SaveX(ctx context.Context) *CardType {
	v, err := ctc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (ctc *CardTypeCreate) Exec(ctx context.Context) error {
	_, err := ctc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (ctc *CardTypeCreate) ExecX(ctx context.Context) {
	if err := ctc.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (ctc *CardTypeCreate) check() error {
	if _, ok := ctc.mutation.Key(); !ok {
		return &ValidationError{Name: "key", err: errors.New(`ent: missing required field "CardType.key"`)}
	}
	if v, ok := ctc.mutation.Key(); ok {
		if err := cardtype.KeyValidator(v); err != nil {
			return &ValidationError{Name: "key", err: fmt.Errorf(`ent: validator failed for field "CardType.key": %w`, err)}
		}
	}
	return nil
}

func (ctc *CardTypeCreate) sqlSave(ctx context.Context) (*CardType, error) {
	if err := ctc.check(); err != nil {
		return nil, err
	}
	_node, _spec := ctc.createSpec()
	if err := sqlgraph.CreateNode(ctx, ctc.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	if _spec.ID.Value != _node.ID {
		id := _spec.ID.Value.(int64)
		_node.ID = int(id)
	}
	ctc.mutation.id = &_node.ID
	ctc.mutation.done = true
	return _node, nil
}

func (ctc *CardTypeCreate) createSpec() (*CardType, *sqlgraph.CreateSpec) {
	var (
		_node = &CardType{config: ctc.config}
		_spec = sqlgraph.NewCreateSpec(cardtype.Table, sqlgraph.NewFieldSpec(cardtype.FieldID, field.TypeInt))
	)
	_spec.OnConflict = ctc.conflict
	if id, ok := ctc.mutation.ID(); ok {
		_node.ID = id
		_spec.ID.Value = id
	}
	if value, ok := ctc.mutation.Key(); ok {
		_spec.SetField(cardtype.FieldKey, field.TypeString, value)
		_node.Key = value
	}
	if value, ok := ctc.mutation.NameJa(); ok {
		_spec.SetField(cardtype.FieldNameJa, field.TypeString, value)
		_node.NameJa = value
	}
	if value, ok := ctc.mutation.NameEn(); ok {
		_spec.SetField(cardtype.FieldNameEn, field.TypeString, value)
		_node.NameEn = value
	}
	if nodes := ctc.mutation.CardsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   cardtype.CardsTable,
			Columns: []string{cardtype.CardsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(card.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	return _node, _spec
}

// OnConflict allows configuring the `ON CONFLICT` / `ON DUPLICATE KEY` clause
// of the `INSERT` statement. For example:
//
//	client.CardType.Create().
//		SetKey(v).
//		OnConflict(
//			// Update the row with the new values
//			// the was proposed for insertion.
//			sql.ResolveWithNewValues(),
//		).
//		// Override some of the fields with custom
//		// update values.
//		Update(func(u *ent.CardTypeUpsert) {
//			SetKey(v+v).
//		}).
//		Exec(ctx)
func (ctc *CardTypeCreate) OnConflict(opts ...sql.ConflictOption) *CardTypeUpsertOne {
	ctc.conflict = opts
	return &CardTypeUpsertOne{
		create: ctc,
	}
}

// OnConflictColumns calls `OnConflict` and configures the columns
// as conflict target. Using this option is equivalent to using:
//
//	client.CardType.Create().
//		OnConflict(sql.ConflictColumns(columns...)).
//		Exec(ctx)
func (ctc *CardTypeCreate) OnConflictColumns(columns ...string) *CardTypeUpsertOne {
	ctc.conflict = append(ctc.conflict, sql.ConflictColumns(columns...))
	return &CardTypeUpsertOne{
		create: ctc,
	}
}

type (
	// CardTypeUpsertOne is the builder for "upsert"-ing
	//  one CardType node.
	CardTypeUpsertOne struct {
		create *CardTypeCreate
	}

	// CardTypeUpsert is the "OnConflict" setter.
	CardTypeUpsert struct {
		*sql.UpdateSet
	}
)

// SetNameJa sets the "name_ja" field.
func (u *CardTypeUpsert) SetNameJa(v string) *CardTypeUpsert {
	u.Set(cardtype.FieldNameJa, v)
	return u
}

// UpdateNameJa sets the "name_ja" field to the value that was provided on create.
func (u *CardTypeUpsert) UpdateNameJa() *CardTypeUpsert {
	u.SetExcluded(cardtype.FieldNameJa)
	return u
}

// ClearNameJa clears the value of the "name_ja" field.
func (u *CardTypeUpsert) ClearNameJa() *CardTypeUpsert {
	u.SetNull(cardtype.FieldNameJa)
	return u
}

// SetNameEn sets the "name_en" field.
func (u *CardTypeUpsert) SetNameEn(v string) *CardTypeUpsert {
	u.Set(cardtype.FieldNameEn, v)
	return u
}

// UpdateNameEn sets the "name_en" field to the value that was provided on create.
func (u *CardTypeUpsert) UpdateNameEn() *CardTypeUpsert {
	u.SetExcluded(cardtype.FieldNameEn)
	return u
}

// ClearNameEn clears the value of the "name_en" field.
func (u *CardTypeUpsert) ClearNameEn() *CardTypeUpsert {
	u.SetNull(cardtype.FieldNameEn)
	return u
}

// UpdateNewValues updates the mutable fields using the new values that were set on create except the ID field.
// Using this option is equivalent to using:
//
//	client.CardType.Create().
//		OnConflict(
//			sql.ResolveWithNewValues(),
//			sql.ResolveWith(func(u *sql.UpdateSet) {
//				u.SetIgnore(cardtype.FieldID)
//			}),
//		).
//		Exec(ctx)
func (u *CardTypeUpsertOne) UpdateNewValues() *CardTypeUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithNewValues())
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(s *sql.UpdateSet) {
		if _, exists := u.create.mutation.ID(); exists {
			s.SetIgnore(cardtype.FieldID)
		}
		if _, exists := u.create.mutation.Key(); exists {
			s.SetIgnore(cardtype.FieldKey)
		}
	}))
	return u
}

// Ignore sets each column to itself in case of conflict.
// Using this option is equivalent to using:
//
//	client.CardType.Create().
//	    OnConflict(sql.ResolveWithIgnore()).
//	    Exec(ctx)
func (u *CardTypeUpsertOne) Ignore() *CardTypeUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithIgnore())
	return u
}

// DoNothing configures the conflict_action to `DO NOTHING`.
// Supported only by SQLite and PostgreSQL.
func (u *CardTypeUpsertOne) DoNothing() *CardTypeUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.DoNothing())
	return u
}

// Update allows overriding fields `UPDATE` values. See the CardTypeCreate.OnConflict
// documentation for more info.
func (u *CardTypeUpsertOne) Update(set func(*CardTypeUpsert)) *CardTypeUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(update *sql.UpdateSet) {
		set(&CardTypeUpsert{UpdateSet: update})
	}))
	return u
}

// SetNameJa sets the "name_ja" field.
func (u *CardTypeUpsertOne) SetNameJa(v string) *CardTypeUpsertOne {
	return u.Update(func(s *CardTypeUpsert) {
		s.SetNameJa(v)
	})
}

// UpdateNameJa sets the "name_ja" field to the value that was provided on create.
func (u *CardTypeUpsertOne) UpdateNameJa() *CardTypeUpsertOne {
	return u.Update(func(s *CardTypeUpsert) {
		s.UpdateNameJa()
	})
}

// ClearNameJa clears the value of the "name_ja" field.
func (u *CardTypeUpsertOne) ClearNameJa() *CardTypeUpsertOne {
	return u.Update(func(s *CardTypeUpsert) {
		s.ClearNameJa()
	})
}

// SetNameEn sets the "name_en" field.
func (u *CardTypeUpsertOne) SetNameEn(v string) *CardTypeUpsertOne {
	return u.Update(func(s *CardTypeUpsert) {
		s.SetNameEn(v)
	})
}

// UpdateNameEn sets the "name_en" field to the value that was provided on create.
func (u *CardTypeUpsertOne) UpdateNameEn() *CardTypeUpsertOne {
	return u.Update(func(s *CardTypeUpsert) {
		s.UpdateNameEn()
	})
}

// ClearNameEn clears the value of the "name_en" field.
func (u *CardTypeUpsertOne) ClearNameEn() *CardTypeUpsertOne {
	return u.Update(func(s *CardTypeUpsert) {
		s.ClearNameEn()
	})
}

// Exec executes the query.
func (u *CardTypeUpsertOne) Exec(ctx context.Context) error {
	if len(u.create.conflict) == 0 {
		return errors.New("ent: missing options for CardTypeCreate.OnConflict")
	}
	return u.create.Exec(ctx)
}

// ExecX is like Exec, but panics if an error occurs.
func (u *CardTypeUpsertOne) ExecX(ctx context.Context) {
	if err := u.create.Exec(ctx); err != nil {
		panic(err)
	}
}

// Exec executes the UPSERT query and returns the inserted/updated ID.
func (u *CardTypeUpsertOne) ID(ctx context.Context) (id int, err error) {
	node, err := u.create.Save(ctx)
	if err != nil {
		return id, err
	}
	return node.ID, nil
}

// IDX is like ID, but panics if an error occurs.
func (u *CardTypeUpsertOne) IDX(ctx context.Context) int {
	id, err := u.ID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// CardTypeCreateBulk is the builder for creating many CardType entities in bulk.
type CardTypeCreateBulk struct {
	config
	err      error
	builders []*CardTypeCreate
	conflict []sql.ConflictOption
}

// Save creates the CardType entities in the database.
func (ctcb *CardTypeCreateBulk) Save(ctx context.Context) ([]*CardType, error) {
	if ctcb.err != nil {
		return nil, ctcb.err
	}
	specs := make([]*sqlgraph.CreateSpec, len(ctcb.builders))
	nodes := make([]*CardType, len(ctcb.builders))
	mutators := make([]Mutator, len(ctcb.builders))
	for i := range ctcb.builders {
		func(i int, root context.Context) {
			builder := ctcb.builders[i]
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*CardTypeMutation)
				if !ok {
					return nil, fmt.Errorf("unexpected mutation type %T", m)
				}
				if err := builder.check(); err != nil {
					return nil, err
				}
				builder.mutation = mutation
				var err error
				nodes[i], specs[i] = builder.createSpec()
				if i < len(mutators)-1 {
					_, err = mutators[i+1].Mutate(root, ctcb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					spec.OnConflict = ctcb.conflict
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, ctcb.driver, spec); err != nil {
						if sqlgraph.IsConstraintError(err) {
							err = &ConstraintError{msg: err.Error(), wrap: err}
						}
					}
				}
				if err != nil {
					return nil, err
				}
				mutation.id = &nodes[i].ID
				if specs[i].ID.Value != nil && nodes[i].ID == 0 {
					id := specs[i].ID.Value.(int64)
					nodes[i].ID = int(id)
				}
				mutation.done = true
				return nodes[i], nil
			})
			for i := len(builder.hooks) - 1; i >= 0; i-- {
				mut = builder.hooks[i](mut)
			}
			mutators[i] = mut
		}(i, ctx)
	}
	if len(mutators) > 0 {
		if _, err := mutators[0].Mutate(ctx, ctcb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (ctcb *CardTypeCreateBulk) SaveX(ctx context.Context) []*CardType {
	v, err := ctcb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (ctcb *CardTypeCreateBulk) Exec(ctx context.Context) error {
	_, err := ctcb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (ctcb *CardTypeCreateBulk) ExecX(ctx context.Context) {
	if err := ctcb.Exec(ctx); err != nil {
		panic(err)
	}
}

// OnConflict allows configuring the `ON CONFLICT` / `ON DUPLICATE KEY` clause
// of the `INSERT` statement. For example:
//
//	client.CardType.CreateBulk(builders...).
//		OnConflict(
//			// Update the row with the new values
//			// the was proposed for insertion.
//			sql.ResolveWithNewValues(),
//		).
//		// Override some of the fields with custom
//		// update values.
//		Update(func(u *ent.CardTypeUpsert) {
//			SetKey(v+v).
//		}).
//		Exec(ctx)
func (ctcb *CardTypeCreateBulk) OnConflict(opts ...sql.ConflictOption) *CardTypeUpsertBulk {
	ctcb.conflict = opts
	return &CardTypeUpsertBulk{
		create: ctcb,
	}
}

// OnConflictColumns calls `OnConflict` and configures the columns
// as conflict target. Using this option is equivalent to using:
//
//	client.CardType.Create().
//		OnConflict(sql.ConflictColumns(columns...)).
//		Exec(ctx)
func (ctcb *CardTypeCreateBulk) OnConflictColumns(columns ...string) *CardTypeUpsertBulk {
	ctcb.conflict = append(ctcb.conflict, sql.ConflictColumns(columns...))
	return &CardTypeUpsertBulk{
		create: ctcb,
	}
}

// CardTypeUpsertBulk is the builder for "upsert"-ing
// a bulk of CardType nodes.
type CardTypeUpsertBulk struct {
	create *CardTypeCreateBulk
}

// UpdateNewValues updates the mutable fields using the new values that
// were set on create. Using this option is equivalent to using:
//
//	client.CardType.Create().
//		OnConflict(
//			sql.ResolveWithNewValues(),
//			sql.ResolveWith(func(u *sql.UpdateSet) {
//				u.SetIgnore(cardtype.FieldID)
//			}),
//		).
//		Exec(ctx)
func (u *CardTypeUpsertBulk) UpdateNewValues() *CardTypeUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithNewValues())
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(s *sql.UpdateSet) {
		for _, b := range u.create.builders {
			if _, exists := b.mutation.ID(); exists {
				s.SetIgnore(cardtype.FieldID)
			}
			if _, exists := b.mutation.Key(); exists {
				s.SetIgnore(cardtype.FieldKey)
			}
		}
	}))
	return u
}

// Ignore sets each column to itself in case of conflict.
// Using this option is equivalent to using:
//
//	client.CardType.Create().
//		OnConflict(sql.ResolveWithIgnore()).
//		Exec(ctx)
func (u *CardTypeUpsertBulk) Ignore() *CardTypeUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithIgnore())
	return u
}

// DoNothing configures the conflict_action to `DO NOTHING`.
// Supported only by SQLite and PostgreSQL.
func (u *CardTypeUpsertBulk) DoNothing() *CardTypeUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.DoNothing())
	return u
}

// Update allows overriding fields `UPDATE` values. See the CardTypeCreateBulk.OnConflict
// documentation for more info.
func (u *CardTypeUpsertBulk) Update(set func(*CardTypeUpsert)) *CardTypeUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(update *sql.UpdateSet) {
		set(&CardTypeUpsert{UpdateSet: update})
	}))
	return u
}

// SetNameJa sets the "name_ja" field.
func (u *CardTypeUpsertBulk) SetNameJa(v string) *CardTypeUpsertBulk {
	return u.Update(func(s *CardTypeUpsert) {
		s.SetNameJa(v)
	})
}

// UpdateNameJa sets the "name_ja" field to the value that was provided on create.
func (u *CardTypeUpsertBulk) UpdateNameJa() *CardTypeUpsertBulk {
	return u.Update(func(s *CardTypeUpsert) {
		s.UpdateNameJa()
	})
}

// ClearNameJa clears the value of the "name_ja" field.
func (u *CardTypeUpsertBulk) ClearNameJa() *CardTypeUpsertBulk {
	return u.Update(func(s *CardTypeUpsert) {
		s.ClearNameJa()
	})
}

// SetNameEn sets the "name_en" field.
func (u *CardTypeUpsertBulk) SetNameEn(v string) *CardTypeUpsertBulk {
	return u.Update(func(s *CardTypeUpsert) {
		s.SetNameEn(v)
	})
}

// UpdateNameEn sets the "name_en" field to the value that was provided on create.
func (u *CardTypeUpsertBulk) UpdateNameEn() *CardTypeUpsertBulk {
	return u.Update(func(s *CardTypeUpsert) {
		s.UpdateNameEn()
	})
}

// ClearNameEn clears the value of the "name_en" field.
func (u *CardTypeUpsertBulk) ClearNameEn() *CardTypeUpsertBulk {
	return u.Update(func(s *CardTypeUpsert) {
		s.ClearNameEn()
	})
}

// Exec executes the query.
func (u *CardTypeUpsertBulk) Exec(ctx context.Context) error {
	if u.create.err != nil {
		return u.create.err
	}
	for i, b := range u.create.builders {
		if len(b.conflict) != 0 {
			return fmt.Errorf("ent: OnConflict was set for builder %d. Set it on the CardTypeCreateBulk instead", i)
		}
	}
	if len(u.create.conflict) == 0 {
		return errors.New("ent: missing options for CardTypeCreateBulk.OnConflict")
	}
	return u.create.Exec(ctx)
}

// ExecX is like Exec, but panics if an error occurs.
func (u *CardTypeUpsertBulk) ExecX(ctx context.Context) {
	if err := u.create.Exec(ctx); err != nil {
		panic(err)
	}
}

// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"

	"entgo.io/ent/dialect"
	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/open-boardgame-stats/backend/internal/ent/match"
	"github.com/open-boardgame-stats/backend/internal/ent/player"
	"github.com/open-boardgame-stats/backend/internal/ent/schema/guidgql"
	"github.com/open-boardgame-stats/backend/internal/ent/statdescription"
	"github.com/open-boardgame-stats/backend/internal/ent/statistic"
)

// StatisticCreate is the builder for creating a Statistic entity.
type StatisticCreate struct {
	config
	mutation *StatisticMutation
	hooks    []Hook
	conflict []sql.ConflictOption
}

// SetValue sets the "value" field.
func (sc *StatisticCreate) SetValue(s string) *StatisticCreate {
	sc.mutation.SetValue(s)
	return sc
}

// SetNillableValue sets the "value" field if the given value is not nil.
func (sc *StatisticCreate) SetNillableValue(s *string) *StatisticCreate {
	if s != nil {
		sc.SetValue(*s)
	}
	return sc
}

// SetID sets the "id" field.
func (sc *StatisticCreate) SetID(gu guidgql.GUID) *StatisticCreate {
	sc.mutation.SetID(gu)
	return sc
}

// SetNillableID sets the "id" field if the given value is not nil.
func (sc *StatisticCreate) SetNillableID(gu *guidgql.GUID) *StatisticCreate {
	if gu != nil {
		sc.SetID(*gu)
	}
	return sc
}

// SetMatchID sets the "match" edge to the Match entity by ID.
func (sc *StatisticCreate) SetMatchID(id guidgql.GUID) *StatisticCreate {
	sc.mutation.SetMatchID(id)
	return sc
}

// SetMatch sets the "match" edge to the Match entity.
func (sc *StatisticCreate) SetMatch(m *Match) *StatisticCreate {
	return sc.SetMatchID(m.ID)
}

// SetStatDescriptionID sets the "stat_description" edge to the StatDescription entity by ID.
func (sc *StatisticCreate) SetStatDescriptionID(id guidgql.GUID) *StatisticCreate {
	sc.mutation.SetStatDescriptionID(id)
	return sc
}

// SetStatDescription sets the "stat_description" edge to the StatDescription entity.
func (sc *StatisticCreate) SetStatDescription(s *StatDescription) *StatisticCreate {
	return sc.SetStatDescriptionID(s.ID)
}

// SetPlayerID sets the "player" edge to the Player entity by ID.
func (sc *StatisticCreate) SetPlayerID(id guidgql.GUID) *StatisticCreate {
	sc.mutation.SetPlayerID(id)
	return sc
}

// SetPlayer sets the "player" edge to the Player entity.
func (sc *StatisticCreate) SetPlayer(p *Player) *StatisticCreate {
	return sc.SetPlayerID(p.ID)
}

// Mutation returns the StatisticMutation object of the builder.
func (sc *StatisticCreate) Mutation() *StatisticMutation {
	return sc.mutation
}

// Save creates the Statistic in the database.
func (sc *StatisticCreate) Save(ctx context.Context) (*Statistic, error) {
	sc.defaults()
	return withHooks[*Statistic, StatisticMutation](ctx, sc.sqlSave, sc.mutation, sc.hooks)
}

// SaveX calls Save and panics if Save returns an error.
func (sc *StatisticCreate) SaveX(ctx context.Context) *Statistic {
	v, err := sc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (sc *StatisticCreate) Exec(ctx context.Context) error {
	_, err := sc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (sc *StatisticCreate) ExecX(ctx context.Context) {
	if err := sc.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (sc *StatisticCreate) defaults() {
	if _, ok := sc.mutation.Value(); !ok {
		v := statistic.DefaultValue
		sc.mutation.SetValue(v)
	}
	if _, ok := sc.mutation.ID(); !ok {
		v := statistic.DefaultID()
		sc.mutation.SetID(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (sc *StatisticCreate) check() error {
	if _, ok := sc.mutation.Value(); !ok {
		return &ValidationError{Name: "value", err: errors.New(`ent: missing required field "Statistic.value"`)}
	}
	if _, ok := sc.mutation.MatchID(); !ok {
		return &ValidationError{Name: "match", err: errors.New(`ent: missing required edge "Statistic.match"`)}
	}
	if _, ok := sc.mutation.StatDescriptionID(); !ok {
		return &ValidationError{Name: "stat_description", err: errors.New(`ent: missing required edge "Statistic.stat_description"`)}
	}
	if _, ok := sc.mutation.PlayerID(); !ok {
		return &ValidationError{Name: "player", err: errors.New(`ent: missing required edge "Statistic.player"`)}
	}
	return nil
}

func (sc *StatisticCreate) sqlSave(ctx context.Context) (*Statistic, error) {
	if err := sc.check(); err != nil {
		return nil, err
	}
	_node, _spec := sc.createSpec()
	if err := sqlgraph.CreateNode(ctx, sc.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	if _spec.ID.Value != nil {
		if id, ok := _spec.ID.Value.(*guidgql.GUID); ok {
			_node.ID = *id
		} else if err := _node.ID.Scan(_spec.ID.Value); err != nil {
			return nil, err
		}
	}
	sc.mutation.id = &_node.ID
	sc.mutation.done = true
	return _node, nil
}

func (sc *StatisticCreate) createSpec() (*Statistic, *sqlgraph.CreateSpec) {
	var (
		_node = &Statistic{config: sc.config}
		_spec = sqlgraph.NewCreateSpec(statistic.Table, sqlgraph.NewFieldSpec(statistic.FieldID, field.TypeString))
	)
	_spec.OnConflict = sc.conflict
	if id, ok := sc.mutation.ID(); ok {
		_node.ID = id
		_spec.ID.Value = &id
	}
	if value, ok := sc.mutation.Value(); ok {
		_spec.SetField(statistic.FieldValue, field.TypeString, value)
		_node.Value = value
	}
	if nodes := sc.mutation.MatchIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   statistic.MatchTable,
			Columns: []string{statistic.MatchColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(match.FieldID, field.TypeString),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_node.match_stats = &nodes[0]
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := sc.mutation.StatDescriptionIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   statistic.StatDescriptionTable,
			Columns: []string{statistic.StatDescriptionColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(statdescription.FieldID, field.TypeString),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_node.stat_description_stats = &nodes[0]
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := sc.mutation.PlayerIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   statistic.PlayerTable,
			Columns: []string{statistic.PlayerColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(player.FieldID, field.TypeString),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_node.player_stats = &nodes[0]
		_spec.Edges = append(_spec.Edges, edge)
	}
	return _node, _spec
}

// OnConflict allows configuring the `ON CONFLICT` / `ON DUPLICATE KEY` clause
// of the `INSERT` statement. For example:
//
//	client.Statistic.Create().
//		SetValue(v).
//		OnConflict(
//			// Update the row with the new values
//			// the was proposed for insertion.
//			sql.ResolveWithNewValues(),
//		).
//		// Override some of the fields with custom
//		// update values.
//		Update(func(u *ent.StatisticUpsert) {
//			SetValue(v+v).
//		}).
//		Exec(ctx)
func (sc *StatisticCreate) OnConflict(opts ...sql.ConflictOption) *StatisticUpsertOne {
	sc.conflict = opts
	return &StatisticUpsertOne{
		create: sc,
	}
}

// OnConflictColumns calls `OnConflict` and configures the columns
// as conflict target. Using this option is equivalent to using:
//
//	client.Statistic.Create().
//		OnConflict(sql.ConflictColumns(columns...)).
//		Exec(ctx)
func (sc *StatisticCreate) OnConflictColumns(columns ...string) *StatisticUpsertOne {
	sc.conflict = append(sc.conflict, sql.ConflictColumns(columns...))
	return &StatisticUpsertOne{
		create: sc,
	}
}

type (
	// StatisticUpsertOne is the builder for "upsert"-ing
	//  one Statistic node.
	StatisticUpsertOne struct {
		create *StatisticCreate
	}

	// StatisticUpsert is the "OnConflict" setter.
	StatisticUpsert struct {
		*sql.UpdateSet
	}
)

// SetValue sets the "value" field.
func (u *StatisticUpsert) SetValue(v string) *StatisticUpsert {
	u.Set(statistic.FieldValue, v)
	return u
}

// UpdateValue sets the "value" field to the value that was provided on create.
func (u *StatisticUpsert) UpdateValue() *StatisticUpsert {
	u.SetExcluded(statistic.FieldValue)
	return u
}

// UpdateNewValues updates the mutable fields using the new values that were set on create except the ID field.
// Using this option is equivalent to using:
//
//	client.Statistic.Create().
//		OnConflict(
//			sql.ResolveWithNewValues(),
//			sql.ResolveWith(func(u *sql.UpdateSet) {
//				u.SetIgnore(statistic.FieldID)
//			}),
//		).
//		Exec(ctx)
func (u *StatisticUpsertOne) UpdateNewValues() *StatisticUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithNewValues())
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(s *sql.UpdateSet) {
		if _, exists := u.create.mutation.ID(); exists {
			s.SetIgnore(statistic.FieldID)
		}
	}))
	return u
}

// Ignore sets each column to itself in case of conflict.
// Using this option is equivalent to using:
//
//	client.Statistic.Create().
//	    OnConflict(sql.ResolveWithIgnore()).
//	    Exec(ctx)
func (u *StatisticUpsertOne) Ignore() *StatisticUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithIgnore())
	return u
}

// DoNothing configures the conflict_action to `DO NOTHING`.
// Supported only by SQLite and PostgreSQL.
func (u *StatisticUpsertOne) DoNothing() *StatisticUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.DoNothing())
	return u
}

// Update allows overriding fields `UPDATE` values. See the StatisticCreate.OnConflict
// documentation for more info.
func (u *StatisticUpsertOne) Update(set func(*StatisticUpsert)) *StatisticUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(update *sql.UpdateSet) {
		set(&StatisticUpsert{UpdateSet: update})
	}))
	return u
}

// SetValue sets the "value" field.
func (u *StatisticUpsertOne) SetValue(v string) *StatisticUpsertOne {
	return u.Update(func(s *StatisticUpsert) {
		s.SetValue(v)
	})
}

// UpdateValue sets the "value" field to the value that was provided on create.
func (u *StatisticUpsertOne) UpdateValue() *StatisticUpsertOne {
	return u.Update(func(s *StatisticUpsert) {
		s.UpdateValue()
	})
}

// Exec executes the query.
func (u *StatisticUpsertOne) Exec(ctx context.Context) error {
	if len(u.create.conflict) == 0 {
		return errors.New("ent: missing options for StatisticCreate.OnConflict")
	}
	return u.create.Exec(ctx)
}

// ExecX is like Exec, but panics if an error occurs.
func (u *StatisticUpsertOne) ExecX(ctx context.Context) {
	if err := u.create.Exec(ctx); err != nil {
		panic(err)
	}
}

// Exec executes the UPSERT query and returns the inserted/updated ID.
func (u *StatisticUpsertOne) ID(ctx context.Context) (id guidgql.GUID, err error) {
	if u.create.driver.Dialect() == dialect.MySQL {
		// In case of "ON CONFLICT", there is no way to get back non-numeric ID
		// fields from the database since MySQL does not support the RETURNING clause.
		return id, errors.New("ent: StatisticUpsertOne.ID is not supported by MySQL driver. Use StatisticUpsertOne.Exec instead")
	}
	node, err := u.create.Save(ctx)
	if err != nil {
		return id, err
	}
	return node.ID, nil
}

// IDX is like ID, but panics if an error occurs.
func (u *StatisticUpsertOne) IDX(ctx context.Context) guidgql.GUID {
	id, err := u.ID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// StatisticCreateBulk is the builder for creating many Statistic entities in bulk.
type StatisticCreateBulk struct {
	config
	builders []*StatisticCreate
	conflict []sql.ConflictOption
}

// Save creates the Statistic entities in the database.
func (scb *StatisticCreateBulk) Save(ctx context.Context) ([]*Statistic, error) {
	specs := make([]*sqlgraph.CreateSpec, len(scb.builders))
	nodes := make([]*Statistic, len(scb.builders))
	mutators := make([]Mutator, len(scb.builders))
	for i := range scb.builders {
		func(i int, root context.Context) {
			builder := scb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*StatisticMutation)
				if !ok {
					return nil, fmt.Errorf("unexpected mutation type %T", m)
				}
				if err := builder.check(); err != nil {
					return nil, err
				}
				builder.mutation = mutation
				nodes[i], specs[i] = builder.createSpec()
				var err error
				if i < len(mutators)-1 {
					_, err = mutators[i+1].Mutate(root, scb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					spec.OnConflict = scb.conflict
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, scb.driver, spec); err != nil {
						if sqlgraph.IsConstraintError(err) {
							err = &ConstraintError{msg: err.Error(), wrap: err}
						}
					}
				}
				if err != nil {
					return nil, err
				}
				mutation.id = &nodes[i].ID
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
		if _, err := mutators[0].Mutate(ctx, scb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (scb *StatisticCreateBulk) SaveX(ctx context.Context) []*Statistic {
	v, err := scb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (scb *StatisticCreateBulk) Exec(ctx context.Context) error {
	_, err := scb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (scb *StatisticCreateBulk) ExecX(ctx context.Context) {
	if err := scb.Exec(ctx); err != nil {
		panic(err)
	}
}

// OnConflict allows configuring the `ON CONFLICT` / `ON DUPLICATE KEY` clause
// of the `INSERT` statement. For example:
//
//	client.Statistic.CreateBulk(builders...).
//		OnConflict(
//			// Update the row with the new values
//			// the was proposed for insertion.
//			sql.ResolveWithNewValues(),
//		).
//		// Override some of the fields with custom
//		// update values.
//		Update(func(u *ent.StatisticUpsert) {
//			SetValue(v+v).
//		}).
//		Exec(ctx)
func (scb *StatisticCreateBulk) OnConflict(opts ...sql.ConflictOption) *StatisticUpsertBulk {
	scb.conflict = opts
	return &StatisticUpsertBulk{
		create: scb,
	}
}

// OnConflictColumns calls `OnConflict` and configures the columns
// as conflict target. Using this option is equivalent to using:
//
//	client.Statistic.Create().
//		OnConflict(sql.ConflictColumns(columns...)).
//		Exec(ctx)
func (scb *StatisticCreateBulk) OnConflictColumns(columns ...string) *StatisticUpsertBulk {
	scb.conflict = append(scb.conflict, sql.ConflictColumns(columns...))
	return &StatisticUpsertBulk{
		create: scb,
	}
}

// StatisticUpsertBulk is the builder for "upsert"-ing
// a bulk of Statistic nodes.
type StatisticUpsertBulk struct {
	create *StatisticCreateBulk
}

// UpdateNewValues updates the mutable fields using the new values that
// were set on create. Using this option is equivalent to using:
//
//	client.Statistic.Create().
//		OnConflict(
//			sql.ResolveWithNewValues(),
//			sql.ResolveWith(func(u *sql.UpdateSet) {
//				u.SetIgnore(statistic.FieldID)
//			}),
//		).
//		Exec(ctx)
func (u *StatisticUpsertBulk) UpdateNewValues() *StatisticUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithNewValues())
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(s *sql.UpdateSet) {
		for _, b := range u.create.builders {
			if _, exists := b.mutation.ID(); exists {
				s.SetIgnore(statistic.FieldID)
			}
		}
	}))
	return u
}

// Ignore sets each column to itself in case of conflict.
// Using this option is equivalent to using:
//
//	client.Statistic.Create().
//		OnConflict(sql.ResolveWithIgnore()).
//		Exec(ctx)
func (u *StatisticUpsertBulk) Ignore() *StatisticUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithIgnore())
	return u
}

// DoNothing configures the conflict_action to `DO NOTHING`.
// Supported only by SQLite and PostgreSQL.
func (u *StatisticUpsertBulk) DoNothing() *StatisticUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.DoNothing())
	return u
}

// Update allows overriding fields `UPDATE` values. See the StatisticCreateBulk.OnConflict
// documentation for more info.
func (u *StatisticUpsertBulk) Update(set func(*StatisticUpsert)) *StatisticUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(update *sql.UpdateSet) {
		set(&StatisticUpsert{UpdateSet: update})
	}))
	return u
}

// SetValue sets the "value" field.
func (u *StatisticUpsertBulk) SetValue(v string) *StatisticUpsertBulk {
	return u.Update(func(s *StatisticUpsert) {
		s.SetValue(v)
	})
}

// UpdateValue sets the "value" field to the value that was provided on create.
func (u *StatisticUpsertBulk) UpdateValue() *StatisticUpsertBulk {
	return u.Update(func(s *StatisticUpsert) {
		s.UpdateValue()
	})
}

// Exec executes the query.
func (u *StatisticUpsertBulk) Exec(ctx context.Context) error {
	for i, b := range u.create.builders {
		if len(b.conflict) != 0 {
			return fmt.Errorf("ent: OnConflict was set for builder %d. Set it on the StatisticCreateBulk instead", i)
		}
	}
	if len(u.create.conflict) == 0 {
		return errors.New("ent: missing options for StatisticCreateBulk.OnConflict")
	}
	return u.create.Exec(ctx)
}

// ExecX is like Exec, but panics if an error occurs.
func (u *StatisticUpsertBulk) ExecX(ctx context.Context) {
	if err := u.create.Exec(ctx); err != nil {
		panic(err)
	}
}

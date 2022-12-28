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
	"github.com/open-boardgame-stats/backend/internal/ent/playersupervisionrequest"
	"github.com/open-boardgame-stats/backend/internal/ent/schema/guidgql"
	"github.com/open-boardgame-stats/backend/internal/ent/statistic"
	"github.com/open-boardgame-stats/backend/internal/ent/user"
)

// PlayerCreate is the builder for creating a Player entity.
type PlayerCreate struct {
	config
	mutation *PlayerMutation
	hooks    []Hook
	conflict []sql.ConflictOption
}

// SetName sets the "name" field.
func (pc *PlayerCreate) SetName(s string) *PlayerCreate {
	pc.mutation.SetName(s)
	return pc
}

// SetNillableName sets the "name" field if the given value is not nil.
func (pc *PlayerCreate) SetNillableName(s *string) *PlayerCreate {
	if s != nil {
		pc.SetName(*s)
	}
	return pc
}

// SetID sets the "id" field.
func (pc *PlayerCreate) SetID(gu guidgql.GUID) *PlayerCreate {
	pc.mutation.SetID(gu)
	return pc
}

// SetNillableID sets the "id" field if the given value is not nil.
func (pc *PlayerCreate) SetNillableID(gu *guidgql.GUID) *PlayerCreate {
	if gu != nil {
		pc.SetID(*gu)
	}
	return pc
}

// SetOwnerID sets the "owner" edge to the User entity by ID.
func (pc *PlayerCreate) SetOwnerID(id guidgql.GUID) *PlayerCreate {
	pc.mutation.SetOwnerID(id)
	return pc
}

// SetNillableOwnerID sets the "owner" edge to the User entity by ID if the given value is not nil.
func (pc *PlayerCreate) SetNillableOwnerID(id *guidgql.GUID) *PlayerCreate {
	if id != nil {
		pc = pc.SetOwnerID(*id)
	}
	return pc
}

// SetOwner sets the "owner" edge to the User entity.
func (pc *PlayerCreate) SetOwner(u *User) *PlayerCreate {
	return pc.SetOwnerID(u.ID)
}

// AddSupervisorIDs adds the "supervisors" edge to the User entity by IDs.
func (pc *PlayerCreate) AddSupervisorIDs(ids ...guidgql.GUID) *PlayerCreate {
	pc.mutation.AddSupervisorIDs(ids...)
	return pc
}

// AddSupervisors adds the "supervisors" edges to the User entity.
func (pc *PlayerCreate) AddSupervisors(u ...*User) *PlayerCreate {
	ids := make([]guidgql.GUID, len(u))
	for i := range u {
		ids[i] = u[i].ID
	}
	return pc.AddSupervisorIDs(ids...)
}

// AddSupervisionRequestIDs adds the "supervision_requests" edge to the PlayerSupervisionRequest entity by IDs.
func (pc *PlayerCreate) AddSupervisionRequestIDs(ids ...guidgql.GUID) *PlayerCreate {
	pc.mutation.AddSupervisionRequestIDs(ids...)
	return pc
}

// AddSupervisionRequests adds the "supervision_requests" edges to the PlayerSupervisionRequest entity.
func (pc *PlayerCreate) AddSupervisionRequests(p ...*PlayerSupervisionRequest) *PlayerCreate {
	ids := make([]guidgql.GUID, len(p))
	for i := range p {
		ids[i] = p[i].ID
	}
	return pc.AddSupervisionRequestIDs(ids...)
}

// AddMatchIDs adds the "matches" edge to the Match entity by IDs.
func (pc *PlayerCreate) AddMatchIDs(ids ...guidgql.GUID) *PlayerCreate {
	pc.mutation.AddMatchIDs(ids...)
	return pc
}

// AddMatches adds the "matches" edges to the Match entity.
func (pc *PlayerCreate) AddMatches(m ...*Match) *PlayerCreate {
	ids := make([]guidgql.GUID, len(m))
	for i := range m {
		ids[i] = m[i].ID
	}
	return pc.AddMatchIDs(ids...)
}

// AddStatIDs adds the "stats" edge to the Statistic entity by IDs.
func (pc *PlayerCreate) AddStatIDs(ids ...guidgql.GUID) *PlayerCreate {
	pc.mutation.AddStatIDs(ids...)
	return pc
}

// AddStats adds the "stats" edges to the Statistic entity.
func (pc *PlayerCreate) AddStats(s ...*Statistic) *PlayerCreate {
	ids := make([]guidgql.GUID, len(s))
	for i := range s {
		ids[i] = s[i].ID
	}
	return pc.AddStatIDs(ids...)
}

// Mutation returns the PlayerMutation object of the builder.
func (pc *PlayerCreate) Mutation() *PlayerMutation {
	return pc.mutation
}

// Save creates the Player in the database.
func (pc *PlayerCreate) Save(ctx context.Context) (*Player, error) {
	var (
		err  error
		node *Player
	)
	pc.defaults()
	if len(pc.hooks) == 0 {
		if err = pc.check(); err != nil {
			return nil, err
		}
		node, err = pc.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*PlayerMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = pc.check(); err != nil {
				return nil, err
			}
			pc.mutation = mutation
			if node, err = pc.sqlSave(ctx); err != nil {
				return nil, err
			}
			mutation.id = &node.ID
			mutation.done = true
			return node, err
		})
		for i := len(pc.hooks) - 1; i >= 0; i-- {
			if pc.hooks[i] == nil {
				return nil, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = pc.hooks[i](mut)
		}
		v, err := mut.Mutate(ctx, pc.mutation)
		if err != nil {
			return nil, err
		}
		nv, ok := v.(*Player)
		if !ok {
			return nil, fmt.Errorf("unexpected node type %T returned from PlayerMutation", v)
		}
		node = nv
	}
	return node, err
}

// SaveX calls Save and panics if Save returns an error.
func (pc *PlayerCreate) SaveX(ctx context.Context) *Player {
	v, err := pc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (pc *PlayerCreate) Exec(ctx context.Context) error {
	_, err := pc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (pc *PlayerCreate) ExecX(ctx context.Context) {
	if err := pc.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (pc *PlayerCreate) defaults() {
	if _, ok := pc.mutation.Name(); !ok {
		v := player.DefaultName
		pc.mutation.SetName(v)
	}
	if _, ok := pc.mutation.ID(); !ok {
		v := player.DefaultID()
		pc.mutation.SetID(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (pc *PlayerCreate) check() error {
	if _, ok := pc.mutation.Name(); !ok {
		return &ValidationError{Name: "name", err: errors.New(`ent: missing required field "Player.name"`)}
	}
	return nil
}

func (pc *PlayerCreate) sqlSave(ctx context.Context) (*Player, error) {
	_node, _spec := pc.createSpec()
	if err := sqlgraph.CreateNode(ctx, pc.driver, _spec); err != nil {
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
	return _node, nil
}

func (pc *PlayerCreate) createSpec() (*Player, *sqlgraph.CreateSpec) {
	var (
		_node = &Player{config: pc.config}
		_spec = &sqlgraph.CreateSpec{
			Table: player.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeString,
				Column: player.FieldID,
			},
		}
	)
	_spec.OnConflict = pc.conflict
	if id, ok := pc.mutation.ID(); ok {
		_node.ID = id
		_spec.ID.Value = &id
	}
	if value, ok := pc.mutation.Name(); ok {
		_spec.SetField(player.FieldName, field.TypeString, value)
		_node.Name = value
	}
	if nodes := pc.mutation.OwnerIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2O,
			Inverse: true,
			Table:   player.OwnerTable,
			Columns: []string{player.OwnerColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeString,
					Column: user.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_node.user_main_player = &nodes[0]
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := pc.mutation.SupervisorsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: true,
			Table:   player.SupervisorsTable,
			Columns: player.SupervisorsPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeString,
					Column: user.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := pc.mutation.SupervisionRequestsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   player.SupervisionRequestsTable,
			Columns: []string{player.SupervisionRequestsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeString,
					Column: playersupervisionrequest.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := pc.mutation.MatchesIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: true,
			Table:   player.MatchesTable,
			Columns: player.MatchesPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeString,
					Column: match.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := pc.mutation.StatsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   player.StatsTable,
			Columns: []string{player.StatsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeString,
					Column: statistic.FieldID,
				},
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
//	client.Player.Create().
//		SetName(v).
//		OnConflict(
//			// Update the row with the new values
//			// the was proposed for insertion.
//			sql.ResolveWithNewValues(),
//		).
//		// Override some of the fields with custom
//		// update values.
//		Update(func(u *ent.PlayerUpsert) {
//			SetName(v+v).
//		}).
//		Exec(ctx)
func (pc *PlayerCreate) OnConflict(opts ...sql.ConflictOption) *PlayerUpsertOne {
	pc.conflict = opts
	return &PlayerUpsertOne{
		create: pc,
	}
}

// OnConflictColumns calls `OnConflict` and configures the columns
// as conflict target. Using this option is equivalent to using:
//
//	client.Player.Create().
//		OnConflict(sql.ConflictColumns(columns...)).
//		Exec(ctx)
func (pc *PlayerCreate) OnConflictColumns(columns ...string) *PlayerUpsertOne {
	pc.conflict = append(pc.conflict, sql.ConflictColumns(columns...))
	return &PlayerUpsertOne{
		create: pc,
	}
}

type (
	// PlayerUpsertOne is the builder for "upsert"-ing
	//  one Player node.
	PlayerUpsertOne struct {
		create *PlayerCreate
	}

	// PlayerUpsert is the "OnConflict" setter.
	PlayerUpsert struct {
		*sql.UpdateSet
	}
)

// SetName sets the "name" field.
func (u *PlayerUpsert) SetName(v string) *PlayerUpsert {
	u.Set(player.FieldName, v)
	return u
}

// UpdateName sets the "name" field to the value that was provided on create.
func (u *PlayerUpsert) UpdateName() *PlayerUpsert {
	u.SetExcluded(player.FieldName)
	return u
}

// UpdateNewValues updates the mutable fields using the new values that were set on create except the ID field.
// Using this option is equivalent to using:
//
//	client.Player.Create().
//		OnConflict(
//			sql.ResolveWithNewValues(),
//			sql.ResolveWith(func(u *sql.UpdateSet) {
//				u.SetIgnore(player.FieldID)
//			}),
//		).
//		Exec(ctx)
func (u *PlayerUpsertOne) UpdateNewValues() *PlayerUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithNewValues())
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(s *sql.UpdateSet) {
		if _, exists := u.create.mutation.ID(); exists {
			s.SetIgnore(player.FieldID)
		}
	}))
	return u
}

// Ignore sets each column to itself in case of conflict.
// Using this option is equivalent to using:
//
//	client.Player.Create().
//	    OnConflict(sql.ResolveWithIgnore()).
//	    Exec(ctx)
func (u *PlayerUpsertOne) Ignore() *PlayerUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithIgnore())
	return u
}

// DoNothing configures the conflict_action to `DO NOTHING`.
// Supported only by SQLite and PostgreSQL.
func (u *PlayerUpsertOne) DoNothing() *PlayerUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.DoNothing())
	return u
}

// Update allows overriding fields `UPDATE` values. See the PlayerCreate.OnConflict
// documentation for more info.
func (u *PlayerUpsertOne) Update(set func(*PlayerUpsert)) *PlayerUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(update *sql.UpdateSet) {
		set(&PlayerUpsert{UpdateSet: update})
	}))
	return u
}

// SetName sets the "name" field.
func (u *PlayerUpsertOne) SetName(v string) *PlayerUpsertOne {
	return u.Update(func(s *PlayerUpsert) {
		s.SetName(v)
	})
}

// UpdateName sets the "name" field to the value that was provided on create.
func (u *PlayerUpsertOne) UpdateName() *PlayerUpsertOne {
	return u.Update(func(s *PlayerUpsert) {
		s.UpdateName()
	})
}

// Exec executes the query.
func (u *PlayerUpsertOne) Exec(ctx context.Context) error {
	if len(u.create.conflict) == 0 {
		return errors.New("ent: missing options for PlayerCreate.OnConflict")
	}
	return u.create.Exec(ctx)
}

// ExecX is like Exec, but panics if an error occurs.
func (u *PlayerUpsertOne) ExecX(ctx context.Context) {
	if err := u.create.Exec(ctx); err != nil {
		panic(err)
	}
}

// Exec executes the UPSERT query and returns the inserted/updated ID.
func (u *PlayerUpsertOne) ID(ctx context.Context) (id guidgql.GUID, err error) {
	if u.create.driver.Dialect() == dialect.MySQL {
		// In case of "ON CONFLICT", there is no way to get back non-numeric ID
		// fields from the database since MySQL does not support the RETURNING clause.
		return id, errors.New("ent: PlayerUpsertOne.ID is not supported by MySQL driver. Use PlayerUpsertOne.Exec instead")
	}
	node, err := u.create.Save(ctx)
	if err != nil {
		return id, err
	}
	return node.ID, nil
}

// IDX is like ID, but panics if an error occurs.
func (u *PlayerUpsertOne) IDX(ctx context.Context) guidgql.GUID {
	id, err := u.ID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// PlayerCreateBulk is the builder for creating many Player entities in bulk.
type PlayerCreateBulk struct {
	config
	builders []*PlayerCreate
	conflict []sql.ConflictOption
}

// Save creates the Player entities in the database.
func (pcb *PlayerCreateBulk) Save(ctx context.Context) ([]*Player, error) {
	specs := make([]*sqlgraph.CreateSpec, len(pcb.builders))
	nodes := make([]*Player, len(pcb.builders))
	mutators := make([]Mutator, len(pcb.builders))
	for i := range pcb.builders {
		func(i int, root context.Context) {
			builder := pcb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*PlayerMutation)
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
					_, err = mutators[i+1].Mutate(root, pcb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					spec.OnConflict = pcb.conflict
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, pcb.driver, spec); err != nil {
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
		if _, err := mutators[0].Mutate(ctx, pcb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (pcb *PlayerCreateBulk) SaveX(ctx context.Context) []*Player {
	v, err := pcb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (pcb *PlayerCreateBulk) Exec(ctx context.Context) error {
	_, err := pcb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (pcb *PlayerCreateBulk) ExecX(ctx context.Context) {
	if err := pcb.Exec(ctx); err != nil {
		panic(err)
	}
}

// OnConflict allows configuring the `ON CONFLICT` / `ON DUPLICATE KEY` clause
// of the `INSERT` statement. For example:
//
//	client.Player.CreateBulk(builders...).
//		OnConflict(
//			// Update the row with the new values
//			// the was proposed for insertion.
//			sql.ResolveWithNewValues(),
//		).
//		// Override some of the fields with custom
//		// update values.
//		Update(func(u *ent.PlayerUpsert) {
//			SetName(v+v).
//		}).
//		Exec(ctx)
func (pcb *PlayerCreateBulk) OnConflict(opts ...sql.ConflictOption) *PlayerUpsertBulk {
	pcb.conflict = opts
	return &PlayerUpsertBulk{
		create: pcb,
	}
}

// OnConflictColumns calls `OnConflict` and configures the columns
// as conflict target. Using this option is equivalent to using:
//
//	client.Player.Create().
//		OnConflict(sql.ConflictColumns(columns...)).
//		Exec(ctx)
func (pcb *PlayerCreateBulk) OnConflictColumns(columns ...string) *PlayerUpsertBulk {
	pcb.conflict = append(pcb.conflict, sql.ConflictColumns(columns...))
	return &PlayerUpsertBulk{
		create: pcb,
	}
}

// PlayerUpsertBulk is the builder for "upsert"-ing
// a bulk of Player nodes.
type PlayerUpsertBulk struct {
	create *PlayerCreateBulk
}

// UpdateNewValues updates the mutable fields using the new values that
// were set on create. Using this option is equivalent to using:
//
//	client.Player.Create().
//		OnConflict(
//			sql.ResolveWithNewValues(),
//			sql.ResolveWith(func(u *sql.UpdateSet) {
//				u.SetIgnore(player.FieldID)
//			}),
//		).
//		Exec(ctx)
func (u *PlayerUpsertBulk) UpdateNewValues() *PlayerUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithNewValues())
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(s *sql.UpdateSet) {
		for _, b := range u.create.builders {
			if _, exists := b.mutation.ID(); exists {
				s.SetIgnore(player.FieldID)
			}
		}
	}))
	return u
}

// Ignore sets each column to itself in case of conflict.
// Using this option is equivalent to using:
//
//	client.Player.Create().
//		OnConflict(sql.ResolveWithIgnore()).
//		Exec(ctx)
func (u *PlayerUpsertBulk) Ignore() *PlayerUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithIgnore())
	return u
}

// DoNothing configures the conflict_action to `DO NOTHING`.
// Supported only by SQLite and PostgreSQL.
func (u *PlayerUpsertBulk) DoNothing() *PlayerUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.DoNothing())
	return u
}

// Update allows overriding fields `UPDATE` values. See the PlayerCreateBulk.OnConflict
// documentation for more info.
func (u *PlayerUpsertBulk) Update(set func(*PlayerUpsert)) *PlayerUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(update *sql.UpdateSet) {
		set(&PlayerUpsert{UpdateSet: update})
	}))
	return u
}

// SetName sets the "name" field.
func (u *PlayerUpsertBulk) SetName(v string) *PlayerUpsertBulk {
	return u.Update(func(s *PlayerUpsert) {
		s.SetName(v)
	})
}

// UpdateName sets the "name" field to the value that was provided on create.
func (u *PlayerUpsertBulk) UpdateName() *PlayerUpsertBulk {
	return u.Update(func(s *PlayerUpsert) {
		s.UpdateName()
	})
}

// Exec executes the query.
func (u *PlayerUpsertBulk) Exec(ctx context.Context) error {
	for i, b := range u.create.builders {
		if len(b.conflict) != 0 {
			return fmt.Errorf("ent: OnConflict was set for builder %d. Set it on the PlayerCreateBulk instead", i)
		}
	}
	if len(u.create.conflict) == 0 {
		return errors.New("ent: missing options for PlayerCreateBulk.OnConflict")
	}
	return u.create.Exec(ctx)
}

// ExecX is like Exec, but panics if an error occurs.
func (u *PlayerUpsertBulk) ExecX(ctx context.Context) {
	if err := u.create.Exec(ctx); err != nil {
		panic(err)
	}
}

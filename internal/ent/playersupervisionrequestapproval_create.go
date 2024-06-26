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
	"github.com/open-boardgame-stats/backend/internal/ent/playersupervisionrequest"
	"github.com/open-boardgame-stats/backend/internal/ent/playersupervisionrequestapproval"
	"github.com/open-boardgame-stats/backend/internal/ent/schema/guidgql"
	"github.com/open-boardgame-stats/backend/internal/ent/user"
)

// PlayerSupervisionRequestApprovalCreate is the builder for creating a PlayerSupervisionRequestApproval entity.
type PlayerSupervisionRequestApprovalCreate struct {
	config
	mutation *PlayerSupervisionRequestApprovalMutation
	hooks    []Hook
	conflict []sql.ConflictOption
}

// SetApproved sets the "approved" field.
func (psrac *PlayerSupervisionRequestApprovalCreate) SetApproved(b bool) *PlayerSupervisionRequestApprovalCreate {
	psrac.mutation.SetApproved(b)
	return psrac
}

// SetNillableApproved sets the "approved" field if the given value is not nil.
func (psrac *PlayerSupervisionRequestApprovalCreate) SetNillableApproved(b *bool) *PlayerSupervisionRequestApprovalCreate {
	if b != nil {
		psrac.SetApproved(*b)
	}
	return psrac
}

// SetID sets the "id" field.
func (psrac *PlayerSupervisionRequestApprovalCreate) SetID(gu guidgql.GUID) *PlayerSupervisionRequestApprovalCreate {
	psrac.mutation.SetID(gu)
	return psrac
}

// SetNillableID sets the "id" field if the given value is not nil.
func (psrac *PlayerSupervisionRequestApprovalCreate) SetNillableID(gu *guidgql.GUID) *PlayerSupervisionRequestApprovalCreate {
	if gu != nil {
		psrac.SetID(*gu)
	}
	return psrac
}

// SetApproverID sets the "approver" edge to the User entity by ID.
func (psrac *PlayerSupervisionRequestApprovalCreate) SetApproverID(id guidgql.GUID) *PlayerSupervisionRequestApprovalCreate {
	psrac.mutation.SetApproverID(id)
	return psrac
}

// SetApprover sets the "approver" edge to the User entity.
func (psrac *PlayerSupervisionRequestApprovalCreate) SetApprover(u *User) *PlayerSupervisionRequestApprovalCreate {
	return psrac.SetApproverID(u.ID)
}

// SetSupervisionRequestID sets the "supervision_request" edge to the PlayerSupervisionRequest entity by ID.
func (psrac *PlayerSupervisionRequestApprovalCreate) SetSupervisionRequestID(id guidgql.GUID) *PlayerSupervisionRequestApprovalCreate {
	psrac.mutation.SetSupervisionRequestID(id)
	return psrac
}

// SetSupervisionRequest sets the "supervision_request" edge to the PlayerSupervisionRequest entity.
func (psrac *PlayerSupervisionRequestApprovalCreate) SetSupervisionRequest(p *PlayerSupervisionRequest) *PlayerSupervisionRequestApprovalCreate {
	return psrac.SetSupervisionRequestID(p.ID)
}

// Mutation returns the PlayerSupervisionRequestApprovalMutation object of the builder.
func (psrac *PlayerSupervisionRequestApprovalCreate) Mutation() *PlayerSupervisionRequestApprovalMutation {
	return psrac.mutation
}

// Save creates the PlayerSupervisionRequestApproval in the database.
func (psrac *PlayerSupervisionRequestApprovalCreate) Save(ctx context.Context) (*PlayerSupervisionRequestApproval, error) {
	psrac.defaults()
	return withHooks(ctx, psrac.sqlSave, psrac.mutation, psrac.hooks)
}

// SaveX calls Save and panics if Save returns an error.
func (psrac *PlayerSupervisionRequestApprovalCreate) SaveX(ctx context.Context) *PlayerSupervisionRequestApproval {
	v, err := psrac.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (psrac *PlayerSupervisionRequestApprovalCreate) Exec(ctx context.Context) error {
	_, err := psrac.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (psrac *PlayerSupervisionRequestApprovalCreate) ExecX(ctx context.Context) {
	if err := psrac.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (psrac *PlayerSupervisionRequestApprovalCreate) defaults() {
	if _, ok := psrac.mutation.ID(); !ok {
		v := playersupervisionrequestapproval.DefaultID()
		psrac.mutation.SetID(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (psrac *PlayerSupervisionRequestApprovalCreate) check() error {
	if _, ok := psrac.mutation.ApproverID(); !ok {
		return &ValidationError{Name: "approver", err: errors.New(`ent: missing required edge "PlayerSupervisionRequestApproval.approver"`)}
	}
	if _, ok := psrac.mutation.SupervisionRequestID(); !ok {
		return &ValidationError{Name: "supervision_request", err: errors.New(`ent: missing required edge "PlayerSupervisionRequestApproval.supervision_request"`)}
	}
	return nil
}

func (psrac *PlayerSupervisionRequestApprovalCreate) sqlSave(ctx context.Context) (*PlayerSupervisionRequestApproval, error) {
	if err := psrac.check(); err != nil {
		return nil, err
	}
	_node, _spec := psrac.createSpec()
	if err := sqlgraph.CreateNode(ctx, psrac.driver, _spec); err != nil {
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
	psrac.mutation.id = &_node.ID
	psrac.mutation.done = true
	return _node, nil
}

func (psrac *PlayerSupervisionRequestApprovalCreate) createSpec() (*PlayerSupervisionRequestApproval, *sqlgraph.CreateSpec) {
	var (
		_node = &PlayerSupervisionRequestApproval{config: psrac.config}
		_spec = sqlgraph.NewCreateSpec(playersupervisionrequestapproval.Table, sqlgraph.NewFieldSpec(playersupervisionrequestapproval.FieldID, field.TypeString))
	)
	_spec.OnConflict = psrac.conflict
	if id, ok := psrac.mutation.ID(); ok {
		_node.ID = id
		_spec.ID.Value = &id
	}
	if value, ok := psrac.mutation.Approved(); ok {
		_spec.SetField(playersupervisionrequestapproval.FieldApproved, field.TypeBool, value)
		_node.Approved = &value
	}
	if nodes := psrac.mutation.ApproverIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   playersupervisionrequestapproval.ApproverTable,
			Columns: []string{playersupervisionrequestapproval.ApproverColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(user.FieldID, field.TypeString),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_node.user_supervision_request_approvals = &nodes[0]
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := psrac.mutation.SupervisionRequestIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   playersupervisionrequestapproval.SupervisionRequestTable,
			Columns: []string{playersupervisionrequestapproval.SupervisionRequestColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(playersupervisionrequest.FieldID, field.TypeString),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_node.player_supervision_request_approvals = &nodes[0]
		_spec.Edges = append(_spec.Edges, edge)
	}
	return _node, _spec
}

// OnConflict allows configuring the `ON CONFLICT` / `ON DUPLICATE KEY` clause
// of the `INSERT` statement. For example:
//
//	client.PlayerSupervisionRequestApproval.Create().
//		SetApproved(v).
//		OnConflict(
//			// Update the row with the new values
//			// the was proposed for insertion.
//			sql.ResolveWithNewValues(),
//		).
//		// Override some of the fields with custom
//		// update values.
//		Update(func(u *ent.PlayerSupervisionRequestApprovalUpsert) {
//			SetApproved(v+v).
//		}).
//		Exec(ctx)
func (psrac *PlayerSupervisionRequestApprovalCreate) OnConflict(opts ...sql.ConflictOption) *PlayerSupervisionRequestApprovalUpsertOne {
	psrac.conflict = opts
	return &PlayerSupervisionRequestApprovalUpsertOne{
		create: psrac,
	}
}

// OnConflictColumns calls `OnConflict` and configures the columns
// as conflict target. Using this option is equivalent to using:
//
//	client.PlayerSupervisionRequestApproval.Create().
//		OnConflict(sql.ConflictColumns(columns...)).
//		Exec(ctx)
func (psrac *PlayerSupervisionRequestApprovalCreate) OnConflictColumns(columns ...string) *PlayerSupervisionRequestApprovalUpsertOne {
	psrac.conflict = append(psrac.conflict, sql.ConflictColumns(columns...))
	return &PlayerSupervisionRequestApprovalUpsertOne{
		create: psrac,
	}
}

type (
	// PlayerSupervisionRequestApprovalUpsertOne is the builder for "upsert"-ing
	//  one PlayerSupervisionRequestApproval node.
	PlayerSupervisionRequestApprovalUpsertOne struct {
		create *PlayerSupervisionRequestApprovalCreate
	}

	// PlayerSupervisionRequestApprovalUpsert is the "OnConflict" setter.
	PlayerSupervisionRequestApprovalUpsert struct {
		*sql.UpdateSet
	}
)

// SetApproved sets the "approved" field.
func (u *PlayerSupervisionRequestApprovalUpsert) SetApproved(v bool) *PlayerSupervisionRequestApprovalUpsert {
	u.Set(playersupervisionrequestapproval.FieldApproved, v)
	return u
}

// UpdateApproved sets the "approved" field to the value that was provided on create.
func (u *PlayerSupervisionRequestApprovalUpsert) UpdateApproved() *PlayerSupervisionRequestApprovalUpsert {
	u.SetExcluded(playersupervisionrequestapproval.FieldApproved)
	return u
}

// ClearApproved clears the value of the "approved" field.
func (u *PlayerSupervisionRequestApprovalUpsert) ClearApproved() *PlayerSupervisionRequestApprovalUpsert {
	u.SetNull(playersupervisionrequestapproval.FieldApproved)
	return u
}

// UpdateNewValues updates the mutable fields using the new values that were set on create except the ID field.
// Using this option is equivalent to using:
//
//	client.PlayerSupervisionRequestApproval.Create().
//		OnConflict(
//			sql.ResolveWithNewValues(),
//			sql.ResolveWith(func(u *sql.UpdateSet) {
//				u.SetIgnore(playersupervisionrequestapproval.FieldID)
//			}),
//		).
//		Exec(ctx)
func (u *PlayerSupervisionRequestApprovalUpsertOne) UpdateNewValues() *PlayerSupervisionRequestApprovalUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithNewValues())
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(s *sql.UpdateSet) {
		if _, exists := u.create.mutation.ID(); exists {
			s.SetIgnore(playersupervisionrequestapproval.FieldID)
		}
	}))
	return u
}

// Ignore sets each column to itself in case of conflict.
// Using this option is equivalent to using:
//
//	client.PlayerSupervisionRequestApproval.Create().
//	    OnConflict(sql.ResolveWithIgnore()).
//	    Exec(ctx)
func (u *PlayerSupervisionRequestApprovalUpsertOne) Ignore() *PlayerSupervisionRequestApprovalUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithIgnore())
	return u
}

// DoNothing configures the conflict_action to `DO NOTHING`.
// Supported only by SQLite and PostgreSQL.
func (u *PlayerSupervisionRequestApprovalUpsertOne) DoNothing() *PlayerSupervisionRequestApprovalUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.DoNothing())
	return u
}

// Update allows overriding fields `UPDATE` values. See the PlayerSupervisionRequestApprovalCreate.OnConflict
// documentation for more info.
func (u *PlayerSupervisionRequestApprovalUpsertOne) Update(set func(*PlayerSupervisionRequestApprovalUpsert)) *PlayerSupervisionRequestApprovalUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(update *sql.UpdateSet) {
		set(&PlayerSupervisionRequestApprovalUpsert{UpdateSet: update})
	}))
	return u
}

// SetApproved sets the "approved" field.
func (u *PlayerSupervisionRequestApprovalUpsertOne) SetApproved(v bool) *PlayerSupervisionRequestApprovalUpsertOne {
	return u.Update(func(s *PlayerSupervisionRequestApprovalUpsert) {
		s.SetApproved(v)
	})
}

// UpdateApproved sets the "approved" field to the value that was provided on create.
func (u *PlayerSupervisionRequestApprovalUpsertOne) UpdateApproved() *PlayerSupervisionRequestApprovalUpsertOne {
	return u.Update(func(s *PlayerSupervisionRequestApprovalUpsert) {
		s.UpdateApproved()
	})
}

// ClearApproved clears the value of the "approved" field.
func (u *PlayerSupervisionRequestApprovalUpsertOne) ClearApproved() *PlayerSupervisionRequestApprovalUpsertOne {
	return u.Update(func(s *PlayerSupervisionRequestApprovalUpsert) {
		s.ClearApproved()
	})
}

// Exec executes the query.
func (u *PlayerSupervisionRequestApprovalUpsertOne) Exec(ctx context.Context) error {
	if len(u.create.conflict) == 0 {
		return errors.New("ent: missing options for PlayerSupervisionRequestApprovalCreate.OnConflict")
	}
	return u.create.Exec(ctx)
}

// ExecX is like Exec, but panics if an error occurs.
func (u *PlayerSupervisionRequestApprovalUpsertOne) ExecX(ctx context.Context) {
	if err := u.create.Exec(ctx); err != nil {
		panic(err)
	}
}

// Exec executes the UPSERT query and returns the inserted/updated ID.
func (u *PlayerSupervisionRequestApprovalUpsertOne) ID(ctx context.Context) (id guidgql.GUID, err error) {
	if u.create.driver.Dialect() == dialect.MySQL {
		// In case of "ON CONFLICT", there is no way to get back non-numeric ID
		// fields from the database since MySQL does not support the RETURNING clause.
		return id, errors.New("ent: PlayerSupervisionRequestApprovalUpsertOne.ID is not supported by MySQL driver. Use PlayerSupervisionRequestApprovalUpsertOne.Exec instead")
	}
	node, err := u.create.Save(ctx)
	if err != nil {
		return id, err
	}
	return node.ID, nil
}

// IDX is like ID, but panics if an error occurs.
func (u *PlayerSupervisionRequestApprovalUpsertOne) IDX(ctx context.Context) guidgql.GUID {
	id, err := u.ID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// PlayerSupervisionRequestApprovalCreateBulk is the builder for creating many PlayerSupervisionRequestApproval entities in bulk.
type PlayerSupervisionRequestApprovalCreateBulk struct {
	config
	err      error
	builders []*PlayerSupervisionRequestApprovalCreate
	conflict []sql.ConflictOption
}

// Save creates the PlayerSupervisionRequestApproval entities in the database.
func (psracb *PlayerSupervisionRequestApprovalCreateBulk) Save(ctx context.Context) ([]*PlayerSupervisionRequestApproval, error) {
	if psracb.err != nil {
		return nil, psracb.err
	}
	specs := make([]*sqlgraph.CreateSpec, len(psracb.builders))
	nodes := make([]*PlayerSupervisionRequestApproval, len(psracb.builders))
	mutators := make([]Mutator, len(psracb.builders))
	for i := range psracb.builders {
		func(i int, root context.Context) {
			builder := psracb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*PlayerSupervisionRequestApprovalMutation)
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
					_, err = mutators[i+1].Mutate(root, psracb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					spec.OnConflict = psracb.conflict
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, psracb.driver, spec); err != nil {
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
		if _, err := mutators[0].Mutate(ctx, psracb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (psracb *PlayerSupervisionRequestApprovalCreateBulk) SaveX(ctx context.Context) []*PlayerSupervisionRequestApproval {
	v, err := psracb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (psracb *PlayerSupervisionRequestApprovalCreateBulk) Exec(ctx context.Context) error {
	_, err := psracb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (psracb *PlayerSupervisionRequestApprovalCreateBulk) ExecX(ctx context.Context) {
	if err := psracb.Exec(ctx); err != nil {
		panic(err)
	}
}

// OnConflict allows configuring the `ON CONFLICT` / `ON DUPLICATE KEY` clause
// of the `INSERT` statement. For example:
//
//	client.PlayerSupervisionRequestApproval.CreateBulk(builders...).
//		OnConflict(
//			// Update the row with the new values
//			// the was proposed for insertion.
//			sql.ResolveWithNewValues(),
//		).
//		// Override some of the fields with custom
//		// update values.
//		Update(func(u *ent.PlayerSupervisionRequestApprovalUpsert) {
//			SetApproved(v+v).
//		}).
//		Exec(ctx)
func (psracb *PlayerSupervisionRequestApprovalCreateBulk) OnConflict(opts ...sql.ConflictOption) *PlayerSupervisionRequestApprovalUpsertBulk {
	psracb.conflict = opts
	return &PlayerSupervisionRequestApprovalUpsertBulk{
		create: psracb,
	}
}

// OnConflictColumns calls `OnConflict` and configures the columns
// as conflict target. Using this option is equivalent to using:
//
//	client.PlayerSupervisionRequestApproval.Create().
//		OnConflict(sql.ConflictColumns(columns...)).
//		Exec(ctx)
func (psracb *PlayerSupervisionRequestApprovalCreateBulk) OnConflictColumns(columns ...string) *PlayerSupervisionRequestApprovalUpsertBulk {
	psracb.conflict = append(psracb.conflict, sql.ConflictColumns(columns...))
	return &PlayerSupervisionRequestApprovalUpsertBulk{
		create: psracb,
	}
}

// PlayerSupervisionRequestApprovalUpsertBulk is the builder for "upsert"-ing
// a bulk of PlayerSupervisionRequestApproval nodes.
type PlayerSupervisionRequestApprovalUpsertBulk struct {
	create *PlayerSupervisionRequestApprovalCreateBulk
}

// UpdateNewValues updates the mutable fields using the new values that
// were set on create. Using this option is equivalent to using:
//
//	client.PlayerSupervisionRequestApproval.Create().
//		OnConflict(
//			sql.ResolveWithNewValues(),
//			sql.ResolveWith(func(u *sql.UpdateSet) {
//				u.SetIgnore(playersupervisionrequestapproval.FieldID)
//			}),
//		).
//		Exec(ctx)
func (u *PlayerSupervisionRequestApprovalUpsertBulk) UpdateNewValues() *PlayerSupervisionRequestApprovalUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithNewValues())
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(s *sql.UpdateSet) {
		for _, b := range u.create.builders {
			if _, exists := b.mutation.ID(); exists {
				s.SetIgnore(playersupervisionrequestapproval.FieldID)
			}
		}
	}))
	return u
}

// Ignore sets each column to itself in case of conflict.
// Using this option is equivalent to using:
//
//	client.PlayerSupervisionRequestApproval.Create().
//		OnConflict(sql.ResolveWithIgnore()).
//		Exec(ctx)
func (u *PlayerSupervisionRequestApprovalUpsertBulk) Ignore() *PlayerSupervisionRequestApprovalUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithIgnore())
	return u
}

// DoNothing configures the conflict_action to `DO NOTHING`.
// Supported only by SQLite and PostgreSQL.
func (u *PlayerSupervisionRequestApprovalUpsertBulk) DoNothing() *PlayerSupervisionRequestApprovalUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.DoNothing())
	return u
}

// Update allows overriding fields `UPDATE` values. See the PlayerSupervisionRequestApprovalCreateBulk.OnConflict
// documentation for more info.
func (u *PlayerSupervisionRequestApprovalUpsertBulk) Update(set func(*PlayerSupervisionRequestApprovalUpsert)) *PlayerSupervisionRequestApprovalUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(update *sql.UpdateSet) {
		set(&PlayerSupervisionRequestApprovalUpsert{UpdateSet: update})
	}))
	return u
}

// SetApproved sets the "approved" field.
func (u *PlayerSupervisionRequestApprovalUpsertBulk) SetApproved(v bool) *PlayerSupervisionRequestApprovalUpsertBulk {
	return u.Update(func(s *PlayerSupervisionRequestApprovalUpsert) {
		s.SetApproved(v)
	})
}

// UpdateApproved sets the "approved" field to the value that was provided on create.
func (u *PlayerSupervisionRequestApprovalUpsertBulk) UpdateApproved() *PlayerSupervisionRequestApprovalUpsertBulk {
	return u.Update(func(s *PlayerSupervisionRequestApprovalUpsert) {
		s.UpdateApproved()
	})
}

// ClearApproved clears the value of the "approved" field.
func (u *PlayerSupervisionRequestApprovalUpsertBulk) ClearApproved() *PlayerSupervisionRequestApprovalUpsertBulk {
	return u.Update(func(s *PlayerSupervisionRequestApprovalUpsert) {
		s.ClearApproved()
	})
}

// Exec executes the query.
func (u *PlayerSupervisionRequestApprovalUpsertBulk) Exec(ctx context.Context) error {
	if u.create.err != nil {
		return u.create.err
	}
	for i, b := range u.create.builders {
		if len(b.conflict) != 0 {
			return fmt.Errorf("ent: OnConflict was set for builder %d. Set it on the PlayerSupervisionRequestApprovalCreateBulk instead", i)
		}
	}
	if len(u.create.conflict) == 0 {
		return errors.New("ent: missing options for PlayerSupervisionRequestApprovalCreateBulk.OnConflict")
	}
	return u.create.Exec(ctx)
}

// ExecX is like Exec, but panics if an error occurs.
func (u *PlayerSupervisionRequestApprovalUpsertBulk) ExecX(ctx context.Context) {
	if err := u.create.Exec(ctx); err != nil {
		panic(err)
	}
}

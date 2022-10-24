// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/open-boardgame-stats/backend/internal/ent/player"
	"github.com/open-boardgame-stats/backend/internal/ent/playersupervisionrequest"
	"github.com/open-boardgame-stats/backend/internal/ent/predicate"
	"github.com/open-boardgame-stats/backend/internal/ent/schema/guidgql"
	"github.com/open-boardgame-stats/backend/internal/ent/user"
)

// PlayerUpdate is the builder for updating Player entities.
type PlayerUpdate struct {
	config
	hooks    []Hook
	mutation *PlayerMutation
}

// Where appends a list predicates to the PlayerUpdate builder.
func (pu *PlayerUpdate) Where(ps ...predicate.Player) *PlayerUpdate {
	pu.mutation.Where(ps...)
	return pu
}

// SetName sets the "name" field.
func (pu *PlayerUpdate) SetName(s string) *PlayerUpdate {
	pu.mutation.SetName(s)
	return pu
}

// SetNillableName sets the "name" field if the given value is not nil.
func (pu *PlayerUpdate) SetNillableName(s *string) *PlayerUpdate {
	if s != nil {
		pu.SetName(*s)
	}
	return pu
}

// SetOwnerID sets the "owner" edge to the User entity by ID.
func (pu *PlayerUpdate) SetOwnerID(id guidgql.GUID) *PlayerUpdate {
	pu.mutation.SetOwnerID(id)
	return pu
}

// SetNillableOwnerID sets the "owner" edge to the User entity by ID if the given value is not nil.
func (pu *PlayerUpdate) SetNillableOwnerID(id *guidgql.GUID) *PlayerUpdate {
	if id != nil {
		pu = pu.SetOwnerID(*id)
	}
	return pu
}

// SetOwner sets the "owner" edge to the User entity.
func (pu *PlayerUpdate) SetOwner(u *User) *PlayerUpdate {
	return pu.SetOwnerID(u.ID)
}

// AddSupervisorIDs adds the "supervisors" edge to the User entity by IDs.
func (pu *PlayerUpdate) AddSupervisorIDs(ids ...guidgql.GUID) *PlayerUpdate {
	pu.mutation.AddSupervisorIDs(ids...)
	return pu
}

// AddSupervisors adds the "supervisors" edges to the User entity.
func (pu *PlayerUpdate) AddSupervisors(u ...*User) *PlayerUpdate {
	ids := make([]guidgql.GUID, len(u))
	for i := range u {
		ids[i] = u[i].ID
	}
	return pu.AddSupervisorIDs(ids...)
}

// AddSupervisionRequestIDs adds the "supervision_requests" edge to the PlayerSupervisionRequest entity by IDs.
func (pu *PlayerUpdate) AddSupervisionRequestIDs(ids ...guidgql.GUID) *PlayerUpdate {
	pu.mutation.AddSupervisionRequestIDs(ids...)
	return pu
}

// AddSupervisionRequests adds the "supervision_requests" edges to the PlayerSupervisionRequest entity.
func (pu *PlayerUpdate) AddSupervisionRequests(p ...*PlayerSupervisionRequest) *PlayerUpdate {
	ids := make([]guidgql.GUID, len(p))
	for i := range p {
		ids[i] = p[i].ID
	}
	return pu.AddSupervisionRequestIDs(ids...)
}

// Mutation returns the PlayerMutation object of the builder.
func (pu *PlayerUpdate) Mutation() *PlayerMutation {
	return pu.mutation
}

// ClearOwner clears the "owner" edge to the User entity.
func (pu *PlayerUpdate) ClearOwner() *PlayerUpdate {
	pu.mutation.ClearOwner()
	return pu
}

// ClearSupervisors clears all "supervisors" edges to the User entity.
func (pu *PlayerUpdate) ClearSupervisors() *PlayerUpdate {
	pu.mutation.ClearSupervisors()
	return pu
}

// RemoveSupervisorIDs removes the "supervisors" edge to User entities by IDs.
func (pu *PlayerUpdate) RemoveSupervisorIDs(ids ...guidgql.GUID) *PlayerUpdate {
	pu.mutation.RemoveSupervisorIDs(ids...)
	return pu
}

// RemoveSupervisors removes "supervisors" edges to User entities.
func (pu *PlayerUpdate) RemoveSupervisors(u ...*User) *PlayerUpdate {
	ids := make([]guidgql.GUID, len(u))
	for i := range u {
		ids[i] = u[i].ID
	}
	return pu.RemoveSupervisorIDs(ids...)
}

// ClearSupervisionRequests clears all "supervision_requests" edges to the PlayerSupervisionRequest entity.
func (pu *PlayerUpdate) ClearSupervisionRequests() *PlayerUpdate {
	pu.mutation.ClearSupervisionRequests()
	return pu
}

// RemoveSupervisionRequestIDs removes the "supervision_requests" edge to PlayerSupervisionRequest entities by IDs.
func (pu *PlayerUpdate) RemoveSupervisionRequestIDs(ids ...guidgql.GUID) *PlayerUpdate {
	pu.mutation.RemoveSupervisionRequestIDs(ids...)
	return pu
}

// RemoveSupervisionRequests removes "supervision_requests" edges to PlayerSupervisionRequest entities.
func (pu *PlayerUpdate) RemoveSupervisionRequests(p ...*PlayerSupervisionRequest) *PlayerUpdate {
	ids := make([]guidgql.GUID, len(p))
	for i := range p {
		ids[i] = p[i].ID
	}
	return pu.RemoveSupervisionRequestIDs(ids...)
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (pu *PlayerUpdate) Save(ctx context.Context) (int, error) {
	var (
		err      error
		affected int
	)
	if len(pu.hooks) == 0 {
		affected, err = pu.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*PlayerMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			pu.mutation = mutation
			affected, err = pu.sqlSave(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(pu.hooks) - 1; i >= 0; i-- {
			if pu.hooks[i] == nil {
				return 0, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = pu.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, pu.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// SaveX is like Save, but panics if an error occurs.
func (pu *PlayerUpdate) SaveX(ctx context.Context) int {
	affected, err := pu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (pu *PlayerUpdate) Exec(ctx context.Context) error {
	_, err := pu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (pu *PlayerUpdate) ExecX(ctx context.Context) {
	if err := pu.Exec(ctx); err != nil {
		panic(err)
	}
}

func (pu *PlayerUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   player.Table,
			Columns: player.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeString,
				Column: player.FieldID,
			},
		},
	}
	if ps := pu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := pu.mutation.Name(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: player.FieldName,
		})
	}
	if pu.mutation.OwnerCleared() {
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
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := pu.mutation.OwnerIDs(); len(nodes) > 0 {
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
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if pu.mutation.SupervisorsCleared() {
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
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := pu.mutation.RemovedSupervisorsIDs(); len(nodes) > 0 && !pu.mutation.SupervisorsCleared() {
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
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := pu.mutation.SupervisorsIDs(); len(nodes) > 0 {
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
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if pu.mutation.SupervisionRequestsCleared() {
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
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := pu.mutation.RemovedSupervisionRequestsIDs(); len(nodes) > 0 && !pu.mutation.SupervisionRequestsCleared() {
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
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := pu.mutation.SupervisionRequestsIDs(); len(nodes) > 0 {
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
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, pu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{player.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	return n, nil
}

// PlayerUpdateOne is the builder for updating a single Player entity.
type PlayerUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *PlayerMutation
}

// SetName sets the "name" field.
func (puo *PlayerUpdateOne) SetName(s string) *PlayerUpdateOne {
	puo.mutation.SetName(s)
	return puo
}

// SetNillableName sets the "name" field if the given value is not nil.
func (puo *PlayerUpdateOne) SetNillableName(s *string) *PlayerUpdateOne {
	if s != nil {
		puo.SetName(*s)
	}
	return puo
}

// SetOwnerID sets the "owner" edge to the User entity by ID.
func (puo *PlayerUpdateOne) SetOwnerID(id guidgql.GUID) *PlayerUpdateOne {
	puo.mutation.SetOwnerID(id)
	return puo
}

// SetNillableOwnerID sets the "owner" edge to the User entity by ID if the given value is not nil.
func (puo *PlayerUpdateOne) SetNillableOwnerID(id *guidgql.GUID) *PlayerUpdateOne {
	if id != nil {
		puo = puo.SetOwnerID(*id)
	}
	return puo
}

// SetOwner sets the "owner" edge to the User entity.
func (puo *PlayerUpdateOne) SetOwner(u *User) *PlayerUpdateOne {
	return puo.SetOwnerID(u.ID)
}

// AddSupervisorIDs adds the "supervisors" edge to the User entity by IDs.
func (puo *PlayerUpdateOne) AddSupervisorIDs(ids ...guidgql.GUID) *PlayerUpdateOne {
	puo.mutation.AddSupervisorIDs(ids...)
	return puo
}

// AddSupervisors adds the "supervisors" edges to the User entity.
func (puo *PlayerUpdateOne) AddSupervisors(u ...*User) *PlayerUpdateOne {
	ids := make([]guidgql.GUID, len(u))
	for i := range u {
		ids[i] = u[i].ID
	}
	return puo.AddSupervisorIDs(ids...)
}

// AddSupervisionRequestIDs adds the "supervision_requests" edge to the PlayerSupervisionRequest entity by IDs.
func (puo *PlayerUpdateOne) AddSupervisionRequestIDs(ids ...guidgql.GUID) *PlayerUpdateOne {
	puo.mutation.AddSupervisionRequestIDs(ids...)
	return puo
}

// AddSupervisionRequests adds the "supervision_requests" edges to the PlayerSupervisionRequest entity.
func (puo *PlayerUpdateOne) AddSupervisionRequests(p ...*PlayerSupervisionRequest) *PlayerUpdateOne {
	ids := make([]guidgql.GUID, len(p))
	for i := range p {
		ids[i] = p[i].ID
	}
	return puo.AddSupervisionRequestIDs(ids...)
}

// Mutation returns the PlayerMutation object of the builder.
func (puo *PlayerUpdateOne) Mutation() *PlayerMutation {
	return puo.mutation
}

// ClearOwner clears the "owner" edge to the User entity.
func (puo *PlayerUpdateOne) ClearOwner() *PlayerUpdateOne {
	puo.mutation.ClearOwner()
	return puo
}

// ClearSupervisors clears all "supervisors" edges to the User entity.
func (puo *PlayerUpdateOne) ClearSupervisors() *PlayerUpdateOne {
	puo.mutation.ClearSupervisors()
	return puo
}

// RemoveSupervisorIDs removes the "supervisors" edge to User entities by IDs.
func (puo *PlayerUpdateOne) RemoveSupervisorIDs(ids ...guidgql.GUID) *PlayerUpdateOne {
	puo.mutation.RemoveSupervisorIDs(ids...)
	return puo
}

// RemoveSupervisors removes "supervisors" edges to User entities.
func (puo *PlayerUpdateOne) RemoveSupervisors(u ...*User) *PlayerUpdateOne {
	ids := make([]guidgql.GUID, len(u))
	for i := range u {
		ids[i] = u[i].ID
	}
	return puo.RemoveSupervisorIDs(ids...)
}

// ClearSupervisionRequests clears all "supervision_requests" edges to the PlayerSupervisionRequest entity.
func (puo *PlayerUpdateOne) ClearSupervisionRequests() *PlayerUpdateOne {
	puo.mutation.ClearSupervisionRequests()
	return puo
}

// RemoveSupervisionRequestIDs removes the "supervision_requests" edge to PlayerSupervisionRequest entities by IDs.
func (puo *PlayerUpdateOne) RemoveSupervisionRequestIDs(ids ...guidgql.GUID) *PlayerUpdateOne {
	puo.mutation.RemoveSupervisionRequestIDs(ids...)
	return puo
}

// RemoveSupervisionRequests removes "supervision_requests" edges to PlayerSupervisionRequest entities.
func (puo *PlayerUpdateOne) RemoveSupervisionRequests(p ...*PlayerSupervisionRequest) *PlayerUpdateOne {
	ids := make([]guidgql.GUID, len(p))
	for i := range p {
		ids[i] = p[i].ID
	}
	return puo.RemoveSupervisionRequestIDs(ids...)
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (puo *PlayerUpdateOne) Select(field string, fields ...string) *PlayerUpdateOne {
	puo.fields = append([]string{field}, fields...)
	return puo
}

// Save executes the query and returns the updated Player entity.
func (puo *PlayerUpdateOne) Save(ctx context.Context) (*Player, error) {
	var (
		err  error
		node *Player
	)
	if len(puo.hooks) == 0 {
		node, err = puo.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*PlayerMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			puo.mutation = mutation
			node, err = puo.sqlSave(ctx)
			mutation.done = true
			return node, err
		})
		for i := len(puo.hooks) - 1; i >= 0; i-- {
			if puo.hooks[i] == nil {
				return nil, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = puo.hooks[i](mut)
		}
		v, err := mut.Mutate(ctx, puo.mutation)
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

// SaveX is like Save, but panics if an error occurs.
func (puo *PlayerUpdateOne) SaveX(ctx context.Context) *Player {
	node, err := puo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (puo *PlayerUpdateOne) Exec(ctx context.Context) error {
	_, err := puo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (puo *PlayerUpdateOne) ExecX(ctx context.Context) {
	if err := puo.Exec(ctx); err != nil {
		panic(err)
	}
}

func (puo *PlayerUpdateOne) sqlSave(ctx context.Context) (_node *Player, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   player.Table,
			Columns: player.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeString,
				Column: player.FieldID,
			},
		},
	}
	id, ok := puo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "Player.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := puo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, player.FieldID)
		for _, f := range fields {
			if !player.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != player.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := puo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := puo.mutation.Name(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: player.FieldName,
		})
	}
	if puo.mutation.OwnerCleared() {
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
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := puo.mutation.OwnerIDs(); len(nodes) > 0 {
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
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if puo.mutation.SupervisorsCleared() {
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
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := puo.mutation.RemovedSupervisorsIDs(); len(nodes) > 0 && !puo.mutation.SupervisorsCleared() {
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
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := puo.mutation.SupervisorsIDs(); len(nodes) > 0 {
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
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if puo.mutation.SupervisionRequestsCleared() {
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
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := puo.mutation.RemovedSupervisionRequestsIDs(); len(nodes) > 0 && !puo.mutation.SupervisionRequestsCleared() {
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
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := puo.mutation.SupervisionRequestsIDs(); len(nodes) > 0 {
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
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	_node = &Player{config: puo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, puo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{player.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	return _node, nil
}

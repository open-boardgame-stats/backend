// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"

	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/google/uuid"
	"github.com/open-boardgame-stats/backend/internal/ent/player"
	"github.com/open-boardgame-stats/backend/internal/ent/playersupervisionrequest"
	"github.com/open-boardgame-stats/backend/internal/ent/user"
)

// PlayerCreate is the builder for creating a Player entity.
type PlayerCreate struct {
	config
	mutation *PlayerMutation
	hooks    []Hook
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
func (pc *PlayerCreate) SetID(u uuid.UUID) *PlayerCreate {
	pc.mutation.SetID(u)
	return pc
}

// SetNillableID sets the "id" field if the given value is not nil.
func (pc *PlayerCreate) SetNillableID(u *uuid.UUID) *PlayerCreate {
	if u != nil {
		pc.SetID(*u)
	}
	return pc
}

// SetOwnerID sets the "owner" edge to the User entity by ID.
func (pc *PlayerCreate) SetOwnerID(id uuid.UUID) *PlayerCreate {
	pc.mutation.SetOwnerID(id)
	return pc
}

// SetNillableOwnerID sets the "owner" edge to the User entity by ID if the given value is not nil.
func (pc *PlayerCreate) SetNillableOwnerID(id *uuid.UUID) *PlayerCreate {
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
func (pc *PlayerCreate) AddSupervisorIDs(ids ...uuid.UUID) *PlayerCreate {
	pc.mutation.AddSupervisorIDs(ids...)
	return pc
}

// AddSupervisors adds the "supervisors" edges to the User entity.
func (pc *PlayerCreate) AddSupervisors(u ...*User) *PlayerCreate {
	ids := make([]uuid.UUID, len(u))
	for i := range u {
		ids[i] = u[i].ID
	}
	return pc.AddSupervisorIDs(ids...)
}

// AddSupervisionRequestIDs adds the "supervision_requests" edge to the PlayerSupervisionRequest entity by IDs.
func (pc *PlayerCreate) AddSupervisionRequestIDs(ids ...uuid.UUID) *PlayerCreate {
	pc.mutation.AddSupervisionRequestIDs(ids...)
	return pc
}

// AddSupervisionRequests adds the "supervision_requests" edges to the PlayerSupervisionRequest entity.
func (pc *PlayerCreate) AddSupervisionRequests(p ...*PlayerSupervisionRequest) *PlayerCreate {
	ids := make([]uuid.UUID, len(p))
	for i := range p {
		ids[i] = p[i].ID
	}
	return pc.AddSupervisionRequestIDs(ids...)
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
		if id, ok := _spec.ID.Value.(*uuid.UUID); ok {
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
				Type:   field.TypeUUID,
				Column: player.FieldID,
			},
		}
	)
	if id, ok := pc.mutation.ID(); ok {
		_node.ID = id
		_spec.ID.Value = &id
	}
	if value, ok := pc.mutation.Name(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: player.FieldName,
		})
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
					Type:   field.TypeUUID,
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
					Type:   field.TypeUUID,
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
					Type:   field.TypeUUID,
					Column: playersupervisionrequest.FieldID,
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

// PlayerCreateBulk is the builder for creating many Player entities in bulk.
type PlayerCreateBulk struct {
	config
	builders []*PlayerCreate
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

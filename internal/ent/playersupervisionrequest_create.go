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
	"github.com/open-boardgame-stats/backend/internal/ent/playersupervisionrequestapproval"
	"github.com/open-boardgame-stats/backend/internal/ent/user"
)

// PlayerSupervisionRequestCreate is the builder for creating a PlayerSupervisionRequest entity.
type PlayerSupervisionRequestCreate struct {
	config
	mutation *PlayerSupervisionRequestMutation
	hooks    []Hook
}

// SetMessage sets the "message" field.
func (psrc *PlayerSupervisionRequestCreate) SetMessage(s string) *PlayerSupervisionRequestCreate {
	psrc.mutation.SetMessage(s)
	return psrc
}

// SetNillableMessage sets the "message" field if the given value is not nil.
func (psrc *PlayerSupervisionRequestCreate) SetNillableMessage(s *string) *PlayerSupervisionRequestCreate {
	if s != nil {
		psrc.SetMessage(*s)
	}
	return psrc
}

// SetID sets the "id" field.
func (psrc *PlayerSupervisionRequestCreate) SetID(u uuid.UUID) *PlayerSupervisionRequestCreate {
	psrc.mutation.SetID(u)
	return psrc
}

// SetNillableID sets the "id" field if the given value is not nil.
func (psrc *PlayerSupervisionRequestCreate) SetNillableID(u *uuid.UUID) *PlayerSupervisionRequestCreate {
	if u != nil {
		psrc.SetID(*u)
	}
	return psrc
}

// SetSenderID sets the "sender" edge to the User entity by ID.
func (psrc *PlayerSupervisionRequestCreate) SetSenderID(id uuid.UUID) *PlayerSupervisionRequestCreate {
	psrc.mutation.SetSenderID(id)
	return psrc
}

// SetSender sets the "sender" edge to the User entity.
func (psrc *PlayerSupervisionRequestCreate) SetSender(u *User) *PlayerSupervisionRequestCreate {
	return psrc.SetSenderID(u.ID)
}

// SetPlayerID sets the "player" edge to the Player entity by ID.
func (psrc *PlayerSupervisionRequestCreate) SetPlayerID(id uuid.UUID) *PlayerSupervisionRequestCreate {
	psrc.mutation.SetPlayerID(id)
	return psrc
}

// SetPlayer sets the "player" edge to the Player entity.
func (psrc *PlayerSupervisionRequestCreate) SetPlayer(p *Player) *PlayerSupervisionRequestCreate {
	return psrc.SetPlayerID(p.ID)
}

// AddApprovalIDs adds the "approvals" edge to the PlayerSupervisionRequestApproval entity by IDs.
func (psrc *PlayerSupervisionRequestCreate) AddApprovalIDs(ids ...uuid.UUID) *PlayerSupervisionRequestCreate {
	psrc.mutation.AddApprovalIDs(ids...)
	return psrc
}

// AddApprovals adds the "approvals" edges to the PlayerSupervisionRequestApproval entity.
func (psrc *PlayerSupervisionRequestCreate) AddApprovals(p ...*PlayerSupervisionRequestApproval) *PlayerSupervisionRequestCreate {
	ids := make([]uuid.UUID, len(p))
	for i := range p {
		ids[i] = p[i].ID
	}
	return psrc.AddApprovalIDs(ids...)
}

// Mutation returns the PlayerSupervisionRequestMutation object of the builder.
func (psrc *PlayerSupervisionRequestCreate) Mutation() *PlayerSupervisionRequestMutation {
	return psrc.mutation
}

// Save creates the PlayerSupervisionRequest in the database.
func (psrc *PlayerSupervisionRequestCreate) Save(ctx context.Context) (*PlayerSupervisionRequest, error) {
	var (
		err  error
		node *PlayerSupervisionRequest
	)
	psrc.defaults()
	if len(psrc.hooks) == 0 {
		if err = psrc.check(); err != nil {
			return nil, err
		}
		node, err = psrc.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*PlayerSupervisionRequestMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = psrc.check(); err != nil {
				return nil, err
			}
			psrc.mutation = mutation
			if node, err = psrc.sqlSave(ctx); err != nil {
				return nil, err
			}
			mutation.id = &node.ID
			mutation.done = true
			return node, err
		})
		for i := len(psrc.hooks) - 1; i >= 0; i-- {
			if psrc.hooks[i] == nil {
				return nil, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = psrc.hooks[i](mut)
		}
		v, err := mut.Mutate(ctx, psrc.mutation)
		if err != nil {
			return nil, err
		}
		nv, ok := v.(*PlayerSupervisionRequest)
		if !ok {
			return nil, fmt.Errorf("unexpected node type %T returned from PlayerSupervisionRequestMutation", v)
		}
		node = nv
	}
	return node, err
}

// SaveX calls Save and panics if Save returns an error.
func (psrc *PlayerSupervisionRequestCreate) SaveX(ctx context.Context) *PlayerSupervisionRequest {
	v, err := psrc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (psrc *PlayerSupervisionRequestCreate) Exec(ctx context.Context) error {
	_, err := psrc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (psrc *PlayerSupervisionRequestCreate) ExecX(ctx context.Context) {
	if err := psrc.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (psrc *PlayerSupervisionRequestCreate) defaults() {
	if _, ok := psrc.mutation.ID(); !ok {
		v := playersupervisionrequest.DefaultID()
		psrc.mutation.SetID(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (psrc *PlayerSupervisionRequestCreate) check() error {
	if _, ok := psrc.mutation.SenderID(); !ok {
		return &ValidationError{Name: "sender", err: errors.New(`ent: missing required edge "PlayerSupervisionRequest.sender"`)}
	}
	if _, ok := psrc.mutation.PlayerID(); !ok {
		return &ValidationError{Name: "player", err: errors.New(`ent: missing required edge "PlayerSupervisionRequest.player"`)}
	}
	return nil
}

func (psrc *PlayerSupervisionRequestCreate) sqlSave(ctx context.Context) (*PlayerSupervisionRequest, error) {
	_node, _spec := psrc.createSpec()
	if err := sqlgraph.CreateNode(ctx, psrc.driver, _spec); err != nil {
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

func (psrc *PlayerSupervisionRequestCreate) createSpec() (*PlayerSupervisionRequest, *sqlgraph.CreateSpec) {
	var (
		_node = &PlayerSupervisionRequest{config: psrc.config}
		_spec = &sqlgraph.CreateSpec{
			Table: playersupervisionrequest.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeUUID,
				Column: playersupervisionrequest.FieldID,
			},
		}
	)
	if id, ok := psrc.mutation.ID(); ok {
		_node.ID = id
		_spec.ID.Value = &id
	}
	if value, ok := psrc.mutation.Message(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: playersupervisionrequest.FieldMessage,
		})
		_node.Message = value
	}
	if nodes := psrc.mutation.SenderIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   playersupervisionrequest.SenderTable,
			Columns: []string{playersupervisionrequest.SenderColumn},
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
		_node.user_sent_supervision_requests = &nodes[0]
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := psrc.mutation.PlayerIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   playersupervisionrequest.PlayerTable,
			Columns: []string{playersupervisionrequest.PlayerColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeUUID,
					Column: player.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_node.player_supervision_requests = &nodes[0]
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := psrc.mutation.ApprovalsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   playersupervisionrequest.ApprovalsTable,
			Columns: []string{playersupervisionrequest.ApprovalsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeUUID,
					Column: playersupervisionrequestapproval.FieldID,
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

// PlayerSupervisionRequestCreateBulk is the builder for creating many PlayerSupervisionRequest entities in bulk.
type PlayerSupervisionRequestCreateBulk struct {
	config
	builders []*PlayerSupervisionRequestCreate
}

// Save creates the PlayerSupervisionRequest entities in the database.
func (psrcb *PlayerSupervisionRequestCreateBulk) Save(ctx context.Context) ([]*PlayerSupervisionRequest, error) {
	specs := make([]*sqlgraph.CreateSpec, len(psrcb.builders))
	nodes := make([]*PlayerSupervisionRequest, len(psrcb.builders))
	mutators := make([]Mutator, len(psrcb.builders))
	for i := range psrcb.builders {
		func(i int, root context.Context) {
			builder := psrcb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*PlayerSupervisionRequestMutation)
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
					_, err = mutators[i+1].Mutate(root, psrcb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, psrcb.driver, spec); err != nil {
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
		if _, err := mutators[0].Mutate(ctx, psrcb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (psrcb *PlayerSupervisionRequestCreateBulk) SaveX(ctx context.Context) []*PlayerSupervisionRequest {
	v, err := psrcb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (psrcb *PlayerSupervisionRequestCreateBulk) Exec(ctx context.Context) error {
	_, err := psrcb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (psrcb *PlayerSupervisionRequestCreateBulk) ExecX(ctx context.Context) {
	if err := psrcb.Exec(ctx); err != nil {
		panic(err)
	}
}

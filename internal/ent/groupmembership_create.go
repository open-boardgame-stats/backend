// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"

	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/open-boardgame-stats/backend/internal/ent/enums"
	"github.com/open-boardgame-stats/backend/internal/ent/group"
	"github.com/open-boardgame-stats/backend/internal/ent/groupmembership"
	"github.com/open-boardgame-stats/backend/internal/ent/schema/guidgql"
	"github.com/open-boardgame-stats/backend/internal/ent/user"
)

// GroupMembershipCreate is the builder for creating a GroupMembership entity.
type GroupMembershipCreate struct {
	config
	mutation *GroupMembershipMutation
	hooks    []Hook
}

// SetRole sets the "role" field.
func (gmc *GroupMembershipCreate) SetRole(e enums.Role) *GroupMembershipCreate {
	gmc.mutation.SetRole(e)
	return gmc
}

// SetID sets the "id" field.
func (gmc *GroupMembershipCreate) SetID(gu guidgql.GUID) *GroupMembershipCreate {
	gmc.mutation.SetID(gu)
	return gmc
}

// SetNillableID sets the "id" field if the given value is not nil.
func (gmc *GroupMembershipCreate) SetNillableID(gu *guidgql.GUID) *GroupMembershipCreate {
	if gu != nil {
		gmc.SetID(*gu)
	}
	return gmc
}

// SetGroupID sets the "group" edge to the Group entity by ID.
func (gmc *GroupMembershipCreate) SetGroupID(id guidgql.GUID) *GroupMembershipCreate {
	gmc.mutation.SetGroupID(id)
	return gmc
}

// SetGroup sets the "group" edge to the Group entity.
func (gmc *GroupMembershipCreate) SetGroup(g *Group) *GroupMembershipCreate {
	return gmc.SetGroupID(g.ID)
}

// SetUserID sets the "user" edge to the User entity by ID.
func (gmc *GroupMembershipCreate) SetUserID(id guidgql.GUID) *GroupMembershipCreate {
	gmc.mutation.SetUserID(id)
	return gmc
}

// SetUser sets the "user" edge to the User entity.
func (gmc *GroupMembershipCreate) SetUser(u *User) *GroupMembershipCreate {
	return gmc.SetUserID(u.ID)
}

// Mutation returns the GroupMembershipMutation object of the builder.
func (gmc *GroupMembershipCreate) Mutation() *GroupMembershipMutation {
	return gmc.mutation
}

// Save creates the GroupMembership in the database.
func (gmc *GroupMembershipCreate) Save(ctx context.Context) (*GroupMembership, error) {
	var (
		err  error
		node *GroupMembership
	)
	gmc.defaults()
	if len(gmc.hooks) == 0 {
		if err = gmc.check(); err != nil {
			return nil, err
		}
		node, err = gmc.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*GroupMembershipMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = gmc.check(); err != nil {
				return nil, err
			}
			gmc.mutation = mutation
			if node, err = gmc.sqlSave(ctx); err != nil {
				return nil, err
			}
			mutation.id = &node.ID
			mutation.done = true
			return node, err
		})
		for i := len(gmc.hooks) - 1; i >= 0; i-- {
			if gmc.hooks[i] == nil {
				return nil, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = gmc.hooks[i](mut)
		}
		v, err := mut.Mutate(ctx, gmc.mutation)
		if err != nil {
			return nil, err
		}
		nv, ok := v.(*GroupMembership)
		if !ok {
			return nil, fmt.Errorf("unexpected node type %T returned from GroupMembershipMutation", v)
		}
		node = nv
	}
	return node, err
}

// SaveX calls Save and panics if Save returns an error.
func (gmc *GroupMembershipCreate) SaveX(ctx context.Context) *GroupMembership {
	v, err := gmc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (gmc *GroupMembershipCreate) Exec(ctx context.Context) error {
	_, err := gmc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (gmc *GroupMembershipCreate) ExecX(ctx context.Context) {
	if err := gmc.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (gmc *GroupMembershipCreate) defaults() {
	if _, ok := gmc.mutation.ID(); !ok {
		v := groupmembership.DefaultID()
		gmc.mutation.SetID(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (gmc *GroupMembershipCreate) check() error {
	if _, ok := gmc.mutation.Role(); !ok {
		return &ValidationError{Name: "role", err: errors.New(`ent: missing required field "GroupMembership.role"`)}
	}
	if v, ok := gmc.mutation.Role(); ok {
		if err := groupmembership.RoleValidator(v); err != nil {
			return &ValidationError{Name: "role", err: fmt.Errorf(`ent: validator failed for field "GroupMembership.role": %w`, err)}
		}
	}
	if _, ok := gmc.mutation.GroupID(); !ok {
		return &ValidationError{Name: "group", err: errors.New(`ent: missing required edge "GroupMembership.group"`)}
	}
	if _, ok := gmc.mutation.UserID(); !ok {
		return &ValidationError{Name: "user", err: errors.New(`ent: missing required edge "GroupMembership.user"`)}
	}
	return nil
}

func (gmc *GroupMembershipCreate) sqlSave(ctx context.Context) (*GroupMembership, error) {
	_node, _spec := gmc.createSpec()
	if err := sqlgraph.CreateNode(ctx, gmc.driver, _spec); err != nil {
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

func (gmc *GroupMembershipCreate) createSpec() (*GroupMembership, *sqlgraph.CreateSpec) {
	var (
		_node = &GroupMembership{config: gmc.config}
		_spec = &sqlgraph.CreateSpec{
			Table: groupmembership.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeString,
				Column: groupmembership.FieldID,
			},
		}
	)
	if id, ok := gmc.mutation.ID(); ok {
		_node.ID = id
		_spec.ID.Value = &id
	}
	if value, ok := gmc.mutation.Role(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeEnum,
			Value:  value,
			Column: groupmembership.FieldRole,
		})
		_node.Role = value
	}
	if nodes := gmc.mutation.GroupIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   groupmembership.GroupTable,
			Columns: []string{groupmembership.GroupColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeString,
					Column: group.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_node.group_members = &nodes[0]
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := gmc.mutation.UserIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   groupmembership.UserTable,
			Columns: []string{groupmembership.UserColumn},
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
		_node.user_group_memberships = &nodes[0]
		_spec.Edges = append(_spec.Edges, edge)
	}
	return _node, _spec
}

// GroupMembershipCreateBulk is the builder for creating many GroupMembership entities in bulk.
type GroupMembershipCreateBulk struct {
	config
	builders []*GroupMembershipCreate
}

// Save creates the GroupMembership entities in the database.
func (gmcb *GroupMembershipCreateBulk) Save(ctx context.Context) ([]*GroupMembership, error) {
	specs := make([]*sqlgraph.CreateSpec, len(gmcb.builders))
	nodes := make([]*GroupMembership, len(gmcb.builders))
	mutators := make([]Mutator, len(gmcb.builders))
	for i := range gmcb.builders {
		func(i int, root context.Context) {
			builder := gmcb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*GroupMembershipMutation)
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
					_, err = mutators[i+1].Mutate(root, gmcb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, gmcb.driver, spec); err != nil {
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
		if _, err := mutators[0].Mutate(ctx, gmcb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (gmcb *GroupMembershipCreateBulk) SaveX(ctx context.Context) []*GroupMembership {
	v, err := gmcb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (gmcb *GroupMembershipCreateBulk) Exec(ctx context.Context) error {
	_, err := gmcb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (gmcb *GroupMembershipCreateBulk) ExecX(ctx context.Context) {
	if err := gmcb.Exec(ctx); err != nil {
		panic(err)
	}
}

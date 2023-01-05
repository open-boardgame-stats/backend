// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"fmt"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/open-boardgame-stats/backend/internal/ent/numericalstat"
	"github.com/open-boardgame-stats/backend/internal/ent/predicate"
)

// NumericalStatDelete is the builder for deleting a NumericalStat entity.
type NumericalStatDelete struct {
	config
	hooks    []Hook
	mutation *NumericalStatMutation
}

// Where appends a list predicates to the NumericalStatDelete builder.
func (nsd *NumericalStatDelete) Where(ps ...predicate.NumericalStat) *NumericalStatDelete {
	nsd.mutation.Where(ps...)
	return nsd
}

// Exec executes the deletion query and returns how many vertices were deleted.
func (nsd *NumericalStatDelete) Exec(ctx context.Context) (int, error) {
	var (
		err      error
		affected int
	)
	if len(nsd.hooks) == 0 {
		affected, err = nsd.sqlExec(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*NumericalStatMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			nsd.mutation = mutation
			affected, err = nsd.sqlExec(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(nsd.hooks) - 1; i >= 0; i-- {
			if nsd.hooks[i] == nil {
				return 0, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = nsd.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, nsd.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// ExecX is like Exec, but panics if an error occurs.
func (nsd *NumericalStatDelete) ExecX(ctx context.Context) int {
	n, err := nsd.Exec(ctx)
	if err != nil {
		panic(err)
	}
	return n
}

func (nsd *NumericalStatDelete) sqlExec(ctx context.Context) (int, error) {
	_spec := &sqlgraph.DeleteSpec{
		Node: &sqlgraph.NodeSpec{
			Table: numericalstat.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeString,
				Column: numericalstat.FieldID,
			},
		},
	}
	if ps := nsd.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	affected, err := sqlgraph.DeleteNodes(ctx, nsd.driver, _spec)
	if err != nil && sqlgraph.IsConstraintError(err) {
		err = &ConstraintError{msg: err.Error(), wrap: err}
	}
	return affected, err
}

// NumericalStatDeleteOne is the builder for deleting a single NumericalStat entity.
type NumericalStatDeleteOne struct {
	nsd *NumericalStatDelete
}

// Exec executes the deletion query.
func (nsdo *NumericalStatDeleteOne) Exec(ctx context.Context) error {
	n, err := nsdo.nsd.Exec(ctx)
	switch {
	case err != nil:
		return err
	case n == 0:
		return &NotFoundError{numericalstat.Label}
	default:
		return nil
	}
}

// ExecX is like Exec, but panics if an error occurs.
func (nsdo *NumericalStatDeleteOne) ExecX(ctx context.Context) {
	nsdo.nsd.ExecX(ctx)
}

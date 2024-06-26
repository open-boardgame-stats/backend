// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/open-boardgame-stats/backend/internal/ent/playersupervisionrequestapproval"
	"github.com/open-boardgame-stats/backend/internal/ent/predicate"
)

// PlayerSupervisionRequestApprovalDelete is the builder for deleting a PlayerSupervisionRequestApproval entity.
type PlayerSupervisionRequestApprovalDelete struct {
	config
	hooks    []Hook
	mutation *PlayerSupervisionRequestApprovalMutation
}

// Where appends a list predicates to the PlayerSupervisionRequestApprovalDelete builder.
func (psrad *PlayerSupervisionRequestApprovalDelete) Where(ps ...predicate.PlayerSupervisionRequestApproval) *PlayerSupervisionRequestApprovalDelete {
	psrad.mutation.Where(ps...)
	return psrad
}

// Exec executes the deletion query and returns how many vertices were deleted.
func (psrad *PlayerSupervisionRequestApprovalDelete) Exec(ctx context.Context) (int, error) {
	return withHooks(ctx, psrad.sqlExec, psrad.mutation, psrad.hooks)
}

// ExecX is like Exec, but panics if an error occurs.
func (psrad *PlayerSupervisionRequestApprovalDelete) ExecX(ctx context.Context) int {
	n, err := psrad.Exec(ctx)
	if err != nil {
		panic(err)
	}
	return n
}

func (psrad *PlayerSupervisionRequestApprovalDelete) sqlExec(ctx context.Context) (int, error) {
	_spec := sqlgraph.NewDeleteSpec(playersupervisionrequestapproval.Table, sqlgraph.NewFieldSpec(playersupervisionrequestapproval.FieldID, field.TypeString))
	if ps := psrad.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	affected, err := sqlgraph.DeleteNodes(ctx, psrad.driver, _spec)
	if err != nil && sqlgraph.IsConstraintError(err) {
		err = &ConstraintError{msg: err.Error(), wrap: err}
	}
	psrad.mutation.done = true
	return affected, err
}

// PlayerSupervisionRequestApprovalDeleteOne is the builder for deleting a single PlayerSupervisionRequestApproval entity.
type PlayerSupervisionRequestApprovalDeleteOne struct {
	psrad *PlayerSupervisionRequestApprovalDelete
}

// Where appends a list predicates to the PlayerSupervisionRequestApprovalDelete builder.
func (psrado *PlayerSupervisionRequestApprovalDeleteOne) Where(ps ...predicate.PlayerSupervisionRequestApproval) *PlayerSupervisionRequestApprovalDeleteOne {
	psrado.psrad.mutation.Where(ps...)
	return psrado
}

// Exec executes the deletion query.
func (psrado *PlayerSupervisionRequestApprovalDeleteOne) Exec(ctx context.Context) error {
	n, err := psrado.psrad.Exec(ctx)
	switch {
	case err != nil:
		return err
	case n == 0:
		return &NotFoundError{playersupervisionrequestapproval.Label}
	default:
		return nil
	}
}

// ExecX is like Exec, but panics if an error occurs.
func (psrado *PlayerSupervisionRequestApprovalDeleteOne) ExecX(ctx context.Context) {
	if err := psrado.Exec(ctx); err != nil {
		panic(err)
	}
}

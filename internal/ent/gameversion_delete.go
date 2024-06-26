// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/open-boardgame-stats/backend/internal/ent/gameversion"
	"github.com/open-boardgame-stats/backend/internal/ent/predicate"
)

// GameVersionDelete is the builder for deleting a GameVersion entity.
type GameVersionDelete struct {
	config
	hooks    []Hook
	mutation *GameVersionMutation
}

// Where appends a list predicates to the GameVersionDelete builder.
func (gvd *GameVersionDelete) Where(ps ...predicate.GameVersion) *GameVersionDelete {
	gvd.mutation.Where(ps...)
	return gvd
}

// Exec executes the deletion query and returns how many vertices were deleted.
func (gvd *GameVersionDelete) Exec(ctx context.Context) (int, error) {
	return withHooks(ctx, gvd.sqlExec, gvd.mutation, gvd.hooks)
}

// ExecX is like Exec, but panics if an error occurs.
func (gvd *GameVersionDelete) ExecX(ctx context.Context) int {
	n, err := gvd.Exec(ctx)
	if err != nil {
		panic(err)
	}
	return n
}

func (gvd *GameVersionDelete) sqlExec(ctx context.Context) (int, error) {
	_spec := sqlgraph.NewDeleteSpec(gameversion.Table, sqlgraph.NewFieldSpec(gameversion.FieldID, field.TypeString))
	if ps := gvd.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	affected, err := sqlgraph.DeleteNodes(ctx, gvd.driver, _spec)
	if err != nil && sqlgraph.IsConstraintError(err) {
		err = &ConstraintError{msg: err.Error(), wrap: err}
	}
	gvd.mutation.done = true
	return affected, err
}

// GameVersionDeleteOne is the builder for deleting a single GameVersion entity.
type GameVersionDeleteOne struct {
	gvd *GameVersionDelete
}

// Where appends a list predicates to the GameVersionDelete builder.
func (gvdo *GameVersionDeleteOne) Where(ps ...predicate.GameVersion) *GameVersionDeleteOne {
	gvdo.gvd.mutation.Where(ps...)
	return gvdo
}

// Exec executes the deletion query.
func (gvdo *GameVersionDeleteOne) Exec(ctx context.Context) error {
	n, err := gvdo.gvd.Exec(ctx)
	switch {
	case err != nil:
		return err
	case n == 0:
		return &NotFoundError{gameversion.Label}
	default:
		return nil
	}
}

// ExecX is like Exec, but panics if an error occurs.
func (gvdo *GameVersionDeleteOne) ExecX(ctx context.Context) {
	if err := gvdo.Exec(ctx); err != nil {
		panic(err)
	}
}

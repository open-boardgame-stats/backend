// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/open-boardgame-stats/backend/internal/ent/gamefavorite"
	"github.com/open-boardgame-stats/backend/internal/ent/predicate"
)

// GameFavoriteDelete is the builder for deleting a GameFavorite entity.
type GameFavoriteDelete struct {
	config
	hooks    []Hook
	mutation *GameFavoriteMutation
}

// Where appends a list predicates to the GameFavoriteDelete builder.
func (gfd *GameFavoriteDelete) Where(ps ...predicate.GameFavorite) *GameFavoriteDelete {
	gfd.mutation.Where(ps...)
	return gfd
}

// Exec executes the deletion query and returns how many vertices were deleted.
func (gfd *GameFavoriteDelete) Exec(ctx context.Context) (int, error) {
	return withHooks(ctx, gfd.sqlExec, gfd.mutation, gfd.hooks)
}

// ExecX is like Exec, but panics if an error occurs.
func (gfd *GameFavoriteDelete) ExecX(ctx context.Context) int {
	n, err := gfd.Exec(ctx)
	if err != nil {
		panic(err)
	}
	return n
}

func (gfd *GameFavoriteDelete) sqlExec(ctx context.Context) (int, error) {
	_spec := sqlgraph.NewDeleteSpec(gamefavorite.Table, sqlgraph.NewFieldSpec(gamefavorite.FieldID, field.TypeString))
	if ps := gfd.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	affected, err := sqlgraph.DeleteNodes(ctx, gfd.driver, _spec)
	if err != nil && sqlgraph.IsConstraintError(err) {
		err = &ConstraintError{msg: err.Error(), wrap: err}
	}
	gfd.mutation.done = true
	return affected, err
}

// GameFavoriteDeleteOne is the builder for deleting a single GameFavorite entity.
type GameFavoriteDeleteOne struct {
	gfd *GameFavoriteDelete
}

// Where appends a list predicates to the GameFavoriteDelete builder.
func (gfdo *GameFavoriteDeleteOne) Where(ps ...predicate.GameFavorite) *GameFavoriteDeleteOne {
	gfdo.gfd.mutation.Where(ps...)
	return gfdo
}

// Exec executes the deletion query.
func (gfdo *GameFavoriteDeleteOne) Exec(ctx context.Context) error {
	n, err := gfdo.gfd.Exec(ctx)
	switch {
	case err != nil:
		return err
	case n == 0:
		return &NotFoundError{gamefavorite.Label}
	default:
		return nil
	}
}

// ExecX is like Exec, but panics if an error occurs.
func (gfdo *GameFavoriteDeleteOne) ExecX(ctx context.Context) {
	if err := gfdo.Exec(ctx); err != nil {
		panic(err)
	}
}

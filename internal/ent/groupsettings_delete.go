// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/open-boardgame-stats/backend/internal/ent/groupsettings"
	"github.com/open-boardgame-stats/backend/internal/ent/predicate"
)

// GroupSettingsDelete is the builder for deleting a GroupSettings entity.
type GroupSettingsDelete struct {
	config
	hooks    []Hook
	mutation *GroupSettingsMutation
}

// Where appends a list predicates to the GroupSettingsDelete builder.
func (gsd *GroupSettingsDelete) Where(ps ...predicate.GroupSettings) *GroupSettingsDelete {
	gsd.mutation.Where(ps...)
	return gsd
}

// Exec executes the deletion query and returns how many vertices were deleted.
func (gsd *GroupSettingsDelete) Exec(ctx context.Context) (int, error) {
	return withHooks(ctx, gsd.sqlExec, gsd.mutation, gsd.hooks)
}

// ExecX is like Exec, but panics if an error occurs.
func (gsd *GroupSettingsDelete) ExecX(ctx context.Context) int {
	n, err := gsd.Exec(ctx)
	if err != nil {
		panic(err)
	}
	return n
}

func (gsd *GroupSettingsDelete) sqlExec(ctx context.Context) (int, error) {
	_spec := sqlgraph.NewDeleteSpec(groupsettings.Table, sqlgraph.NewFieldSpec(groupsettings.FieldID, field.TypeString))
	if ps := gsd.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	affected, err := sqlgraph.DeleteNodes(ctx, gsd.driver, _spec)
	if err != nil && sqlgraph.IsConstraintError(err) {
		err = &ConstraintError{msg: err.Error(), wrap: err}
	}
	gsd.mutation.done = true
	return affected, err
}

// GroupSettingsDeleteOne is the builder for deleting a single GroupSettings entity.
type GroupSettingsDeleteOne struct {
	gsd *GroupSettingsDelete
}

// Where appends a list predicates to the GroupSettingsDelete builder.
func (gsdo *GroupSettingsDeleteOne) Where(ps ...predicate.GroupSettings) *GroupSettingsDeleteOne {
	gsdo.gsd.mutation.Where(ps...)
	return gsdo
}

// Exec executes the deletion query.
func (gsdo *GroupSettingsDeleteOne) Exec(ctx context.Context) error {
	n, err := gsdo.gsd.Exec(ctx)
	switch {
	case err != nil:
		return err
	case n == 0:
		return &NotFoundError{groupsettings.Label}
	default:
		return nil
	}
}

// ExecX is like Exec, but panics if an error occurs.
func (gsdo *GroupSettingsDeleteOne) ExecX(ctx context.Context) {
	if err := gsdo.Exec(ctx); err != nil {
		panic(err)
	}
}

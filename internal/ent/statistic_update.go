// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/open-boardgame-stats/backend/internal/ent/match"
	"github.com/open-boardgame-stats/backend/internal/ent/player"
	"github.com/open-boardgame-stats/backend/internal/ent/predicate"
	"github.com/open-boardgame-stats/backend/internal/ent/schema/guidgql"
	"github.com/open-boardgame-stats/backend/internal/ent/statdescription"
	"github.com/open-boardgame-stats/backend/internal/ent/statistic"
)

// StatisticUpdate is the builder for updating Statistic entities.
type StatisticUpdate struct {
	config
	hooks    []Hook
	mutation *StatisticMutation
}

// Where appends a list predicates to the StatisticUpdate builder.
func (su *StatisticUpdate) Where(ps ...predicate.Statistic) *StatisticUpdate {
	su.mutation.Where(ps...)
	return su
}

// SetValue sets the "value" field.
func (su *StatisticUpdate) SetValue(s string) *StatisticUpdate {
	su.mutation.SetValue(s)
	return su
}

// SetNillableValue sets the "value" field if the given value is not nil.
func (su *StatisticUpdate) SetNillableValue(s *string) *StatisticUpdate {
	if s != nil {
		su.SetValue(*s)
	}
	return su
}

// SetMatchID sets the "match" edge to the Match entity by ID.
func (su *StatisticUpdate) SetMatchID(id guidgql.GUID) *StatisticUpdate {
	su.mutation.SetMatchID(id)
	return su
}

// SetMatch sets the "match" edge to the Match entity.
func (su *StatisticUpdate) SetMatch(m *Match) *StatisticUpdate {
	return su.SetMatchID(m.ID)
}

// SetStatDescriptionID sets the "stat_description" edge to the StatDescription entity by ID.
func (su *StatisticUpdate) SetStatDescriptionID(id guidgql.GUID) *StatisticUpdate {
	su.mutation.SetStatDescriptionID(id)
	return su
}

// SetStatDescription sets the "stat_description" edge to the StatDescription entity.
func (su *StatisticUpdate) SetStatDescription(s *StatDescription) *StatisticUpdate {
	return su.SetStatDescriptionID(s.ID)
}

// SetPlayerID sets the "player" edge to the Player entity by ID.
func (su *StatisticUpdate) SetPlayerID(id guidgql.GUID) *StatisticUpdate {
	su.mutation.SetPlayerID(id)
	return su
}

// SetPlayer sets the "player" edge to the Player entity.
func (su *StatisticUpdate) SetPlayer(p *Player) *StatisticUpdate {
	return su.SetPlayerID(p.ID)
}

// Mutation returns the StatisticMutation object of the builder.
func (su *StatisticUpdate) Mutation() *StatisticMutation {
	return su.mutation
}

// ClearMatch clears the "match" edge to the Match entity.
func (su *StatisticUpdate) ClearMatch() *StatisticUpdate {
	su.mutation.ClearMatch()
	return su
}

// ClearStatDescription clears the "stat_description" edge to the StatDescription entity.
func (su *StatisticUpdate) ClearStatDescription() *StatisticUpdate {
	su.mutation.ClearStatDescription()
	return su
}

// ClearPlayer clears the "player" edge to the Player entity.
func (su *StatisticUpdate) ClearPlayer() *StatisticUpdate {
	su.mutation.ClearPlayer()
	return su
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (su *StatisticUpdate) Save(ctx context.Context) (int, error) {
	return withHooks(ctx, su.sqlSave, su.mutation, su.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (su *StatisticUpdate) SaveX(ctx context.Context) int {
	affected, err := su.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (su *StatisticUpdate) Exec(ctx context.Context) error {
	_, err := su.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (su *StatisticUpdate) ExecX(ctx context.Context) {
	if err := su.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (su *StatisticUpdate) check() error {
	if _, ok := su.mutation.MatchID(); su.mutation.MatchCleared() && !ok {
		return errors.New(`ent: clearing a required unique edge "Statistic.match"`)
	}
	if _, ok := su.mutation.StatDescriptionID(); su.mutation.StatDescriptionCleared() && !ok {
		return errors.New(`ent: clearing a required unique edge "Statistic.stat_description"`)
	}
	if _, ok := su.mutation.PlayerID(); su.mutation.PlayerCleared() && !ok {
		return errors.New(`ent: clearing a required unique edge "Statistic.player"`)
	}
	return nil
}

func (su *StatisticUpdate) sqlSave(ctx context.Context) (n int, err error) {
	if err := su.check(); err != nil {
		return n, err
	}
	_spec := sqlgraph.NewUpdateSpec(statistic.Table, statistic.Columns, sqlgraph.NewFieldSpec(statistic.FieldID, field.TypeString))
	if ps := su.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := su.mutation.Value(); ok {
		_spec.SetField(statistic.FieldValue, field.TypeString, value)
	}
	if su.mutation.MatchCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   statistic.MatchTable,
			Columns: []string{statistic.MatchColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(match.FieldID, field.TypeString),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := su.mutation.MatchIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   statistic.MatchTable,
			Columns: []string{statistic.MatchColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(match.FieldID, field.TypeString),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if su.mutation.StatDescriptionCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   statistic.StatDescriptionTable,
			Columns: []string{statistic.StatDescriptionColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(statdescription.FieldID, field.TypeString),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := su.mutation.StatDescriptionIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   statistic.StatDescriptionTable,
			Columns: []string{statistic.StatDescriptionColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(statdescription.FieldID, field.TypeString),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if su.mutation.PlayerCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   statistic.PlayerTable,
			Columns: []string{statistic.PlayerColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(player.FieldID, field.TypeString),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := su.mutation.PlayerIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   statistic.PlayerTable,
			Columns: []string{statistic.PlayerColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(player.FieldID, field.TypeString),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, su.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{statistic.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	su.mutation.done = true
	return n, nil
}

// StatisticUpdateOne is the builder for updating a single Statistic entity.
type StatisticUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *StatisticMutation
}

// SetValue sets the "value" field.
func (suo *StatisticUpdateOne) SetValue(s string) *StatisticUpdateOne {
	suo.mutation.SetValue(s)
	return suo
}

// SetNillableValue sets the "value" field if the given value is not nil.
func (suo *StatisticUpdateOne) SetNillableValue(s *string) *StatisticUpdateOne {
	if s != nil {
		suo.SetValue(*s)
	}
	return suo
}

// SetMatchID sets the "match" edge to the Match entity by ID.
func (suo *StatisticUpdateOne) SetMatchID(id guidgql.GUID) *StatisticUpdateOne {
	suo.mutation.SetMatchID(id)
	return suo
}

// SetMatch sets the "match" edge to the Match entity.
func (suo *StatisticUpdateOne) SetMatch(m *Match) *StatisticUpdateOne {
	return suo.SetMatchID(m.ID)
}

// SetStatDescriptionID sets the "stat_description" edge to the StatDescription entity by ID.
func (suo *StatisticUpdateOne) SetStatDescriptionID(id guidgql.GUID) *StatisticUpdateOne {
	suo.mutation.SetStatDescriptionID(id)
	return suo
}

// SetStatDescription sets the "stat_description" edge to the StatDescription entity.
func (suo *StatisticUpdateOne) SetStatDescription(s *StatDescription) *StatisticUpdateOne {
	return suo.SetStatDescriptionID(s.ID)
}

// SetPlayerID sets the "player" edge to the Player entity by ID.
func (suo *StatisticUpdateOne) SetPlayerID(id guidgql.GUID) *StatisticUpdateOne {
	suo.mutation.SetPlayerID(id)
	return suo
}

// SetPlayer sets the "player" edge to the Player entity.
func (suo *StatisticUpdateOne) SetPlayer(p *Player) *StatisticUpdateOne {
	return suo.SetPlayerID(p.ID)
}

// Mutation returns the StatisticMutation object of the builder.
func (suo *StatisticUpdateOne) Mutation() *StatisticMutation {
	return suo.mutation
}

// ClearMatch clears the "match" edge to the Match entity.
func (suo *StatisticUpdateOne) ClearMatch() *StatisticUpdateOne {
	suo.mutation.ClearMatch()
	return suo
}

// ClearStatDescription clears the "stat_description" edge to the StatDescription entity.
func (suo *StatisticUpdateOne) ClearStatDescription() *StatisticUpdateOne {
	suo.mutation.ClearStatDescription()
	return suo
}

// ClearPlayer clears the "player" edge to the Player entity.
func (suo *StatisticUpdateOne) ClearPlayer() *StatisticUpdateOne {
	suo.mutation.ClearPlayer()
	return suo
}

// Where appends a list predicates to the StatisticUpdate builder.
func (suo *StatisticUpdateOne) Where(ps ...predicate.Statistic) *StatisticUpdateOne {
	suo.mutation.Where(ps...)
	return suo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (suo *StatisticUpdateOne) Select(field string, fields ...string) *StatisticUpdateOne {
	suo.fields = append([]string{field}, fields...)
	return suo
}

// Save executes the query and returns the updated Statistic entity.
func (suo *StatisticUpdateOne) Save(ctx context.Context) (*Statistic, error) {
	return withHooks(ctx, suo.sqlSave, suo.mutation, suo.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (suo *StatisticUpdateOne) SaveX(ctx context.Context) *Statistic {
	node, err := suo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (suo *StatisticUpdateOne) Exec(ctx context.Context) error {
	_, err := suo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (suo *StatisticUpdateOne) ExecX(ctx context.Context) {
	if err := suo.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (suo *StatisticUpdateOne) check() error {
	if _, ok := suo.mutation.MatchID(); suo.mutation.MatchCleared() && !ok {
		return errors.New(`ent: clearing a required unique edge "Statistic.match"`)
	}
	if _, ok := suo.mutation.StatDescriptionID(); suo.mutation.StatDescriptionCleared() && !ok {
		return errors.New(`ent: clearing a required unique edge "Statistic.stat_description"`)
	}
	if _, ok := suo.mutation.PlayerID(); suo.mutation.PlayerCleared() && !ok {
		return errors.New(`ent: clearing a required unique edge "Statistic.player"`)
	}
	return nil
}

func (suo *StatisticUpdateOne) sqlSave(ctx context.Context) (_node *Statistic, err error) {
	if err := suo.check(); err != nil {
		return _node, err
	}
	_spec := sqlgraph.NewUpdateSpec(statistic.Table, statistic.Columns, sqlgraph.NewFieldSpec(statistic.FieldID, field.TypeString))
	id, ok := suo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "Statistic.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := suo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, statistic.FieldID)
		for _, f := range fields {
			if !statistic.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != statistic.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := suo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := suo.mutation.Value(); ok {
		_spec.SetField(statistic.FieldValue, field.TypeString, value)
	}
	if suo.mutation.MatchCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   statistic.MatchTable,
			Columns: []string{statistic.MatchColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(match.FieldID, field.TypeString),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := suo.mutation.MatchIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   statistic.MatchTable,
			Columns: []string{statistic.MatchColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(match.FieldID, field.TypeString),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if suo.mutation.StatDescriptionCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   statistic.StatDescriptionTable,
			Columns: []string{statistic.StatDescriptionColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(statdescription.FieldID, field.TypeString),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := suo.mutation.StatDescriptionIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   statistic.StatDescriptionTable,
			Columns: []string{statistic.StatDescriptionColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(statdescription.FieldID, field.TypeString),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if suo.mutation.PlayerCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   statistic.PlayerTable,
			Columns: []string{statistic.PlayerColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(player.FieldID, field.TypeString),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := suo.mutation.PlayerIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   statistic.PlayerTable,
			Columns: []string{statistic.PlayerColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(player.FieldID, field.TypeString),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	_node = &Statistic{config: suo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, suo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{statistic.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	suo.mutation.done = true
	return _node, nil
}

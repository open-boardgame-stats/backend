// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/open-boardgame-stats/backend/internal/ent/game"
	"github.com/open-boardgame-stats/backend/internal/ent/gameversion"
	"github.com/open-boardgame-stats/backend/internal/ent/match"
	"github.com/open-boardgame-stats/backend/internal/ent/predicate"
	"github.com/open-boardgame-stats/backend/internal/ent/schema/guidgql"
	"github.com/open-boardgame-stats/backend/internal/ent/statdescription"
)

// GameVersionUpdate is the builder for updating GameVersion entities.
type GameVersionUpdate struct {
	config
	hooks    []Hook
	mutation *GameVersionMutation
}

// Where appends a list predicates to the GameVersionUpdate builder.
func (gvu *GameVersionUpdate) Where(ps ...predicate.GameVersion) *GameVersionUpdate {
	gvu.mutation.Where(ps...)
	return gvu
}

// SetVersionNumber sets the "version_number" field.
func (gvu *GameVersionUpdate) SetVersionNumber(i int) *GameVersionUpdate {
	gvu.mutation.ResetVersionNumber()
	gvu.mutation.SetVersionNumber(i)
	return gvu
}

// SetNillableVersionNumber sets the "version_number" field if the given value is not nil.
func (gvu *GameVersionUpdate) SetNillableVersionNumber(i *int) *GameVersionUpdate {
	if i != nil {
		gvu.SetVersionNumber(*i)
	}
	return gvu
}

// AddVersionNumber adds i to the "version_number" field.
func (gvu *GameVersionUpdate) AddVersionNumber(i int) *GameVersionUpdate {
	gvu.mutation.AddVersionNumber(i)
	return gvu
}

// SetGameID sets the "game" edge to the Game entity by ID.
func (gvu *GameVersionUpdate) SetGameID(id guidgql.GUID) *GameVersionUpdate {
	gvu.mutation.SetGameID(id)
	return gvu
}

// SetGame sets the "game" edge to the Game entity.
func (gvu *GameVersionUpdate) SetGame(g *Game) *GameVersionUpdate {
	return gvu.SetGameID(g.ID)
}

// AddStatDescriptionIDs adds the "stat_descriptions" edge to the StatDescription entity by IDs.
func (gvu *GameVersionUpdate) AddStatDescriptionIDs(ids ...guidgql.GUID) *GameVersionUpdate {
	gvu.mutation.AddStatDescriptionIDs(ids...)
	return gvu
}

// AddStatDescriptions adds the "stat_descriptions" edges to the StatDescription entity.
func (gvu *GameVersionUpdate) AddStatDescriptions(s ...*StatDescription) *GameVersionUpdate {
	ids := make([]guidgql.GUID, len(s))
	for i := range s {
		ids[i] = s[i].ID
	}
	return gvu.AddStatDescriptionIDs(ids...)
}

// AddMatchIDs adds the "matches" edge to the Match entity by IDs.
func (gvu *GameVersionUpdate) AddMatchIDs(ids ...guidgql.GUID) *GameVersionUpdate {
	gvu.mutation.AddMatchIDs(ids...)
	return gvu
}

// AddMatches adds the "matches" edges to the Match entity.
func (gvu *GameVersionUpdate) AddMatches(m ...*Match) *GameVersionUpdate {
	ids := make([]guidgql.GUID, len(m))
	for i := range m {
		ids[i] = m[i].ID
	}
	return gvu.AddMatchIDs(ids...)
}

// Mutation returns the GameVersionMutation object of the builder.
func (gvu *GameVersionUpdate) Mutation() *GameVersionMutation {
	return gvu.mutation
}

// ClearGame clears the "game" edge to the Game entity.
func (gvu *GameVersionUpdate) ClearGame() *GameVersionUpdate {
	gvu.mutation.ClearGame()
	return gvu
}

// ClearStatDescriptions clears all "stat_descriptions" edges to the StatDescription entity.
func (gvu *GameVersionUpdate) ClearStatDescriptions() *GameVersionUpdate {
	gvu.mutation.ClearStatDescriptions()
	return gvu
}

// RemoveStatDescriptionIDs removes the "stat_descriptions" edge to StatDescription entities by IDs.
func (gvu *GameVersionUpdate) RemoveStatDescriptionIDs(ids ...guidgql.GUID) *GameVersionUpdate {
	gvu.mutation.RemoveStatDescriptionIDs(ids...)
	return gvu
}

// RemoveStatDescriptions removes "stat_descriptions" edges to StatDescription entities.
func (gvu *GameVersionUpdate) RemoveStatDescriptions(s ...*StatDescription) *GameVersionUpdate {
	ids := make([]guidgql.GUID, len(s))
	for i := range s {
		ids[i] = s[i].ID
	}
	return gvu.RemoveStatDescriptionIDs(ids...)
}

// ClearMatches clears all "matches" edges to the Match entity.
func (gvu *GameVersionUpdate) ClearMatches() *GameVersionUpdate {
	gvu.mutation.ClearMatches()
	return gvu
}

// RemoveMatchIDs removes the "matches" edge to Match entities by IDs.
func (gvu *GameVersionUpdate) RemoveMatchIDs(ids ...guidgql.GUID) *GameVersionUpdate {
	gvu.mutation.RemoveMatchIDs(ids...)
	return gvu
}

// RemoveMatches removes "matches" edges to Match entities.
func (gvu *GameVersionUpdate) RemoveMatches(m ...*Match) *GameVersionUpdate {
	ids := make([]guidgql.GUID, len(m))
	for i := range m {
		ids[i] = m[i].ID
	}
	return gvu.RemoveMatchIDs(ids...)
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (gvu *GameVersionUpdate) Save(ctx context.Context) (int, error) {
	return withHooks[int, GameVersionMutation](ctx, gvu.sqlSave, gvu.mutation, gvu.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (gvu *GameVersionUpdate) SaveX(ctx context.Context) int {
	affected, err := gvu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (gvu *GameVersionUpdate) Exec(ctx context.Context) error {
	_, err := gvu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (gvu *GameVersionUpdate) ExecX(ctx context.Context) {
	if err := gvu.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (gvu *GameVersionUpdate) check() error {
	if _, ok := gvu.mutation.GameID(); gvu.mutation.GameCleared() && !ok {
		return errors.New(`ent: clearing a required unique edge "GameVersion.game"`)
	}
	return nil
}

func (gvu *GameVersionUpdate) sqlSave(ctx context.Context) (n int, err error) {
	if err := gvu.check(); err != nil {
		return n, err
	}
	_spec := sqlgraph.NewUpdateSpec(gameversion.Table, gameversion.Columns, sqlgraph.NewFieldSpec(gameversion.FieldID, field.TypeString))
	if ps := gvu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := gvu.mutation.VersionNumber(); ok {
		_spec.SetField(gameversion.FieldVersionNumber, field.TypeInt, value)
	}
	if value, ok := gvu.mutation.AddedVersionNumber(); ok {
		_spec.AddField(gameversion.FieldVersionNumber, field.TypeInt, value)
	}
	if gvu.mutation.GameCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: false,
			Table:   gameversion.GameTable,
			Columns: []string{gameversion.GameColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(game.FieldID, field.TypeString),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := gvu.mutation.GameIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: false,
			Table:   gameversion.GameTable,
			Columns: []string{gameversion.GameColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(game.FieldID, field.TypeString),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if gvu.mutation.StatDescriptionsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: true,
			Table:   gameversion.StatDescriptionsTable,
			Columns: gameversion.StatDescriptionsPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(statdescription.FieldID, field.TypeString),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := gvu.mutation.RemovedStatDescriptionsIDs(); len(nodes) > 0 && !gvu.mutation.StatDescriptionsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: true,
			Table:   gameversion.StatDescriptionsTable,
			Columns: gameversion.StatDescriptionsPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(statdescription.FieldID, field.TypeString),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := gvu.mutation.StatDescriptionsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: true,
			Table:   gameversion.StatDescriptionsTable,
			Columns: gameversion.StatDescriptionsPrimaryKey,
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
	if gvu.mutation.MatchesCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   gameversion.MatchesTable,
			Columns: []string{gameversion.MatchesColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(match.FieldID, field.TypeString),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := gvu.mutation.RemovedMatchesIDs(); len(nodes) > 0 && !gvu.mutation.MatchesCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   gameversion.MatchesTable,
			Columns: []string{gameversion.MatchesColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(match.FieldID, field.TypeString),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := gvu.mutation.MatchesIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   gameversion.MatchesTable,
			Columns: []string{gameversion.MatchesColumn},
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
	if n, err = sqlgraph.UpdateNodes(ctx, gvu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{gameversion.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	gvu.mutation.done = true
	return n, nil
}

// GameVersionUpdateOne is the builder for updating a single GameVersion entity.
type GameVersionUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *GameVersionMutation
}

// SetVersionNumber sets the "version_number" field.
func (gvuo *GameVersionUpdateOne) SetVersionNumber(i int) *GameVersionUpdateOne {
	gvuo.mutation.ResetVersionNumber()
	gvuo.mutation.SetVersionNumber(i)
	return gvuo
}

// SetNillableVersionNumber sets the "version_number" field if the given value is not nil.
func (gvuo *GameVersionUpdateOne) SetNillableVersionNumber(i *int) *GameVersionUpdateOne {
	if i != nil {
		gvuo.SetVersionNumber(*i)
	}
	return gvuo
}

// AddVersionNumber adds i to the "version_number" field.
func (gvuo *GameVersionUpdateOne) AddVersionNumber(i int) *GameVersionUpdateOne {
	gvuo.mutation.AddVersionNumber(i)
	return gvuo
}

// SetGameID sets the "game" edge to the Game entity by ID.
func (gvuo *GameVersionUpdateOne) SetGameID(id guidgql.GUID) *GameVersionUpdateOne {
	gvuo.mutation.SetGameID(id)
	return gvuo
}

// SetGame sets the "game" edge to the Game entity.
func (gvuo *GameVersionUpdateOne) SetGame(g *Game) *GameVersionUpdateOne {
	return gvuo.SetGameID(g.ID)
}

// AddStatDescriptionIDs adds the "stat_descriptions" edge to the StatDescription entity by IDs.
func (gvuo *GameVersionUpdateOne) AddStatDescriptionIDs(ids ...guidgql.GUID) *GameVersionUpdateOne {
	gvuo.mutation.AddStatDescriptionIDs(ids...)
	return gvuo
}

// AddStatDescriptions adds the "stat_descriptions" edges to the StatDescription entity.
func (gvuo *GameVersionUpdateOne) AddStatDescriptions(s ...*StatDescription) *GameVersionUpdateOne {
	ids := make([]guidgql.GUID, len(s))
	for i := range s {
		ids[i] = s[i].ID
	}
	return gvuo.AddStatDescriptionIDs(ids...)
}

// AddMatchIDs adds the "matches" edge to the Match entity by IDs.
func (gvuo *GameVersionUpdateOne) AddMatchIDs(ids ...guidgql.GUID) *GameVersionUpdateOne {
	gvuo.mutation.AddMatchIDs(ids...)
	return gvuo
}

// AddMatches adds the "matches" edges to the Match entity.
func (gvuo *GameVersionUpdateOne) AddMatches(m ...*Match) *GameVersionUpdateOne {
	ids := make([]guidgql.GUID, len(m))
	for i := range m {
		ids[i] = m[i].ID
	}
	return gvuo.AddMatchIDs(ids...)
}

// Mutation returns the GameVersionMutation object of the builder.
func (gvuo *GameVersionUpdateOne) Mutation() *GameVersionMutation {
	return gvuo.mutation
}

// ClearGame clears the "game" edge to the Game entity.
func (gvuo *GameVersionUpdateOne) ClearGame() *GameVersionUpdateOne {
	gvuo.mutation.ClearGame()
	return gvuo
}

// ClearStatDescriptions clears all "stat_descriptions" edges to the StatDescription entity.
func (gvuo *GameVersionUpdateOne) ClearStatDescriptions() *GameVersionUpdateOne {
	gvuo.mutation.ClearStatDescriptions()
	return gvuo
}

// RemoveStatDescriptionIDs removes the "stat_descriptions" edge to StatDescription entities by IDs.
func (gvuo *GameVersionUpdateOne) RemoveStatDescriptionIDs(ids ...guidgql.GUID) *GameVersionUpdateOne {
	gvuo.mutation.RemoveStatDescriptionIDs(ids...)
	return gvuo
}

// RemoveStatDescriptions removes "stat_descriptions" edges to StatDescription entities.
func (gvuo *GameVersionUpdateOne) RemoveStatDescriptions(s ...*StatDescription) *GameVersionUpdateOne {
	ids := make([]guidgql.GUID, len(s))
	for i := range s {
		ids[i] = s[i].ID
	}
	return gvuo.RemoveStatDescriptionIDs(ids...)
}

// ClearMatches clears all "matches" edges to the Match entity.
func (gvuo *GameVersionUpdateOne) ClearMatches() *GameVersionUpdateOne {
	gvuo.mutation.ClearMatches()
	return gvuo
}

// RemoveMatchIDs removes the "matches" edge to Match entities by IDs.
func (gvuo *GameVersionUpdateOne) RemoveMatchIDs(ids ...guidgql.GUID) *GameVersionUpdateOne {
	gvuo.mutation.RemoveMatchIDs(ids...)
	return gvuo
}

// RemoveMatches removes "matches" edges to Match entities.
func (gvuo *GameVersionUpdateOne) RemoveMatches(m ...*Match) *GameVersionUpdateOne {
	ids := make([]guidgql.GUID, len(m))
	for i := range m {
		ids[i] = m[i].ID
	}
	return gvuo.RemoveMatchIDs(ids...)
}

// Where appends a list predicates to the GameVersionUpdate builder.
func (gvuo *GameVersionUpdateOne) Where(ps ...predicate.GameVersion) *GameVersionUpdateOne {
	gvuo.mutation.Where(ps...)
	return gvuo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (gvuo *GameVersionUpdateOne) Select(field string, fields ...string) *GameVersionUpdateOne {
	gvuo.fields = append([]string{field}, fields...)
	return gvuo
}

// Save executes the query and returns the updated GameVersion entity.
func (gvuo *GameVersionUpdateOne) Save(ctx context.Context) (*GameVersion, error) {
	return withHooks[*GameVersion, GameVersionMutation](ctx, gvuo.sqlSave, gvuo.mutation, gvuo.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (gvuo *GameVersionUpdateOne) SaveX(ctx context.Context) *GameVersion {
	node, err := gvuo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (gvuo *GameVersionUpdateOne) Exec(ctx context.Context) error {
	_, err := gvuo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (gvuo *GameVersionUpdateOne) ExecX(ctx context.Context) {
	if err := gvuo.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (gvuo *GameVersionUpdateOne) check() error {
	if _, ok := gvuo.mutation.GameID(); gvuo.mutation.GameCleared() && !ok {
		return errors.New(`ent: clearing a required unique edge "GameVersion.game"`)
	}
	return nil
}

func (gvuo *GameVersionUpdateOne) sqlSave(ctx context.Context) (_node *GameVersion, err error) {
	if err := gvuo.check(); err != nil {
		return _node, err
	}
	_spec := sqlgraph.NewUpdateSpec(gameversion.Table, gameversion.Columns, sqlgraph.NewFieldSpec(gameversion.FieldID, field.TypeString))
	id, ok := gvuo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "GameVersion.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := gvuo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, gameversion.FieldID)
		for _, f := range fields {
			if !gameversion.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != gameversion.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := gvuo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := gvuo.mutation.VersionNumber(); ok {
		_spec.SetField(gameversion.FieldVersionNumber, field.TypeInt, value)
	}
	if value, ok := gvuo.mutation.AddedVersionNumber(); ok {
		_spec.AddField(gameversion.FieldVersionNumber, field.TypeInt, value)
	}
	if gvuo.mutation.GameCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: false,
			Table:   gameversion.GameTable,
			Columns: []string{gameversion.GameColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(game.FieldID, field.TypeString),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := gvuo.mutation.GameIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: false,
			Table:   gameversion.GameTable,
			Columns: []string{gameversion.GameColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(game.FieldID, field.TypeString),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if gvuo.mutation.StatDescriptionsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: true,
			Table:   gameversion.StatDescriptionsTable,
			Columns: gameversion.StatDescriptionsPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(statdescription.FieldID, field.TypeString),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := gvuo.mutation.RemovedStatDescriptionsIDs(); len(nodes) > 0 && !gvuo.mutation.StatDescriptionsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: true,
			Table:   gameversion.StatDescriptionsTable,
			Columns: gameversion.StatDescriptionsPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(statdescription.FieldID, field.TypeString),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := gvuo.mutation.StatDescriptionsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: true,
			Table:   gameversion.StatDescriptionsTable,
			Columns: gameversion.StatDescriptionsPrimaryKey,
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
	if gvuo.mutation.MatchesCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   gameversion.MatchesTable,
			Columns: []string{gameversion.MatchesColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(match.FieldID, field.TypeString),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := gvuo.mutation.RemovedMatchesIDs(); len(nodes) > 0 && !gvuo.mutation.MatchesCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   gameversion.MatchesTable,
			Columns: []string{gameversion.MatchesColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(match.FieldID, field.TypeString),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := gvuo.mutation.MatchesIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   gameversion.MatchesTable,
			Columns: []string{gameversion.MatchesColumn},
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
	_node = &GameVersion{config: gvuo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, gvuo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{gameversion.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	gvuo.mutation.done = true
	return _node, nil
}

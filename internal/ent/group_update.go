// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/open-boardgame-stats/backend/internal/ent/group"
	"github.com/open-boardgame-stats/backend/internal/ent/groupmembership"
	"github.com/open-boardgame-stats/backend/internal/ent/groupmembershipapplication"
	"github.com/open-boardgame-stats/backend/internal/ent/groupsettings"
	"github.com/open-boardgame-stats/backend/internal/ent/predicate"
	"github.com/open-boardgame-stats/backend/internal/ent/schema/guidgql"
)

// GroupUpdate is the builder for updating Group entities.
type GroupUpdate struct {
	config
	hooks    []Hook
	mutation *GroupMutation
}

// Where appends a list predicates to the GroupUpdate builder.
func (gu *GroupUpdate) Where(ps ...predicate.Group) *GroupUpdate {
	gu.mutation.Where(ps...)
	return gu
}

// SetName sets the "name" field.
func (gu *GroupUpdate) SetName(s string) *GroupUpdate {
	gu.mutation.SetName(s)
	return gu
}

// SetDescription sets the "description" field.
func (gu *GroupUpdate) SetDescription(s string) *GroupUpdate {
	gu.mutation.SetDescription(s)
	return gu
}

// SetNillableDescription sets the "description" field if the given value is not nil.
func (gu *GroupUpdate) SetNillableDescription(s *string) *GroupUpdate {
	if s != nil {
		gu.SetDescription(*s)
	}
	return gu
}

// SetLogoURL sets the "logo_url" field.
func (gu *GroupUpdate) SetLogoURL(s string) *GroupUpdate {
	gu.mutation.SetLogoURL(s)
	return gu
}

// SetSettingsID sets the "settings" edge to the GroupSettings entity by ID.
func (gu *GroupUpdate) SetSettingsID(id guidgql.GUID) *GroupUpdate {
	gu.mutation.SetSettingsID(id)
	return gu
}

// SetSettings sets the "settings" edge to the GroupSettings entity.
func (gu *GroupUpdate) SetSettings(g *GroupSettings) *GroupUpdate {
	return gu.SetSettingsID(g.ID)
}

// AddMemberIDs adds the "members" edge to the GroupMembership entity by IDs.
func (gu *GroupUpdate) AddMemberIDs(ids ...guidgql.GUID) *GroupUpdate {
	gu.mutation.AddMemberIDs(ids...)
	return gu
}

// AddMembers adds the "members" edges to the GroupMembership entity.
func (gu *GroupUpdate) AddMembers(g ...*GroupMembership) *GroupUpdate {
	ids := make([]guidgql.GUID, len(g))
	for i := range g {
		ids[i] = g[i].ID
	}
	return gu.AddMemberIDs(ids...)
}

// AddApplicationIDs adds the "applications" edge to the GroupMembershipApplication entity by IDs.
func (gu *GroupUpdate) AddApplicationIDs(ids ...guidgql.GUID) *GroupUpdate {
	gu.mutation.AddApplicationIDs(ids...)
	return gu
}

// AddApplications adds the "applications" edges to the GroupMembershipApplication entity.
func (gu *GroupUpdate) AddApplications(g ...*GroupMembershipApplication) *GroupUpdate {
	ids := make([]guidgql.GUID, len(g))
	for i := range g {
		ids[i] = g[i].ID
	}
	return gu.AddApplicationIDs(ids...)
}

// Mutation returns the GroupMutation object of the builder.
func (gu *GroupUpdate) Mutation() *GroupMutation {
	return gu.mutation
}

// ClearSettings clears the "settings" edge to the GroupSettings entity.
func (gu *GroupUpdate) ClearSettings() *GroupUpdate {
	gu.mutation.ClearSettings()
	return gu
}

// ClearMembers clears all "members" edges to the GroupMembership entity.
func (gu *GroupUpdate) ClearMembers() *GroupUpdate {
	gu.mutation.ClearMembers()
	return gu
}

// RemoveMemberIDs removes the "members" edge to GroupMembership entities by IDs.
func (gu *GroupUpdate) RemoveMemberIDs(ids ...guidgql.GUID) *GroupUpdate {
	gu.mutation.RemoveMemberIDs(ids...)
	return gu
}

// RemoveMembers removes "members" edges to GroupMembership entities.
func (gu *GroupUpdate) RemoveMembers(g ...*GroupMembership) *GroupUpdate {
	ids := make([]guidgql.GUID, len(g))
	for i := range g {
		ids[i] = g[i].ID
	}
	return gu.RemoveMemberIDs(ids...)
}

// ClearApplications clears all "applications" edges to the GroupMembershipApplication entity.
func (gu *GroupUpdate) ClearApplications() *GroupUpdate {
	gu.mutation.ClearApplications()
	return gu
}

// RemoveApplicationIDs removes the "applications" edge to GroupMembershipApplication entities by IDs.
func (gu *GroupUpdate) RemoveApplicationIDs(ids ...guidgql.GUID) *GroupUpdate {
	gu.mutation.RemoveApplicationIDs(ids...)
	return gu
}

// RemoveApplications removes "applications" edges to GroupMembershipApplication entities.
func (gu *GroupUpdate) RemoveApplications(g ...*GroupMembershipApplication) *GroupUpdate {
	ids := make([]guidgql.GUID, len(g))
	for i := range g {
		ids[i] = g[i].ID
	}
	return gu.RemoveApplicationIDs(ids...)
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (gu *GroupUpdate) Save(ctx context.Context) (int, error) {
	return withHooks[int, GroupMutation](ctx, gu.sqlSave, gu.mutation, gu.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (gu *GroupUpdate) SaveX(ctx context.Context) int {
	affected, err := gu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (gu *GroupUpdate) Exec(ctx context.Context) error {
	_, err := gu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (gu *GroupUpdate) ExecX(ctx context.Context) {
	if err := gu.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (gu *GroupUpdate) check() error {
	if v, ok := gu.mutation.Name(); ok {
		if err := group.NameValidator(v); err != nil {
			return &ValidationError{Name: "name", err: fmt.Errorf(`ent: validator failed for field "Group.name": %w`, err)}
		}
	}
	if v, ok := gu.mutation.LogoURL(); ok {
		if err := group.LogoURLValidator(v); err != nil {
			return &ValidationError{Name: "logo_url", err: fmt.Errorf(`ent: validator failed for field "Group.logo_url": %w`, err)}
		}
	}
	if _, ok := gu.mutation.SettingsID(); gu.mutation.SettingsCleared() && !ok {
		return errors.New(`ent: clearing a required unique edge "Group.settings"`)
	}
	return nil
}

func (gu *GroupUpdate) sqlSave(ctx context.Context) (n int, err error) {
	if err := gu.check(); err != nil {
		return n, err
	}
	_spec := sqlgraph.NewUpdateSpec(group.Table, group.Columns, sqlgraph.NewFieldSpec(group.FieldID, field.TypeString))
	if ps := gu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := gu.mutation.Name(); ok {
		_spec.SetField(group.FieldName, field.TypeString, value)
	}
	if value, ok := gu.mutation.Description(); ok {
		_spec.SetField(group.FieldDescription, field.TypeString, value)
	}
	if value, ok := gu.mutation.LogoURL(); ok {
		_spec.SetField(group.FieldLogoURL, field.TypeString, value)
	}
	if gu.mutation.SettingsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2O,
			Inverse: false,
			Table:   group.SettingsTable,
			Columns: []string{group.SettingsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(groupsettings.FieldID, field.TypeString),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := gu.mutation.SettingsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2O,
			Inverse: false,
			Table:   group.SettingsTable,
			Columns: []string{group.SettingsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(groupsettings.FieldID, field.TypeString),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if gu.mutation.MembersCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   group.MembersTable,
			Columns: []string{group.MembersColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(groupmembership.FieldID, field.TypeString),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := gu.mutation.RemovedMembersIDs(); len(nodes) > 0 && !gu.mutation.MembersCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   group.MembersTable,
			Columns: []string{group.MembersColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(groupmembership.FieldID, field.TypeString),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := gu.mutation.MembersIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   group.MembersTable,
			Columns: []string{group.MembersColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(groupmembership.FieldID, field.TypeString),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if gu.mutation.ApplicationsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   group.ApplicationsTable,
			Columns: []string{group.ApplicationsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(groupmembershipapplication.FieldID, field.TypeString),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := gu.mutation.RemovedApplicationsIDs(); len(nodes) > 0 && !gu.mutation.ApplicationsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   group.ApplicationsTable,
			Columns: []string{group.ApplicationsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(groupmembershipapplication.FieldID, field.TypeString),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := gu.mutation.ApplicationsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   group.ApplicationsTable,
			Columns: []string{group.ApplicationsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(groupmembershipapplication.FieldID, field.TypeString),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, gu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{group.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	gu.mutation.done = true
	return n, nil
}

// GroupUpdateOne is the builder for updating a single Group entity.
type GroupUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *GroupMutation
}

// SetName sets the "name" field.
func (guo *GroupUpdateOne) SetName(s string) *GroupUpdateOne {
	guo.mutation.SetName(s)
	return guo
}

// SetDescription sets the "description" field.
func (guo *GroupUpdateOne) SetDescription(s string) *GroupUpdateOne {
	guo.mutation.SetDescription(s)
	return guo
}

// SetNillableDescription sets the "description" field if the given value is not nil.
func (guo *GroupUpdateOne) SetNillableDescription(s *string) *GroupUpdateOne {
	if s != nil {
		guo.SetDescription(*s)
	}
	return guo
}

// SetLogoURL sets the "logo_url" field.
func (guo *GroupUpdateOne) SetLogoURL(s string) *GroupUpdateOne {
	guo.mutation.SetLogoURL(s)
	return guo
}

// SetSettingsID sets the "settings" edge to the GroupSettings entity by ID.
func (guo *GroupUpdateOne) SetSettingsID(id guidgql.GUID) *GroupUpdateOne {
	guo.mutation.SetSettingsID(id)
	return guo
}

// SetSettings sets the "settings" edge to the GroupSettings entity.
func (guo *GroupUpdateOne) SetSettings(g *GroupSettings) *GroupUpdateOne {
	return guo.SetSettingsID(g.ID)
}

// AddMemberIDs adds the "members" edge to the GroupMembership entity by IDs.
func (guo *GroupUpdateOne) AddMemberIDs(ids ...guidgql.GUID) *GroupUpdateOne {
	guo.mutation.AddMemberIDs(ids...)
	return guo
}

// AddMembers adds the "members" edges to the GroupMembership entity.
func (guo *GroupUpdateOne) AddMembers(g ...*GroupMembership) *GroupUpdateOne {
	ids := make([]guidgql.GUID, len(g))
	for i := range g {
		ids[i] = g[i].ID
	}
	return guo.AddMemberIDs(ids...)
}

// AddApplicationIDs adds the "applications" edge to the GroupMembershipApplication entity by IDs.
func (guo *GroupUpdateOne) AddApplicationIDs(ids ...guidgql.GUID) *GroupUpdateOne {
	guo.mutation.AddApplicationIDs(ids...)
	return guo
}

// AddApplications adds the "applications" edges to the GroupMembershipApplication entity.
func (guo *GroupUpdateOne) AddApplications(g ...*GroupMembershipApplication) *GroupUpdateOne {
	ids := make([]guidgql.GUID, len(g))
	for i := range g {
		ids[i] = g[i].ID
	}
	return guo.AddApplicationIDs(ids...)
}

// Mutation returns the GroupMutation object of the builder.
func (guo *GroupUpdateOne) Mutation() *GroupMutation {
	return guo.mutation
}

// ClearSettings clears the "settings" edge to the GroupSettings entity.
func (guo *GroupUpdateOne) ClearSettings() *GroupUpdateOne {
	guo.mutation.ClearSettings()
	return guo
}

// ClearMembers clears all "members" edges to the GroupMembership entity.
func (guo *GroupUpdateOne) ClearMembers() *GroupUpdateOne {
	guo.mutation.ClearMembers()
	return guo
}

// RemoveMemberIDs removes the "members" edge to GroupMembership entities by IDs.
func (guo *GroupUpdateOne) RemoveMemberIDs(ids ...guidgql.GUID) *GroupUpdateOne {
	guo.mutation.RemoveMemberIDs(ids...)
	return guo
}

// RemoveMembers removes "members" edges to GroupMembership entities.
func (guo *GroupUpdateOne) RemoveMembers(g ...*GroupMembership) *GroupUpdateOne {
	ids := make([]guidgql.GUID, len(g))
	for i := range g {
		ids[i] = g[i].ID
	}
	return guo.RemoveMemberIDs(ids...)
}

// ClearApplications clears all "applications" edges to the GroupMembershipApplication entity.
func (guo *GroupUpdateOne) ClearApplications() *GroupUpdateOne {
	guo.mutation.ClearApplications()
	return guo
}

// RemoveApplicationIDs removes the "applications" edge to GroupMembershipApplication entities by IDs.
func (guo *GroupUpdateOne) RemoveApplicationIDs(ids ...guidgql.GUID) *GroupUpdateOne {
	guo.mutation.RemoveApplicationIDs(ids...)
	return guo
}

// RemoveApplications removes "applications" edges to GroupMembershipApplication entities.
func (guo *GroupUpdateOne) RemoveApplications(g ...*GroupMembershipApplication) *GroupUpdateOne {
	ids := make([]guidgql.GUID, len(g))
	for i := range g {
		ids[i] = g[i].ID
	}
	return guo.RemoveApplicationIDs(ids...)
}

// Where appends a list predicates to the GroupUpdate builder.
func (guo *GroupUpdateOne) Where(ps ...predicate.Group) *GroupUpdateOne {
	guo.mutation.Where(ps...)
	return guo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (guo *GroupUpdateOne) Select(field string, fields ...string) *GroupUpdateOne {
	guo.fields = append([]string{field}, fields...)
	return guo
}

// Save executes the query and returns the updated Group entity.
func (guo *GroupUpdateOne) Save(ctx context.Context) (*Group, error) {
	return withHooks[*Group, GroupMutation](ctx, guo.sqlSave, guo.mutation, guo.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (guo *GroupUpdateOne) SaveX(ctx context.Context) *Group {
	node, err := guo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (guo *GroupUpdateOne) Exec(ctx context.Context) error {
	_, err := guo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (guo *GroupUpdateOne) ExecX(ctx context.Context) {
	if err := guo.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (guo *GroupUpdateOne) check() error {
	if v, ok := guo.mutation.Name(); ok {
		if err := group.NameValidator(v); err != nil {
			return &ValidationError{Name: "name", err: fmt.Errorf(`ent: validator failed for field "Group.name": %w`, err)}
		}
	}
	if v, ok := guo.mutation.LogoURL(); ok {
		if err := group.LogoURLValidator(v); err != nil {
			return &ValidationError{Name: "logo_url", err: fmt.Errorf(`ent: validator failed for field "Group.logo_url": %w`, err)}
		}
	}
	if _, ok := guo.mutation.SettingsID(); guo.mutation.SettingsCleared() && !ok {
		return errors.New(`ent: clearing a required unique edge "Group.settings"`)
	}
	return nil
}

func (guo *GroupUpdateOne) sqlSave(ctx context.Context) (_node *Group, err error) {
	if err := guo.check(); err != nil {
		return _node, err
	}
	_spec := sqlgraph.NewUpdateSpec(group.Table, group.Columns, sqlgraph.NewFieldSpec(group.FieldID, field.TypeString))
	id, ok := guo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "Group.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := guo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, group.FieldID)
		for _, f := range fields {
			if !group.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != group.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := guo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := guo.mutation.Name(); ok {
		_spec.SetField(group.FieldName, field.TypeString, value)
	}
	if value, ok := guo.mutation.Description(); ok {
		_spec.SetField(group.FieldDescription, field.TypeString, value)
	}
	if value, ok := guo.mutation.LogoURL(); ok {
		_spec.SetField(group.FieldLogoURL, field.TypeString, value)
	}
	if guo.mutation.SettingsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2O,
			Inverse: false,
			Table:   group.SettingsTable,
			Columns: []string{group.SettingsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(groupsettings.FieldID, field.TypeString),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := guo.mutation.SettingsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2O,
			Inverse: false,
			Table:   group.SettingsTable,
			Columns: []string{group.SettingsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(groupsettings.FieldID, field.TypeString),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if guo.mutation.MembersCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   group.MembersTable,
			Columns: []string{group.MembersColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(groupmembership.FieldID, field.TypeString),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := guo.mutation.RemovedMembersIDs(); len(nodes) > 0 && !guo.mutation.MembersCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   group.MembersTable,
			Columns: []string{group.MembersColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(groupmembership.FieldID, field.TypeString),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := guo.mutation.MembersIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   group.MembersTable,
			Columns: []string{group.MembersColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(groupmembership.FieldID, field.TypeString),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if guo.mutation.ApplicationsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   group.ApplicationsTable,
			Columns: []string{group.ApplicationsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(groupmembershipapplication.FieldID, field.TypeString),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := guo.mutation.RemovedApplicationsIDs(); len(nodes) > 0 && !guo.mutation.ApplicationsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   group.ApplicationsTable,
			Columns: []string{group.ApplicationsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(groupmembershipapplication.FieldID, field.TypeString),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := guo.mutation.ApplicationsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   group.ApplicationsTable,
			Columns: []string{group.ApplicationsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(groupmembershipapplication.FieldID, field.TypeString),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	_node = &Group{config: guo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, guo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{group.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	guo.mutation.done = true
	return _node, nil
}

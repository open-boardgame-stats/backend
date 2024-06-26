// Code generated by ent, DO NOT EDIT.

package ent

import (
	"fmt"
	"strings"

	"entgo.io/ent"
	"entgo.io/ent/dialect/sql"
	"github.com/open-boardgame-stats/backend/internal/ent/enums"
	"github.com/open-boardgame-stats/backend/internal/ent/group"
	"github.com/open-boardgame-stats/backend/internal/ent/groupsettings"
	"github.com/open-boardgame-stats/backend/internal/ent/schema/guidgql"
)

// GroupSettings is the model entity for the GroupSettings schema.
type GroupSettings struct {
	config `json:"-"`
	// ID of the ent.
	ID guidgql.GUID `json:"id,omitempty"`
	// Visibility holds the value of the "visibility" field.
	Visibility groupsettings.Visibility `json:"visibility,omitempty"`
	// JoinPolicy holds the value of the "join_policy" field.
	JoinPolicy groupsettings.JoinPolicy `json:"join_policy,omitempty"`
	// MinimumRoleToInvite holds the value of the "minimum_role_to_invite" field.
	MinimumRoleToInvite *enums.Role `json:"minimum_role_to_invite,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the GroupSettingsQuery when eager-loading is set.
	Edges          GroupSettingsEdges `json:"edges"`
	group_settings *guidgql.GUID
	selectValues   sql.SelectValues
}

// GroupSettingsEdges holds the relations/edges for other nodes in the graph.
type GroupSettingsEdges struct {
	// Group holds the value of the group edge.
	Group *Group `json:"group,omitempty"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [1]bool
}

// GroupOrErr returns the Group value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e GroupSettingsEdges) GroupOrErr() (*Group, error) {
	if e.Group != nil {
		return e.Group, nil
	} else if e.loadedTypes[0] {
		return nil, &NotFoundError{label: group.Label}
	}
	return nil, &NotLoadedError{edge: "group"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*GroupSettings) scanValues(columns []string) ([]any, error) {
	values := make([]any, len(columns))
	for i := range columns {
		switch columns[i] {
		case groupsettings.FieldID:
			values[i] = new(guidgql.GUID)
		case groupsettings.FieldVisibility, groupsettings.FieldJoinPolicy, groupsettings.FieldMinimumRoleToInvite:
			values[i] = new(sql.NullString)
		case groupsettings.ForeignKeys[0]: // group_settings
			values[i] = &sql.NullScanner{S: new(guidgql.GUID)}
		default:
			values[i] = new(sql.UnknownType)
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the GroupSettings fields.
func (gs *GroupSettings) assignValues(columns []string, values []any) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case groupsettings.FieldID:
			if value, ok := values[i].(*guidgql.GUID); !ok {
				return fmt.Errorf("unexpected type %T for field id", values[i])
			} else if value != nil {
				gs.ID = *value
			}
		case groupsettings.FieldVisibility:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field visibility", values[i])
			} else if value.Valid {
				gs.Visibility = groupsettings.Visibility(value.String)
			}
		case groupsettings.FieldJoinPolicy:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field join_policy", values[i])
			} else if value.Valid {
				gs.JoinPolicy = groupsettings.JoinPolicy(value.String)
			}
		case groupsettings.FieldMinimumRoleToInvite:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field minimum_role_to_invite", values[i])
			} else if value.Valid {
				gs.MinimumRoleToInvite = new(enums.Role)
				*gs.MinimumRoleToInvite = enums.Role(value.String)
			}
		case groupsettings.ForeignKeys[0]:
			if value, ok := values[i].(*sql.NullScanner); !ok {
				return fmt.Errorf("unexpected type %T for field group_settings", values[i])
			} else if value.Valid {
				gs.group_settings = new(guidgql.GUID)
				*gs.group_settings = *value.S.(*guidgql.GUID)
			}
		default:
			gs.selectValues.Set(columns[i], values[i])
		}
	}
	return nil
}

// Value returns the ent.Value that was dynamically selected and assigned to the GroupSettings.
// This includes values selected through modifiers, order, etc.
func (gs *GroupSettings) Value(name string) (ent.Value, error) {
	return gs.selectValues.Get(name)
}

// QueryGroup queries the "group" edge of the GroupSettings entity.
func (gs *GroupSettings) QueryGroup() *GroupQuery {
	return NewGroupSettingsClient(gs.config).QueryGroup(gs)
}

// Update returns a builder for updating this GroupSettings.
// Note that you need to call GroupSettings.Unwrap() before calling this method if this GroupSettings
// was returned from a transaction, and the transaction was committed or rolled back.
func (gs *GroupSettings) Update() *GroupSettingsUpdateOne {
	return NewGroupSettingsClient(gs.config).UpdateOne(gs)
}

// Unwrap unwraps the GroupSettings entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (gs *GroupSettings) Unwrap() *GroupSettings {
	_tx, ok := gs.config.driver.(*txDriver)
	if !ok {
		panic("ent: GroupSettings is not a transactional entity")
	}
	gs.config.driver = _tx.drv
	return gs
}

// String implements the fmt.Stringer.
func (gs *GroupSettings) String() string {
	var builder strings.Builder
	builder.WriteString("GroupSettings(")
	builder.WriteString(fmt.Sprintf("id=%v, ", gs.ID))
	builder.WriteString("visibility=")
	builder.WriteString(fmt.Sprintf("%v", gs.Visibility))
	builder.WriteString(", ")
	builder.WriteString("join_policy=")
	builder.WriteString(fmt.Sprintf("%v", gs.JoinPolicy))
	builder.WriteString(", ")
	if v := gs.MinimumRoleToInvite; v != nil {
		builder.WriteString("minimum_role_to_invite=")
		builder.WriteString(fmt.Sprintf("%v", *v))
	}
	builder.WriteByte(')')
	return builder.String()
}

// GroupSettingsSlice is a parsable slice of GroupSettings.
type GroupSettingsSlice []*GroupSettings

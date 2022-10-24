// Code generated by ent, DO NOT EDIT.

package groupsettings

import (
	"fmt"
	"io"
	"strconv"

	"github.com/99designs/gqlgen/graphql"
	"github.com/open-boardgame-stats/backend/internal/ent/enums"
	"github.com/open-boardgame-stats/backend/internal/ent/schema/guidgql"
)

const (
	// Label holds the string label denoting the groupsettings type in the database.
	Label = "group_settings"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldVisibility holds the string denoting the visibility field in the database.
	FieldVisibility = "visibility"
	// FieldJoinPolicy holds the string denoting the join_policy field in the database.
	FieldJoinPolicy = "join_policy"
	// FieldMinimumRoleToInvite holds the string denoting the minimum_role_to_invite field in the database.
	FieldMinimumRoleToInvite = "minimum_role_to_invite"
	// EdgeGroup holds the string denoting the group edge name in mutations.
	EdgeGroup = "group"
	// Table holds the table name of the groupsettings in the database.
	Table = "group_settings"
	// GroupTable is the table that holds the group relation/edge.
	GroupTable = "group_settings"
	// GroupInverseTable is the table name for the Group entity.
	// It exists in this package in order to avoid circular dependency with the "group" package.
	GroupInverseTable = "groups"
	// GroupColumn is the table column denoting the group relation/edge.
	GroupColumn = "group_settings"
)

// Columns holds all SQL columns for groupsettings fields.
var Columns = []string{
	FieldID,
	FieldVisibility,
	FieldJoinPolicy,
	FieldMinimumRoleToInvite,
}

// ForeignKeys holds the SQL foreign-keys that are owned by the "group_settings"
// table and are not defined as standalone fields in the schema.
var ForeignKeys = []string{
	"group_settings",
}

// ValidColumn reports if the column name is valid (part of the table columns).
func ValidColumn(column string) bool {
	for i := range Columns {
		if column == Columns[i] {
			return true
		}
	}
	for i := range ForeignKeys {
		if column == ForeignKeys[i] {
			return true
		}
	}
	return false
}

var (
	// DefaultID holds the default value on creation for the "id" field.
	DefaultID func() guidgql.GUID
)

// Visibility defines the type for the "visibility" enum field.
type Visibility string

// VisibilityPublic is the default value of the Visibility enum.
const DefaultVisibility = VisibilityPublic

// Visibility values.
const (
	VisibilityPublic  Visibility = "PUBLIC"
	VisibilityPrivate Visibility = "PRIVATE"
)

func (v Visibility) String() string {
	return string(v)
}

// VisibilityValidator is a validator for the "visibility" field enum values. It is called by the builders before save.
func VisibilityValidator(v Visibility) error {
	switch v {
	case VisibilityPublic, VisibilityPrivate:
		return nil
	default:
		return fmt.Errorf("groupsettings: invalid enum value for visibility field: %q", v)
	}
}

// JoinPolicy defines the type for the "join_policy" enum field.
type JoinPolicy string

// JoinPolicyOpen is the default value of the JoinPolicy enum.
const DefaultJoinPolicy = JoinPolicyOpen

// JoinPolicy values.
const (
	JoinPolicyOpen                JoinPolicy = "OPEN"
	JoinPolicyInviteOnly          JoinPolicy = "INVITE_ONLY"
	JoinPolicyApplicationOnly     JoinPolicy = "APPLICATION_ONLY"
	JoinPolicyInviteOrApplication JoinPolicy = "INVITE_OR_APPLICATION"
)

func (jp JoinPolicy) String() string {
	return string(jp)
}

// JoinPolicyValidator is a validator for the "join_policy" field enum values. It is called by the builders before save.
func JoinPolicyValidator(jp JoinPolicy) error {
	switch jp {
	case JoinPolicyOpen, JoinPolicyInviteOnly, JoinPolicyApplicationOnly, JoinPolicyInviteOrApplication:
		return nil
	default:
		return fmt.Errorf("groupsettings: invalid enum value for join_policy field: %q", jp)
	}
}

// MinimumRoleToInviteValidator is a validator for the "minimum_role_to_invite" field enum values. It is called by the builders before save.
func MinimumRoleToInviteValidator(mrti enums.Role) error {
	switch mrti.String() {
	case "owner", "admin", "member":
		return nil
	default:
		return fmt.Errorf("groupsettings: invalid enum value for minimum_role_to_invite field: %q", mrti)
	}
}

// MarshalGQL implements graphql.Marshaler interface.
func (e Visibility) MarshalGQL(w io.Writer) {
	io.WriteString(w, strconv.Quote(e.String()))
}

// UnmarshalGQL implements graphql.Unmarshaler interface.
func (e *Visibility) UnmarshalGQL(val interface{}) error {
	str, ok := val.(string)
	if !ok {
		return fmt.Errorf("enum %T must be a string", val)
	}
	*e = Visibility(str)
	if err := VisibilityValidator(*e); err != nil {
		return fmt.Errorf("%s is not a valid Visibility", str)
	}
	return nil
}

// MarshalGQL implements graphql.Marshaler interface.
func (e JoinPolicy) MarshalGQL(w io.Writer) {
	io.WriteString(w, strconv.Quote(e.String()))
}

// UnmarshalGQL implements graphql.Unmarshaler interface.
func (e *JoinPolicy) UnmarshalGQL(val interface{}) error {
	str, ok := val.(string)
	if !ok {
		return fmt.Errorf("enum %T must be a string", val)
	}
	*e = JoinPolicy(str)
	if err := JoinPolicyValidator(*e); err != nil {
		return fmt.Errorf("%s is not a valid JoinPolicy", str)
	}
	return nil
}

var (
	// enums.Role must implement graphql.Marshaler.
	_ graphql.Marshaler = enums.Role("")
	// enums.Role must implement graphql.Unmarshaler.
	_ graphql.Unmarshaler = (*enums.Role)(nil)
)

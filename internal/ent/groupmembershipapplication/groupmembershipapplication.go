// Code generated by ent, DO NOT EDIT.

package groupmembershipapplication

import (
	"github.com/open-boardgame-stats/backend/internal/ent/schema/guidgql"
)

const (
	// Label holds the string label denoting the groupmembershipapplication type in the database.
	Label = "group_membership_application"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldMessage holds the string denoting the message field in the database.
	FieldMessage = "message"
	// EdgeUser holds the string denoting the user edge name in mutations.
	EdgeUser = "user"
	// EdgeGroup holds the string denoting the group edge name in mutations.
	EdgeGroup = "group"
	// Table holds the table name of the groupmembershipapplication in the database.
	Table = "group_membership_applications"
	// UserTable is the table that holds the user relation/edge. The primary key declared below.
	UserTable = "user_group_membership_applications"
	// UserInverseTable is the table name for the User entity.
	// It exists in this package in order to avoid circular dependency with the "user" package.
	UserInverseTable = "users"
	// GroupTable is the table that holds the group relation/edge. The primary key declared below.
	GroupTable = "group_applications"
	// GroupInverseTable is the table name for the Group entity.
	// It exists in this package in order to avoid circular dependency with the "group" package.
	GroupInverseTable = "groups"
)

// Columns holds all SQL columns for groupmembershipapplication fields.
var Columns = []string{
	FieldID,
	FieldMessage,
}

var (
	// UserPrimaryKey and UserColumn2 are the table columns denoting the
	// primary key for the user relation (M2M).
	UserPrimaryKey = []string{"user_id", "group_membership_application_id"}
	// GroupPrimaryKey and GroupColumn2 are the table columns denoting the
	// primary key for the group relation (M2M).
	GroupPrimaryKey = []string{"group_id", "group_membership_application_id"}
)

// ValidColumn reports if the column name is valid (part of the table columns).
func ValidColumn(column string) bool {
	for i := range Columns {
		if column == Columns[i] {
			return true
		}
	}
	return false
}

var (
	// DefaultMessage holds the default value on creation for the "message" field.
	DefaultMessage string
	// DefaultID holds the default value on creation for the "id" field.
	DefaultID func() guidgql.GUID
)

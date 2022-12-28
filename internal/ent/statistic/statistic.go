// Code generated by ent, DO NOT EDIT.

package statistic

import (
	"github.com/open-boardgame-stats/backend/internal/ent/schema/guidgql"
)

const (
	// Label holds the string label denoting the statistic type in the database.
	Label = "statistic"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldValue holds the string denoting the value field in the database.
	FieldValue = "value"
	// EdgeMatch holds the string denoting the match edge name in mutations.
	EdgeMatch = "match"
	// EdgeStatDescription holds the string denoting the stat_description edge name in mutations.
	EdgeStatDescription = "stat_description"
	// EdgePlayer holds the string denoting the player edge name in mutations.
	EdgePlayer = "player"
	// Table holds the table name of the statistic in the database.
	Table = "statistics"
	// MatchTable is the table that holds the match relation/edge.
	MatchTable = "statistics"
	// MatchInverseTable is the table name for the Match entity.
	// It exists in this package in order to avoid circular dependency with the "match" package.
	MatchInverseTable = "matches"
	// MatchColumn is the table column denoting the match relation/edge.
	MatchColumn = "match_stats"
	// StatDescriptionTable is the table that holds the stat_description relation/edge.
	StatDescriptionTable = "statistics"
	// StatDescriptionInverseTable is the table name for the StatDescription entity.
	// It exists in this package in order to avoid circular dependency with the "statdescription" package.
	StatDescriptionInverseTable = "stat_descriptions"
	// StatDescriptionColumn is the table column denoting the stat_description relation/edge.
	StatDescriptionColumn = "stat_description_stats"
	// PlayerTable is the table that holds the player relation/edge.
	PlayerTable = "statistics"
	// PlayerInverseTable is the table name for the Player entity.
	// It exists in this package in order to avoid circular dependency with the "player" package.
	PlayerInverseTable = "players"
	// PlayerColumn is the table column denoting the player relation/edge.
	PlayerColumn = "player_stats"
)

// Columns holds all SQL columns for statistic fields.
var Columns = []string{
	FieldID,
	FieldValue,
}

// ForeignKeys holds the SQL foreign-keys that are owned by the "statistics"
// table and are not defined as standalone fields in the schema.
var ForeignKeys = []string{
	"match_stats",
	"player_stats",
	"stat_description_stats",
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
	// DefaultValue holds the default value on creation for the "value" field.
	DefaultValue string
	// DefaultID holds the default value on creation for the "id" field.
	DefaultID func() guidgql.GUID
)

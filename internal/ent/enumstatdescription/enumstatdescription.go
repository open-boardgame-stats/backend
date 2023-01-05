// Code generated by ent, DO NOT EDIT.

package enumstatdescription

import (
	"github.com/open-boardgame-stats/backend/internal/ent/schema/guidgql"
)

const (
	// Label holds the string label denoting the enumstatdescription type in the database.
	Label = "enum_stat_description"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldName holds the string denoting the name field in the database.
	FieldName = "name"
	// FieldDescription holds the string denoting the description field in the database.
	FieldDescription = "description"
	// FieldPossibleValues holds the string denoting the possible_values field in the database.
	FieldPossibleValues = "possible_values"
	// EdgeGame holds the string denoting the game edge name in mutations.
	EdgeGame = "game"
	// EdgeEnumStats holds the string denoting the enum_stats edge name in mutations.
	EdgeEnumStats = "enum_stats"
	// Table holds the table name of the enumstatdescription in the database.
	Table = "enum_stat_descriptions"
	// GameTable is the table that holds the game relation/edge. The primary key declared below.
	GameTable = "enum_stat_description_game"
	// GameInverseTable is the table name for the Game entity.
	// It exists in this package in order to avoid circular dependency with the "game" package.
	GameInverseTable = "games"
	// EnumStatsTable is the table that holds the enum_stats relation/edge.
	EnumStatsTable = "enum_stats"
	// EnumStatsInverseTable is the table name for the EnumStat entity.
	// It exists in this package in order to avoid circular dependency with the "enumstat" package.
	EnumStatsInverseTable = "enum_stats"
	// EnumStatsColumn is the table column denoting the enum_stats relation/edge.
	EnumStatsColumn = "enum_stat_description_enum_stats"
)

// Columns holds all SQL columns for enumstatdescription fields.
var Columns = []string{
	FieldID,
	FieldName,
	FieldDescription,
	FieldPossibleValues,
}

var (
	// GamePrimaryKey and GameColumn2 are the table columns denoting the
	// primary key for the game relation (M2M).
	GamePrimaryKey = []string{"enum_stat_description_id", "game_id"}
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
	// NameValidator is a validator for the "name" field. It is called by the builders before save.
	NameValidator func(string) error
	// DefaultID holds the default value on creation for the "id" field.
	DefaultID func() guidgql.GUID
)

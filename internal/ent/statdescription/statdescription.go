// Code generated by ent, DO NOT EDIT.

package statdescription

import (
	"fmt"

	"github.com/99designs/gqlgen/graphql"
	"github.com/open-boardgame-stats/backend/internal/ent/schema/guidgql"
	"github.com/open-boardgame-stats/backend/internal/ent/schema/stat"
)

const (
	// Label holds the string label denoting the statdescription type in the database.
	Label = "stat_description"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldType holds the string denoting the type field in the database.
	FieldType = "type"
	// FieldName holds the string denoting the name field in the database.
	FieldName = "name"
	// FieldDescription holds the string denoting the description field in the database.
	FieldDescription = "description"
	// FieldMetadata holds the string denoting the metadata field in the database.
	FieldMetadata = "metadata"
	// FieldOrderNumber holds the string denoting the order_number field in the database.
	FieldOrderNumber = "order_number"
	// EdgeGame holds the string denoting the game edge name in mutations.
	EdgeGame = "game"
	// EdgeStats holds the string denoting the stats edge name in mutations.
	EdgeStats = "stats"
	// Table holds the table name of the statdescription in the database.
	Table = "stat_descriptions"
	// GameTable is the table that holds the game relation/edge. The primary key declared below.
	GameTable = "stat_description_game"
	// GameInverseTable is the table name for the Game entity.
	// It exists in this package in order to avoid circular dependency with the "game" package.
	GameInverseTable = "games"
	// StatsTable is the table that holds the stats relation/edge.
	StatsTable = "statistics"
	// StatsInverseTable is the table name for the Statistic entity.
	// It exists in this package in order to avoid circular dependency with the "statistic" package.
	StatsInverseTable = "statistics"
	// StatsColumn is the table column denoting the stats relation/edge.
	StatsColumn = "stat_description_stats"
)

// Columns holds all SQL columns for statdescription fields.
var Columns = []string{
	FieldID,
	FieldType,
	FieldName,
	FieldDescription,
	FieldMetadata,
	FieldOrderNumber,
}

var (
	// GamePrimaryKey and GameColumn2 are the table columns denoting the
	// primary key for the game relation (M2M).
	GamePrimaryKey = []string{"stat_description_id", "game_id"}
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
	// DefaultDescription holds the default value on creation for the "description" field.
	DefaultDescription string
	// DefaultID holds the default value on creation for the "id" field.
	DefaultID func() guidgql.GUID
)

// TypeValidator is a validator for the "type" field enum values. It is called by the builders before save.
func TypeValidator(_type stat.StatType) error {
	switch _type.String() {
	case "numeric", "enum", "aggregate":
		return nil
	default:
		return fmt.Errorf("statdescription: invalid enum value for type field: %q", _type)
	}
}

var (
	// stat.StatType must implement graphql.Marshaler.
	_ graphql.Marshaler = (*stat.StatType)(nil)
	// stat.StatType must implement graphql.Unmarshaler.
	_ graphql.Unmarshaler = (*stat.StatType)(nil)
)

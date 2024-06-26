// Code generated by ent, DO NOT EDIT.

package gameversion

import (
	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"github.com/open-boardgame-stats/backend/internal/ent/schema/guidgql"
)

const (
	// Label holds the string label denoting the gameversion type in the database.
	Label = "game_version"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldVersionNumber holds the string denoting the version_number field in the database.
	FieldVersionNumber = "version_number"
	// EdgeGame holds the string denoting the game edge name in mutations.
	EdgeGame = "game"
	// EdgeStatDescriptions holds the string denoting the stat_descriptions edge name in mutations.
	EdgeStatDescriptions = "stat_descriptions"
	// EdgeMatches holds the string denoting the matches edge name in mutations.
	EdgeMatches = "matches"
	// Table holds the table name of the gameversion in the database.
	Table = "game_versions"
	// GameTable is the table that holds the game relation/edge.
	GameTable = "game_versions"
	// GameInverseTable is the table name for the Game entity.
	// It exists in this package in order to avoid circular dependency with the "game" package.
	GameInverseTable = "games"
	// GameColumn is the table column denoting the game relation/edge.
	GameColumn = "game_version_game"
	// StatDescriptionsTable is the table that holds the stat_descriptions relation/edge. The primary key declared below.
	StatDescriptionsTable = "stat_description_game_version"
	// StatDescriptionsInverseTable is the table name for the StatDescription entity.
	// It exists in this package in order to avoid circular dependency with the "statdescription" package.
	StatDescriptionsInverseTable = "stat_descriptions"
	// MatchesTable is the table that holds the matches relation/edge.
	MatchesTable = "matches"
	// MatchesInverseTable is the table name for the Match entity.
	// It exists in this package in order to avoid circular dependency with the "match" package.
	MatchesInverseTable = "matches"
	// MatchesColumn is the table column denoting the matches relation/edge.
	MatchesColumn = "game_version_matches"
)

// Columns holds all SQL columns for gameversion fields.
var Columns = []string{
	FieldID,
	FieldVersionNumber,
}

// ForeignKeys holds the SQL foreign-keys that are owned by the "game_versions"
// table and are not defined as standalone fields in the schema.
var ForeignKeys = []string{
	"game_version_game",
}

var (
	// StatDescriptionsPrimaryKey and StatDescriptionsColumn2 are the table columns denoting the
	// primary key for the stat_descriptions relation (M2M).
	StatDescriptionsPrimaryKey = []string{"stat_description_id", "game_version_id"}
)

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
	// DefaultVersionNumber holds the default value on creation for the "version_number" field.
	DefaultVersionNumber int
	// DefaultID holds the default value on creation for the "id" field.
	DefaultID func() guidgql.GUID
)

// OrderOption defines the ordering options for the GameVersion queries.
type OrderOption func(*sql.Selector)

// ByID orders the results by the id field.
func ByID(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldID, opts...).ToFunc()
}

// ByVersionNumber orders the results by the version_number field.
func ByVersionNumber(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldVersionNumber, opts...).ToFunc()
}

// ByGameField orders the results by game field.
func ByGameField(field string, opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newGameStep(), sql.OrderByField(field, opts...))
	}
}

// ByStatDescriptionsCount orders the results by stat_descriptions count.
func ByStatDescriptionsCount(opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborsCount(s, newStatDescriptionsStep(), opts...)
	}
}

// ByStatDescriptions orders the results by stat_descriptions terms.
func ByStatDescriptions(term sql.OrderTerm, terms ...sql.OrderTerm) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newStatDescriptionsStep(), append([]sql.OrderTerm{term}, terms...)...)
	}
}

// ByMatchesCount orders the results by matches count.
func ByMatchesCount(opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborsCount(s, newMatchesStep(), opts...)
	}
}

// ByMatches orders the results by matches terms.
func ByMatches(term sql.OrderTerm, terms ...sql.OrderTerm) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newMatchesStep(), append([]sql.OrderTerm{term}, terms...)...)
	}
}
func newGameStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(GameInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.M2O, false, GameTable, GameColumn),
	)
}
func newStatDescriptionsStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(StatDescriptionsInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.M2M, true, StatDescriptionsTable, StatDescriptionsPrimaryKey...),
	)
}
func newMatchesStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(MatchesInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.O2M, false, MatchesTable, MatchesColumn),
	)
}

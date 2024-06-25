// Code generated by ent, DO NOT EDIT.

package match

import (
	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"github.com/open-boardgame-stats/backend/internal/ent/schema/guidgql"
)

const (
	// Label holds the string label denoting the match type in the database.
	Label = "match"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// EdgeGameVersion holds the string denoting the game_version edge name in mutations.
	EdgeGameVersion = "game_version"
	// EdgePlayers holds the string denoting the players edge name in mutations.
	EdgePlayers = "players"
	// EdgeStats holds the string denoting the stats edge name in mutations.
	EdgeStats = "stats"
	// Table holds the table name of the match in the database.
	Table = "matches"
	// GameVersionTable is the table that holds the game_version relation/edge.
	GameVersionTable = "matches"
	// GameVersionInverseTable is the table name for the GameVersion entity.
	// It exists in this package in order to avoid circular dependency with the "gameversion" package.
	GameVersionInverseTable = "game_versions"
	// GameVersionColumn is the table column denoting the game_version relation/edge.
	GameVersionColumn = "game_version_matches"
	// PlayersTable is the table that holds the players relation/edge. The primary key declared below.
	PlayersTable = "match_players"
	// PlayersInverseTable is the table name for the Player entity.
	// It exists in this package in order to avoid circular dependency with the "player" package.
	PlayersInverseTable = "players"
	// StatsTable is the table that holds the stats relation/edge.
	StatsTable = "statistics"
	// StatsInverseTable is the table name for the Statistic entity.
	// It exists in this package in order to avoid circular dependency with the "statistic" package.
	StatsInverseTable = "statistics"
	// StatsColumn is the table column denoting the stats relation/edge.
	StatsColumn = "match_stats"
)

// Columns holds all SQL columns for match fields.
var Columns = []string{
	FieldID,
}

// ForeignKeys holds the SQL foreign-keys that are owned by the "matches"
// table and are not defined as standalone fields in the schema.
var ForeignKeys = []string{
	"game_version_matches",
}

var (
	// PlayersPrimaryKey and PlayersColumn2 are the table columns denoting the
	// primary key for the players relation (M2M).
	PlayersPrimaryKey = []string{"match_id", "player_id"}
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
	// DefaultID holds the default value on creation for the "id" field.
	DefaultID func() guidgql.GUID
)

// OrderOption defines the ordering options for the Match queries.
type OrderOption func(*sql.Selector)

// ByID orders the results by the id field.
func ByID(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldID, opts...).ToFunc()
}

// ByGameVersionField orders the results by game_version field.
func ByGameVersionField(field string, opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newGameVersionStep(), sql.OrderByField(field, opts...))
	}
}

// ByPlayersCount orders the results by players count.
func ByPlayersCount(opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborsCount(s, newPlayersStep(), opts...)
	}
}

// ByPlayers orders the results by players terms.
func ByPlayers(term sql.OrderTerm, terms ...sql.OrderTerm) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newPlayersStep(), append([]sql.OrderTerm{term}, terms...)...)
	}
}

// ByStatsCount orders the results by stats count.
func ByStatsCount(opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborsCount(s, newStatsStep(), opts...)
	}
}

// ByStats orders the results by stats terms.
func ByStats(term sql.OrderTerm, terms ...sql.OrderTerm) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newStatsStep(), append([]sql.OrderTerm{term}, terms...)...)
	}
}
func newGameVersionStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(GameVersionInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.M2O, true, GameVersionTable, GameVersionColumn),
	)
}
func newPlayersStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(PlayersInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.M2M, false, PlayersTable, PlayersPrimaryKey...),
	)
}
func newStatsStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(StatsInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.O2M, false, StatsTable, StatsColumn),
	)
}

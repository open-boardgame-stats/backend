// Code generated by ent, DO NOT EDIT.

package ent

import (
	"fmt"
	"strings"

	"entgo.io/ent/dialect/sql"
	"github.com/open-boardgame-stats/backend/internal/ent/player"
	"github.com/open-boardgame-stats/backend/internal/ent/schema/guidgql"
	"github.com/open-boardgame-stats/backend/internal/ent/user"
)

// Player is the model entity for the Player schema.
type Player struct {
	config `json:"-"`
	// ID of the ent.
	ID guidgql.GUID `json:"id,omitempty"`
	// Name holds the value of the "name" field.
	Name string `json:"name,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the PlayerQuery when eager-loading is set.
	Edges            PlayerEdges `json:"edges"`
	user_main_player *guidgql.GUID
}

// PlayerEdges holds the relations/edges for other nodes in the graph.
type PlayerEdges struct {
	// Owner holds the value of the owner edge.
	Owner *User `json:"owner,omitempty"`
	// Supervisors holds the value of the supervisors edge.
	Supervisors []*User `json:"supervisors,omitempty"`
	// SupervisionRequests holds the value of the supervision_requests edge.
	SupervisionRequests []*PlayerSupervisionRequest `json:"supervision_requests,omitempty"`
	// Matches holds the value of the matches edge.
	Matches []*Match `json:"matches,omitempty"`
	// NumericalStats holds the value of the numerical_stats edge.
	NumericalStats []*NumericalStat `json:"numerical_stats,omitempty"`
	// EnumStats holds the value of the enum_stats edge.
	EnumStats []*EnumStat `json:"enum_stats,omitempty"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [6]bool
	// totalCount holds the count of the edges above.
	totalCount [4]map[string]int

	namedSupervisors         map[string][]*User
	namedSupervisionRequests map[string][]*PlayerSupervisionRequest
	namedMatches             map[string][]*Match
	namedNumericalStats      map[string][]*NumericalStat
	namedEnumStats           map[string][]*EnumStat
}

// OwnerOrErr returns the Owner value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e PlayerEdges) OwnerOrErr() (*User, error) {
	if e.loadedTypes[0] {
		if e.Owner == nil {
			// Edge was loaded but was not found.
			return nil, &NotFoundError{label: user.Label}
		}
		return e.Owner, nil
	}
	return nil, &NotLoadedError{edge: "owner"}
}

// SupervisorsOrErr returns the Supervisors value or an error if the edge
// was not loaded in eager-loading.
func (e PlayerEdges) SupervisorsOrErr() ([]*User, error) {
	if e.loadedTypes[1] {
		return e.Supervisors, nil
	}
	return nil, &NotLoadedError{edge: "supervisors"}
}

// SupervisionRequestsOrErr returns the SupervisionRequests value or an error if the edge
// was not loaded in eager-loading.
func (e PlayerEdges) SupervisionRequestsOrErr() ([]*PlayerSupervisionRequest, error) {
	if e.loadedTypes[2] {
		return e.SupervisionRequests, nil
	}
	return nil, &NotLoadedError{edge: "supervision_requests"}
}

// MatchesOrErr returns the Matches value or an error if the edge
// was not loaded in eager-loading.
func (e PlayerEdges) MatchesOrErr() ([]*Match, error) {
	if e.loadedTypes[3] {
		return e.Matches, nil
	}
	return nil, &NotLoadedError{edge: "matches"}
}

// NumericalStatsOrErr returns the NumericalStats value or an error if the edge
// was not loaded in eager-loading.
func (e PlayerEdges) NumericalStatsOrErr() ([]*NumericalStat, error) {
	if e.loadedTypes[4] {
		return e.NumericalStats, nil
	}
	return nil, &NotLoadedError{edge: "numerical_stats"}
}

// EnumStatsOrErr returns the EnumStats value or an error if the edge
// was not loaded in eager-loading.
func (e PlayerEdges) EnumStatsOrErr() ([]*EnumStat, error) {
	if e.loadedTypes[5] {
		return e.EnumStats, nil
	}
	return nil, &NotLoadedError{edge: "enum_stats"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*Player) scanValues(columns []string) ([]any, error) {
	values := make([]any, len(columns))
	for i := range columns {
		switch columns[i] {
		case player.FieldID:
			values[i] = new(guidgql.GUID)
		case player.FieldName:
			values[i] = new(sql.NullString)
		case player.ForeignKeys[0]: // user_main_player
			values[i] = &sql.NullScanner{S: new(guidgql.GUID)}
		default:
			return nil, fmt.Errorf("unexpected column %q for type Player", columns[i])
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the Player fields.
func (pl *Player) assignValues(columns []string, values []any) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case player.FieldID:
			if value, ok := values[i].(*guidgql.GUID); !ok {
				return fmt.Errorf("unexpected type %T for field id", values[i])
			} else if value != nil {
				pl.ID = *value
			}
		case player.FieldName:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field name", values[i])
			} else if value.Valid {
				pl.Name = value.String
			}
		case player.ForeignKeys[0]:
			if value, ok := values[i].(*sql.NullScanner); !ok {
				return fmt.Errorf("unexpected type %T for field user_main_player", values[i])
			} else if value.Valid {
				pl.user_main_player = new(guidgql.GUID)
				*pl.user_main_player = *value.S.(*guidgql.GUID)
			}
		}
	}
	return nil
}

// QueryOwner queries the "owner" edge of the Player entity.
func (pl *Player) QueryOwner() *UserQuery {
	return (&PlayerClient{config: pl.config}).QueryOwner(pl)
}

// QuerySupervisors queries the "supervisors" edge of the Player entity.
func (pl *Player) QuerySupervisors() *UserQuery {
	return (&PlayerClient{config: pl.config}).QuerySupervisors(pl)
}

// QuerySupervisionRequests queries the "supervision_requests" edge of the Player entity.
func (pl *Player) QuerySupervisionRequests() *PlayerSupervisionRequestQuery {
	return (&PlayerClient{config: pl.config}).QuerySupervisionRequests(pl)
}

// QueryMatches queries the "matches" edge of the Player entity.
func (pl *Player) QueryMatches() *MatchQuery {
	return (&PlayerClient{config: pl.config}).QueryMatches(pl)
}

// QueryNumericalStats queries the "numerical_stats" edge of the Player entity.
func (pl *Player) QueryNumericalStats() *NumericalStatQuery {
	return (&PlayerClient{config: pl.config}).QueryNumericalStats(pl)
}

// QueryEnumStats queries the "enum_stats" edge of the Player entity.
func (pl *Player) QueryEnumStats() *EnumStatQuery {
	return (&PlayerClient{config: pl.config}).QueryEnumStats(pl)
}

// Update returns a builder for updating this Player.
// Note that you need to call Player.Unwrap() before calling this method if this Player
// was returned from a transaction, and the transaction was committed or rolled back.
func (pl *Player) Update() *PlayerUpdateOne {
	return (&PlayerClient{config: pl.config}).UpdateOne(pl)
}

// Unwrap unwraps the Player entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (pl *Player) Unwrap() *Player {
	_tx, ok := pl.config.driver.(*txDriver)
	if !ok {
		panic("ent: Player is not a transactional entity")
	}
	pl.config.driver = _tx.drv
	return pl
}

// String implements the fmt.Stringer.
func (pl *Player) String() string {
	var builder strings.Builder
	builder.WriteString("Player(")
	builder.WriteString(fmt.Sprintf("id=%v, ", pl.ID))
	builder.WriteString("name=")
	builder.WriteString(pl.Name)
	builder.WriteByte(')')
	return builder.String()
}

// NamedSupervisors returns the Supervisors named value or an error if the edge was not
// loaded in eager-loading with this name.
func (pl *Player) NamedSupervisors(name string) ([]*User, error) {
	if pl.Edges.namedSupervisors == nil {
		return nil, &NotLoadedError{edge: name}
	}
	nodes, ok := pl.Edges.namedSupervisors[name]
	if !ok {
		return nil, &NotLoadedError{edge: name}
	}
	return nodes, nil
}

func (pl *Player) appendNamedSupervisors(name string, edges ...*User) {
	if pl.Edges.namedSupervisors == nil {
		pl.Edges.namedSupervisors = make(map[string][]*User)
	}
	if len(edges) == 0 {
		pl.Edges.namedSupervisors[name] = []*User{}
	} else {
		pl.Edges.namedSupervisors[name] = append(pl.Edges.namedSupervisors[name], edges...)
	}
}

// NamedSupervisionRequests returns the SupervisionRequests named value or an error if the edge was not
// loaded in eager-loading with this name.
func (pl *Player) NamedSupervisionRequests(name string) ([]*PlayerSupervisionRequest, error) {
	if pl.Edges.namedSupervisionRequests == nil {
		return nil, &NotLoadedError{edge: name}
	}
	nodes, ok := pl.Edges.namedSupervisionRequests[name]
	if !ok {
		return nil, &NotLoadedError{edge: name}
	}
	return nodes, nil
}

func (pl *Player) appendNamedSupervisionRequests(name string, edges ...*PlayerSupervisionRequest) {
	if pl.Edges.namedSupervisionRequests == nil {
		pl.Edges.namedSupervisionRequests = make(map[string][]*PlayerSupervisionRequest)
	}
	if len(edges) == 0 {
		pl.Edges.namedSupervisionRequests[name] = []*PlayerSupervisionRequest{}
	} else {
		pl.Edges.namedSupervisionRequests[name] = append(pl.Edges.namedSupervisionRequests[name], edges...)
	}
}

// NamedMatches returns the Matches named value or an error if the edge was not
// loaded in eager-loading with this name.
func (pl *Player) NamedMatches(name string) ([]*Match, error) {
	if pl.Edges.namedMatches == nil {
		return nil, &NotLoadedError{edge: name}
	}
	nodes, ok := pl.Edges.namedMatches[name]
	if !ok {
		return nil, &NotLoadedError{edge: name}
	}
	return nodes, nil
}

func (pl *Player) appendNamedMatches(name string, edges ...*Match) {
	if pl.Edges.namedMatches == nil {
		pl.Edges.namedMatches = make(map[string][]*Match)
	}
	if len(edges) == 0 {
		pl.Edges.namedMatches[name] = []*Match{}
	} else {
		pl.Edges.namedMatches[name] = append(pl.Edges.namedMatches[name], edges...)
	}
}

// NamedNumericalStats returns the NumericalStats named value or an error if the edge was not
// loaded in eager-loading with this name.
func (pl *Player) NamedNumericalStats(name string) ([]*NumericalStat, error) {
	if pl.Edges.namedNumericalStats == nil {
		return nil, &NotLoadedError{edge: name}
	}
	nodes, ok := pl.Edges.namedNumericalStats[name]
	if !ok {
		return nil, &NotLoadedError{edge: name}
	}
	return nodes, nil
}

func (pl *Player) appendNamedNumericalStats(name string, edges ...*NumericalStat) {
	if pl.Edges.namedNumericalStats == nil {
		pl.Edges.namedNumericalStats = make(map[string][]*NumericalStat)
	}
	if len(edges) == 0 {
		pl.Edges.namedNumericalStats[name] = []*NumericalStat{}
	} else {
		pl.Edges.namedNumericalStats[name] = append(pl.Edges.namedNumericalStats[name], edges...)
	}
}

// NamedEnumStats returns the EnumStats named value or an error if the edge was not
// loaded in eager-loading with this name.
func (pl *Player) NamedEnumStats(name string) ([]*EnumStat, error) {
	if pl.Edges.namedEnumStats == nil {
		return nil, &NotLoadedError{edge: name}
	}
	nodes, ok := pl.Edges.namedEnumStats[name]
	if !ok {
		return nil, &NotLoadedError{edge: name}
	}
	return nodes, nil
}

func (pl *Player) appendNamedEnumStats(name string, edges ...*EnumStat) {
	if pl.Edges.namedEnumStats == nil {
		pl.Edges.namedEnumStats = make(map[string][]*EnumStat)
	}
	if len(edges) == 0 {
		pl.Edges.namedEnumStats[name] = []*EnumStat{}
	} else {
		pl.Edges.namedEnumStats[name] = append(pl.Edges.namedEnumStats[name], edges...)
	}
}

// Players is a parsable slice of Player.
type Players []*Player

func (pl Players) config(cfg config) {
	for _i := range pl {
		pl[_i].config = cfg
	}
}

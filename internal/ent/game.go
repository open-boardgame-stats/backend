// Code generated by ent, DO NOT EDIT.

package ent

import (
	"fmt"
	"strings"

	"entgo.io/ent/dialect/sql"
	"github.com/open-boardgame-stats/backend/internal/ent/game"
	"github.com/open-boardgame-stats/backend/internal/ent/schema/guidgql"
	"github.com/open-boardgame-stats/backend/internal/ent/user"
)

// Game is the model entity for the Game schema.
type Game struct {
	config `json:"-"`
	// ID of the ent.
	ID guidgql.GUID `json:"id,omitempty"`
	// Name holds the value of the "name" field.
	Name string `json:"name,omitempty"`
	// MinPlayers holds the value of the "min_players" field.
	MinPlayers int `json:"min_players,omitempty"`
	// MaxPlayers holds the value of the "max_players" field.
	MaxPlayers int `json:"max_players,omitempty"`
	// Description holds the value of the "description" field.
	Description string `json:"description,omitempty"`
	// BoardgamegeekURL holds the value of the "boardgamegeek_url" field.
	BoardgamegeekURL string `json:"boardgamegeek_url,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the GameQuery when eager-loading is set.
	Edges      GameEdges `json:"edges"`
	user_games *guidgql.GUID
}

// GameEdges holds the relations/edges for other nodes in the graph.
type GameEdges struct {
	// Author holds the value of the author edge.
	Author *User `json:"author,omitempty"`
	// Favorites holds the value of the favorites edge.
	Favorites []*GameFavorite `json:"favorites,omitempty"`
	// StatDescriptions holds the value of the stat_descriptions edge.
	StatDescriptions []*StatDescription `json:"stat_descriptions,omitempty"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [3]bool
	// totalCount holds the count of the edges above.
	totalCount [2]map[string]int

	namedFavorites        map[string][]*GameFavorite
	namedStatDescriptions map[string][]*StatDescription
}

// AuthorOrErr returns the Author value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e GameEdges) AuthorOrErr() (*User, error) {
	if e.loadedTypes[0] {
		if e.Author == nil {
			// Edge was loaded but was not found.
			return nil, &NotFoundError{label: user.Label}
		}
		return e.Author, nil
	}
	return nil, &NotLoadedError{edge: "author"}
}

// FavoritesOrErr returns the Favorites value or an error if the edge
// was not loaded in eager-loading.
func (e GameEdges) FavoritesOrErr() ([]*GameFavorite, error) {
	if e.loadedTypes[1] {
		return e.Favorites, nil
	}
	return nil, &NotLoadedError{edge: "favorites"}
}

// StatDescriptionsOrErr returns the StatDescriptions value or an error if the edge
// was not loaded in eager-loading.
func (e GameEdges) StatDescriptionsOrErr() ([]*StatDescription, error) {
	if e.loadedTypes[2] {
		return e.StatDescriptions, nil
	}
	return nil, &NotLoadedError{edge: "stat_descriptions"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*Game) scanValues(columns []string) ([]any, error) {
	values := make([]any, len(columns))
	for i := range columns {
		switch columns[i] {
		case game.FieldID:
			values[i] = new(guidgql.GUID)
		case game.FieldMinPlayers, game.FieldMaxPlayers:
			values[i] = new(sql.NullInt64)
		case game.FieldName, game.FieldDescription, game.FieldBoardgamegeekURL:
			values[i] = new(sql.NullString)
		case game.ForeignKeys[0]: // user_games
			values[i] = &sql.NullScanner{S: new(guidgql.GUID)}
		default:
			return nil, fmt.Errorf("unexpected column %q for type Game", columns[i])
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the Game fields.
func (ga *Game) assignValues(columns []string, values []any) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case game.FieldID:
			if value, ok := values[i].(*guidgql.GUID); !ok {
				return fmt.Errorf("unexpected type %T for field id", values[i])
			} else if value != nil {
				ga.ID = *value
			}
		case game.FieldName:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field name", values[i])
			} else if value.Valid {
				ga.Name = value.String
			}
		case game.FieldMinPlayers:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field min_players", values[i])
			} else if value.Valid {
				ga.MinPlayers = int(value.Int64)
			}
		case game.FieldMaxPlayers:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field max_players", values[i])
			} else if value.Valid {
				ga.MaxPlayers = int(value.Int64)
			}
		case game.FieldDescription:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field description", values[i])
			} else if value.Valid {
				ga.Description = value.String
			}
		case game.FieldBoardgamegeekURL:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field boardgamegeek_url", values[i])
			} else if value.Valid {
				ga.BoardgamegeekURL = value.String
			}
		case game.ForeignKeys[0]:
			if value, ok := values[i].(*sql.NullScanner); !ok {
				return fmt.Errorf("unexpected type %T for field user_games", values[i])
			} else if value.Valid {
				ga.user_games = new(guidgql.GUID)
				*ga.user_games = *value.S.(*guidgql.GUID)
			}
		}
	}
	return nil
}

// QueryAuthor queries the "author" edge of the Game entity.
func (ga *Game) QueryAuthor() *UserQuery {
	return (&GameClient{config: ga.config}).QueryAuthor(ga)
}

// QueryFavorites queries the "favorites" edge of the Game entity.
func (ga *Game) QueryFavorites() *GameFavoriteQuery {
	return (&GameClient{config: ga.config}).QueryFavorites(ga)
}

// QueryStatDescriptions queries the "stat_descriptions" edge of the Game entity.
func (ga *Game) QueryStatDescriptions() *StatDescriptionQuery {
	return (&GameClient{config: ga.config}).QueryStatDescriptions(ga)
}

// Update returns a builder for updating this Game.
// Note that you need to call Game.Unwrap() before calling this method if this Game
// was returned from a transaction, and the transaction was committed or rolled back.
func (ga *Game) Update() *GameUpdateOne {
	return (&GameClient{config: ga.config}).UpdateOne(ga)
}

// Unwrap unwraps the Game entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (ga *Game) Unwrap() *Game {
	_tx, ok := ga.config.driver.(*txDriver)
	if !ok {
		panic("ent: Game is not a transactional entity")
	}
	ga.config.driver = _tx.drv
	return ga
}

// String implements the fmt.Stringer.
func (ga *Game) String() string {
	var builder strings.Builder
	builder.WriteString("Game(")
	builder.WriteString(fmt.Sprintf("id=%v, ", ga.ID))
	builder.WriteString("name=")
	builder.WriteString(ga.Name)
	builder.WriteString(", ")
	builder.WriteString("min_players=")
	builder.WriteString(fmt.Sprintf("%v", ga.MinPlayers))
	builder.WriteString(", ")
	builder.WriteString("max_players=")
	builder.WriteString(fmt.Sprintf("%v", ga.MaxPlayers))
	builder.WriteString(", ")
	builder.WriteString("description=")
	builder.WriteString(ga.Description)
	builder.WriteString(", ")
	builder.WriteString("boardgamegeek_url=")
	builder.WriteString(ga.BoardgamegeekURL)
	builder.WriteByte(')')
	return builder.String()
}

// NamedFavorites returns the Favorites named value or an error if the edge was not
// loaded in eager-loading with this name.
func (ga *Game) NamedFavorites(name string) ([]*GameFavorite, error) {
	if ga.Edges.namedFavorites == nil {
		return nil, &NotLoadedError{edge: name}
	}
	nodes, ok := ga.Edges.namedFavorites[name]
	if !ok {
		return nil, &NotLoadedError{edge: name}
	}
	return nodes, nil
}

func (ga *Game) appendNamedFavorites(name string, edges ...*GameFavorite) {
	if ga.Edges.namedFavorites == nil {
		ga.Edges.namedFavorites = make(map[string][]*GameFavorite)
	}
	if len(edges) == 0 {
		ga.Edges.namedFavorites[name] = []*GameFavorite{}
	} else {
		ga.Edges.namedFavorites[name] = append(ga.Edges.namedFavorites[name], edges...)
	}
}

// NamedStatDescriptions returns the StatDescriptions named value or an error if the edge was not
// loaded in eager-loading with this name.
func (ga *Game) NamedStatDescriptions(name string) ([]*StatDescription, error) {
	if ga.Edges.namedStatDescriptions == nil {
		return nil, &NotLoadedError{edge: name}
	}
	nodes, ok := ga.Edges.namedStatDescriptions[name]
	if !ok {
		return nil, &NotLoadedError{edge: name}
	}
	return nodes, nil
}

func (ga *Game) appendNamedStatDescriptions(name string, edges ...*StatDescription) {
	if ga.Edges.namedStatDescriptions == nil {
		ga.Edges.namedStatDescriptions = make(map[string][]*StatDescription)
	}
	if len(edges) == 0 {
		ga.Edges.namedStatDescriptions[name] = []*StatDescription{}
	} else {
		ga.Edges.namedStatDescriptions[name] = append(ga.Edges.namedStatDescriptions[name], edges...)
	}
}

// Games is a parsable slice of Game.
type Games []*Game

func (ga Games) config(cfg config) {
	for _i := range ga {
		ga[_i].config = cfg
	}
}
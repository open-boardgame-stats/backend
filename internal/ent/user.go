// Code generated by ent, DO NOT EDIT.

package ent

import (
	"fmt"
	"strings"

	"entgo.io/ent"
	"entgo.io/ent/dialect/sql"
	"github.com/open-boardgame-stats/backend/internal/ent/player"
	"github.com/open-boardgame-stats/backend/internal/ent/schema/guidgql"
	"github.com/open-boardgame-stats/backend/internal/ent/user"
)

// User is the model entity for the User schema.
type User struct {
	config `json:"-"`
	// ID of the ent.
	ID guidgql.GUID `json:"id,omitempty"`
	// Name holds the value of the "name" field.
	Name string `json:"name,omitempty"`
	// Email holds the value of the "email" field.
	Email string `json:"email,omitempty"`
	// Password holds the value of the "password" field.
	Password string `json:"-"`
	// AvatarURL holds the value of the "avatar_url" field.
	AvatarURL string `json:"avatar_url,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the UserQuery when eager-loading is set.
	Edges        UserEdges `json:"edges"`
	selectValues sql.SelectValues
}

// UserEdges holds the relations/edges for other nodes in the graph.
type UserEdges struct {
	// Players holds the value of the players edge.
	Players []*Player `json:"players,omitempty"`
	// MainPlayer holds the value of the main_player edge.
	MainPlayer *Player `json:"main_player,omitempty"`
	// SentSupervisionRequests holds the value of the sent_supervision_requests edge.
	SentSupervisionRequests []*PlayerSupervisionRequest `json:"sent_supervision_requests,omitempty"`
	// SupervisionRequestApprovals holds the value of the supervision_request_approvals edge.
	SupervisionRequestApprovals []*PlayerSupervisionRequestApproval `json:"supervision_request_approvals,omitempty"`
	// GroupMemberships holds the value of the group_memberships edge.
	GroupMemberships []*GroupMembership `json:"group_memberships,omitempty"`
	// GroupMembershipApplications holds the value of the group_membership_applications edge.
	GroupMembershipApplications []*GroupMembershipApplication `json:"group_membership_applications,omitempty"`
	// Games holds the value of the games edge.
	Games []*Game `json:"games,omitempty"`
	// FavoriteGames holds the value of the favorite_games edge.
	FavoriteGames []*GameFavorite `json:"favorite_games,omitempty"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [8]bool
	// totalCount holds the count of the edges above.
	totalCount [5]map[string]int

	namedPlayers                     map[string][]*Player
	namedSentSupervisionRequests     map[string][]*PlayerSupervisionRequest
	namedSupervisionRequestApprovals map[string][]*PlayerSupervisionRequestApproval
	namedGroupMemberships            map[string][]*GroupMembership
	namedGroupMembershipApplications map[string][]*GroupMembershipApplication
	namedGames                       map[string][]*Game
	namedFavoriteGames               map[string][]*GameFavorite
}

// PlayersOrErr returns the Players value or an error if the edge
// was not loaded in eager-loading.
func (e UserEdges) PlayersOrErr() ([]*Player, error) {
	if e.loadedTypes[0] {
		return e.Players, nil
	}
	return nil, &NotLoadedError{edge: "players"}
}

// MainPlayerOrErr returns the MainPlayer value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e UserEdges) MainPlayerOrErr() (*Player, error) {
	if e.MainPlayer != nil {
		return e.MainPlayer, nil
	} else if e.loadedTypes[1] {
		return nil, &NotFoundError{label: player.Label}
	}
	return nil, &NotLoadedError{edge: "main_player"}
}

// SentSupervisionRequestsOrErr returns the SentSupervisionRequests value or an error if the edge
// was not loaded in eager-loading.
func (e UserEdges) SentSupervisionRequestsOrErr() ([]*PlayerSupervisionRequest, error) {
	if e.loadedTypes[2] {
		return e.SentSupervisionRequests, nil
	}
	return nil, &NotLoadedError{edge: "sent_supervision_requests"}
}

// SupervisionRequestApprovalsOrErr returns the SupervisionRequestApprovals value or an error if the edge
// was not loaded in eager-loading.
func (e UserEdges) SupervisionRequestApprovalsOrErr() ([]*PlayerSupervisionRequestApproval, error) {
	if e.loadedTypes[3] {
		return e.SupervisionRequestApprovals, nil
	}
	return nil, &NotLoadedError{edge: "supervision_request_approvals"}
}

// GroupMembershipsOrErr returns the GroupMemberships value or an error if the edge
// was not loaded in eager-loading.
func (e UserEdges) GroupMembershipsOrErr() ([]*GroupMembership, error) {
	if e.loadedTypes[4] {
		return e.GroupMemberships, nil
	}
	return nil, &NotLoadedError{edge: "group_memberships"}
}

// GroupMembershipApplicationsOrErr returns the GroupMembershipApplications value or an error if the edge
// was not loaded in eager-loading.
func (e UserEdges) GroupMembershipApplicationsOrErr() ([]*GroupMembershipApplication, error) {
	if e.loadedTypes[5] {
		return e.GroupMembershipApplications, nil
	}
	return nil, &NotLoadedError{edge: "group_membership_applications"}
}

// GamesOrErr returns the Games value or an error if the edge
// was not loaded in eager-loading.
func (e UserEdges) GamesOrErr() ([]*Game, error) {
	if e.loadedTypes[6] {
		return e.Games, nil
	}
	return nil, &NotLoadedError{edge: "games"}
}

// FavoriteGamesOrErr returns the FavoriteGames value or an error if the edge
// was not loaded in eager-loading.
func (e UserEdges) FavoriteGamesOrErr() ([]*GameFavorite, error) {
	if e.loadedTypes[7] {
		return e.FavoriteGames, nil
	}
	return nil, &NotLoadedError{edge: "favorite_games"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*User) scanValues(columns []string) ([]any, error) {
	values := make([]any, len(columns))
	for i := range columns {
		switch columns[i] {
		case user.FieldID:
			values[i] = new(guidgql.GUID)
		case user.FieldName, user.FieldEmail, user.FieldPassword, user.FieldAvatarURL:
			values[i] = new(sql.NullString)
		default:
			values[i] = new(sql.UnknownType)
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the User fields.
func (u *User) assignValues(columns []string, values []any) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case user.FieldID:
			if value, ok := values[i].(*guidgql.GUID); !ok {
				return fmt.Errorf("unexpected type %T for field id", values[i])
			} else if value != nil {
				u.ID = *value
			}
		case user.FieldName:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field name", values[i])
			} else if value.Valid {
				u.Name = value.String
			}
		case user.FieldEmail:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field email", values[i])
			} else if value.Valid {
				u.Email = value.String
			}
		case user.FieldPassword:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field password", values[i])
			} else if value.Valid {
				u.Password = value.String
			}
		case user.FieldAvatarURL:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field avatar_url", values[i])
			} else if value.Valid {
				u.AvatarURL = value.String
			}
		default:
			u.selectValues.Set(columns[i], values[i])
		}
	}
	return nil
}

// Value returns the ent.Value that was dynamically selected and assigned to the User.
// This includes values selected through modifiers, order, etc.
func (u *User) Value(name string) (ent.Value, error) {
	return u.selectValues.Get(name)
}

// QueryPlayers queries the "players" edge of the User entity.
func (u *User) QueryPlayers() *PlayerQuery {
	return NewUserClient(u.config).QueryPlayers(u)
}

// QueryMainPlayer queries the "main_player" edge of the User entity.
func (u *User) QueryMainPlayer() *PlayerQuery {
	return NewUserClient(u.config).QueryMainPlayer(u)
}

// QuerySentSupervisionRequests queries the "sent_supervision_requests" edge of the User entity.
func (u *User) QuerySentSupervisionRequests() *PlayerSupervisionRequestQuery {
	return NewUserClient(u.config).QuerySentSupervisionRequests(u)
}

// QuerySupervisionRequestApprovals queries the "supervision_request_approvals" edge of the User entity.
func (u *User) QuerySupervisionRequestApprovals() *PlayerSupervisionRequestApprovalQuery {
	return NewUserClient(u.config).QuerySupervisionRequestApprovals(u)
}

// QueryGroupMemberships queries the "group_memberships" edge of the User entity.
func (u *User) QueryGroupMemberships() *GroupMembershipQuery {
	return NewUserClient(u.config).QueryGroupMemberships(u)
}

// QueryGroupMembershipApplications queries the "group_membership_applications" edge of the User entity.
func (u *User) QueryGroupMembershipApplications() *GroupMembershipApplicationQuery {
	return NewUserClient(u.config).QueryGroupMembershipApplications(u)
}

// QueryGames queries the "games" edge of the User entity.
func (u *User) QueryGames() *GameQuery {
	return NewUserClient(u.config).QueryGames(u)
}

// QueryFavoriteGames queries the "favorite_games" edge of the User entity.
func (u *User) QueryFavoriteGames() *GameFavoriteQuery {
	return NewUserClient(u.config).QueryFavoriteGames(u)
}

// Update returns a builder for updating this User.
// Note that you need to call User.Unwrap() before calling this method if this User
// was returned from a transaction, and the transaction was committed or rolled back.
func (u *User) Update() *UserUpdateOne {
	return NewUserClient(u.config).UpdateOne(u)
}

// Unwrap unwraps the User entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (u *User) Unwrap() *User {
	_tx, ok := u.config.driver.(*txDriver)
	if !ok {
		panic("ent: User is not a transactional entity")
	}
	u.config.driver = _tx.drv
	return u
}

// String implements the fmt.Stringer.
func (u *User) String() string {
	var builder strings.Builder
	builder.WriteString("User(")
	builder.WriteString(fmt.Sprintf("id=%v, ", u.ID))
	builder.WriteString("name=")
	builder.WriteString(u.Name)
	builder.WriteString(", ")
	builder.WriteString("email=")
	builder.WriteString(u.Email)
	builder.WriteString(", ")
	builder.WriteString("password=<sensitive>")
	builder.WriteString(", ")
	builder.WriteString("avatar_url=")
	builder.WriteString(u.AvatarURL)
	builder.WriteByte(')')
	return builder.String()
}

// NamedPlayers returns the Players named value or an error if the edge was not
// loaded in eager-loading with this name.
func (u *User) NamedPlayers(name string) ([]*Player, error) {
	if u.Edges.namedPlayers == nil {
		return nil, &NotLoadedError{edge: name}
	}
	nodes, ok := u.Edges.namedPlayers[name]
	if !ok {
		return nil, &NotLoadedError{edge: name}
	}
	return nodes, nil
}

func (u *User) appendNamedPlayers(name string, edges ...*Player) {
	if u.Edges.namedPlayers == nil {
		u.Edges.namedPlayers = make(map[string][]*Player)
	}
	if len(edges) == 0 {
		u.Edges.namedPlayers[name] = []*Player{}
	} else {
		u.Edges.namedPlayers[name] = append(u.Edges.namedPlayers[name], edges...)
	}
}

// NamedSentSupervisionRequests returns the SentSupervisionRequests named value or an error if the edge was not
// loaded in eager-loading with this name.
func (u *User) NamedSentSupervisionRequests(name string) ([]*PlayerSupervisionRequest, error) {
	if u.Edges.namedSentSupervisionRequests == nil {
		return nil, &NotLoadedError{edge: name}
	}
	nodes, ok := u.Edges.namedSentSupervisionRequests[name]
	if !ok {
		return nil, &NotLoadedError{edge: name}
	}
	return nodes, nil
}

func (u *User) appendNamedSentSupervisionRequests(name string, edges ...*PlayerSupervisionRequest) {
	if u.Edges.namedSentSupervisionRequests == nil {
		u.Edges.namedSentSupervisionRequests = make(map[string][]*PlayerSupervisionRequest)
	}
	if len(edges) == 0 {
		u.Edges.namedSentSupervisionRequests[name] = []*PlayerSupervisionRequest{}
	} else {
		u.Edges.namedSentSupervisionRequests[name] = append(u.Edges.namedSentSupervisionRequests[name], edges...)
	}
}

// NamedSupervisionRequestApprovals returns the SupervisionRequestApprovals named value or an error if the edge was not
// loaded in eager-loading with this name.
func (u *User) NamedSupervisionRequestApprovals(name string) ([]*PlayerSupervisionRequestApproval, error) {
	if u.Edges.namedSupervisionRequestApprovals == nil {
		return nil, &NotLoadedError{edge: name}
	}
	nodes, ok := u.Edges.namedSupervisionRequestApprovals[name]
	if !ok {
		return nil, &NotLoadedError{edge: name}
	}
	return nodes, nil
}

func (u *User) appendNamedSupervisionRequestApprovals(name string, edges ...*PlayerSupervisionRequestApproval) {
	if u.Edges.namedSupervisionRequestApprovals == nil {
		u.Edges.namedSupervisionRequestApprovals = make(map[string][]*PlayerSupervisionRequestApproval)
	}
	if len(edges) == 0 {
		u.Edges.namedSupervisionRequestApprovals[name] = []*PlayerSupervisionRequestApproval{}
	} else {
		u.Edges.namedSupervisionRequestApprovals[name] = append(u.Edges.namedSupervisionRequestApprovals[name], edges...)
	}
}

// NamedGroupMemberships returns the GroupMemberships named value or an error if the edge was not
// loaded in eager-loading with this name.
func (u *User) NamedGroupMemberships(name string) ([]*GroupMembership, error) {
	if u.Edges.namedGroupMemberships == nil {
		return nil, &NotLoadedError{edge: name}
	}
	nodes, ok := u.Edges.namedGroupMemberships[name]
	if !ok {
		return nil, &NotLoadedError{edge: name}
	}
	return nodes, nil
}

func (u *User) appendNamedGroupMemberships(name string, edges ...*GroupMembership) {
	if u.Edges.namedGroupMemberships == nil {
		u.Edges.namedGroupMemberships = make(map[string][]*GroupMembership)
	}
	if len(edges) == 0 {
		u.Edges.namedGroupMemberships[name] = []*GroupMembership{}
	} else {
		u.Edges.namedGroupMemberships[name] = append(u.Edges.namedGroupMemberships[name], edges...)
	}
}

// NamedGroupMembershipApplications returns the GroupMembershipApplications named value or an error if the edge was not
// loaded in eager-loading with this name.
func (u *User) NamedGroupMembershipApplications(name string) ([]*GroupMembershipApplication, error) {
	if u.Edges.namedGroupMembershipApplications == nil {
		return nil, &NotLoadedError{edge: name}
	}
	nodes, ok := u.Edges.namedGroupMembershipApplications[name]
	if !ok {
		return nil, &NotLoadedError{edge: name}
	}
	return nodes, nil
}

func (u *User) appendNamedGroupMembershipApplications(name string, edges ...*GroupMembershipApplication) {
	if u.Edges.namedGroupMembershipApplications == nil {
		u.Edges.namedGroupMembershipApplications = make(map[string][]*GroupMembershipApplication)
	}
	if len(edges) == 0 {
		u.Edges.namedGroupMembershipApplications[name] = []*GroupMembershipApplication{}
	} else {
		u.Edges.namedGroupMembershipApplications[name] = append(u.Edges.namedGroupMembershipApplications[name], edges...)
	}
}

// NamedGames returns the Games named value or an error if the edge was not
// loaded in eager-loading with this name.
func (u *User) NamedGames(name string) ([]*Game, error) {
	if u.Edges.namedGames == nil {
		return nil, &NotLoadedError{edge: name}
	}
	nodes, ok := u.Edges.namedGames[name]
	if !ok {
		return nil, &NotLoadedError{edge: name}
	}
	return nodes, nil
}

func (u *User) appendNamedGames(name string, edges ...*Game) {
	if u.Edges.namedGames == nil {
		u.Edges.namedGames = make(map[string][]*Game)
	}
	if len(edges) == 0 {
		u.Edges.namedGames[name] = []*Game{}
	} else {
		u.Edges.namedGames[name] = append(u.Edges.namedGames[name], edges...)
	}
}

// NamedFavoriteGames returns the FavoriteGames named value or an error if the edge was not
// loaded in eager-loading with this name.
func (u *User) NamedFavoriteGames(name string) ([]*GameFavorite, error) {
	if u.Edges.namedFavoriteGames == nil {
		return nil, &NotLoadedError{edge: name}
	}
	nodes, ok := u.Edges.namedFavoriteGames[name]
	if !ok {
		return nil, &NotLoadedError{edge: name}
	}
	return nodes, nil
}

func (u *User) appendNamedFavoriteGames(name string, edges ...*GameFavorite) {
	if u.Edges.namedFavoriteGames == nil {
		u.Edges.namedFavoriteGames = make(map[string][]*GameFavorite)
	}
	if len(edges) == 0 {
		u.Edges.namedFavoriteGames[name] = []*GameFavorite{}
	} else {
		u.Edges.namedFavoriteGames[name] = append(u.Edges.namedFavoriteGames[name], edges...)
	}
}

// Users is a parsable slice of User.
type Users []*User

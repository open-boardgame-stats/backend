// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"encoding/json"
	"fmt"

	"entgo.io/contrib/entgql"
	"github.com/99designs/gqlgen/graphql"
	"github.com/hashicorp/go-multierror"
	"github.com/open-boardgame-stats/backend/internal/ent/game"
	"github.com/open-boardgame-stats/backend/internal/ent/group"
	"github.com/open-boardgame-stats/backend/internal/ent/groupmembership"
	"github.com/open-boardgame-stats/backend/internal/ent/groupmembershipapplication"
	"github.com/open-boardgame-stats/backend/internal/ent/groupsettings"
	"github.com/open-boardgame-stats/backend/internal/ent/player"
	"github.com/open-boardgame-stats/backend/internal/ent/playersupervisionrequest"
	"github.com/open-boardgame-stats/backend/internal/ent/playersupervisionrequestapproval"
	"github.com/open-boardgame-stats/backend/internal/ent/schema/guidgql"
	"github.com/open-boardgame-stats/backend/internal/ent/statdescription"
	"github.com/open-boardgame-stats/backend/internal/ent/user"
)

// Noder wraps the basic Node method.
type Noder interface {
	Node(context.Context) (*Node, error)
}

// Node in the graph.
type Node struct {
	ID     guidgql.GUID `json:"id,omitempty"`     // node id.
	Type   string       `json:"type,omitempty"`   // node type.
	Fields []*Field     `json:"fields,omitempty"` // node fields.
	Edges  []*Edge      `json:"edges,omitempty"`  // node edges.
}

// Field of a node.
type Field struct {
	Type  string `json:"type,omitempty"`  // field type.
	Name  string `json:"name,omitempty"`  // field name (as in struct).
	Value string `json:"value,omitempty"` // stringified value.
}

// Edges between two nodes.
type Edge struct {
	Type string         `json:"type,omitempty"` // edge type.
	Name string         `json:"name,omitempty"` // edge name.
	IDs  []guidgql.GUID `json:"ids,omitempty"`  // node ids (where this edge point to).
}

func (ga *Game) Node(ctx context.Context) (node *Node, err error) {
	node = &Node{
		ID:     ga.ID,
		Type:   "Game",
		Fields: make([]*Field, 5),
		Edges:  make([]*Edge, 2),
	}
	var buf []byte
	if buf, err = json.Marshal(ga.Name); err != nil {
		return nil, err
	}
	node.Fields[0] = &Field{
		Type:  "string",
		Name:  "name",
		Value: string(buf),
	}
	if buf, err = json.Marshal(ga.MinPlayers); err != nil {
		return nil, err
	}
	node.Fields[1] = &Field{
		Type:  "int",
		Name:  "min_players",
		Value: string(buf),
	}
	if buf, err = json.Marshal(ga.MaxPlayers); err != nil {
		return nil, err
	}
	node.Fields[2] = &Field{
		Type:  "int",
		Name:  "max_players",
		Value: string(buf),
	}
	if buf, err = json.Marshal(ga.Description); err != nil {
		return nil, err
	}
	node.Fields[3] = &Field{
		Type:  "string",
		Name:  "description",
		Value: string(buf),
	}
	if buf, err = json.Marshal(ga.BoardgamegeekURL); err != nil {
		return nil, err
	}
	node.Fields[4] = &Field{
		Type:  "string",
		Name:  "boardgamegeek_url",
		Value: string(buf),
	}
	node.Edges[0] = &Edge{
		Type: "User",
		Name: "author",
	}
	err = ga.QueryAuthor().
		Select(user.FieldID).
		Scan(ctx, &node.Edges[0].IDs)
	if err != nil {
		return nil, err
	}
	node.Edges[1] = &Edge{
		Type: "StatDescription",
		Name: "stat_descriptions",
	}
	err = ga.QueryStatDescriptions().
		Select(statdescription.FieldID).
		Scan(ctx, &node.Edges[1].IDs)
	if err != nil {
		return nil, err
	}
	return node, nil
}

func (gr *Group) Node(ctx context.Context) (node *Node, err error) {
	node = &Node{
		ID:     gr.ID,
		Type:   "Group",
		Fields: make([]*Field, 3),
		Edges:  make([]*Edge, 3),
	}
	var buf []byte
	if buf, err = json.Marshal(gr.Name); err != nil {
		return nil, err
	}
	node.Fields[0] = &Field{
		Type:  "string",
		Name:  "name",
		Value: string(buf),
	}
	if buf, err = json.Marshal(gr.Description); err != nil {
		return nil, err
	}
	node.Fields[1] = &Field{
		Type:  "string",
		Name:  "description",
		Value: string(buf),
	}
	if buf, err = json.Marshal(gr.LogoURL); err != nil {
		return nil, err
	}
	node.Fields[2] = &Field{
		Type:  "string",
		Name:  "logo_url",
		Value: string(buf),
	}
	node.Edges[0] = &Edge{
		Type: "GroupSettings",
		Name: "settings",
	}
	err = gr.QuerySettings().
		Select(groupsettings.FieldID).
		Scan(ctx, &node.Edges[0].IDs)
	if err != nil {
		return nil, err
	}
	node.Edges[1] = &Edge{
		Type: "GroupMembership",
		Name: "members",
	}
	err = gr.QueryMembers().
		Select(groupmembership.FieldID).
		Scan(ctx, &node.Edges[1].IDs)
	if err != nil {
		return nil, err
	}
	node.Edges[2] = &Edge{
		Type: "GroupMembershipApplication",
		Name: "applications",
	}
	err = gr.QueryApplications().
		Select(groupmembershipapplication.FieldID).
		Scan(ctx, &node.Edges[2].IDs)
	if err != nil {
		return nil, err
	}
	return node, nil
}

func (gm *GroupMembership) Node(ctx context.Context) (node *Node, err error) {
	node = &Node{
		ID:     gm.ID,
		Type:   "GroupMembership",
		Fields: make([]*Field, 1),
		Edges:  make([]*Edge, 2),
	}
	var buf []byte
	if buf, err = json.Marshal(gm.Role); err != nil {
		return nil, err
	}
	node.Fields[0] = &Field{
		Type:  "enums.Role",
		Name:  "role",
		Value: string(buf),
	}
	node.Edges[0] = &Edge{
		Type: "Group",
		Name: "group",
	}
	err = gm.QueryGroup().
		Select(group.FieldID).
		Scan(ctx, &node.Edges[0].IDs)
	if err != nil {
		return nil, err
	}
	node.Edges[1] = &Edge{
		Type: "User",
		Name: "user",
	}
	err = gm.QueryUser().
		Select(user.FieldID).
		Scan(ctx, &node.Edges[1].IDs)
	if err != nil {
		return nil, err
	}
	return node, nil
}

func (gma *GroupMembershipApplication) Node(ctx context.Context) (node *Node, err error) {
	node = &Node{
		ID:     gma.ID,
		Type:   "GroupMembershipApplication",
		Fields: make([]*Field, 1),
		Edges:  make([]*Edge, 2),
	}
	var buf []byte
	if buf, err = json.Marshal(gma.Message); err != nil {
		return nil, err
	}
	node.Fields[0] = &Field{
		Type:  "string",
		Name:  "message",
		Value: string(buf),
	}
	node.Edges[0] = &Edge{
		Type: "User",
		Name: "user",
	}
	err = gma.QueryUser().
		Select(user.FieldID).
		Scan(ctx, &node.Edges[0].IDs)
	if err != nil {
		return nil, err
	}
	node.Edges[1] = &Edge{
		Type: "Group",
		Name: "group",
	}
	err = gma.QueryGroup().
		Select(group.FieldID).
		Scan(ctx, &node.Edges[1].IDs)
	if err != nil {
		return nil, err
	}
	return node, nil
}

func (gs *GroupSettings) Node(ctx context.Context) (node *Node, err error) {
	node = &Node{
		ID:     gs.ID,
		Type:   "GroupSettings",
		Fields: make([]*Field, 3),
		Edges:  make([]*Edge, 0),
	}
	var buf []byte
	if buf, err = json.Marshal(gs.Visibility); err != nil {
		return nil, err
	}
	node.Fields[0] = &Field{
		Type:  "groupsettings.Visibility",
		Name:  "visibility",
		Value: string(buf),
	}
	if buf, err = json.Marshal(gs.JoinPolicy); err != nil {
		return nil, err
	}
	node.Fields[1] = &Field{
		Type:  "groupsettings.JoinPolicy",
		Name:  "join_policy",
		Value: string(buf),
	}
	if buf, err = json.Marshal(gs.MinimumRoleToInvite); err != nil {
		return nil, err
	}
	node.Fields[2] = &Field{
		Type:  "enums.Role",
		Name:  "minimum_role_to_invite",
		Value: string(buf),
	}
	return node, nil
}

func (pl *Player) Node(ctx context.Context) (node *Node, err error) {
	node = &Node{
		ID:     pl.ID,
		Type:   "Player",
		Fields: make([]*Field, 1),
		Edges:  make([]*Edge, 3),
	}
	var buf []byte
	if buf, err = json.Marshal(pl.Name); err != nil {
		return nil, err
	}
	node.Fields[0] = &Field{
		Type:  "string",
		Name:  "name",
		Value: string(buf),
	}
	node.Edges[0] = &Edge{
		Type: "User",
		Name: "owner",
	}
	err = pl.QueryOwner().
		Select(user.FieldID).
		Scan(ctx, &node.Edges[0].IDs)
	if err != nil {
		return nil, err
	}
	node.Edges[1] = &Edge{
		Type: "User",
		Name: "supervisors",
	}
	err = pl.QuerySupervisors().
		Select(user.FieldID).
		Scan(ctx, &node.Edges[1].IDs)
	if err != nil {
		return nil, err
	}
	node.Edges[2] = &Edge{
		Type: "PlayerSupervisionRequest",
		Name: "supervision_requests",
	}
	err = pl.QuerySupervisionRequests().
		Select(playersupervisionrequest.FieldID).
		Scan(ctx, &node.Edges[2].IDs)
	if err != nil {
		return nil, err
	}
	return node, nil
}

func (psr *PlayerSupervisionRequest) Node(ctx context.Context) (node *Node, err error) {
	node = &Node{
		ID:     psr.ID,
		Type:   "PlayerSupervisionRequest",
		Fields: make([]*Field, 1),
		Edges:  make([]*Edge, 3),
	}
	var buf []byte
	if buf, err = json.Marshal(psr.Message); err != nil {
		return nil, err
	}
	node.Fields[0] = &Field{
		Type:  "string",
		Name:  "message",
		Value: string(buf),
	}
	node.Edges[0] = &Edge{
		Type: "User",
		Name: "sender",
	}
	err = psr.QuerySender().
		Select(user.FieldID).
		Scan(ctx, &node.Edges[0].IDs)
	if err != nil {
		return nil, err
	}
	node.Edges[1] = &Edge{
		Type: "Player",
		Name: "player",
	}
	err = psr.QueryPlayer().
		Select(player.FieldID).
		Scan(ctx, &node.Edges[1].IDs)
	if err != nil {
		return nil, err
	}
	node.Edges[2] = &Edge{
		Type: "PlayerSupervisionRequestApproval",
		Name: "approvals",
	}
	err = psr.QueryApprovals().
		Select(playersupervisionrequestapproval.FieldID).
		Scan(ctx, &node.Edges[2].IDs)
	if err != nil {
		return nil, err
	}
	return node, nil
}

func (psra *PlayerSupervisionRequestApproval) Node(ctx context.Context) (node *Node, err error) {
	node = &Node{
		ID:     psra.ID,
		Type:   "PlayerSupervisionRequestApproval",
		Fields: make([]*Field, 1),
		Edges:  make([]*Edge, 2),
	}
	var buf []byte
	if buf, err = json.Marshal(psra.Approved); err != nil {
		return nil, err
	}
	node.Fields[0] = &Field{
		Type:  "bool",
		Name:  "approved",
		Value: string(buf),
	}
	node.Edges[0] = &Edge{
		Type: "User",
		Name: "approver",
	}
	err = psra.QueryApprover().
		Select(user.FieldID).
		Scan(ctx, &node.Edges[0].IDs)
	if err != nil {
		return nil, err
	}
	node.Edges[1] = &Edge{
		Type: "PlayerSupervisionRequest",
		Name: "supervision_request",
	}
	err = psra.QuerySupervisionRequest().
		Select(playersupervisionrequest.FieldID).
		Scan(ctx, &node.Edges[1].IDs)
	if err != nil {
		return nil, err
	}
	return node, nil
}

func (sd *StatDescription) Node(ctx context.Context) (node *Node, err error) {
	node = &Node{
		ID:     sd.ID,
		Type:   "StatDescription",
		Fields: make([]*Field, 4),
		Edges:  make([]*Edge, 0),
	}
	var buf []byte
	if buf, err = json.Marshal(sd.Type); err != nil {
		return nil, err
	}
	node.Fields[0] = &Field{
		Type:  "stat.StatType",
		Name:  "type",
		Value: string(buf),
	}
	if buf, err = json.Marshal(sd.Name); err != nil {
		return nil, err
	}
	node.Fields[1] = &Field{
		Type:  "string",
		Name:  "name",
		Value: string(buf),
	}
	if buf, err = json.Marshal(sd.Description); err != nil {
		return nil, err
	}
	node.Fields[2] = &Field{
		Type:  "string",
		Name:  "description",
		Value: string(buf),
	}
	if buf, err = json.Marshal(sd.PossibleValues); err != nil {
		return nil, err
	}
	node.Fields[3] = &Field{
		Type:  "[]string",
		Name:  "possible_values",
		Value: string(buf),
	}
	return node, nil
}

func (u *User) Node(ctx context.Context) (node *Node, err error) {
	node = &Node{
		ID:     u.ID,
		Type:   "User",
		Fields: make([]*Field, 3),
		Edges:  make([]*Edge, 5),
	}
	var buf []byte
	if buf, err = json.Marshal(u.Name); err != nil {
		return nil, err
	}
	node.Fields[0] = &Field{
		Type:  "string",
		Name:  "name",
		Value: string(buf),
	}
	if buf, err = json.Marshal(u.Email); err != nil {
		return nil, err
	}
	node.Fields[1] = &Field{
		Type:  "string",
		Name:  "email",
		Value: string(buf),
	}
	if buf, err = json.Marshal(u.AvatarURL); err != nil {
		return nil, err
	}
	node.Fields[2] = &Field{
		Type:  "string",
		Name:  "avatar_url",
		Value: string(buf),
	}
	node.Edges[0] = &Edge{
		Type: "Player",
		Name: "players",
	}
	err = u.QueryPlayers().
		Select(player.FieldID).
		Scan(ctx, &node.Edges[0].IDs)
	if err != nil {
		return nil, err
	}
	node.Edges[1] = &Edge{
		Type: "Player",
		Name: "main_player",
	}
	err = u.QueryMainPlayer().
		Select(player.FieldID).
		Scan(ctx, &node.Edges[1].IDs)
	if err != nil {
		return nil, err
	}
	node.Edges[2] = &Edge{
		Type: "GroupMembership",
		Name: "group_memberships",
	}
	err = u.QueryGroupMemberships().
		Select(groupmembership.FieldID).
		Scan(ctx, &node.Edges[2].IDs)
	if err != nil {
		return nil, err
	}
	node.Edges[3] = &Edge{
		Type: "GroupMembershipApplication",
		Name: "group_membership_applications",
	}
	err = u.QueryGroupMembershipApplications().
		Select(groupmembershipapplication.FieldID).
		Scan(ctx, &node.Edges[3].IDs)
	if err != nil {
		return nil, err
	}
	node.Edges[4] = &Edge{
		Type: "Game",
		Name: "games",
	}
	err = u.QueryGames().
		Select(game.FieldID).
		Scan(ctx, &node.Edges[4].IDs)
	if err != nil {
		return nil, err
	}
	return node, nil
}

func (c *Client) Node(ctx context.Context, id guidgql.GUID) (*Node, error) {
	n, err := c.Noder(ctx, id)
	if err != nil {
		return nil, err
	}
	return n.Node(ctx)
}

var errNodeInvalidID = &NotFoundError{"node"}

// NodeOption allows configuring the Noder execution using functional options.
type NodeOption func(*nodeOptions)

// WithNodeType sets the node Type resolver function (i.e. the table to query).
// If was not provided, the table will be derived from the universal-id
// configuration as described in: https://entgo.io/docs/migrate/#universal-ids.
func WithNodeType(f func(context.Context, guidgql.GUID) (string, error)) NodeOption {
	return func(o *nodeOptions) {
		o.nodeType = f
	}
}

// WithFixedNodeType sets the Type of the node to a fixed value.
func WithFixedNodeType(t string) NodeOption {
	return WithNodeType(func(context.Context, guidgql.GUID) (string, error) {
		return t, nil
	})
}

type nodeOptions struct {
	nodeType func(context.Context, guidgql.GUID) (string, error)
}

func (c *Client) newNodeOpts(opts []NodeOption) *nodeOptions {
	nopts := &nodeOptions{}
	for _, opt := range opts {
		opt(nopts)
	}
	if nopts.nodeType == nil {
		nopts.nodeType = func(ctx context.Context, id guidgql.GUID) (string, error) {
			return "", fmt.Errorf("cannot resolve noder (%v) without its type", id)
		}
	}
	return nopts
}

// Noder returns a Node by its id. If the NodeType was not provided, it will
// be derived from the id value according to the universal-id configuration.
//
//	c.Noder(ctx, id)
//	c.Noder(ctx, id, ent.WithNodeType(typeResolver))
func (c *Client) Noder(ctx context.Context, id guidgql.GUID, opts ...NodeOption) (_ Noder, err error) {
	defer func() {
		if IsNotFound(err) {
			err = multierror.Append(err, entgql.ErrNodeNotFound(id))
		}
	}()
	table, err := c.newNodeOpts(opts).nodeType(ctx, id)
	if err != nil {
		return nil, err
	}
	return c.noder(ctx, table, id)
}

func (c *Client) noder(ctx context.Context, table string, id guidgql.GUID) (Noder, error) {
	switch table {
	case game.Table:
		var uid guidgql.GUID
		if err := uid.UnmarshalGQL(id); err != nil {
			return nil, err
		}
		query := c.Game.Query().
			Where(game.ID(uid))
		query, err := query.CollectFields(ctx, "Game")
		if err != nil {
			return nil, err
		}
		n, err := query.Only(ctx)
		if err != nil {
			return nil, err
		}
		return n, nil
	case group.Table:
		var uid guidgql.GUID
		if err := uid.UnmarshalGQL(id); err != nil {
			return nil, err
		}
		query := c.Group.Query().
			Where(group.ID(uid))
		query, err := query.CollectFields(ctx, "Group")
		if err != nil {
			return nil, err
		}
		n, err := query.Only(ctx)
		if err != nil {
			return nil, err
		}
		return n, nil
	case groupmembership.Table:
		var uid guidgql.GUID
		if err := uid.UnmarshalGQL(id); err != nil {
			return nil, err
		}
		query := c.GroupMembership.Query().
			Where(groupmembership.ID(uid))
		query, err := query.CollectFields(ctx, "GroupMembership")
		if err != nil {
			return nil, err
		}
		n, err := query.Only(ctx)
		if err != nil {
			return nil, err
		}
		return n, nil
	case groupmembershipapplication.Table:
		var uid guidgql.GUID
		if err := uid.UnmarshalGQL(id); err != nil {
			return nil, err
		}
		query := c.GroupMembershipApplication.Query().
			Where(groupmembershipapplication.ID(uid))
		query, err := query.CollectFields(ctx, "GroupMembershipApplication")
		if err != nil {
			return nil, err
		}
		n, err := query.Only(ctx)
		if err != nil {
			return nil, err
		}
		return n, nil
	case groupsettings.Table:
		var uid guidgql.GUID
		if err := uid.UnmarshalGQL(id); err != nil {
			return nil, err
		}
		query := c.GroupSettings.Query().
			Where(groupsettings.ID(uid))
		query, err := query.CollectFields(ctx, "GroupSettings")
		if err != nil {
			return nil, err
		}
		n, err := query.Only(ctx)
		if err != nil {
			return nil, err
		}
		return n, nil
	case player.Table:
		var uid guidgql.GUID
		if err := uid.UnmarshalGQL(id); err != nil {
			return nil, err
		}
		query := c.Player.Query().
			Where(player.ID(uid))
		query, err := query.CollectFields(ctx, "Player")
		if err != nil {
			return nil, err
		}
		n, err := query.Only(ctx)
		if err != nil {
			return nil, err
		}
		return n, nil
	case playersupervisionrequest.Table:
		var uid guidgql.GUID
		if err := uid.UnmarshalGQL(id); err != nil {
			return nil, err
		}
		query := c.PlayerSupervisionRequest.Query().
			Where(playersupervisionrequest.ID(uid))
		query, err := query.CollectFields(ctx, "PlayerSupervisionRequest")
		if err != nil {
			return nil, err
		}
		n, err := query.Only(ctx)
		if err != nil {
			return nil, err
		}
		return n, nil
	case playersupervisionrequestapproval.Table:
		var uid guidgql.GUID
		if err := uid.UnmarshalGQL(id); err != nil {
			return nil, err
		}
		query := c.PlayerSupervisionRequestApproval.Query().
			Where(playersupervisionrequestapproval.ID(uid))
		query, err := query.CollectFields(ctx, "PlayerSupervisionRequestApproval")
		if err != nil {
			return nil, err
		}
		n, err := query.Only(ctx)
		if err != nil {
			return nil, err
		}
		return n, nil
	case statdescription.Table:
		var uid guidgql.GUID
		if err := uid.UnmarshalGQL(id); err != nil {
			return nil, err
		}
		query := c.StatDescription.Query().
			Where(statdescription.ID(uid))
		query, err := query.CollectFields(ctx, "StatDescription")
		if err != nil {
			return nil, err
		}
		n, err := query.Only(ctx)
		if err != nil {
			return nil, err
		}
		return n, nil
	case user.Table:
		var uid guidgql.GUID
		if err := uid.UnmarshalGQL(id); err != nil {
			return nil, err
		}
		query := c.User.Query().
			Where(user.ID(uid))
		query, err := query.CollectFields(ctx, "User")
		if err != nil {
			return nil, err
		}
		n, err := query.Only(ctx)
		if err != nil {
			return nil, err
		}
		return n, nil
	default:
		return nil, fmt.Errorf("cannot resolve noder from table %q: %w", table, errNodeInvalidID)
	}
}

func (c *Client) Noders(ctx context.Context, ids []guidgql.GUID, opts ...NodeOption) ([]Noder, error) {
	switch len(ids) {
	case 1:
		noder, err := c.Noder(ctx, ids[0], opts...)
		if err != nil {
			return nil, err
		}
		return []Noder{noder}, nil
	case 0:
		return []Noder{}, nil
	}

	noders := make([]Noder, len(ids))
	errors := make([]error, len(ids))
	tables := make(map[string][]guidgql.GUID)
	id2idx := make(map[guidgql.GUID][]int, len(ids))
	nopts := c.newNodeOpts(opts)
	for i, id := range ids {
		table, err := nopts.nodeType(ctx, id)
		if err != nil {
			errors[i] = err
			continue
		}
		tables[table] = append(tables[table], id)
		id2idx[id] = append(id2idx[id], i)
	}

	for table, ids := range tables {
		nodes, err := c.noders(ctx, table, ids)
		if err != nil {
			for _, id := range ids {
				for _, idx := range id2idx[id] {
					errors[idx] = err
				}
			}
		} else {
			for i, id := range ids {
				for _, idx := range id2idx[id] {
					noders[idx] = nodes[i]
				}
			}
		}
	}

	for i, id := range ids {
		if errors[i] == nil {
			if noders[i] != nil {
				continue
			}
			errors[i] = entgql.ErrNodeNotFound(id)
		} else if IsNotFound(errors[i]) {
			errors[i] = multierror.Append(errors[i], entgql.ErrNodeNotFound(id))
		}
		ctx := graphql.WithPathContext(ctx,
			graphql.NewPathWithIndex(i),
		)
		graphql.AddError(ctx, errors[i])
	}
	return noders, nil
}

func (c *Client) noders(ctx context.Context, table string, ids []guidgql.GUID) ([]Noder, error) {
	noders := make([]Noder, len(ids))
	idmap := make(map[guidgql.GUID][]*Noder, len(ids))
	for i, id := range ids {
		idmap[id] = append(idmap[id], &noders[i])
	}
	switch table {
	case game.Table:
		query := c.Game.Query().
			Where(game.IDIn(ids...))
		query, err := query.CollectFields(ctx, "Game")
		if err != nil {
			return nil, err
		}
		nodes, err := query.All(ctx)
		if err != nil {
			return nil, err
		}
		for _, node := range nodes {
			for _, noder := range idmap[node.ID] {
				*noder = node
			}
		}
	case group.Table:
		query := c.Group.Query().
			Where(group.IDIn(ids...))
		query, err := query.CollectFields(ctx, "Group")
		if err != nil {
			return nil, err
		}
		nodes, err := query.All(ctx)
		if err != nil {
			return nil, err
		}
		for _, node := range nodes {
			for _, noder := range idmap[node.ID] {
				*noder = node
			}
		}
	case groupmembership.Table:
		query := c.GroupMembership.Query().
			Where(groupmembership.IDIn(ids...))
		query, err := query.CollectFields(ctx, "GroupMembership")
		if err != nil {
			return nil, err
		}
		nodes, err := query.All(ctx)
		if err != nil {
			return nil, err
		}
		for _, node := range nodes {
			for _, noder := range idmap[node.ID] {
				*noder = node
			}
		}
	case groupmembershipapplication.Table:
		query := c.GroupMembershipApplication.Query().
			Where(groupmembershipapplication.IDIn(ids...))
		query, err := query.CollectFields(ctx, "GroupMembershipApplication")
		if err != nil {
			return nil, err
		}
		nodes, err := query.All(ctx)
		if err != nil {
			return nil, err
		}
		for _, node := range nodes {
			for _, noder := range idmap[node.ID] {
				*noder = node
			}
		}
	case groupsettings.Table:
		query := c.GroupSettings.Query().
			Where(groupsettings.IDIn(ids...))
		query, err := query.CollectFields(ctx, "GroupSettings")
		if err != nil {
			return nil, err
		}
		nodes, err := query.All(ctx)
		if err != nil {
			return nil, err
		}
		for _, node := range nodes {
			for _, noder := range idmap[node.ID] {
				*noder = node
			}
		}
	case player.Table:
		query := c.Player.Query().
			Where(player.IDIn(ids...))
		query, err := query.CollectFields(ctx, "Player")
		if err != nil {
			return nil, err
		}
		nodes, err := query.All(ctx)
		if err != nil {
			return nil, err
		}
		for _, node := range nodes {
			for _, noder := range idmap[node.ID] {
				*noder = node
			}
		}
	case playersupervisionrequest.Table:
		query := c.PlayerSupervisionRequest.Query().
			Where(playersupervisionrequest.IDIn(ids...))
		query, err := query.CollectFields(ctx, "PlayerSupervisionRequest")
		if err != nil {
			return nil, err
		}
		nodes, err := query.All(ctx)
		if err != nil {
			return nil, err
		}
		for _, node := range nodes {
			for _, noder := range idmap[node.ID] {
				*noder = node
			}
		}
	case playersupervisionrequestapproval.Table:
		query := c.PlayerSupervisionRequestApproval.Query().
			Where(playersupervisionrequestapproval.IDIn(ids...))
		query, err := query.CollectFields(ctx, "PlayerSupervisionRequestApproval")
		if err != nil {
			return nil, err
		}
		nodes, err := query.All(ctx)
		if err != nil {
			return nil, err
		}
		for _, node := range nodes {
			for _, noder := range idmap[node.ID] {
				*noder = node
			}
		}
	case statdescription.Table:
		query := c.StatDescription.Query().
			Where(statdescription.IDIn(ids...))
		query, err := query.CollectFields(ctx, "StatDescription")
		if err != nil {
			return nil, err
		}
		nodes, err := query.All(ctx)
		if err != nil {
			return nil, err
		}
		for _, node := range nodes {
			for _, noder := range idmap[node.ID] {
				*noder = node
			}
		}
	case user.Table:
		query := c.User.Query().
			Where(user.IDIn(ids...))
		query, err := query.CollectFields(ctx, "User")
		if err != nil {
			return nil, err
		}
		nodes, err := query.All(ctx)
		if err != nil {
			return nil, err
		}
		for _, node := range nodes {
			for _, noder := range idmap[node.ID] {
				*noder = node
			}
		}
	default:
		return nil, fmt.Errorf("cannot resolve noders from table %q: %w", table, errNodeInvalidID)
	}
	return noders, nil
}

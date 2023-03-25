// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"fmt"

	"entgo.io/contrib/entgql"
	"github.com/99designs/gqlgen/graphql"
	"github.com/hashicorp/go-multierror"
	"github.com/open-boardgame-stats/backend/internal/ent/game"
	"github.com/open-boardgame-stats/backend/internal/ent/group"
	"github.com/open-boardgame-stats/backend/internal/ent/groupmembership"
	"github.com/open-boardgame-stats/backend/internal/ent/groupmembershipapplication"
	"github.com/open-boardgame-stats/backend/internal/ent/groupsettings"
	"github.com/open-boardgame-stats/backend/internal/ent/match"
	"github.com/open-boardgame-stats/backend/internal/ent/player"
	"github.com/open-boardgame-stats/backend/internal/ent/playersupervisionrequest"
	"github.com/open-boardgame-stats/backend/internal/ent/playersupervisionrequestapproval"
	"github.com/open-boardgame-stats/backend/internal/ent/schema/guidgql"
	"github.com/open-boardgame-stats/backend/internal/ent/statdescription"
	"github.com/open-boardgame-stats/backend/internal/ent/statistic"
	"github.com/open-boardgame-stats/backend/internal/ent/user"
)

// Noder wraps the basic Node method.
type Noder interface {
	IsNode()
}

// IsNode implements the Node interface check for GQLGen.
func (n *Game) IsNode() {}

// IsNode implements the Node interface check for GQLGen.
func (n *Group) IsNode() {}

// IsNode implements the Node interface check for GQLGen.
func (n *GroupMembership) IsNode() {}

// IsNode implements the Node interface check for GQLGen.
func (n *GroupMembershipApplication) IsNode() {}

// IsNode implements the Node interface check for GQLGen.
func (n *GroupSettings) IsNode() {}

// IsNode implements the Node interface check for GQLGen.
func (n *Match) IsNode() {}

// IsNode implements the Node interface check for GQLGen.
func (n *Player) IsNode() {}

// IsNode implements the Node interface check for GQLGen.
func (n *PlayerSupervisionRequest) IsNode() {}

// IsNode implements the Node interface check for GQLGen.
func (n *PlayerSupervisionRequestApproval) IsNode() {}

// IsNode implements the Node interface check for GQLGen.
func (n *StatDescription) IsNode() {}

// IsNode implements the Node interface check for GQLGen.
func (n *Statistic) IsNode() {}

// IsNode implements the Node interface check for GQLGen.
func (n *User) IsNode() {}

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
	case match.Table:
		var uid guidgql.GUID
		if err := uid.UnmarshalGQL(id); err != nil {
			return nil, err
		}
		query := c.Match.Query().
			Where(match.ID(uid))
		query, err := query.CollectFields(ctx, "Match")
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
	case statistic.Table:
		var uid guidgql.GUID
		if err := uid.UnmarshalGQL(id); err != nil {
			return nil, err
		}
		query := c.Statistic.Query().
			Where(statistic.ID(uid))
		query, err := query.CollectFields(ctx, "Statistic")
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
	case match.Table:
		query := c.Match.Query().
			Where(match.IDIn(ids...))
		query, err := query.CollectFields(ctx, "Match")
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
	case statistic.Table:
		query := c.Statistic.Query().
			Where(statistic.IDIn(ids...))
		query, err := query.CollectFields(ctx, "Statistic")
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

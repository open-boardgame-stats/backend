// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"

	"github.com/99designs/gqlgen/graphql"
)

func (ga *Game) Author(ctx context.Context) (*User, error) {
	result, err := ga.Edges.AuthorOrErr()
	if IsNotLoaded(err) {
		result, err = ga.QueryAuthor().Only(ctx)
	}
	return result, err
}

func (gr *Group) Settings(ctx context.Context) (*GroupSettings, error) {
	result, err := gr.Edges.SettingsOrErr()
	if IsNotLoaded(err) {
		result, err = gr.QuerySettings().Only(ctx)
	}
	return result, err
}

func (gr *Group) Members(
	ctx context.Context, after *Cursor, first *int, before *Cursor, last *int, where *GroupMembershipWhereInput,
) (*GroupMembershipConnection, error) {
	opts := []GroupMembershipPaginateOption{
		WithGroupMembershipFilter(where.Filter),
	}
	alias := graphql.GetFieldContext(ctx).Field.Alias
	totalCount, hasTotalCount := gr.Edges.totalCount[1][alias]
	if nodes, err := gr.NamedMembers(alias); err == nil || hasTotalCount {
		pager, err := newGroupMembershipPager(opts)
		if err != nil {
			return nil, err
		}
		conn := &GroupMembershipConnection{Edges: []*GroupMembershipEdge{}, TotalCount: totalCount}
		conn.build(nodes, pager, after, first, before, last)
		return conn, nil
	}
	return gr.QueryMembers().Paginate(ctx, after, first, before, last, opts...)
}

func (gr *Group) Applications(ctx context.Context) (result []*GroupMembershipApplication, err error) {
	if fc := graphql.GetFieldContext(ctx); fc != nil && fc.Field.Alias != "" {
		result, err = gr.NamedApplications(graphql.GetFieldContext(ctx).Field.Alias)
	} else {
		result, err = gr.Edges.ApplicationsOrErr()
	}
	if IsNotLoaded(err) {
		result, err = gr.QueryApplications().All(ctx)
	}
	return result, err
}

func (gm *GroupMembership) Group(ctx context.Context) (*Group, error) {
	result, err := gm.Edges.GroupOrErr()
	if IsNotLoaded(err) {
		result, err = gm.QueryGroup().Only(ctx)
	}
	return result, err
}

func (gm *GroupMembership) User(ctx context.Context) (*User, error) {
	result, err := gm.Edges.UserOrErr()
	if IsNotLoaded(err) {
		result, err = gm.QueryUser().Only(ctx)
	}
	return result, err
}

func (gma *GroupMembershipApplication) User(ctx context.Context) (*User, error) {
	result, err := gma.Edges.UserOrErr()
	if IsNotLoaded(err) {
		result, err = gma.QueryUser().Only(ctx)
	}
	return result, err
}

func (gma *GroupMembershipApplication) Group(ctx context.Context) (*Group, error) {
	result, err := gma.Edges.GroupOrErr()
	if IsNotLoaded(err) {
		result, err = gma.QueryGroup().Only(ctx)
	}
	return result, err
}

func (pl *Player) Owner(ctx context.Context) (*User, error) {
	result, err := pl.Edges.OwnerOrErr()
	if IsNotLoaded(err) {
		result, err = pl.QueryOwner().Only(ctx)
	}
	return result, MaskNotFound(err)
}

func (pl *Player) Supervisors(ctx context.Context) (result []*User, err error) {
	if fc := graphql.GetFieldContext(ctx); fc != nil && fc.Field.Alias != "" {
		result, err = pl.NamedSupervisors(graphql.GetFieldContext(ctx).Field.Alias)
	} else {
		result, err = pl.Edges.SupervisorsOrErr()
	}
	if IsNotLoaded(err) {
		result, err = pl.QuerySupervisors().All(ctx)
	}
	return result, err
}

func (pl *Player) SupervisionRequests(ctx context.Context) (result []*PlayerSupervisionRequest, err error) {
	if fc := graphql.GetFieldContext(ctx); fc != nil && fc.Field.Alias != "" {
		result, err = pl.NamedSupervisionRequests(graphql.GetFieldContext(ctx).Field.Alias)
	} else {
		result, err = pl.Edges.SupervisionRequestsOrErr()
	}
	if IsNotLoaded(err) {
		result, err = pl.QuerySupervisionRequests().All(ctx)
	}
	return result, err
}

func (psr *PlayerSupervisionRequest) Sender(ctx context.Context) (*User, error) {
	result, err := psr.Edges.SenderOrErr()
	if IsNotLoaded(err) {
		result, err = psr.QuerySender().Only(ctx)
	}
	return result, err
}

func (psr *PlayerSupervisionRequest) Player(ctx context.Context) (*Player, error) {
	result, err := psr.Edges.PlayerOrErr()
	if IsNotLoaded(err) {
		result, err = psr.QueryPlayer().Only(ctx)
	}
	return result, err
}

func (psr *PlayerSupervisionRequest) Approvals(ctx context.Context) (result []*PlayerSupervisionRequestApproval, err error) {
	if fc := graphql.GetFieldContext(ctx); fc != nil && fc.Field.Alias != "" {
		result, err = psr.NamedApprovals(graphql.GetFieldContext(ctx).Field.Alias)
	} else {
		result, err = psr.Edges.ApprovalsOrErr()
	}
	if IsNotLoaded(err) {
		result, err = psr.QueryApprovals().All(ctx)
	}
	return result, err
}

func (psra *PlayerSupervisionRequestApproval) Approver(ctx context.Context) (*User, error) {
	result, err := psra.Edges.ApproverOrErr()
	if IsNotLoaded(err) {
		result, err = psra.QueryApprover().Only(ctx)
	}
	return result, err
}

func (psra *PlayerSupervisionRequestApproval) SupervisionRequest(ctx context.Context) (*PlayerSupervisionRequest, error) {
	result, err := psra.Edges.SupervisionRequestOrErr()
	if IsNotLoaded(err) {
		result, err = psra.QuerySupervisionRequest().Only(ctx)
	}
	return result, err
}

func (u *User) Players(ctx context.Context) (result []*Player, err error) {
	if fc := graphql.GetFieldContext(ctx); fc != nil && fc.Field.Alias != "" {
		result, err = u.NamedPlayers(graphql.GetFieldContext(ctx).Field.Alias)
	} else {
		result, err = u.Edges.PlayersOrErr()
	}
	if IsNotLoaded(err) {
		result, err = u.QueryPlayers().All(ctx)
	}
	return result, err
}

func (u *User) MainPlayer(ctx context.Context) (*Player, error) {
	result, err := u.Edges.MainPlayerOrErr()
	if IsNotLoaded(err) {
		result, err = u.QueryMainPlayer().Only(ctx)
	}
	return result, MaskNotFound(err)
}

func (u *User) GroupMemberships(ctx context.Context) (result []*GroupMembership, err error) {
	if fc := graphql.GetFieldContext(ctx); fc != nil && fc.Field.Alias != "" {
		result, err = u.NamedGroupMemberships(graphql.GetFieldContext(ctx).Field.Alias)
	} else {
		result, err = u.Edges.GroupMembershipsOrErr()
	}
	if IsNotLoaded(err) {
		result, err = u.QueryGroupMemberships().All(ctx)
	}
	return result, err
}

func (u *User) GroupMembershipApplications(ctx context.Context) (result []*GroupMembershipApplication, err error) {
	if fc := graphql.GetFieldContext(ctx); fc != nil && fc.Field.Alias != "" {
		result, err = u.NamedGroupMembershipApplications(graphql.GetFieldContext(ctx).Field.Alias)
	} else {
		result, err = u.Edges.GroupMembershipApplicationsOrErr()
	}
	if IsNotLoaded(err) {
		result, err = u.QueryGroupMembershipApplications().All(ctx)
	}
	return result, err
}

func (u *User) Games(ctx context.Context) (result []*Game, err error) {
	if fc := graphql.GetFieldContext(ctx); fc != nil && fc.Field.Alias != "" {
		result, err = u.NamedGames(graphql.GetFieldContext(ctx).Field.Alias)
	} else {
		result, err = u.Edges.GamesOrErr()
	}
	if IsNotLoaded(err) {
		result, err = u.QueryGames().All(ctx)
	}
	return result, err
}

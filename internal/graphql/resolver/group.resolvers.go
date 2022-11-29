package resolver

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"fmt"

	"github.com/open-boardgame-stats/backend/internal/auth"
	"github.com/open-boardgame-stats/backend/internal/ent"
	"github.com/open-boardgame-stats/backend/internal/ent/enums"
	"github.com/open-boardgame-stats/backend/internal/ent/group"
	"github.com/open-boardgame-stats/backend/internal/ent/groupmembership"
	"github.com/open-boardgame-stats/backend/internal/ent/groupmembershipapplication"
	"github.com/open-boardgame-stats/backend/internal/ent/groupsettings"
	"github.com/open-boardgame-stats/backend/internal/ent/schema/guidgql"
	"github.com/open-boardgame-stats/backend/internal/ent/user"
	"github.com/open-boardgame-stats/backend/internal/graphql/generated"
	"github.com/open-boardgame-stats/backend/internal/graphql/model"
)

// Role is the resolver for the role field.
func (r *groupResolver) Role(ctx context.Context, obj *ent.Group) (*enums.Role, error) {
	u, _ := auth.UserFromContext(ctx)
	if u == nil {
		return nil, nil
	}

	membership, err := r.client.GroupMembership.Query().Where(
		groupmembership.HasUserWith(user.ID(u.ID)),
		groupmembership.HasGroupWith(group.ID(obj.ID)),
	).Only(ctx)
	if err != nil && !ent.IsNotFound(err) {
		return nil, err
	}
	if membership == nil {
		return nil, nil
	}

	return &membership.Role, nil
}

// Applied is the resolver for the applied field.
func (r *groupResolver) Applied(ctx context.Context, obj *ent.Group) (*bool, error) {
	res := false
	u, _ := auth.UserFromContext(ctx)
	if u == nil {
		return &res, nil
	}

	a, err := r.client.GroupMembershipApplication.Query().Where(
		groupmembershipapplication.HasUserWith(user.ID(u.ID)),
		groupmembershipapplication.HasGroupWith(group.ID(obj.ID)),
	).Only(ctx)
	if err != nil && !ent.IsNotFound(err) {
		return nil, err
	}

	if a != nil {
		res = true
	}

	return &res, nil
}

// CreateOrUpdateGroup is the resolver for the createOrUpdateGroup field.
func (r *mutationResolver) CreateOrUpdateGroup(ctx context.Context, input model.CreateOrUpdateGroupInput) (*ent.Group, error) {
	var g *ent.Group
	var err error
	if input.ID == nil {
		g, err = createGroup(ctx, r.client, input)
	} else {
		g, err = updateGroup(ctx, r.client, input)
	}

	if err != nil {
		return nil, err
	}

	return g, nil
}

// JoinGroup is the resolver for the joinGroup field.
func (r *mutationResolver) JoinGroup(ctx context.Context, groupID guidgql.GUID) (bool, error) {
	u, err := auth.UserFromContext(ctx)
	if err != nil {
		return false, err
	}

	g, err := r.client.Group.Query().Where(group.ID(groupID)).Only(ctx)
	if err != nil {
		return false, err
	}

	// check if the user is already a member
	membership, err := r.client.GroupMembership.Query().Where(
		groupmembership.HasUserWith(user.ID(u.ID)),
		groupmembership.HasGroupWith(group.ID(g.ID)),
	).Only(ctx)
	if err != nil && !ent.IsNotFound(err) {
		return false, err
	}
	if membership != nil {
		return false, fmt.Errorf("user is already a member")
	}

	// check group's join policy
	s, err := g.QuerySettings().Where(
		groupsettings.HasGroupWith(group.ID(g.ID)),
	).Only(ctx)
	if err != nil {
		return false, err
	}
	if s.JoinPolicy != groupsettings.JoinPolicyOpen {
		return false, fmt.Errorf("group is not open for joining")
	}

	// create the membership
	_, err = r.client.GroupMembership.Create().
		SetGroup(g).
		SetUser(u).
		SetRole(enums.RoleMember).
		Save(ctx)

	return err == nil, err
}

// ApplyToGroup is the resolver for the applyToGroup field.
func (r *mutationResolver) ApplyToGroup(ctx context.Context, input model.GroupApplicationInput) (*ent.GroupMembershipApplication, error) {
	u, err := auth.UserFromContext(ctx)
	if err != nil {
		return nil, err
	}

	g, err := r.client.Group.Query().Where(group.ID(input.GroupID)).Only(ctx)
	if err != nil {
		return nil, err
	}

	// check if there is already an application
	application, err := r.client.GroupMembershipApplication.Query().Where(
		groupmembershipapplication.HasUserWith(user.ID(u.ID)),
		groupmembershipapplication.HasGroupWith(group.ID(g.ID)),
	).Only(ctx)
	if err != nil && !ent.IsNotFound(err) {
		return nil, err
	}
	if application != nil {
		return nil, fmt.Errorf("user already applied to this group")
	}

	// check if the user is already a member
	membership, err := r.client.GroupMembership.Query().Where(
		groupmembership.HasUserWith(user.ID(u.ID)),
		groupmembership.HasGroupWith(group.ID(g.ID)),
	).Only(ctx)
	if err != nil && !ent.IsNotFound(err) {
		return nil, err
	}
	if membership != nil {
		return nil, fmt.Errorf("user is already a member")
	}

	// check group's join policy
	s, err := g.QuerySettings().Where(
		groupsettings.HasGroupWith(group.ID(g.ID)),
	).Only(ctx)
	if err != nil {
		return nil, err
	}
	if s.JoinPolicy != groupsettings.JoinPolicyApplicationOnly && s.JoinPolicy != groupsettings.JoinPolicyInviteOrApplication {
		return nil, fmt.Errorf("group is not open for applications")
	}

	// create the application
	a, err := r.client.GroupMembershipApplication.Create().
		SetMessage(*input.Message).
		SetUser(u).
		SetGroup(g).
		Save(ctx)

	return a, err
}

// ResolveGroupMembershipApplication is the resolver for the resolveGroupMembershipApplication field.
func (r *mutationResolver) ResolveGroupMembershipApplication(ctx context.Context, applicationID guidgql.GUID, accepted bool) (bool, error) {
	u, err := auth.UserFromContext(ctx)
	if err != nil {
		return false, err
	}

	a, err := r.client.GroupMembershipApplication.Query().
		Where(groupmembershipapplication.ID(applicationID)).
		WithGroup().
		WithUser().
		Only(ctx)
	if err != nil {
		return false, err
	}

	// check if the user is a group admin or owner
	membership, err := r.client.GroupMembership.Query().
		Where(
			groupmembership.HasUserWith(user.ID(u.ID)),
			groupmembership.HasGroupWith(group.ID(a.Edges.Group.ID)),
		).
		Only(ctx)
	if err != nil {
		return false, err
	}
	if membership.Role != enums.RoleAdmin && membership.Role != enums.RoleOwner {
		return false, fmt.Errorf("you don't have permission to resolve this application")
	}

	// remove the application
	err = r.client.GroupMembershipApplication.DeleteOne(a).Exec(ctx)
	if err != nil {
		return false, err
	}

	// if the application was accepted, create the membership
	if accepted {
		_, err = r.client.GroupMembership.Create().
			SetGroup(a.Edges.Group).
			SetUser(a.Edges.User).
			SetRole(enums.RoleMember).
			Save(ctx)
		if err != nil {
			return false, err
		}
	}

	return true, nil
}

// ChangeUserGroupMembershipRole is the resolver for the changeUserGroupMembershipRole field.
func (r *mutationResolver) ChangeUserGroupMembershipRole(ctx context.Context, groupID guidgql.GUID, userID guidgql.GUID, role enums.Role) (bool, error) {
	u, _ := auth.UserFromContext(ctx)

	// check if the current user is a group admin or owner
	m, err := r.client.GroupMembership.Query().
		Where(
			groupmembership.HasUserWith(user.ID(u.ID)),
			groupmembership.HasGroupWith(group.ID(groupID)),
		).
		Only(ctx)
	if err != nil {
		return false, err
	}
	if m.Role != enums.RoleAdmin && m.Role != enums.RoleOwner {
		return false, fmt.Errorf("you don't have permission to change user roles")
	}

	// get the target user's membership
	m, err = r.client.GroupMembership.Query().
		Where(
			groupmembership.HasUserWith(user.ID(userID)),
			groupmembership.HasGroupWith(group.ID(groupID)),
		).
		Only(ctx)
	if err != nil {
		return false, err
	}

	// update the role
	err = r.client.GroupMembership.UpdateOne(m).
		SetRole(role).
		Exec(ctx)

	return err == nil, err
}

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

type mutationResolver struct{ *Resolver }

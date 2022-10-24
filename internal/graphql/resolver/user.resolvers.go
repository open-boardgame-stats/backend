package resolver

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"

	"github.com/open-boardgame-stats/backend/internal/auth"
	"github.com/open-boardgame-stats/backend/internal/ent"
	"github.com/open-boardgame-stats/backend/internal/ent/playersupervisionrequest"
	"github.com/open-boardgame-stats/backend/internal/ent/playersupervisionrequestapproval"
	"github.com/open-boardgame-stats/backend/internal/ent/schema/guidgql"
	"github.com/open-boardgame-stats/backend/internal/ent/user"
)

// UpdateUser is the resolver for the updateUser field.
func (r *mutationResolver) UpdateUser(ctx context.Context, id guidgql.GUID, input ent.UpdateUserInput) (*ent.User, error) {
	return ent.FromContext(ctx).User.UpdateOneID(id).SetInput(input).Save(ctx)
}

// Me is the resolver for the me field.
func (r *queryResolver) Me(ctx context.Context) (*ent.User, error) {
	u, _ := auth.UserFromContext(ctx)
	return u, nil
}

// SentSupervisionRequests is the resolver for the sentSupervisionRequests field.
func (r *userResolver) SentSupervisionRequests(ctx context.Context, obj *ent.User) ([]*ent.PlayerSupervisionRequest, error) {
	u, err := auth.UserFromContext(ctx)
	if err != nil {
		return nil, err
	}

	return r.client.User.QuerySentSupervisionRequests(u).All(ctx)
}

// ReceivedSupervisionRequests is the resolver for the receivedSupervisionRequests field.
func (r *userResolver) ReceivedSupervisionRequests(ctx context.Context, obj *ent.User) ([]*ent.PlayerSupervisionRequest, error) {
	u, err := auth.UserFromContext(ctx)
	if err != nil {
		return nil, err
	}

	return r.client.PlayerSupervisionRequest.Query().Where(
		playersupervisionrequest.HasApprovalsWith(
			playersupervisionrequestapproval.HasApproverWith(
				user.ID(u.ID),
			),
		),
	).All(ctx)
}

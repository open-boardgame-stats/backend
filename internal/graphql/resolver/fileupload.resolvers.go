package resolver

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"

	"github.com/open-boardgame-stats/backend/internal/graphql/model"
)

// PreSignUploadURL is the resolver for the preSignUploadURL field.
func (r *queryResolver) PreSignUploadURL(ctx context.Context) (*model.UploadURL, error) {
	return r.filestorage.SignUploadURL(ctx)
}

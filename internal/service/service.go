package service

import (
	"context"

	"github.com/spf13/cast"

	"github.com/google/wire"
)

// ProviderSet is repo providers.
var ProviderSet = wire.NewSet(
	NewUserServiceServer,
	NewRelationServiceServer,
	NewPostServiceServer,
	NewCommentServiceServer,
	NewLikeServiceServer,
)

// GetCurrentUserID get current user's id
func GetCurrentUserID(c context.Context) int64 {
	if c == nil {
		return 0
	}

	return cast.ToInt64(c.Value("uid"))
}

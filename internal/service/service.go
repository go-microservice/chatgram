package service

import "github.com/google/wire"

// ProviderSet is repo providers.
var ProviderSet = wire.NewSet(
	NewUserServiceServer,
	NewRelationServiceServer,
	NewPostServiceServer,
	NewCommentServiceServer,
)

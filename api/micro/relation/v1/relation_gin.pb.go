// Code generated protoc-gen-go-gin. DO NOT EDIT.
// protoc-gen-go-gin 0.0.6

package v1

import (
	context "context"
	gin "github.com/gin-gonic/gin"
	app "github.com/go-eagle/eagle/pkg/app"
	errcode "github.com/go-eagle/eagle/pkg/errcode"
	metadata "google.golang.org/grpc/metadata"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the eagle package it is being compiled against.

// context.
// metadata.
// gin.app.errcode.

type RelationServiceHTTPServer interface {
	Follow(context.Context, *FollowRequest) (*FollowReply, error)
	GetFollowerUserList(context.Context, *GetFollowerUserListRequest) (*GetFollowerUserListReply, error)
	GetFollowingUserList(context.Context, *GetFollowingUserListRequest) (*GetFollowingUserListReply, error)
	Unfollow(context.Context, *UnfollowRequest) (*UnfollowReply, error)
}

func RegisterRelationServiceHTTPServer(r gin.IRouter, srv RelationServiceHTTPServer) {
	s := RelationService{
		server: srv,
		router: r,
	}
	s.RegisterService()
}

type RelationService struct {
	server RelationServiceHTTPServer
	router gin.IRouter
}

func (s *RelationService) Follow_0(ctx *gin.Context) {
	var in FollowRequest

	if err := ctx.ShouldBindJSON(&in); err != nil {
		app.Error(ctx, errcode.ErrInvalidParam.WithDetails(err.Error()))
		return
	}

	md := metadata.New(nil)
	for k, v := range ctx.Request.Header {
		md.Set(k, v...)
	}
	newCtx := metadata.NewIncomingContext(ctx, md)
	out, err := s.server.(RelationServiceHTTPServer).Follow(newCtx, &in)
	if err != nil {
		app.Error(ctx, err)
		return
	}

	app.Success(ctx, out)
}

func (s *RelationService) Unfollow_0(ctx *gin.Context) {
	var in UnfollowRequest

	if err := ctx.ShouldBindJSON(&in); err != nil {
		app.Error(ctx, errcode.ErrInvalidParam.WithDetails(err.Error()))
		return
	}

	md := metadata.New(nil)
	for k, v := range ctx.Request.Header {
		md.Set(k, v...)
	}
	newCtx := metadata.NewIncomingContext(ctx, md)
	out, err := s.server.(RelationServiceHTTPServer).Unfollow(newCtx, &in)
	if err != nil {
		app.Error(ctx, err)
		return
	}

	app.Success(ctx, out)
}

func (s *RelationService) GetFollowingUserList_0(ctx *gin.Context) {
	var in GetFollowingUserListRequest

	if err := ctx.ShouldBindUri(&in); err != nil {
		app.Error(ctx, errcode.ErrInvalidParam.WithDetails(err.Error()))
		return
	}

	md := metadata.New(nil)
	for k, v := range ctx.Request.Header {
		md.Set(k, v...)
	}
	newCtx := metadata.NewIncomingContext(ctx, md)
	out, err := s.server.(RelationServiceHTTPServer).GetFollowingUserList(newCtx, &in)
	if err != nil {
		app.Error(ctx, err)
		return
	}

	app.Success(ctx, out)
}

func (s *RelationService) GetFollowerUserList_0(ctx *gin.Context) {
	var in GetFollowerUserListRequest

	if err := ctx.ShouldBindUri(&in); err != nil {
		app.Error(ctx, errcode.ErrInvalidParam.WithDetails(err.Error()))
		return
	}

	md := metadata.New(nil)
	for k, v := range ctx.Request.Header {
		md.Set(k, v...)
	}
	newCtx := metadata.NewIncomingContext(ctx, md)
	out, err := s.server.(RelationServiceHTTPServer).GetFollowerUserList(newCtx, &in)
	if err != nil {
		app.Error(ctx, err)
		return
	}

	app.Success(ctx, out)
}

func (s *RelationService) RegisterService() {
	s.router.Handle("POST", "/v1/relation/follow", s.Follow_0)
	s.router.Handle("POST", "/v1/relation/unfollow", s.Unfollow_0)
	s.router.Handle("GET", "/v1/relation/:id/following", s.GetFollowingUserList_0)
	s.router.Handle("GET", "/v1/relation/:id/follower", s.GetFollowerUserList_0)
}

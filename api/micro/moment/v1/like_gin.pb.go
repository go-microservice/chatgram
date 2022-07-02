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

type LikeServiceHTTPServer interface {
	CreateCommentLike(context.Context, *CreateCommentLikeRequest) (*CreateCommentLikeReply, error)
	CreatePostLike(context.Context, *CreatePostLikeRequest) (*CreatePostLikeReply, error)
	DeleteCommentLike(context.Context, *DeleteCommentLikeRequest) (*DeleteCommentLikeReply, error)
	DeletePostLike(context.Context, *DeletePostLikeRequest) (*DeletePostLikeReply, error)
	ListCommentLike(context.Context, *ListCommentLikeRequest) (*ListCommentLikeReply, error)
	ListPostLike(context.Context, *ListPostLikeRequest) (*ListPostLikeReply, error)
}

func RegisterLikeServiceHTTPServer(r gin.IRouter, srv LikeServiceHTTPServer) {
	s := LikeService{
		server: srv,
		router: r,
	}
	s.RegisterService()
}

type LikeService struct {
	server LikeServiceHTTPServer
	router gin.IRouter
}

func (s *LikeService) CreatePostLike_0(ctx *gin.Context) {
	var in CreatePostLikeRequest

	if err := ctx.ShouldBindJSON(&in); err != nil {
		app.Error(ctx, errcode.ErrInvalidParam.WithDetails(err.Error()))
		return
	}

	md := metadata.New(nil)
	for k, v := range ctx.Request.Header {
		md.Set(k, v...)
	}
	newCtx := metadata.NewIncomingContext(ctx, md)
	out, err := s.server.(LikeServiceHTTPServer).CreatePostLike(newCtx, &in)
	if err != nil {
		app.Error(ctx, err)
		return
	}

	app.Success(ctx, out)
}

func (s *LikeService) DeletePostLike_0(ctx *gin.Context) {
	var in DeletePostLikeRequest

	if err := ctx.ShouldBindJSON(&in); err != nil {
		app.Error(ctx, errcode.ErrInvalidParam.WithDetails(err.Error()))
		return
	}

	md := metadata.New(nil)
	for k, v := range ctx.Request.Header {
		md.Set(k, v...)
	}
	newCtx := metadata.NewIncomingContext(ctx, md)
	out, err := s.server.(LikeServiceHTTPServer).DeletePostLike(newCtx, &in)
	if err != nil {
		app.Error(ctx, err)
		return
	}

	app.Success(ctx, out)
}

func (s *LikeService) ListPostLike_0(ctx *gin.Context) {
	var in ListPostLikeRequest

	if err := ctx.ShouldBindQuery(&in); err != nil {
		app.Error(ctx, errcode.ErrInvalidParam.WithDetails(err.Error()))
		return
	}

	md := metadata.New(nil)
	for k, v := range ctx.Request.Header {
		md.Set(k, v...)
	}
	newCtx := metadata.NewIncomingContext(ctx, md)
	out, err := s.server.(LikeServiceHTTPServer).ListPostLike(newCtx, &in)
	if err != nil {
		app.Error(ctx, err)
		return
	}

	app.Success(ctx, out)
}

func (s *LikeService) CreateCommentLike_0(ctx *gin.Context) {
	var in CreateCommentLikeRequest

	if err := ctx.ShouldBindJSON(&in); err != nil {
		app.Error(ctx, errcode.ErrInvalidParam.WithDetails(err.Error()))
		return
	}

	md := metadata.New(nil)
	for k, v := range ctx.Request.Header {
		md.Set(k, v...)
	}
	newCtx := metadata.NewIncomingContext(ctx, md)
	out, err := s.server.(LikeServiceHTTPServer).CreateCommentLike(newCtx, &in)
	if err != nil {
		app.Error(ctx, err)
		return
	}

	app.Success(ctx, out)
}

func (s *LikeService) DeleteCommentLike_0(ctx *gin.Context) {
	var in DeleteCommentLikeRequest

	if err := ctx.ShouldBindJSON(&in); err != nil {
		app.Error(ctx, errcode.ErrInvalidParam.WithDetails(err.Error()))
		return
	}

	md := metadata.New(nil)
	for k, v := range ctx.Request.Header {
		md.Set(k, v...)
	}
	newCtx := metadata.NewIncomingContext(ctx, md)
	out, err := s.server.(LikeServiceHTTPServer).DeleteCommentLike(newCtx, &in)
	if err != nil {
		app.Error(ctx, err)
		return
	}

	app.Success(ctx, out)
}

func (s *LikeService) ListCommentLike_0(ctx *gin.Context) {
	var in ListCommentLikeRequest

	if err := ctx.ShouldBindQuery(&in); err != nil {
		app.Error(ctx, errcode.ErrInvalidParam.WithDetails(err.Error()))
		return
	}

	md := metadata.New(nil)
	for k, v := range ctx.Request.Header {
		md.Set(k, v...)
	}
	newCtx := metadata.NewIncomingContext(ctx, md)
	out, err := s.server.(LikeServiceHTTPServer).ListCommentLike(newCtx, &in)
	if err != nil {
		app.Error(ctx, err)
		return
	}

	app.Success(ctx, out)
}

func (s *LikeService) RegisterService() {
	s.router.Handle("POST", "/v1/posts/like", s.CreatePostLike_0)
	s.router.Handle("POST", "/v1/posts/dislike", s.DeletePostLike_0)
	s.router.Handle("GET", "/v1/posts/likes", s.ListPostLike_0)
	s.router.Handle("POST", "/v1/comments/like", s.CreateCommentLike_0)
	s.router.Handle("POST", "/v1/comments/dislike", s.DeleteCommentLike_0)
	s.router.Handle("GET", "/v1/comments/likes", s.ListCommentLike_0)
}

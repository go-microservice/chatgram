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
	CreateLike(context.Context, *CreateLikeRequest) (*CreateLikeReply, error)
	DeleteLike(context.Context, *DeleteLikeRequest) (*DeleteLikeReply, error)
	ListLike(context.Context, *ListLikeRequest) (*ListLikeReply, error)
	UpdateLike(context.Context, *UpdateLikeRequest) (*UpdateLikeReply, error)
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

func (s *LikeService) CreateLike_0(ctx *gin.Context) {
	var in CreateLikeRequest

	if err := ctx.ShouldBindJSON(&in); err != nil {
		app.Error(ctx, errcode.ErrInvalidParam.WithDetails(err.Error()))
		return
	}

	md := metadata.New(nil)
	for k, v := range ctx.Request.Header {
		md.Set(k, v...)
	}
	newCtx := metadata.NewIncomingContext(ctx, md)
	out, err := s.server.(LikeServiceHTTPServer).CreateLike(newCtx, &in)
	if err != nil {
		app.Error(ctx, err)
		return
	}

	app.Success(ctx, out)
}

func (s *LikeService) UpdateLike_0(ctx *gin.Context) {
	var in UpdateLikeRequest

	if err := ctx.ShouldBindJSON(&in); err != nil {
		app.Error(ctx, errcode.ErrInvalidParam.WithDetails(err.Error()))
		return
	}

	md := metadata.New(nil)
	for k, v := range ctx.Request.Header {
		md.Set(k, v...)
	}
	newCtx := metadata.NewIncomingContext(ctx, md)
	out, err := s.server.(LikeServiceHTTPServer).UpdateLike(newCtx, &in)
	if err != nil {
		app.Error(ctx, err)
		return
	}

	app.Success(ctx, out)
}

func (s *LikeService) DeleteLike_0(ctx *gin.Context) {
	var in DeleteLikeRequest

	if err := ctx.ShouldBindJSON(&in); err != nil {
		app.Error(ctx, errcode.ErrInvalidParam.WithDetails(err.Error()))
		return
	}

	md := metadata.New(nil)
	for k, v := range ctx.Request.Header {
		md.Set(k, v...)
	}
	newCtx := metadata.NewIncomingContext(ctx, md)
	out, err := s.server.(LikeServiceHTTPServer).DeleteLike(newCtx, &in)
	if err != nil {
		app.Error(ctx, err)
		return
	}

	app.Success(ctx, out)
}

func (s *LikeService) ListLike_0(ctx *gin.Context) {
	var in ListLikeRequest

	if err := ctx.ShouldBindQuery(&in); err != nil {
		app.Error(ctx, errcode.ErrInvalidParam.WithDetails(err.Error()))
		return
	}

	md := metadata.New(nil)
	for k, v := range ctx.Request.Header {
		md.Set(k, v...)
	}
	newCtx := metadata.NewIncomingContext(ctx, md)
	out, err := s.server.(LikeServiceHTTPServer).ListLike(newCtx, &in)
	if err != nil {
		app.Error(ctx, err)
		return
	}

	app.Success(ctx, out)
}

func (s *LikeService) RegisterService() {
	s.router.Handle("POST", "/v1/likes", s.CreateLike_0)
	s.router.Handle("PATCH", "/v1/likes", s.UpdateLike_0)
	s.router.Handle("DELETE", "/v1/likes", s.DeleteLike_0)
	s.router.Handle("GET", "/v1/likes", s.ListLike_0)
}

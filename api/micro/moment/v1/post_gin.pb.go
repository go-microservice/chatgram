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

type PostServiceHTTPServer interface {
	CreatePost(context.Context, *CreatePostRequest) (*CreatePostReply, error)
	DeletePost(context.Context, *DeletePostRequest) (*DeletePostReply, error)
	GetPost(context.Context, *GetPostRequest) (*GetPostReply, error)
	ListPost(context.Context, *ListPostRequest) (*ListPostReply, error)
	UpdatePost(context.Context, *UpdatePostRequest) (*UpdatePostReply, error)
}

func RegisterPostServiceHTTPServer(r gin.IRouter, srv PostServiceHTTPServer) {
	s := PostService{
		server: srv,
		router: r,
	}
	s.RegisterService()
}

type PostService struct {
	server PostServiceHTTPServer
	router gin.IRouter
}

func (s *PostService) CreatePost_0(ctx *gin.Context) {
	var in CreatePostRequest

	if err := ctx.ShouldBindJSON(&in); err != nil {
		app.Error(ctx, errcode.ErrInvalidParam.WithDetails(err.Error()))
		return
	}

	md := metadata.New(nil)
	for k, v := range ctx.Request.Header {
		md.Set(k, v...)
	}
	newCtx := metadata.NewIncomingContext(ctx, md)
	out, err := s.server.(PostServiceHTTPServer).CreatePost(newCtx, &in)
	if err != nil {
		app.Error(ctx, err)
		return
	}

	app.Success(ctx, out)
}

func (s *PostService) UpdatePost_0(ctx *gin.Context) {
	var in UpdatePostRequest

	if err := ctx.ShouldBindJSON(&in); err != nil {
		app.Error(ctx, errcode.ErrInvalidParam.WithDetails(err.Error()))
		return
	}

	md := metadata.New(nil)
	for k, v := range ctx.Request.Header {
		md.Set(k, v...)
	}
	newCtx := metadata.NewIncomingContext(ctx, md)
	out, err := s.server.(PostServiceHTTPServer).UpdatePost(newCtx, &in)
	if err != nil {
		app.Error(ctx, err)
		return
	}

	app.Success(ctx, out)
}

func (s *PostService) DeletePost_0(ctx *gin.Context) {
	var in DeletePostRequest

	if err := ctx.ShouldBindJSON(&in); err != nil {
		app.Error(ctx, errcode.ErrInvalidParam.WithDetails(err.Error()))
		return
	}

	md := metadata.New(nil)
	for k, v := range ctx.Request.Header {
		md.Set(k, v...)
	}
	newCtx := metadata.NewIncomingContext(ctx, md)
	out, err := s.server.(PostServiceHTTPServer).DeletePost(newCtx, &in)
	if err != nil {
		app.Error(ctx, err)
		return
	}

	app.Success(ctx, out)
}

func (s *PostService) GetPost_0(ctx *gin.Context) {
	var in GetPostRequest

	if err := ctx.ShouldBindUri(&in); err != nil {
		app.Error(ctx, errcode.ErrInvalidParam.WithDetails(err.Error()))
		return
	}

	md := metadata.New(nil)
	for k, v := range ctx.Request.Header {
		md.Set(k, v...)
	}
	newCtx := metadata.NewIncomingContext(ctx, md)
	out, err := s.server.(PostServiceHTTPServer).GetPost(newCtx, &in)
	if err != nil {
		app.Error(ctx, err)
		return
	}

	app.Success(ctx, out)
}

func (s *PostService) ListPost_0(ctx *gin.Context) {
	var in ListPostRequest

	if err := ctx.ShouldBindQuery(&in); err != nil {
		app.Error(ctx, errcode.ErrInvalidParam.WithDetails(err.Error()))
		return
	}

	md := metadata.New(nil)
	for k, v := range ctx.Request.Header {
		md.Set(k, v...)
	}
	newCtx := metadata.NewIncomingContext(ctx, md)
	out, err := s.server.(PostServiceHTTPServer).ListPost(newCtx, &in)
	if err != nil {
		app.Error(ctx, err)
		return
	}

	app.Success(ctx, out)
}

func (s *PostService) RegisterService() {
	s.router.Handle("POST", "/v1/posts", s.CreatePost_0)
	s.router.Handle("PATCH", "/v1/posts", s.UpdatePost_0)
	s.router.Handle("DELETE", "/v1/posts", s.DeletePost_0)
	s.router.Handle("GET", "/v1/posts/:id", s.GetPost_0)
	s.router.Handle("GET", "/v1/posts", s.ListPost_0)
}

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

type CommentServiceHTTPServer interface {
	CreateComment(context.Context, *CreateCommentRequest) (*CreateCommentReply, error)
	DeleteComment(context.Context, *DeleteCommentRequest) (*DeleteCommentReply, error)
	GetComment(context.Context, *GetCommentRequest) (*GetCommentReply, error)
	ListComment(context.Context, *ListCommentRequest) (*ListCommentReply, error)
	UpdateComment(context.Context, *UpdateCommentRequest) (*UpdateCommentReply, error)
}

func RegisterCommentServiceHTTPServer(r gin.IRouter, srv CommentServiceHTTPServer) {
	s := CommentService{
		server: srv,
		router: r,
	}
	s.RegisterService()
}

type CommentService struct {
	server CommentServiceHTTPServer
	router gin.IRouter
}

func (s *CommentService) CreateComment_0(ctx *gin.Context) {
	var in CreateCommentRequest

	if err := ctx.ShouldBindJSON(&in); err != nil {
		app.Error(ctx, errcode.ErrInvalidParam.WithDetails(err.Error()))
		return
	}

	md := metadata.New(nil)
	for k, v := range ctx.Request.Header {
		md.Set(k, v...)
	}
	newCtx := metadata.NewIncomingContext(ctx, md)
	out, err := s.server.(CommentServiceHTTPServer).CreateComment(newCtx, &in)
	if err != nil {
		app.Error(ctx, err)
		return
	}

	app.Success(ctx, out)
}

func (s *CommentService) UpdateComment_0(ctx *gin.Context) {
	var in UpdateCommentRequest

	if err := ctx.ShouldBindJSON(&in); err != nil {
		app.Error(ctx, errcode.ErrInvalidParam.WithDetails(err.Error()))
		return
	}

	md := metadata.New(nil)
	for k, v := range ctx.Request.Header {
		md.Set(k, v...)
	}
	newCtx := metadata.NewIncomingContext(ctx, md)
	out, err := s.server.(CommentServiceHTTPServer).UpdateComment(newCtx, &in)
	if err != nil {
		app.Error(ctx, err)
		return
	}

	app.Success(ctx, out)
}

func (s *CommentService) DeleteComment_0(ctx *gin.Context) {
	var in DeleteCommentRequest

	if err := ctx.ShouldBindJSON(&in); err != nil {
		app.Error(ctx, errcode.ErrInvalidParam.WithDetails(err.Error()))
		return
	}

	md := metadata.New(nil)
	for k, v := range ctx.Request.Header {
		md.Set(k, v...)
	}
	newCtx := metadata.NewIncomingContext(ctx, md)
	out, err := s.server.(CommentServiceHTTPServer).DeleteComment(newCtx, &in)
	if err != nil {
		app.Error(ctx, err)
		return
	}

	app.Success(ctx, out)
}

func (s *CommentService) GetComment_0(ctx *gin.Context) {
	var in GetCommentRequest

	if err := ctx.ShouldBindUri(&in); err != nil {
		app.Error(ctx, errcode.ErrInvalidParam.WithDetails(err.Error()))
		return
	}

	md := metadata.New(nil)
	for k, v := range ctx.Request.Header {
		md.Set(k, v...)
	}
	newCtx := metadata.NewIncomingContext(ctx, md)
	out, err := s.server.(CommentServiceHTTPServer).GetComment(newCtx, &in)
	if err != nil {
		app.Error(ctx, err)
		return
	}

	app.Success(ctx, out)
}

func (s *CommentService) ListComment_0(ctx *gin.Context) {
	var in ListCommentRequest

	if err := ctx.ShouldBindQuery(&in); err != nil {
		app.Error(ctx, errcode.ErrInvalidParam.WithDetails(err.Error()))
		return
	}

	md := metadata.New(nil)
	for k, v := range ctx.Request.Header {
		md.Set(k, v...)
	}
	newCtx := metadata.NewIncomingContext(ctx, md)
	out, err := s.server.(CommentServiceHTTPServer).ListComment(newCtx, &in)
	if err != nil {
		app.Error(ctx, err)
		return
	}

	app.Success(ctx, out)
}

func (s *CommentService) RegisterService() {
	s.router.Handle("POST", "/v1/comments", s.CreateComment_0)
	s.router.Handle("PATCH", "/v1/comments", s.UpdateComment_0)
	s.router.Handle("DELETE", "/v1/comments", s.DeleteComment_0)
	s.router.Handle("GET", "/v1/comments/:post_id", s.GetComment_0)
	s.router.Handle("GET", "/v1/comments", s.ListComment_0)
}

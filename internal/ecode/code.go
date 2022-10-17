package ecode

import "github.com/go-eagle/eagle/pkg/errcode"

var (
	// post
	ErrPostNotFound = errcode.NewError(20101, "Post Not found")

	// comment
	ErrCommentNotFound = errcode.NewError(20201, "Comment Not found")
)

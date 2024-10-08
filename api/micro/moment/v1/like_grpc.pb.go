// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.18.1
// source: api/micro/moment/v1/like.proto

package v1

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

// LikeServiceClient is the client API for LikeService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type LikeServiceClient interface {
	CreatePostLike(ctx context.Context, in *CreatePostLikeRequest, opts ...grpc.CallOption) (*CreatePostLikeReply, error)
	DeletePostLike(ctx context.Context, in *DeletePostLikeRequest, opts ...grpc.CallOption) (*DeletePostLikeReply, error)
	ListPostLike(ctx context.Context, in *ListPostLikeRequest, opts ...grpc.CallOption) (*ListPostLikeReply, error)
	CreateCommentLike(ctx context.Context, in *CreateCommentLikeRequest, opts ...grpc.CallOption) (*CreateCommentLikeReply, error)
	DeleteCommentLike(ctx context.Context, in *DeleteCommentLikeRequest, opts ...grpc.CallOption) (*DeleteCommentLikeReply, error)
	ListCommentLike(ctx context.Context, in *ListCommentLikeRequest, opts ...grpc.CallOption) (*ListCommentLikeReply, error)
}

type likeServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewLikeServiceClient(cc grpc.ClientConnInterface) LikeServiceClient {
	return &likeServiceClient{cc}
}

func (c *likeServiceClient) CreatePostLike(ctx context.Context, in *CreatePostLikeRequest, opts ...grpc.CallOption) (*CreatePostLikeReply, error) {
	out := new(CreatePostLikeReply)
	err := c.cc.Invoke(ctx, "/api.micro.moment.v1.LikeService/CreatePostLike", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *likeServiceClient) DeletePostLike(ctx context.Context, in *DeletePostLikeRequest, opts ...grpc.CallOption) (*DeletePostLikeReply, error) {
	out := new(DeletePostLikeReply)
	err := c.cc.Invoke(ctx, "/api.micro.moment.v1.LikeService/DeletePostLike", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *likeServiceClient) ListPostLike(ctx context.Context, in *ListPostLikeRequest, opts ...grpc.CallOption) (*ListPostLikeReply, error) {
	out := new(ListPostLikeReply)
	err := c.cc.Invoke(ctx, "/api.micro.moment.v1.LikeService/ListPostLike", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *likeServiceClient) CreateCommentLike(ctx context.Context, in *CreateCommentLikeRequest, opts ...grpc.CallOption) (*CreateCommentLikeReply, error) {
	out := new(CreateCommentLikeReply)
	err := c.cc.Invoke(ctx, "/api.micro.moment.v1.LikeService/CreateCommentLike", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *likeServiceClient) DeleteCommentLike(ctx context.Context, in *DeleteCommentLikeRequest, opts ...grpc.CallOption) (*DeleteCommentLikeReply, error) {
	out := new(DeleteCommentLikeReply)
	err := c.cc.Invoke(ctx, "/api.micro.moment.v1.LikeService/DeleteCommentLike", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *likeServiceClient) ListCommentLike(ctx context.Context, in *ListCommentLikeRequest, opts ...grpc.CallOption) (*ListCommentLikeReply, error) {
	out := new(ListCommentLikeReply)
	err := c.cc.Invoke(ctx, "/api.micro.moment.v1.LikeService/ListCommentLike", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// LikeServiceServer is the server API for LikeService service.
// All implementations must embed UnimplementedLikeServiceServer
// for forward compatibility
type LikeServiceServer interface {
	CreatePostLike(context.Context, *CreatePostLikeRequest) (*CreatePostLikeReply, error)
	DeletePostLike(context.Context, *DeletePostLikeRequest) (*DeletePostLikeReply, error)
	ListPostLike(context.Context, *ListPostLikeRequest) (*ListPostLikeReply, error)
	CreateCommentLike(context.Context, *CreateCommentLikeRequest) (*CreateCommentLikeReply, error)
	DeleteCommentLike(context.Context, *DeleteCommentLikeRequest) (*DeleteCommentLikeReply, error)
	ListCommentLike(context.Context, *ListCommentLikeRequest) (*ListCommentLikeReply, error)
	mustEmbedUnimplementedLikeServiceServer()
}

// UnimplementedLikeServiceServer must be embedded to have forward compatible implementations.
type UnimplementedLikeServiceServer struct {
}

func (UnimplementedLikeServiceServer) CreatePostLike(context.Context, *CreatePostLikeRequest) (*CreatePostLikeReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreatePostLike not implemented")
}
func (UnimplementedLikeServiceServer) DeletePostLike(context.Context, *DeletePostLikeRequest) (*DeletePostLikeReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeletePostLike not implemented")
}
func (UnimplementedLikeServiceServer) ListPostLike(context.Context, *ListPostLikeRequest) (*ListPostLikeReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListPostLike not implemented")
}
func (UnimplementedLikeServiceServer) CreateCommentLike(context.Context, *CreateCommentLikeRequest) (*CreateCommentLikeReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateCommentLike not implemented")
}
func (UnimplementedLikeServiceServer) DeleteCommentLike(context.Context, *DeleteCommentLikeRequest) (*DeleteCommentLikeReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteCommentLike not implemented")
}
func (UnimplementedLikeServiceServer) ListCommentLike(context.Context, *ListCommentLikeRequest) (*ListCommentLikeReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListCommentLike not implemented")
}
func (UnimplementedLikeServiceServer) mustEmbedUnimplementedLikeServiceServer() {}

// UnsafeLikeServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to LikeServiceServer will
// result in compilation errors.
type UnsafeLikeServiceServer interface {
	mustEmbedUnimplementedLikeServiceServer()
}

func RegisterLikeServiceServer(s grpc.ServiceRegistrar, srv LikeServiceServer) {
	s.RegisterService(&LikeService_ServiceDesc, srv)
}

func _LikeService_CreatePostLike_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreatePostLikeRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(LikeServiceServer).CreatePostLike(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.micro.moment.v1.LikeService/CreatePostLike",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(LikeServiceServer).CreatePostLike(ctx, req.(*CreatePostLikeRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _LikeService_DeletePostLike_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeletePostLikeRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(LikeServiceServer).DeletePostLike(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.micro.moment.v1.LikeService/DeletePostLike",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(LikeServiceServer).DeletePostLike(ctx, req.(*DeletePostLikeRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _LikeService_ListPostLike_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListPostLikeRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(LikeServiceServer).ListPostLike(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.micro.moment.v1.LikeService/ListPostLike",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(LikeServiceServer).ListPostLike(ctx, req.(*ListPostLikeRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _LikeService_CreateCommentLike_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateCommentLikeRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(LikeServiceServer).CreateCommentLike(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.micro.moment.v1.LikeService/CreateCommentLike",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(LikeServiceServer).CreateCommentLike(ctx, req.(*CreateCommentLikeRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _LikeService_DeleteCommentLike_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteCommentLikeRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(LikeServiceServer).DeleteCommentLike(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.micro.moment.v1.LikeService/DeleteCommentLike",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(LikeServiceServer).DeleteCommentLike(ctx, req.(*DeleteCommentLikeRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _LikeService_ListCommentLike_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListCommentLikeRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(LikeServiceServer).ListCommentLike(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.micro.moment.v1.LikeService/ListCommentLike",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(LikeServiceServer).ListCommentLike(ctx, req.(*ListCommentLikeRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// LikeService_ServiceDesc is the grpc.ServiceDesc for LikeService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var LikeService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "api.micro.moment.v1.LikeService",
	HandlerType: (*LikeServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreatePostLike",
			Handler:    _LikeService_CreatePostLike_Handler,
		},
		{
			MethodName: "DeletePostLike",
			Handler:    _LikeService_DeletePostLike_Handler,
		},
		{
			MethodName: "ListPostLike",
			Handler:    _LikeService_ListPostLike_Handler,
		},
		{
			MethodName: "CreateCommentLike",
			Handler:    _LikeService_CreateCommentLike_Handler,
		},
		{
			MethodName: "DeleteCommentLike",
			Handler:    _LikeService_DeleteCommentLike_Handler,
		},
		{
			MethodName: "ListCommentLike",
			Handler:    _LikeService_ListCommentLike_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "api/micro/moment/v1/like.proto",
}

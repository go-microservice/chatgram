package service

import (
	"context"

	pb "github.com/go-microservice/ins-api/api/micro/relation/v1"
	relationV1 "github.com/go-microservice/relation-service/api/relation/v1"
)

var (
	_ pb.RelationServiceHTTPServer = (*RelationServiceServer)(nil)
)

type RelationServiceServer struct {
	relationRPC relationV1.RelationServiceClient
}

func NewRelationServiceServer() *RelationServiceServer {
	return &RelationServiceServer{}
}

func (s *RelationServiceServer) Follow(ctx context.Context, req *pb.FollowRequest) (*pb.FollowReply, error) {
	in := &relationV1.FollowRequest{
		UserId:      req.UserId,
		FollowedUid: req.FollowedUid,
	}
	_, err := s.relationRPC.Follow(ctx, in)
	if err != nil {
		return nil, err
	}
	return &pb.FollowReply{}, nil
}
func (s *RelationServiceServer) Unfollow(ctx context.Context, req *pb.UnfollowRequest) (*pb.UnfollowReply, error) {
	in := &relationV1.UnfollowRequest{
		UserId:      req.UserId,
		FollowedUid: req.FollowedUid,
	}
	_, err := s.relationRPC.Unfollow(ctx, in)
	if err != nil {
		return nil, err
	}
	return &pb.UnfollowReply{}, nil
}
func (s *RelationServiceServer) GetFollowingUserList(ctx context.Context, req *pb.GetFollowingUserListRequest) (*pb.GetFollowingUserListReply, error) {
	in := &relationV1.FollowingListRequest{
		UserId: req.GetUserId(),
		LastId: req.GetLastId(),
		Limit:  int32(req.GetLimit()),
	}
	_, err := s.relationRPC.GetFollowingList(ctx, in)
	if err != nil {
		return nil, err
	}
	return &pb.GetFollowingUserListReply{}, nil
}
func (s *RelationServiceServer) GetFollowerUserList(ctx context.Context, req *pb.GetFollowerUserListRequest) (*pb.GetFollowerUserListReply, error) {
	in := &relationV1.FollowerListRequest{
		UserId: req.GetUserId(),
		LastId: req.GetLastId(),
		Limit:  int32(req.GetLimit()),
	}
	_, err := s.relationRPC.GetFollowerList(ctx, in)
	if err != nil {
		return nil, err
	}
	return &pb.GetFollowerUserListReply{}, nil
}

package service

import (
	"context"

	relationV1 "github.com/go-microservice/relation-service/api/relation/v1"
	userv1 "github.com/go-microservice/user-service/api/user/v1"
	"github.com/jinzhu/copier"
	"github.com/spf13/cast"

	pb "github.com/go-microservice/ins-api/api/micro/relation/v1"
	pbuser "github.com/go-microservice/ins-api/api/micro/user/v1"
)

var (
	_ pb.RelationServiceHTTPServer = (*RelationServiceServer)(nil)
)

type RelationServiceServer struct {
	relationRPC relationV1.RelationServiceClient
	userRPC     userv1.UserServiceClient
}

func NewRelationServiceServer(repo relationV1.RelationServiceClient, userRepo userv1.UserServiceClient) *RelationServiceServer {
	return &RelationServiceServer{
		relationRPC: repo,
		userRPC:     userRepo,
	}
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
		UserId: cast.ToInt64(req.GetUserId()),
		LastId: cast.ToInt64(req.GetLastId()),
		Limit:  cast.ToInt32(req.GetLimit()),
	}
	ret, err := s.relationRPC.GetFollowingList(ctx, in)
	if err != nil {
		return nil, err
	}

	out, err := s.userRPC.BatchGetUsers(ctx, &userv1.BatchGetUsersRequest{Ids: ret.UserIds})
	if err != nil {
		return nil, err
	}

	var users []*pbuser.User
	for _, v := range out.Users {
		user := pbuser.User{}
		err = copier.Copy(&user, &v)
		if err != nil {
			continue
		}
		users = append(users, &user)
	}

	return &pb.GetFollowingUserListReply{
		Users: users,
	}, nil
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

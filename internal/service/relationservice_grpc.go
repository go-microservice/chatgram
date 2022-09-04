package service

import (
	"context"

	"github.com/go-eagle/eagle/pkg/errcode"

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
	// check user is exist
	user, err := s.userRPC.GetUser(ctx, &userv1.GetUserRequest{Id: req.GetUserId()})
	if err != nil {
		return nil, errcode.ErrInvalidParam
	}
	if user == nil {
		return nil, errcode.ErrInvalidParam
	}
	user, err = s.userRPC.GetUser(ctx, &userv1.GetUserRequest{Id: req.GetFollowedUid()})
	if err != nil {
		return nil, errcode.ErrInvalidParam
	}
	if user == nil {
		return nil, errcode.ErrInvalidParam
	}

	// exec follow logic
	in := &relationV1.FollowRequest{
		UserId:      req.UserId,
		FollowedUid: req.FollowedUid,
	}
	_, err = s.relationRPC.Follow(ctx, in)
	if err != nil {
		return nil, err
	}
	return &pb.FollowReply{}, nil
}

func (s *RelationServiceServer) Unfollow(ctx context.Context, req *pb.UnfollowRequest) (*pb.UnfollowReply, error) {
	// check user is exist
	user, err := s.userRPC.GetUser(ctx, &userv1.GetUserRequest{Id: req.GetUserId()})
	if err != nil {
		return nil, errcode.ErrInvalidParam
	}
	if user == nil {
		return nil, errcode.ErrInvalidParam
	}
	user, err = s.userRPC.GetUser(ctx, &userv1.GetUserRequest{Id: req.GetFollowedUid()})
	if err != nil {
		return nil, errcode.ErrInvalidParam
	}
	if user == nil {
		return nil, errcode.ErrInvalidParam
	}

	in := &relationV1.UnfollowRequest{
		UserId:      req.UserId,
		FollowedUid: req.FollowedUid,
	}
	_, err = s.relationRPC.Unfollow(ctx, in)
	if err != nil {
		return nil, err
	}
	return &pb.UnfollowReply{}, nil
}

func (s *RelationServiceServer) GetFollowingUserList(ctx context.Context, req *pb.GetFollowingUserListRequest) (*pb.GetFollowingUserListReply, error) {
	if req.GetId() == "" {
		return nil, errcode.ErrInvalidParam
	}
	// get relation data, support pagination
	limit := cast.ToInt32(req.GetLimit())
	if limit == 0 {
		limit = 10
	}

	in := &relationV1.FollowingListRequest{
		UserId: cast.ToInt64(req.GetId()),
		LastId: cast.ToInt64(req.GetLastId()),
		Limit:  limit + 1,
	}
	ret, err := s.relationRPC.GetFollowingList(ctx, in)
	if err != nil {
		return nil, err
	}

	relationRet := ret.GetResult()
	var (
		hasMore int32
		lastId  string
		userIDs []int64
	)
	if int32(len(relationRet)) > limit {
		hasMore = 1
		lastId = cast.ToString(relationRet[len(relationRet)-1].Id)
		relationRet = relationRet[0 : len(relationRet)-1]
	}
	for _, v := range relationRet {
		userIDs = append(userIDs, v.GetFollowedUid())
	}

	// get user info
	out, err := s.userRPC.BatchGetUsers(ctx, &userv1.BatchGetUsersRequest{Ids: userIDs})
	if err != nil {
		return nil, err
	}

	// convert to pb user
	var users []*pbuser.User
	for _, v := range out.GetUsers() {
		user := pbuser.User{}
		err = copier.Copy(&user, &v)
		if err != nil {
			continue
		}
		users = append(users, &user)
	}

	return &pb.GetFollowingUserListReply{
		HasMore: hasMore,
		LastId:  lastId,
		Items:   users,
	}, nil
}
func (s *RelationServiceServer) GetFollowerUserList(ctx context.Context, req *pb.GetFollowerUserListRequest) (*pb.GetFollowerUserListReply, error) {
	if req.GetUserId() == "" {
		return nil, errcode.ErrInvalidParam
	}
	// get relation data, support pagination
	limit := cast.ToInt32(req.GetLimit())
	if limit == 0 {
		limit = 10
	}
	in := &relationV1.FollowerListRequest{
		UserId: cast.ToInt64(req.GetUserId()),
		LastId: cast.ToInt64(req.GetLastId()),
		Limit:  limit + 1,
	}
	ret, err := s.relationRPC.GetFollowerList(ctx, in)
	if err != nil {
		return nil, err
	}

	relationRet := ret.GetResult()
	var (
		hasMore int32
		lastId  string
		userIDs []int64
	)
	if int32(len(relationRet)) > limit {
		hasMore = 1
		lastId = cast.ToString(relationRet[len(relationRet)-1].Id)
		relationRet = relationRet[0 : len(relationRet)-1]
	}
	for _, v := range relationRet {
		userIDs = append(userIDs, v.GetFollowerUid())
	}

	// get user info
	out, err := s.userRPC.BatchGetUsers(ctx, &userv1.BatchGetUsersRequest{Ids: userIDs})
	if err != nil {
		return nil, err
	}

	// convert to pb user
	var users []*pbuser.User
	for _, v := range out.GetUsers() {
		user := pbuser.User{}
		err = copier.Copy(&user, &v)
		if err != nil {
			continue
		}
		users = append(users, &user)
	}

	return &pb.GetFollowerUserListReply{
		HasMore: hasMore,
		LastId:  lastId,
		Items:   users,
	}, nil
}

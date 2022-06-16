package service

import (
	"context"

	userv1 "github.com/go-microservice/user-service/api/user/v1"
	"github.com/jinzhu/copier"
	"github.com/spf13/cast"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	pb "github.com/go-microservice/ins-api/api/micro/user/v1"
)

var (
	_ pb.UserServiceHTTPServer = (*UserServiceServer)(nil)
)

type UserServiceServer struct {
	userRPC userv1.UserServiceClient
}

func NewUserServiceServer(repo userv1.UserServiceClient) *UserServiceServer {
	return &UserServiceServer{
		userRPC: repo,
	}
}

func (s *UserServiceServer) Register(ctx context.Context, req *pb.RegisterRequest) (*pb.RegisterReply, error) {
	in := &userv1.RegisterRequest{
		Username: req.Username,
		Email:    req.Email,
		Password: req.Password,
	}
	out, err := s.userRPC.Register(ctx, in)
	if err != nil {
		// check client if deadline exceeded
		statusErr, ok := status.FromError(err)
		if ok && statusErr.Code() == codes.DeadlineExceeded {
			return nil, status.Error(codes.DeadlineExceeded, "deadline exceeded")
		}

		return nil, err
	}

	return &pb.RegisterReply{
		Id:       out.Id,
		Username: out.Username,
	}, nil
}

func (s *UserServiceServer) Login(ctx context.Context, req *pb.LoginRequest) (*pb.LoginReply, error) {
	in := &userv1.LoginRequest{
		Username: req.Username,
		Email:    req.Email,
		Password: req.Password,
	}
	out, err := s.userRPC.Login(ctx, in)
	if err != nil {
		// check client if deadline exceeded
		statusErr, ok := status.FromError(err)
		if ok && statusErr.Code() == codes.DeadlineExceeded {
			return nil, status.Error(codes.DeadlineExceeded, "deadline exceeded")
		}
		return nil, err
	}
	return &pb.LoginReply{
		Token: out.GetToken(),
	}, nil
}

func (s *UserServiceServer) Logout(ctx context.Context, req *pb.LogoutRequest) (*pb.LogoutReply, error) {
	in := &userv1.LogoutRequest{
		Id:    req.Id,
		Token: req.Token,
	}
	_, err := s.userRPC.Logout(ctx, in)
	if err != nil {
		// check client if deadline exceeded
		statusErr, ok := status.FromError(err)
		if ok && statusErr.Code() == codes.DeadlineExceeded {
			return nil, status.Error(codes.DeadlineExceeded, "deadline exceeded")
		}
		return nil, err
	}
	return &pb.LogoutReply{}, nil
}

func (s *UserServiceServer) GetUser(ctx context.Context, req *pb.GetUserRequest) (*pb.GetUserReply, error) {
	in := &userv1.GetUserRequest{
		Id: cast.ToInt64(req.GetId()),
	}
	out, err := s.userRPC.GetUser(ctx, in)
	if err != nil {
		// check client if deadline exceeded
		statusErr, ok := status.FromError(err)
		if ok && statusErr.Code() == codes.DeadlineExceeded {
			return nil, status.Error(codes.DeadlineExceeded, "deadline exceeded")
		}
		return nil, err
	}
	user, err := convertUser(out.User)
	if err != nil {
		return nil, err
	}
	return &pb.GetUserReply{
		User: user,
	}, nil
}

func (s *UserServiceServer) UpdateUser(ctx context.Context, req *pb.UpdateUserRequest) (*pb.UpdateUserReply, error) {
	return &pb.UpdateUserReply{}, nil
}

func (s *UserServiceServer) UpdatePassword(ctx context.Context, req *pb.UpdatePasswordRequest) (*pb.UpdatePasswordReply, error) {
	return &pb.UpdatePasswordReply{}, nil
}

func convertUser(u *userv1.User) (*pb.User, error) {
	user := pb.User{}
	err := copier.Copy(&user, &u)
	if err != nil {
		return nil, err
	}
	return &user, nil
}

package service

import (
	"context"

	"github.com/spf13/cast"

	"github.com/jinzhu/copier"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	momentv1 "github.com/go-microservice/moment-service/api/moment/v1"

	pb "github.com/go-microservice/ins-api/api/micro/moment/v1"
)

var (
	_ pb.PostServiceServer = (*PostServiceServer)(nil)
)

type PostServiceServer struct {
	pb.UnimplementedPostServiceServer

	momentRPC momentv1.PostServiceClient
}

func NewPostServiceServer(repo momentv1.PostServiceClient) *PostServiceServer {
	return &PostServiceServer{
		momentRPC: repo,
	}
}

func (s *PostServiceServer) CreatePost(ctx context.Context, req *pb.CreatePostRequest) (*pb.CreatePostReply, error) {
	in := &momentv1.CreatePostRequest{
		UserId:        req.UserId,
		Title:         req.Title,
		Text:          req.Text,
		PicKeys:       req.PicKeys,
		VideoKey:      req.VideoKey,
		VideoDuration: req.VideoDuration,
		CoverKey:      req.CoverKey,
		CoverWidth:    req.CoverWidth,
		CoverHeight:   req.CoverHeight,
		Longitude:     req.Longitude,
		Latitude:      req.Latitude,
		Position:      req.Position,
	}
	out, err := s.momentRPC.CreatePost(ctx, in)
	if err != nil {
		// check client if deadline exceeded
		statusErr, ok := status.FromError(err)
		if ok && statusErr.Code() == codes.DeadlineExceeded {
			return nil, status.Error(codes.DeadlineExceeded, "deadline exceeded")
		}

		return nil, err
	}

	post := pb.Post{}
	err = copier.Copy(&post, &out.Post)
	if err != nil {
		return nil, err
	}

	return &pb.CreatePostReply{
		Post: &post,
	}, nil
}

func (s *PostServiceServer) UpdatePost(ctx context.Context, req *pb.UpdatePostRequest) (*pb.UpdatePostReply, error) {
	return &pb.UpdatePostReply{}, nil
}
func (s *PostServiceServer) DeletePost(ctx context.Context, req *pb.DeletePostRequest) (*pb.DeletePostReply, error) {
	return &pb.DeletePostReply{}, nil
}
func (s *PostServiceServer) GetPost(ctx context.Context, req *pb.GetPostRequest) (*pb.GetPostReply, error) {
	in := &momentv1.GetPostRequest{
		Id: req.GetId(),
	}
	out, err := s.momentRPC.GetPost(ctx, in)
	if err != nil {
		// check client if deadline exceeded
		statusErr, ok := status.FromError(err)
		if ok && statusErr.Code() == codes.DeadlineExceeded {
			return nil, status.Error(codes.DeadlineExceeded, "deadline exceeded")
		}
		return nil, err
	}
	post := pb.Post{}
	err = copier.Copy(&post, &out.Post)
	if err != nil {
		return nil, err
	}
	return &pb.GetPostReply{
		Post: &post,
	}, nil
}
func (s *PostServiceServer) ListPost(ctx context.Context, req *pb.ListPostRequest) (*pb.ListPostReply, error) {
	// get data, support pagination
	limit := cast.ToInt32(req.GetLimit())
	in := &momentv1.ListHotPostRequest{
		LastId: cast.ToInt64(req.GetLastId()),
		Limit:  limit + 1,
	}
	ret, err := s.momentRPC.ListHotPost(ctx, in)
	if err != nil {
		return nil, err
	}

	out := ret.GetItems()
	var (
		hasMore bool
		lastId  string
	)
	if int32(len(out)) > limit {
		hasMore = true
		lastId = cast.ToString(out[len(out)-1].Id)
		out = out[0 : len(out)-1]
	}

	// convert to pb user
	var posts []*pb.Post
	for _, v := range ret.GetItems() {
		post := pb.Post{}
		err = copier.Copy(&post, &v)
		if err != nil {
			continue
		}
		posts = append(posts, &post)
	}
	return &pb.ListPostReply{
		HasMore: hasMore,
		LastId:  lastId,
		Items:   posts,
	}, nil
}

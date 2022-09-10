package service

import (
	"context"
	"fmt"
	"sync"
	"time"

	"github.com/go-eagle/eagle/pkg/log"
	momentv1 "github.com/go-microservice/moment-service/api/moment/v1"
	userv1 "github.com/go-microservice/user-service/api/user/v1"
	"github.com/jinzhu/copier"
	"github.com/spf13/cast"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	pb "github.com/go-microservice/ins-api/api/micro/moment/v1"
)

var (
	_ pb.PostServiceServer = (*PostServiceServer)(nil)
)

type PostServiceServer struct {
	pb.UnimplementedPostServiceServer

	momentRPC momentv1.PostServiceClient
	userRPC   userv1.UserServiceClient
}

func NewPostServiceServer(repo momentv1.PostServiceClient, userRepo userv1.UserServiceClient) *PostServiceServer {
	return &PostServiceServer{
		momentRPC: repo,
		userRPC:   userRepo,
	}
}

func (s *PostServiceServer) CreatePost(ctx context.Context, req *pb.CreatePostRequest) (*pb.CreatePostReply, error) {
	in := &momentv1.CreatePostRequest{
		UserId:        GetCurrentUserID(ctx),
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

	// ger user info
	userIn := &userv1.GetUserRequest{Id: GetCurrentUserID(ctx)}
	user, err := s.userRPC.GetUser(ctx, userIn)
	if err != nil {
		return nil, err
	}
	post.User, err = convertUser(user.GetUser())
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
	in := &momentv1.DeletePostRequest{
		Id:      req.GetId(),
		UserId:  req.GetUserId(),
		DelFlag: req.GetDelFlag(),
	}
	_, err := s.momentRPC.DeletePost(ctx, in)
	if err != nil {
		// check client if deadline exceeded
		statusErr, ok := status.FromError(err)
		if ok && statusErr.Code() == codes.DeadlineExceeded {
			return nil, status.Error(codes.DeadlineExceeded, "deadline exceeded")
		}
		return nil, err
	}
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
	post, err := convertPost(out.Post)
	if err != nil {
		return nil, err
	}

	return &pb.GetPostReply{
		Post: post,
	}, nil
}

func (s *PostServiceServer) ListHotPost(ctx context.Context, req *pb.ListPostRequest) (*pb.ListPostReply, error) {
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

	posts := ret.GetItems()
	var (
		hasMore bool
		lastId  string
	)
	if int32(len(posts)) > limit {
		hasMore = true
		lastId = cast.ToString(posts[len(posts)-1].Id)
		posts = posts[0 : len(posts)-1]
	}
	pbPosts, err := s.assembleData(ctx, posts)
	if err != nil {
		return nil, err
	}

	return &pb.ListPostReply{
		HasMore: hasMore,
		LastId:  lastId,
		Items:   pbPosts,
	}, nil
}

func (s *PostServiceServer) ListLatestPost(ctx context.Context, req *pb.ListPostRequest) (*pb.ListPostReply, error) {
	// get data, support pagination
	limit := cast.ToInt32(req.GetLimit())
	in := &momentv1.ListLatestPostRequest{
		LastId: cast.ToInt64(req.GetLastId()),
		Limit:  limit + 1,
	}
	ret, err := s.momentRPC.ListLatestPost(ctx, in)
	if err != nil {
		return nil, err
	}

	posts := ret.GetItems()
	var (
		hasMore bool
		lastId  string
	)
	if int32(len(posts)) > limit {
		hasMore = true
		lastId = cast.ToString(posts[len(posts)-1].Id)
		posts = posts[0 : len(posts)-1]
	}

	pbPosts, err := s.assembleData(ctx, posts)
	if err != nil {
		return nil, err
	}

	return &pb.ListPostReply{
		HasMore: hasMore,
		LastId:  lastId,
		Items:   pbPosts,
	}, nil
}

func (s *PostServiceServer) assembleData(ctx context.Context, posts []*momentv1.Post) ([]*pb.Post, error) {
	// batch get user data
	var (
		userIDs []int64
	)
	for _, v := range posts {
		userIDs = append(userIDs, v.UserId)
	}

	userReply, err := s.userRPC.BatchGetUsers(ctx, &userv1.BatchGetUsersRequest{Ids: userIDs})
	if err != nil {
		return nil, err
	}
	users := userReply.GetUsers()
	// to map
	userMap := make(map[int64]*userv1.User)
	for _, v := range users {
		userMap[v.Id] = v
	}

	var (
		pbPosts []*pb.Post
		m       sync.Map
		mu      sync.Mutex
	)

	wg := sync.WaitGroup{}
	errChan := make(chan error, 1)
	finished := make(chan bool, 1)

	go func() {
		select {
		case <-finished:
			return
		case err := <-errChan:
			if err != nil {
				// NOTE: if need, record log to file
			}
		case <-time.After(3 * time.Second):
			log.Warn(fmt.Errorf("list users timeout after 3 seconds"))
			return
		}
	}()

	for _, post := range posts {
		wg.Add(1)
		post := post
		go func(info *momentv1.Post) {
			defer func() {
				wg.Done()
			}()

			mu.Lock()
			defer mu.Unlock()

			pbPost, err := convertPost(info)
			if err != nil {
				return
			}
			// user
			user, ok := userMap[post.UserId]
			if !ok {
				return
			}
			pbPost.User, err = convertUser(user)
			if err != nil {
				errChan <- err
			}

			m.Store(info.Id, pbPost)
		}(post)

	}

	wg.Wait()
	close(errChan)
	close(finished)

	for _, pid := range posts {
		post, _ := m.Load(pid.Id)
		if post == nil {
			continue
		}
		pbPosts = append(pbPosts, post.(*pb.Post))
	}

	return pbPosts, nil
}

func convertPost(p *momentv1.Post) (*pb.Post, error) {
	post := pb.Post{}
	err := copier.Copy(&post, &p)
	if err != nil {
		return nil, err
	}
	return &post, nil
}

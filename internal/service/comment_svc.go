package service

import (
	"context"
	"fmt"
	"sync"
	"time"

	"github.com/go-microservice/ins-api/internal/ecode"

	"github.com/go-eagle/eagle/pkg/log"
	"github.com/spf13/cast"

	momentv1 "github.com/go-microservice/moment-service/api/moment/v1"
	userv1 "github.com/go-microservice/user-service/api/user/v1"
	"github.com/jinzhu/copier"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	pb "github.com/go-microservice/ins-api/api/micro/moment/v1"
)

var (
	_ pb.CommentServiceServer = (*CommentServiceServer)(nil)
)

type CommentServiceServer struct {
	pb.UnimplementedCommentServiceServer

	commentRPC momentv1.CommentServiceClient
	postRPC    momentv1.PostServiceClient
	userRPC    userv1.UserServiceClient
}

func NewCommentServiceServer(
	repo momentv1.CommentServiceClient,
	postRPC momentv1.PostServiceClient,
	userRepo userv1.UserServiceClient,
) *CommentServiceServer {
	return &CommentServiceServer{
		commentRPC: repo,
		postRPC:    postRPC,
		userRPC:    userRepo,
	}
}

func (s *CommentServiceServer) CreateComment(ctx context.Context, req *pb.CreateCommentRequest) (*pb.CreateCommentReply, error) {
	// check post if exist
	postReq := &momentv1.GetPostRequest{
		Id: req.GetPostId(),
	}
	_, err := s.postRPC.GetPost(ctx, postReq)
	if err != nil {
		// check err if not found
		statusErr, ok := status.FromError(err)
		if ok && statusErr.Code() == codes.NotFound {
			return nil, ecode.ErrPostNotFound
		}
		return nil, err
	}

	in := &momentv1.CreateCommentRequest{
		PostId:  req.PostId,
		UserId:  GetCurrentUserID(ctx),
		Content: req.Content,
	}
	out, err := s.commentRPC.CreateComment(ctx, in)
	if err != nil {
		// check client if deadline exceeded
		statusErr, ok := status.FromError(err)
		if ok && statusErr.Code() == codes.DeadlineExceeded {
			return nil, status.Error(codes.DeadlineExceeded, "deadline exceeded")
		}

		return nil, err
	}

	comment := pb.Comment{}
	err = copier.Copy(&comment, &out.Comment)
	if err != nil {
		return nil, err
	}

	// ger user info
	userIn := &userv1.GetUserRequest{Id: GetCurrentUserID(ctx)}
	user, err := s.userRPC.GetUser(ctx, userIn)
	if err != nil {
		return nil, err
	}
	comment.User, err = convertUser(user.GetUser())
	if err != nil {
		return nil, err
	}

	return &pb.CreateCommentReply{
		Comment: &comment,
	}, nil
}

func (s *CommentServiceServer) DeleteComment(ctx context.Context, req *pb.DeleteCommentRequest) (*pb.DeleteCommentReply, error) {
	in := &momentv1.DeleteCommentRequest{
		Id:      cast.ToInt64(req.GetId()),
		UserId:  GetCurrentUserID(ctx),
		DelFlag: req.GetDelFlag(),
	}
	_, err := s.commentRPC.DeleteComment(ctx, in)
	if err != nil {
		// check client if deadline exceeded
		statusErr, ok := status.FromError(err)
		if ok && statusErr.Code() == codes.DeadlineExceeded {
			return nil, status.Error(codes.DeadlineExceeded, "deadline exceeded")
		}
		return nil, err
	}

	return &pb.DeleteCommentReply{}, nil
}
func (s *CommentServiceServer) GetComment(ctx context.Context, req *pb.GetCommentRequest) (*pb.GetCommentReply, error) {
	in := &momentv1.GetCommentRequest{
		Id: cast.ToInt64(req.GetId()),
	}
	out, err := s.commentRPC.GetComment(ctx, in)
	if err != nil {
		// check client if deadline exceeded
		statusErr, ok := status.FromError(err)
		if ok && statusErr.Code() == codes.DeadlineExceeded {
			return nil, status.Error(codes.DeadlineExceeded, "deadline exceeded")
		}
		return nil, err
	}
	comment, err := convertComment(out.Comment)
	if err != nil {
		return nil, err
	}

	// ger user info
	userIn := &userv1.GetUserRequest{Id: out.GetComment().UserId}
	user, err := s.userRPC.GetUser(ctx, userIn)
	if err != nil {
		return nil, err
	}
	comment.User, err = convertUser(user.GetUser())
	if err != nil {
		return nil, err
	}

	return &pb.GetCommentReply{
		Comment: comment,
	}, nil
}
func (s *CommentServiceServer) ListHotComment(ctx context.Context, req *pb.ListCommentRequest) (*pb.ListCommentReply, error) {
	// check post if exist
	postReq := &momentv1.GetPostRequest{
		Id: req.GetPostId(),
	}
	_, err := s.postRPC.GetPost(ctx, postReq)
	if err != nil {
		// check err if not found
		statusErr, ok := status.FromError(err)
		if ok && statusErr.Code() == codes.NotFound {
			return nil, ecode.ErrPostNotFound
		}
		return nil, err
	}

	// get data, support pagination
	limit := cast.ToInt32(req.GetLimit())
	if limit == 0 {
		limit = 10
	}
	in := &momentv1.ListCommentRequest{
		PostId: req.GetPostId(),
		LastId: cast.ToInt64(req.GetLastId()),
		Limit:  limit + 1,
	}
	ret, err := s.commentRPC.ListHotComment(ctx, in)
	if err != nil {
		return nil, err
	}

	comments := ret.GetItems()
	var (
		hasMore bool
		lastId  string
	)
	if int32(len(comments)) > limit {
		hasMore = true
		lastId = cast.ToString(comments[len(comments)-1].Id)
		comments = comments[0 : len(comments)-1]
	}
	pbComments, err := s.assembleData(ctx, comments)
	if err != nil {
		return nil, err
	}

	return &pb.ListCommentReply{
		HasMore: hasMore,
		LastId:  lastId,
		Items:   pbComments,
	}, nil
}
func (s *CommentServiceServer) ListLatestComment(ctx context.Context, req *pb.ListCommentRequest) (*pb.ListCommentReply, error) {
	// check post if exist
	postReq := &momentv1.GetPostRequest{
		Id: req.GetPostId(),
	}
	_, err := s.postRPC.GetPost(ctx, postReq)
	if err != nil {
		// check err if not found
		statusErr, ok := status.FromError(err)
		if ok && statusErr.Code() == codes.NotFound {
			return nil, ecode.ErrPostNotFound
		}
		return nil, err
	}

	// get data, support pagination
	limit := cast.ToInt32(req.GetLimit())
	if limit == 0 {
		limit = 10
	}
	in := &momentv1.ListCommentRequest{
		PostId: req.GetPostId(),
		LastId: cast.ToInt64(req.GetLastId()),
		Limit:  limit + 1,
	}
	ret, err := s.commentRPC.ListLatestComment(ctx, in)
	if err != nil {
		return nil, err
	}

	comments := ret.GetItems()
	var (
		hasMore bool
		lastId  string
	)
	if int32(len(comments)) > limit {
		hasMore = true
		lastId = cast.ToString(comments[len(comments)-1].Id)
		comments = comments[0 : len(comments)-1]
	}
	pbComments, err := s.assembleData(ctx, comments)
	if err != nil {
		return nil, err
	}

	return &pb.ListCommentReply{
		HasMore: hasMore,
		LastId:  lastId,
		Items:   pbComments,
	}, nil
}

func (s *CommentServiceServer) ReplyComment(ctx context.Context, req *pb.ReplyCommentRequest) (*pb.ReplyCommentReply, error) {
	in := &momentv1.ReplyCommentRequest{
		CommentId:  req.GetCommentId(),
		RootId:     req.GetRootId(),
		UserId:     GetCurrentUserID(ctx),
		Content:    req.GetContent(),
		DeviceType: req.GetDeviceType(),
		Ip:         "",
	}
	out, err := s.commentRPC.ReplyComment(ctx, in)
	if err != nil {
		// check client if deadline exceeded
		statusErr, ok := status.FromError(err)
		if ok && statusErr.Code() == codes.DeadlineExceeded {
			return nil, status.Error(codes.DeadlineExceeded, "deadline exceeded")
		}

		return nil, err
	}

	comment := pb.Comment{}
	err = copier.Copy(&comment, &out.Comment)
	if err != nil {
		return nil, err
	}
	return &pb.ReplyCommentReply{
		Comment: &comment,
	}, nil
}

func (s *CommentServiceServer) ListReply(ctx context.Context, req *pb.ListReplyRequest) (*pb.ListReplyReply, error) {
	// get data, support pagination
	limit := cast.ToInt32(req.GetLimit())
	if limit == 0 {
		limit = 10
	}
	in := &momentv1.ListReplyCommentRequest{
		CommentId: req.GetCommentId(),
		LastId:    cast.ToInt64(req.GetLastId()),
		Limit:     limit + 1,
	}
	ret, err := s.commentRPC.ListReplyComment(ctx, in)
	if err != nil {
		return nil, err
	}

	comments := ret.GetItems()
	var (
		hasMore bool
		lastId  string
	)
	if int32(len(comments)) > limit {
		hasMore = true
		lastId = cast.ToString(comments[len(comments)-1].Id)
		comments = comments[0 : len(comments)-1]
	}
	pbComments, err := s.assembleData(ctx, comments)
	if err != nil {
		return nil, err
	}

	return &pb.ListReplyReply{
		HasMore: hasMore,
		LastId:  lastId,
		Items:   pbComments,
	}, nil
}

func (s *CommentServiceServer) assembleData(ctx context.Context, comments []*momentv1.Comment) ([]*pb.Comment, error) {
	// batch get user data
	var (
		userIDs []int64
	)
	for _, v := range comments {
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
		pbComments []*pb.Comment
		m          sync.Map
		mu         sync.Mutex
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

	for _, comment := range comments {
		wg.Add(1)
		comment := comment
		go func(info *momentv1.Comment) {
			defer func() {
				wg.Done()
			}()

			mu.Lock()
			defer mu.Unlock()

			pbPost, err := convertComment(info)
			if err != nil {
				return
			}
			// user
			user, ok := userMap[comment.UserId]
			if !ok {
				return
			}
			pbPost.User, err = convertUser(user)
			if err != nil {
				errChan <- err
			}

			m.Store(info.Id, pbPost)
		}(comment)

	}

	wg.Wait()
	close(errChan)
	close(finished)

	for _, pid := range comments {
		comment, _ := m.Load(pid.Id)
		if comment == nil {
			continue
		}
		pbComments = append(pbComments, comment.(*pb.Comment))
	}

	return pbComments, nil
}

func convertComment(p *momentv1.Comment) (*pb.Comment, error) {
	comment := pb.Comment{}
	err := copier.Copy(&comment, &p)
	if err != nil {
		return nil, err
	}
	comment.Id = cast.ToString(p.Id)
	return &comment, nil
}

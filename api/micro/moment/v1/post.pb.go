// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.0
// 	protoc        v3.12.3
// source: api/micro/moment/v1/post.proto

package v1

import (
	v1 "github.com/go-microservice/ins-api/api/micro/user/v1"
	_ "google.golang.org/genproto/googleapis/api/annotations"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type Post struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id           int64    `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	PostType     int32    `protobuf:"varint,2,opt,name=post_type,json=postType,proto3" json:"post_type,omitempty"`
	User         *v1.User `protobuf:"bytes,3,opt,name=user,proto3" json:"user,omitempty"`
	Title        string   `protobuf:"bytes,4,opt,name=title,proto3" json:"title,omitempty"`
	Content      string   `protobuf:"bytes,5,opt,name=content,proto3" json:"content,omitempty"`
	ViewCount    int64    `protobuf:"varint,6,opt,name=view_count,json=viewCount,proto3" json:"view_count,omitempty"`
	LikeCount    int64    `protobuf:"varint,7,opt,name=like_count,json=likeCount,proto3" json:"like_count,omitempty"`
	CommentCount int64    `protobuf:"varint,8,opt,name=comment_count,json=commentCount,proto3" json:"comment_count,omitempty"`
	CollectCount int64    `protobuf:"varint,9,opt,name=collect_count,json=collectCount,proto3" json:"collect_count,omitempty"`
	ShareCount   int64    `protobuf:"varint,10,opt,name=share_count,json=shareCount,proto3" json:"share_count,omitempty"`
	Longitude    float32  `protobuf:"fixed32,11,opt,name=longitude,proto3" json:"longitude,omitempty"`
	Latitude     float32  `protobuf:"fixed32,12,opt,name=latitude,proto3" json:"latitude,omitempty"`
	Position     float32  `protobuf:"fixed32,13,opt,name=position,proto3" json:"position,omitempty"`
	CreatedAt    int64    `protobuf:"varint,14,opt,name=created_at,json=createdAt,proto3" json:"created_at,omitempty"`
	UpdatedAt    int64    `protobuf:"varint,15,opt,name=updated_at,json=updatedAt,proto3" json:"updated_at,omitempty"`
}

func (x *Post) Reset() {
	*x = Post{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_micro_moment_v1_post_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Post) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Post) ProtoMessage() {}

func (x *Post) ProtoReflect() protoreflect.Message {
	mi := &file_api_micro_moment_v1_post_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Post.ProtoReflect.Descriptor instead.
func (*Post) Descriptor() ([]byte, []int) {
	return file_api_micro_moment_v1_post_proto_rawDescGZIP(), []int{0}
}

func (x *Post) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *Post) GetPostType() int32 {
	if x != nil {
		return x.PostType
	}
	return 0
}

func (x *Post) GetUser() *v1.User {
	if x != nil {
		return x.User
	}
	return nil
}

func (x *Post) GetTitle() string {
	if x != nil {
		return x.Title
	}
	return ""
}

func (x *Post) GetContent() string {
	if x != nil {
		return x.Content
	}
	return ""
}

func (x *Post) GetViewCount() int64 {
	if x != nil {
		return x.ViewCount
	}
	return 0
}

func (x *Post) GetLikeCount() int64 {
	if x != nil {
		return x.LikeCount
	}
	return 0
}

func (x *Post) GetCommentCount() int64 {
	if x != nil {
		return x.CommentCount
	}
	return 0
}

func (x *Post) GetCollectCount() int64 {
	if x != nil {
		return x.CollectCount
	}
	return 0
}

func (x *Post) GetShareCount() int64 {
	if x != nil {
		return x.ShareCount
	}
	return 0
}

func (x *Post) GetLongitude() float32 {
	if x != nil {
		return x.Longitude
	}
	return 0
}

func (x *Post) GetLatitude() float32 {
	if x != nil {
		return x.Latitude
	}
	return 0
}

func (x *Post) GetPosition() float32 {
	if x != nil {
		return x.Position
	}
	return 0
}

func (x *Post) GetCreatedAt() int64 {
	if x != nil {
		return x.CreatedAt
	}
	return 0
}

func (x *Post) GetUpdatedAt() int64 {
	if x != nil {
		return x.UpdatedAt
	}
	return 0
}

type CreatePostRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserId        int64   `protobuf:"varint,1,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	Title         string  `protobuf:"bytes,2,opt,name=title,proto3" json:"title,omitempty"`
	Text          string  `protobuf:"bytes,3,opt,name=text,proto3" json:"text,omitempty"`
	PicKeys       string  `protobuf:"bytes,4,opt,name=pic_keys,json=picKeys,proto3" json:"pic_keys,omitempty"`
	VideoKey      string  `protobuf:"bytes,5,opt,name=video_key,json=videoKey,proto3" json:"video_key,omitempty"`
	VideoDuration int32   `protobuf:"varint,6,opt,name=video_duration,json=videoDuration,proto3" json:"video_duration,omitempty"`
	CoverKey      string  `protobuf:"bytes,7,opt,name=cover_key,json=coverKey,proto3" json:"cover_key,omitempty"`
	CoverWidth    int32   `protobuf:"varint,8,opt,name=cover_width,json=coverWidth,proto3" json:"cover_width,omitempty"`
	CoverHeight   int32   `protobuf:"varint,9,opt,name=cover_height,json=coverHeight,proto3" json:"cover_height,omitempty"`
	Longitude     float32 `protobuf:"fixed32,10,opt,name=longitude,proto3" json:"longitude,omitempty"`
	Latitude      float32 `protobuf:"fixed32,11,opt,name=latitude,proto3" json:"latitude,omitempty"`
	Position      string  `protobuf:"bytes,12,opt,name=position,proto3" json:"position,omitempty"`
}

func (x *CreatePostRequest) Reset() {
	*x = CreatePostRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_micro_moment_v1_post_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreatePostRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreatePostRequest) ProtoMessage() {}

func (x *CreatePostRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_micro_moment_v1_post_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreatePostRequest.ProtoReflect.Descriptor instead.
func (*CreatePostRequest) Descriptor() ([]byte, []int) {
	return file_api_micro_moment_v1_post_proto_rawDescGZIP(), []int{1}
}

func (x *CreatePostRequest) GetUserId() int64 {
	if x != nil {
		return x.UserId
	}
	return 0
}

func (x *CreatePostRequest) GetTitle() string {
	if x != nil {
		return x.Title
	}
	return ""
}

func (x *CreatePostRequest) GetText() string {
	if x != nil {
		return x.Text
	}
	return ""
}

func (x *CreatePostRequest) GetPicKeys() string {
	if x != nil {
		return x.PicKeys
	}
	return ""
}

func (x *CreatePostRequest) GetVideoKey() string {
	if x != nil {
		return x.VideoKey
	}
	return ""
}

func (x *CreatePostRequest) GetVideoDuration() int32 {
	if x != nil {
		return x.VideoDuration
	}
	return 0
}

func (x *CreatePostRequest) GetCoverKey() string {
	if x != nil {
		return x.CoverKey
	}
	return ""
}

func (x *CreatePostRequest) GetCoverWidth() int32 {
	if x != nil {
		return x.CoverWidth
	}
	return 0
}

func (x *CreatePostRequest) GetCoverHeight() int32 {
	if x != nil {
		return x.CoverHeight
	}
	return 0
}

func (x *CreatePostRequest) GetLongitude() float32 {
	if x != nil {
		return x.Longitude
	}
	return 0
}

func (x *CreatePostRequest) GetLatitude() float32 {
	if x != nil {
		return x.Latitude
	}
	return 0
}

func (x *CreatePostRequest) GetPosition() string {
	if x != nil {
		return x.Position
	}
	return ""
}

type CreatePostReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Post *Post `protobuf:"bytes,1,opt,name=post,proto3" json:"post,omitempty"`
}

func (x *CreatePostReply) Reset() {
	*x = CreatePostReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_micro_moment_v1_post_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreatePostReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreatePostReply) ProtoMessage() {}

func (x *CreatePostReply) ProtoReflect() protoreflect.Message {
	mi := &file_api_micro_moment_v1_post_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreatePostReply.ProtoReflect.Descriptor instead.
func (*CreatePostReply) Descriptor() ([]byte, []int) {
	return file_api_micro_moment_v1_post_proto_rawDescGZIP(), []int{2}
}

func (x *CreatePostReply) GetPost() *Post {
	if x != nil {
		return x.Post
	}
	return nil
}

type UpdatePostRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *UpdatePostRequest) Reset() {
	*x = UpdatePostRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_micro_moment_v1_post_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdatePostRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdatePostRequest) ProtoMessage() {}

func (x *UpdatePostRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_micro_moment_v1_post_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdatePostRequest.ProtoReflect.Descriptor instead.
func (*UpdatePostRequest) Descriptor() ([]byte, []int) {
	return file_api_micro_moment_v1_post_proto_rawDescGZIP(), []int{3}
}

type UpdatePostReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *UpdatePostReply) Reset() {
	*x = UpdatePostReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_micro_moment_v1_post_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdatePostReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdatePostReply) ProtoMessage() {}

func (x *UpdatePostReply) ProtoReflect() protoreflect.Message {
	mi := &file_api_micro_moment_v1_post_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdatePostReply.ProtoReflect.Descriptor instead.
func (*UpdatePostReply) Descriptor() ([]byte, []int) {
	return file_api_micro_moment_v1_post_proto_rawDescGZIP(), []int{4}
}

type DeletePostRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *DeletePostRequest) Reset() {
	*x = DeletePostRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_micro_moment_v1_post_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeletePostRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeletePostRequest) ProtoMessage() {}

func (x *DeletePostRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_micro_moment_v1_post_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeletePostRequest.ProtoReflect.Descriptor instead.
func (*DeletePostRequest) Descriptor() ([]byte, []int) {
	return file_api_micro_moment_v1_post_proto_rawDescGZIP(), []int{5}
}

type DeletePostReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *DeletePostReply) Reset() {
	*x = DeletePostReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_micro_moment_v1_post_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeletePostReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeletePostReply) ProtoMessage() {}

func (x *DeletePostReply) ProtoReflect() protoreflect.Message {
	mi := &file_api_micro_moment_v1_post_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeletePostReply.ProtoReflect.Descriptor instead.
func (*DeletePostReply) Descriptor() ([]byte, []int) {
	return file_api_micro_moment_v1_post_proto_rawDescGZIP(), []int{6}
}

type GetPostRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// @gotags: uri:"id"
	Id int64 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *GetPostRequest) Reset() {
	*x = GetPostRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_micro_moment_v1_post_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetPostRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetPostRequest) ProtoMessage() {}

func (x *GetPostRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_micro_moment_v1_post_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetPostRequest.ProtoReflect.Descriptor instead.
func (*GetPostRequest) Descriptor() ([]byte, []int) {
	return file_api_micro_moment_v1_post_proto_rawDescGZIP(), []int{7}
}

func (x *GetPostRequest) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

type GetPostReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Post *Post `protobuf:"bytes,1,opt,name=post,proto3" json:"post,omitempty"`
}

func (x *GetPostReply) Reset() {
	*x = GetPostReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_micro_moment_v1_post_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetPostReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetPostReply) ProtoMessage() {}

func (x *GetPostReply) ProtoReflect() protoreflect.Message {
	mi := &file_api_micro_moment_v1_post_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetPostReply.ProtoReflect.Descriptor instead.
func (*GetPostReply) Descriptor() ([]byte, []int) {
	return file_api_micro_moment_v1_post_proto_rawDescGZIP(), []int{8}
}

func (x *GetPostReply) GetPost() *Post {
	if x != nil {
		return x.Post
	}
	return nil
}

type ListPostRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// @gotags: form:"last_id"
	LastId int64 `protobuf:"varint,1,opt,name=last_id,json=lastId,proto3" json:"last_id,omitempty"`
	// @gotags: form:"limit"
	Limit int32 `protobuf:"varint,2,opt,name=limit,proto3" json:"limit,omitempty"`
}

func (x *ListPostRequest) Reset() {
	*x = ListPostRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_micro_moment_v1_post_proto_msgTypes[9]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListPostRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListPostRequest) ProtoMessage() {}

func (x *ListPostRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_micro_moment_v1_post_proto_msgTypes[9]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListPostRequest.ProtoReflect.Descriptor instead.
func (*ListPostRequest) Descriptor() ([]byte, []int) {
	return file_api_micro_moment_v1_post_proto_rawDescGZIP(), []int{9}
}

func (x *ListPostRequest) GetLastId() int64 {
	if x != nil {
		return x.LastId
	}
	return 0
}

func (x *ListPostRequest) GetLimit() int32 {
	if x != nil {
		return x.Limit
	}
	return 0
}

type ListPostReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Items   []*Post `protobuf:"bytes,1,rep,name=items,proto3" json:"items,omitempty"`
	Count   int64   `protobuf:"varint,2,opt,name=count,proto3" json:"count,omitempty"`
	HasMore bool    `protobuf:"varint,3,opt,name=has_more,json=hasMore,proto3" json:"has_more,omitempty"`
	LastId  string  `protobuf:"bytes,4,opt,name=last_id,json=lastId,proto3" json:"last_id,omitempty"`
}

func (x *ListPostReply) Reset() {
	*x = ListPostReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_micro_moment_v1_post_proto_msgTypes[10]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListPostReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListPostReply) ProtoMessage() {}

func (x *ListPostReply) ProtoReflect() protoreflect.Message {
	mi := &file_api_micro_moment_v1_post_proto_msgTypes[10]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListPostReply.ProtoReflect.Descriptor instead.
func (*ListPostReply) Descriptor() ([]byte, []int) {
	return file_api_micro_moment_v1_post_proto_rawDescGZIP(), []int{10}
}

func (x *ListPostReply) GetItems() []*Post {
	if x != nil {
		return x.Items
	}
	return nil
}

func (x *ListPostReply) GetCount() int64 {
	if x != nil {
		return x.Count
	}
	return 0
}

func (x *ListPostReply) GetHasMore() bool {
	if x != nil {
		return x.HasMore
	}
	return false
}

func (x *ListPostReply) GetLastId() string {
	if x != nil {
		return x.LastId
	}
	return ""
}

var File_api_micro_moment_v1_post_proto protoreflect.FileDescriptor

var file_api_micro_moment_v1_post_proto_rawDesc = []byte{
	0x0a, 0x1e, 0x61, 0x70, 0x69, 0x2f, 0x6d, 0x69, 0x63, 0x72, 0x6f, 0x2f, 0x6d, 0x6f, 0x6d, 0x65,
	0x6e, 0x74, 0x2f, 0x76, 0x31, 0x2f, 0x70, 0x6f, 0x73, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x12, 0x13, 0x61, 0x70, 0x69, 0x2e, 0x6d, 0x69, 0x63, 0x72, 0x6f, 0x2e, 0x6d, 0x6f, 0x6d, 0x65,
	0x6e, 0x74, 0x2e, 0x76, 0x31, 0x1a, 0x1c, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70,
	0x69, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x1a, 0x1c, 0x61, 0x70, 0x69, 0x2f, 0x6d, 0x69, 0x63, 0x72, 0x6f, 0x2f, 0x75,
	0x73, 0x65, 0x72, 0x2f, 0x76, 0x31, 0x2f, 0x75, 0x73, 0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x22, 0xcd, 0x03, 0x0a, 0x04, 0x50, 0x6f, 0x73, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x02, 0x69, 0x64, 0x12, 0x1b, 0x0a, 0x09, 0x70, 0x6f,
	0x73, 0x74, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x08, 0x70,
	0x6f, 0x73, 0x74, 0x54, 0x79, 0x70, 0x65, 0x12, 0x2b, 0x0a, 0x04, 0x75, 0x73, 0x65, 0x72, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x17, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x6d, 0x69, 0x63, 0x72,
	0x6f, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x2e, 0x76, 0x31, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x52, 0x04,
	0x75, 0x73, 0x65, 0x72, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x18, 0x04, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x05, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x63, 0x6f,
	0x6e, 0x74, 0x65, 0x6e, 0x74, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x63, 0x6f, 0x6e,
	0x74, 0x65, 0x6e, 0x74, 0x12, 0x1d, 0x0a, 0x0a, 0x76, 0x69, 0x65, 0x77, 0x5f, 0x63, 0x6f, 0x75,
	0x6e, 0x74, 0x18, 0x06, 0x20, 0x01, 0x28, 0x03, 0x52, 0x09, 0x76, 0x69, 0x65, 0x77, 0x43, 0x6f,
	0x75, 0x6e, 0x74, 0x12, 0x1d, 0x0a, 0x0a, 0x6c, 0x69, 0x6b, 0x65, 0x5f, 0x63, 0x6f, 0x75, 0x6e,
	0x74, 0x18, 0x07, 0x20, 0x01, 0x28, 0x03, 0x52, 0x09, 0x6c, 0x69, 0x6b, 0x65, 0x43, 0x6f, 0x75,
	0x6e, 0x74, 0x12, 0x23, 0x0a, 0x0d, 0x63, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x5f, 0x63, 0x6f,
	0x75, 0x6e, 0x74, 0x18, 0x08, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0c, 0x63, 0x6f, 0x6d, 0x6d, 0x65,
	0x6e, 0x74, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x23, 0x0a, 0x0d, 0x63, 0x6f, 0x6c, 0x6c, 0x65,
	0x63, 0x74, 0x5f, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x09, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0c,
	0x63, 0x6f, 0x6c, 0x6c, 0x65, 0x63, 0x74, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x1f, 0x0a, 0x0b,
	0x73, 0x68, 0x61, 0x72, 0x65, 0x5f, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x0a, 0x20, 0x01, 0x28,
	0x03, 0x52, 0x0a, 0x73, 0x68, 0x61, 0x72, 0x65, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x1c, 0x0a,
	0x09, 0x6c, 0x6f, 0x6e, 0x67, 0x69, 0x74, 0x75, 0x64, 0x65, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x02,
	0x52, 0x09, 0x6c, 0x6f, 0x6e, 0x67, 0x69, 0x74, 0x75, 0x64, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x6c,
	0x61, 0x74, 0x69, 0x74, 0x75, 0x64, 0x65, 0x18, 0x0c, 0x20, 0x01, 0x28, 0x02, 0x52, 0x08, 0x6c,
	0x61, 0x74, 0x69, 0x74, 0x75, 0x64, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x70, 0x6f, 0x73, 0x69, 0x74,
	0x69, 0x6f, 0x6e, 0x18, 0x0d, 0x20, 0x01, 0x28, 0x02, 0x52, 0x08, 0x70, 0x6f, 0x73, 0x69, 0x74,
	0x69, 0x6f, 0x6e, 0x12, 0x1d, 0x0a, 0x0a, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x61,
	0x74, 0x18, 0x0e, 0x20, 0x01, 0x28, 0x03, 0x52, 0x09, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64,
	0x41, 0x74, 0x12, 0x1d, 0x0a, 0x0a, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x61, 0x74,
	0x18, 0x0f, 0x20, 0x01, 0x28, 0x03, 0x52, 0x09, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x41,
	0x74, 0x22, 0xec, 0x02, 0x0a, 0x11, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x50, 0x6f, 0x73, 0x74,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x17, 0x0a, 0x07, 0x75, 0x73, 0x65, 0x72, 0x5f,
	0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64,
	0x12, 0x14, 0x0a, 0x05, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x05, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x74, 0x65, 0x78, 0x74, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x74, 0x65, 0x78, 0x74, 0x12, 0x19, 0x0a, 0x08, 0x70, 0x69,
	0x63, 0x5f, 0x6b, 0x65, 0x79, 0x73, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x70, 0x69,
	0x63, 0x4b, 0x65, 0x79, 0x73, 0x12, 0x1b, 0x0a, 0x09, 0x76, 0x69, 0x64, 0x65, 0x6f, 0x5f, 0x6b,
	0x65, 0x79, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x76, 0x69, 0x64, 0x65, 0x6f, 0x4b,
	0x65, 0x79, 0x12, 0x25, 0x0a, 0x0e, 0x76, 0x69, 0x64, 0x65, 0x6f, 0x5f, 0x64, 0x75, 0x72, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x18, 0x06, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0d, 0x76, 0x69, 0x64, 0x65,
	0x6f, 0x44, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x1b, 0x0a, 0x09, 0x63, 0x6f, 0x76,
	0x65, 0x72, 0x5f, 0x6b, 0x65, 0x79, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x63, 0x6f,
	0x76, 0x65, 0x72, 0x4b, 0x65, 0x79, 0x12, 0x1f, 0x0a, 0x0b, 0x63, 0x6f, 0x76, 0x65, 0x72, 0x5f,
	0x77, 0x69, 0x64, 0x74, 0x68, 0x18, 0x08, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0a, 0x63, 0x6f, 0x76,
	0x65, 0x72, 0x57, 0x69, 0x64, 0x74, 0x68, 0x12, 0x21, 0x0a, 0x0c, 0x63, 0x6f, 0x76, 0x65, 0x72,
	0x5f, 0x68, 0x65, 0x69, 0x67, 0x68, 0x74, 0x18, 0x09, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0b, 0x63,
	0x6f, 0x76, 0x65, 0x72, 0x48, 0x65, 0x69, 0x67, 0x68, 0x74, 0x12, 0x1c, 0x0a, 0x09, 0x6c, 0x6f,
	0x6e, 0x67, 0x69, 0x74, 0x75, 0x64, 0x65, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x02, 0x52, 0x09, 0x6c,
	0x6f, 0x6e, 0x67, 0x69, 0x74, 0x75, 0x64, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x6c, 0x61, 0x74, 0x69,
	0x74, 0x75, 0x64, 0x65, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x02, 0x52, 0x08, 0x6c, 0x61, 0x74, 0x69,
	0x74, 0x75, 0x64, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x70, 0x6f, 0x73, 0x69, 0x74, 0x69, 0x6f, 0x6e,
	0x18, 0x0c, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x70, 0x6f, 0x73, 0x69, 0x74, 0x69, 0x6f, 0x6e,
	0x22, 0x40, 0x0a, 0x0f, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x50, 0x6f, 0x73, 0x74, 0x52, 0x65,
	0x70, 0x6c, 0x79, 0x12, 0x2d, 0x0a, 0x04, 0x70, 0x6f, 0x73, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x19, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x6d, 0x69, 0x63, 0x72, 0x6f, 0x2e, 0x6d, 0x6f,
	0x6d, 0x65, 0x6e, 0x74, 0x2e, 0x76, 0x31, 0x2e, 0x50, 0x6f, 0x73, 0x74, 0x52, 0x04, 0x70, 0x6f,
	0x73, 0x74, 0x22, 0x13, 0x0a, 0x11, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x50, 0x6f, 0x73, 0x74,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x22, 0x11, 0x0a, 0x0f, 0x55, 0x70, 0x64, 0x61, 0x74,
	0x65, 0x50, 0x6f, 0x73, 0x74, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x22, 0x13, 0x0a, 0x11, 0x44, 0x65,
	0x6c, 0x65, 0x74, 0x65, 0x50, 0x6f, 0x73, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x22,
	0x11, 0x0a, 0x0f, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x50, 0x6f, 0x73, 0x74, 0x52, 0x65, 0x70,
	0x6c, 0x79, 0x22, 0x20, 0x0a, 0x0e, 0x47, 0x65, 0x74, 0x50, 0x6f, 0x73, 0x74, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03,
	0x52, 0x02, 0x69, 0x64, 0x22, 0x3d, 0x0a, 0x0c, 0x47, 0x65, 0x74, 0x50, 0x6f, 0x73, 0x74, 0x52,
	0x65, 0x70, 0x6c, 0x79, 0x12, 0x2d, 0x0a, 0x04, 0x70, 0x6f, 0x73, 0x74, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x19, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x6d, 0x69, 0x63, 0x72, 0x6f, 0x2e, 0x6d,
	0x6f, 0x6d, 0x65, 0x6e, 0x74, 0x2e, 0x76, 0x31, 0x2e, 0x50, 0x6f, 0x73, 0x74, 0x52, 0x04, 0x70,
	0x6f, 0x73, 0x74, 0x22, 0x40, 0x0a, 0x0f, 0x4c, 0x69, 0x73, 0x74, 0x50, 0x6f, 0x73, 0x74, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x17, 0x0a, 0x07, 0x6c, 0x61, 0x73, 0x74, 0x5f, 0x69,
	0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x06, 0x6c, 0x61, 0x73, 0x74, 0x49, 0x64, 0x12,
	0x14, 0x0a, 0x05, 0x6c, 0x69, 0x6d, 0x69, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x05,
	0x6c, 0x69, 0x6d, 0x69, 0x74, 0x22, 0x8a, 0x01, 0x0a, 0x0d, 0x4c, 0x69, 0x73, 0x74, 0x50, 0x6f,
	0x73, 0x74, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x12, 0x2f, 0x0a, 0x05, 0x69, 0x74, 0x65, 0x6d, 0x73,
	0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x19, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x6d, 0x69, 0x63,
	0x72, 0x6f, 0x2e, 0x6d, 0x6f, 0x6d, 0x65, 0x6e, 0x74, 0x2e, 0x76, 0x31, 0x2e, 0x50, 0x6f, 0x73,
	0x74, 0x52, 0x05, 0x69, 0x74, 0x65, 0x6d, 0x73, 0x12, 0x14, 0x0a, 0x05, 0x63, 0x6f, 0x75, 0x6e,
	0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x05, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x19,
	0x0a, 0x08, 0x68, 0x61, 0x73, 0x5f, 0x6d, 0x6f, 0x72, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x08,
	0x52, 0x07, 0x68, 0x61, 0x73, 0x4d, 0x6f, 0x72, 0x65, 0x12, 0x17, 0x0a, 0x07, 0x6c, 0x61, 0x73,
	0x74, 0x5f, 0x69, 0x64, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x6c, 0x61, 0x73, 0x74,
	0x49, 0x64, 0x32, 0xb4, 0x05, 0x0a, 0x0b, 0x50, 0x6f, 0x73, 0x74, 0x53, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x12, 0x70, 0x0a, 0x0a, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x50, 0x6f, 0x73, 0x74,
	0x12, 0x26, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x6d, 0x69, 0x63, 0x72, 0x6f, 0x2e, 0x6d, 0x6f, 0x6d,
	0x65, 0x6e, 0x74, 0x2e, 0x76, 0x31, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x50, 0x6f, 0x73,
	0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x24, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x6d,
	0x69, 0x63, 0x72, 0x6f, 0x2e, 0x6d, 0x6f, 0x6d, 0x65, 0x6e, 0x74, 0x2e, 0x76, 0x31, 0x2e, 0x43,
	0x72, 0x65, 0x61, 0x74, 0x65, 0x50, 0x6f, 0x73, 0x74, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x22, 0x14,
	0x82, 0xd3, 0xe4, 0x93, 0x02, 0x0e, 0x22, 0x09, 0x2f, 0x76, 0x31, 0x2f, 0x70, 0x6f, 0x73, 0x74,
	0x73, 0x3a, 0x01, 0x2a, 0x12, 0x70, 0x0a, 0x0a, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x50, 0x6f,
	0x73, 0x74, 0x12, 0x26, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x6d, 0x69, 0x63, 0x72, 0x6f, 0x2e, 0x6d,
	0x6f, 0x6d, 0x65, 0x6e, 0x74, 0x2e, 0x76, 0x31, 0x2e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x50,
	0x6f, 0x73, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x24, 0x2e, 0x61, 0x70, 0x69,
	0x2e, 0x6d, 0x69, 0x63, 0x72, 0x6f, 0x2e, 0x6d, 0x6f, 0x6d, 0x65, 0x6e, 0x74, 0x2e, 0x76, 0x31,
	0x2e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x50, 0x6f, 0x73, 0x74, 0x52, 0x65, 0x70, 0x6c, 0x79,
	0x22, 0x14, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x0e, 0x32, 0x09, 0x2f, 0x76, 0x31, 0x2f, 0x70, 0x6f,
	0x73, 0x74, 0x73, 0x3a, 0x01, 0x2a, 0x12, 0x70, 0x0a, 0x0a, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65,
	0x50, 0x6f, 0x73, 0x74, 0x12, 0x26, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x6d, 0x69, 0x63, 0x72, 0x6f,
	0x2e, 0x6d, 0x6f, 0x6d, 0x65, 0x6e, 0x74, 0x2e, 0x76, 0x31, 0x2e, 0x44, 0x65, 0x6c, 0x65, 0x74,
	0x65, 0x50, 0x6f, 0x73, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x24, 0x2e, 0x61,
	0x70, 0x69, 0x2e, 0x6d, 0x69, 0x63, 0x72, 0x6f, 0x2e, 0x6d, 0x6f, 0x6d, 0x65, 0x6e, 0x74, 0x2e,
	0x76, 0x31, 0x2e, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x50, 0x6f, 0x73, 0x74, 0x52, 0x65, 0x70,
	0x6c, 0x79, 0x22, 0x14, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x0e, 0x2a, 0x09, 0x2f, 0x76, 0x31, 0x2f,
	0x70, 0x6f, 0x73, 0x74, 0x73, 0x3a, 0x01, 0x2a, 0x12, 0x69, 0x0a, 0x07, 0x47, 0x65, 0x74, 0x50,
	0x6f, 0x73, 0x74, 0x12, 0x23, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x6d, 0x69, 0x63, 0x72, 0x6f, 0x2e,
	0x6d, 0x6f, 0x6d, 0x65, 0x6e, 0x74, 0x2e, 0x76, 0x31, 0x2e, 0x47, 0x65, 0x74, 0x50, 0x6f, 0x73,
	0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x21, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x6d,
	0x69, 0x63, 0x72, 0x6f, 0x2e, 0x6d, 0x6f, 0x6d, 0x65, 0x6e, 0x74, 0x2e, 0x76, 0x31, 0x2e, 0x47,
	0x65, 0x74, 0x50, 0x6f, 0x73, 0x74, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x22, 0x16, 0x82, 0xd3, 0xe4,
	0x93, 0x02, 0x10, 0x12, 0x0e, 0x2f, 0x76, 0x31, 0x2f, 0x70, 0x6f, 0x73, 0x74, 0x73, 0x2f, 0x7b,
	0x69, 0x64, 0x7d, 0x12, 0x6e, 0x0a, 0x0b, 0x4c, 0x69, 0x73, 0x74, 0x48, 0x6f, 0x74, 0x50, 0x6f,
	0x73, 0x74, 0x12, 0x24, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x6d, 0x69, 0x63, 0x72, 0x6f, 0x2e, 0x6d,
	0x6f, 0x6d, 0x65, 0x6e, 0x74, 0x2e, 0x76, 0x31, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x50, 0x6f, 0x73,
	0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x22, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x6d,
	0x69, 0x63, 0x72, 0x6f, 0x2e, 0x6d, 0x6f, 0x6d, 0x65, 0x6e, 0x74, 0x2e, 0x76, 0x31, 0x2e, 0x4c,
	0x69, 0x73, 0x74, 0x50, 0x6f, 0x73, 0x74, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x22, 0x15, 0x82, 0xd3,
	0xe4, 0x93, 0x02, 0x0f, 0x12, 0x0d, 0x2f, 0x76, 0x31, 0x2f, 0x70, 0x6f, 0x73, 0x74, 0x73, 0x2f,
	0x68, 0x6f, 0x74, 0x12, 0x74, 0x0a, 0x0e, 0x4c, 0x69, 0x73, 0x74, 0x4c, 0x61, 0x74, 0x65, 0x73,
	0x74, 0x50, 0x6f, 0x73, 0x74, 0x12, 0x24, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x6d, 0x69, 0x63, 0x72,
	0x6f, 0x2e, 0x6d, 0x6f, 0x6d, 0x65, 0x6e, 0x74, 0x2e, 0x76, 0x31, 0x2e, 0x4c, 0x69, 0x73, 0x74,
	0x50, 0x6f, 0x73, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x22, 0x2e, 0x61, 0x70,
	0x69, 0x2e, 0x6d, 0x69, 0x63, 0x72, 0x6f, 0x2e, 0x6d, 0x6f, 0x6d, 0x65, 0x6e, 0x74, 0x2e, 0x76,
	0x31, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x50, 0x6f, 0x73, 0x74, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x22,
	0x18, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x12, 0x12, 0x10, 0x2f, 0x76, 0x31, 0x2f, 0x70, 0x6f, 0x73,
	0x74, 0x73, 0x2f, 0x6c, 0x61, 0x74, 0x65, 0x73, 0x74, 0x42, 0x52, 0x0a, 0x13, 0x61, 0x70, 0x69,
	0x2e, 0x6d, 0x69, 0x63, 0x72, 0x6f, 0x2e, 0x6d, 0x6f, 0x6d, 0x65, 0x6e, 0x74, 0x2e, 0x76, 0x31,
	0x50, 0x01, 0x5a, 0x39, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x67,
	0x6f, 0x2d, 0x6d, 0x69, 0x63, 0x72, 0x6f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2f, 0x69,
	0x6e, 0x73, 0x2d, 0x61, 0x70, 0x69, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x6d, 0x69, 0x63, 0x72, 0x6f,
	0x2f, 0x6d, 0x6f, 0x6d, 0x65, 0x6e, 0x74, 0x2f, 0x76, 0x31, 0x3b, 0x76, 0x31, 0x62, 0x06, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_api_micro_moment_v1_post_proto_rawDescOnce sync.Once
	file_api_micro_moment_v1_post_proto_rawDescData = file_api_micro_moment_v1_post_proto_rawDesc
)

func file_api_micro_moment_v1_post_proto_rawDescGZIP() []byte {
	file_api_micro_moment_v1_post_proto_rawDescOnce.Do(func() {
		file_api_micro_moment_v1_post_proto_rawDescData = protoimpl.X.CompressGZIP(file_api_micro_moment_v1_post_proto_rawDescData)
	})
	return file_api_micro_moment_v1_post_proto_rawDescData
}

var file_api_micro_moment_v1_post_proto_msgTypes = make([]protoimpl.MessageInfo, 11)
var file_api_micro_moment_v1_post_proto_goTypes = []interface{}{
	(*Post)(nil),              // 0: api.micro.moment.v1.Post
	(*CreatePostRequest)(nil), // 1: api.micro.moment.v1.CreatePostRequest
	(*CreatePostReply)(nil),   // 2: api.micro.moment.v1.CreatePostReply
	(*UpdatePostRequest)(nil), // 3: api.micro.moment.v1.UpdatePostRequest
	(*UpdatePostReply)(nil),   // 4: api.micro.moment.v1.UpdatePostReply
	(*DeletePostRequest)(nil), // 5: api.micro.moment.v1.DeletePostRequest
	(*DeletePostReply)(nil),   // 6: api.micro.moment.v1.DeletePostReply
	(*GetPostRequest)(nil),    // 7: api.micro.moment.v1.GetPostRequest
	(*GetPostReply)(nil),      // 8: api.micro.moment.v1.GetPostReply
	(*ListPostRequest)(nil),   // 9: api.micro.moment.v1.ListPostRequest
	(*ListPostReply)(nil),     // 10: api.micro.moment.v1.ListPostReply
	(*v1.User)(nil),           // 11: api.micro.user.v1.User
}
var file_api_micro_moment_v1_post_proto_depIdxs = []int32{
	11, // 0: api.micro.moment.v1.Post.user:type_name -> api.micro.user.v1.User
	0,  // 1: api.micro.moment.v1.CreatePostReply.post:type_name -> api.micro.moment.v1.Post
	0,  // 2: api.micro.moment.v1.GetPostReply.post:type_name -> api.micro.moment.v1.Post
	0,  // 3: api.micro.moment.v1.ListPostReply.items:type_name -> api.micro.moment.v1.Post
	1,  // 4: api.micro.moment.v1.PostService.CreatePost:input_type -> api.micro.moment.v1.CreatePostRequest
	3,  // 5: api.micro.moment.v1.PostService.UpdatePost:input_type -> api.micro.moment.v1.UpdatePostRequest
	5,  // 6: api.micro.moment.v1.PostService.DeletePost:input_type -> api.micro.moment.v1.DeletePostRequest
	7,  // 7: api.micro.moment.v1.PostService.GetPost:input_type -> api.micro.moment.v1.GetPostRequest
	9,  // 8: api.micro.moment.v1.PostService.ListHotPost:input_type -> api.micro.moment.v1.ListPostRequest
	9,  // 9: api.micro.moment.v1.PostService.ListLatestPost:input_type -> api.micro.moment.v1.ListPostRequest
	2,  // 10: api.micro.moment.v1.PostService.CreatePost:output_type -> api.micro.moment.v1.CreatePostReply
	4,  // 11: api.micro.moment.v1.PostService.UpdatePost:output_type -> api.micro.moment.v1.UpdatePostReply
	6,  // 12: api.micro.moment.v1.PostService.DeletePost:output_type -> api.micro.moment.v1.DeletePostReply
	8,  // 13: api.micro.moment.v1.PostService.GetPost:output_type -> api.micro.moment.v1.GetPostReply
	10, // 14: api.micro.moment.v1.PostService.ListHotPost:output_type -> api.micro.moment.v1.ListPostReply
	10, // 15: api.micro.moment.v1.PostService.ListLatestPost:output_type -> api.micro.moment.v1.ListPostReply
	10, // [10:16] is the sub-list for method output_type
	4,  // [4:10] is the sub-list for method input_type
	4,  // [4:4] is the sub-list for extension type_name
	4,  // [4:4] is the sub-list for extension extendee
	0,  // [0:4] is the sub-list for field type_name
}

func init() { file_api_micro_moment_v1_post_proto_init() }
func file_api_micro_moment_v1_post_proto_init() {
	if File_api_micro_moment_v1_post_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_api_micro_moment_v1_post_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Post); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_api_micro_moment_v1_post_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreatePostRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_api_micro_moment_v1_post_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreatePostReply); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_api_micro_moment_v1_post_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdatePostRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_api_micro_moment_v1_post_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdatePostReply); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_api_micro_moment_v1_post_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeletePostRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_api_micro_moment_v1_post_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeletePostReply); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_api_micro_moment_v1_post_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetPostRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_api_micro_moment_v1_post_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetPostReply); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_api_micro_moment_v1_post_proto_msgTypes[9].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListPostRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_api_micro_moment_v1_post_proto_msgTypes[10].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListPostReply); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_api_micro_moment_v1_post_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   11,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_api_micro_moment_v1_post_proto_goTypes,
		DependencyIndexes: file_api_micro_moment_v1_post_proto_depIdxs,
		MessageInfos:      file_api_micro_moment_v1_post_proto_msgTypes,
	}.Build()
	File_api_micro_moment_v1_post_proto = out.File
	file_api_micro_moment_v1_post_proto_rawDesc = nil
	file_api_micro_moment_v1_post_proto_goTypes = nil
	file_api_micro_moment_v1_post_proto_depIdxs = nil
}

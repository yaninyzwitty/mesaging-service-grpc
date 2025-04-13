package controllers

import (
	"context"

	"github.com/gocql/gocql"
	"github.com/yaninyzwitty/go-grpc-messaging/pb"
)

type MessagingController struct {
	session *gocql.Session
	pb.UnimplementedMessagingServiceServer
}

func NewMessagingController(session *gocql.Session) *MessagingController {
	return &MessagingController{
		session: session,
	}
}

func (c *MessagingController) GetUser(ctx context.Context, req *pb.GetUserRequest) (*pb.GetUserResponse, error) {
	return &pb.GetUserResponse{}, nil
}

func (c *MessagingController) GetUserByEmail(ctx context.Context, req *pb.GetUserRequestByEmail) (*pb.GetUserResponseByEmail, error) {
	return &pb.GetUserResponseByEmail{}, nil
}

func (c *MessagingController) GetUserPreferences(ctx context.Context, req *pb.GetUserPreferencesRequest) (*pb.GetUserPreferencesResponse, error) {
	return &pb.GetUserPreferencesResponse{}, nil
}

func (c *MessagingController) CreateChannel(ctx context.Context, req *pb.CreateChannelRequest) (*pb.CreateChannelResponse, error) {
	return &pb.CreateChannelResponse{}, nil
}

func (c *MessagingController) ListChannelsByCreatorId(ctx context.Context, req *pb.ListChannelsRequestByCreatorId) (*pb.ListChannelsResponseByCreatorId, error) {
	return &pb.ListChannelsResponseByCreatorId{}, nil
}

func (c *MessagingController) PostMessage(ctx context.Context, req *pb.PostMessageRequest) (*pb.PostMessageResponse, error) {
	return &pb.PostMessageResponse{}, nil
}

func (c *MessagingController) GetRecentMessages(ctx context.Context, req *pb.GetRecentMessagesRequest) (*pb.GetRecentMessagesResponse, error) {
	return &pb.GetRecentMessagesResponse{}, nil
}

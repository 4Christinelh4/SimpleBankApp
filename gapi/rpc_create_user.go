package gapi

import (
	"context"

	"github.com/kristine/simplebank/pb"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (server *Server) CreateUser(ctx context.Context, request *pb.CreateUserRequest) (*pb.CreateUserResponse, error){
	email := request.GetEmail()
	username := request.GetUsername()

	rsp := &pb.CreateUserResponse{
		User: &pb.User{
			Username: username,
			Email: email,
		},
	}

	return rsp, status.Errorf(codes.Unimplemented, "method not implemented")
}

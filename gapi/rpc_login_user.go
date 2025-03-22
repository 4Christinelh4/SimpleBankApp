package gapi

import (
	"context"

	"github.com/kristine/simplebank/pb"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)


func (server *Server) LoginUser(ctx context.Context, in *pb.LoginUserRequest) (*pb.LoginUserResponse, error) {
	user, err := server.store.GetUser(ctx, in.GetUsername())
	if err != nil {
		return nil, status.Errorf(codes.NotFound, "user not found")
	}

	// check if password matches
	session, err := server.store.CreateSession()
	if err != nil {
		return nil,  status.Errorf(codes.Internal, "unable to create session")
	}

	return &pb.LoginUserResponse{
		
	}, nil
}

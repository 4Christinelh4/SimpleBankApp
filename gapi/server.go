package gapi

import (
	db "github.com/kristine/simplebank/db/sqlc"
	"github.com/kristine/simplebank/pb"
	"github.com/kristine/simplebank/util"
)

type Server struct {
	pb.UnimplementedSimpleBankServer
	config util.Config
	store db.Store
}

func NewServer(config util.Config, store db.Store) (*Server, error) {
	return &Server {
		config: config,
		store: store,
	}, nil
}

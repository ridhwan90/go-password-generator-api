package grpcapi

import (
	"fmt"

	db "github.com/ridhwan90/Golang-grpc/db/sqlc"
	"github.com/ridhwan90/Golang-grpc/pb"
	"github.com/ridhwan90/Golang-grpc/token"
)

type Server struct {
	pb.UnimplementedGeneratePasswordServer
	db_query   *db.Queries
	tokenMaker token.Maker
}

func NewServer(db_query *db.Queries) (*Server, error) {
	tokenMaker, err := token.NewPasetoMaker("12345678901234567890123456789012")
	if err != nil {
		return nil, fmt.Errorf("cannot create token maker: %w", err)
	}
	server := &Server{
		db_query:   db_query,
		tokenMaker: tokenMaker,
	}

	return server, nil
}

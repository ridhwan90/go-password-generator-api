package grpcapi

import (
	"context"
	"database/sql"
	"time"

	"github.com/ridhwan90/Golang-grpc/pb"
	"github.com/ridhwan90/Golang-grpc/util"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (server *Server) LoginUser(ctx context.Context, req *pb.LoginUserRequest) (*pb.LoginUserResponse, error) {
	var duration time.Duration = 15 * time.Minute

	user, err := server.db_query.GetUserbyUsername(ctx, req.GetUsername())
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, status.Errorf(codes.NotFound, "username not found:", err)
		}
		return nil, status.Errorf(codes.Internal, "error on query username:", err)
	}

	err = util.CheckPassword(req.GetPassword(), user.HashedPassword)
	if err != nil {
		return nil, status.Errorf(codes.PermissionDenied, "Incorrect password")
	}

	accessToken, err := server.tokenMaker.CreateToken(user.Username, duration)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "Error generating access token:", err)
	}

	rsp := &pb.LoginUserResponse{
		AccessToken: accessToken,
		Username:    user.Username,
	}

	return rsp, nil
}

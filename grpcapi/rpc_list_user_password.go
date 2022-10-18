package grpcapi

import (
	"context"
	"database/sql"

	"github.com/google/uuid"
	"github.com/ridhwan90/Golang-grpc/pb"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/timestamppb"
)

func (server *Server) ListPasswordbyUser(ctx context.Context, req *pb.UsernameParams) (*pb.ListUserPassword, error) {

	user, err := server.db_query.GetUserbyUsername(ctx, req.GetUsername())
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, status.Errorf(codes.NotFound, "user not found")

		}
		return nil, status.Errorf(codes.Internal, "error query user:%s", err)
	}

	listUserPassword, err := server.db_query.ListPasswordbyUser(ctx, uuid.NullUUID{UUID: user.ID, Valid: true})
	if err != nil {
		return nil, status.Errorf(codes.Internal, "error query user:%s", err)
	}
	var resListUserPassword []*pb.UserPassword

	for _, userPassword := range listUserPassword {
		resUserPassword := pb.UserPassword{
			Id:                userPassword.ID.String(),
			UserUuid:          userPassword.UserUuid.UUID.String(),
			Site:              userPassword.Site,
			SiteUsername:      userPassword.SiteUsername.String,
			SiteEmail:         userPassword.SiteEmail.String,
			GeneratedPassword: userPassword.GeneratedPassword,
			CreatedAt:         timestamppb.New(userPassword.CreatedAt),
		}
		resListUserPassword = append(resListUserPassword, &resUserPassword)
	}

	rsp := &pb.ListUserPassword{
		UserPassword: resListUserPassword,
	}
	return rsp, nil
}

package grpcapi

import (
	"context"
	"database/sql"
	"math/rand"
	"time"

	"github.com/google/uuid"
	db "github.com/ridhwan90/Golang-grpc/db/sqlc"
	"github.com/ridhwan90/Golang-grpc/pb"
	"github.com/ridhwan90/Golang-grpc/util"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (server *Server) CreatePassword(ctx context.Context, req *pb.CreatePasswordRequest) (*pb.CreatePasswordResponse, error) {

	user, err := server.db_query.GetUserbyUsername(ctx, req.GetUsername())
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, status.Errorf(codes.NotFound, "username not found")
		}
		return nil, status.Errorf(codes.Internal, "error query username:", err)

	}

	rand.Seed(time.Now().Unix())
	generatedPassword := util.GeneratePassword(20, 2, 2, 2)

	arg := db.CreatePasswordParams{
		ID:                uuid.New(),
		UserUuid:          uuid.NullUUID{UUID: user.ID, Valid: true},
		Site:              req.GetSite(),
		SiteUsername:      sql.NullString{String: req.GetSiteUsername(), Valid: true},
		SiteEmail:         sql.NullString{String: req.GetSiteEmail(), Valid: true},
		GeneratedPassword: generatedPassword,
	}

	userPassword, err := server.db_query.CreatePassword(ctx, arg)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "error inserting user-password:", err)
	}

	var resSiteUsername string
	if userPassword.SiteUsername.Valid {
		resSiteUsername = userPassword.SiteUsername.String
	} else {
		resSiteUsername = ""
	}

	var resSiteEmail string
	if userPassword.SiteEmail.Valid {
		resSiteEmail = userPassword.SiteEmail.String
	} else {
		resSiteEmail = ""
	}

	rsp := &pb.CreatePasswordResponse{
		Username:          user.Username,
		Site:              userPassword.Site,
		SiteUsername:      resSiteUsername,
		SiteEmail:         resSiteEmail,
		GeneratedPassword: userPassword.GeneratedPassword,
	}

	return rsp, nil
}

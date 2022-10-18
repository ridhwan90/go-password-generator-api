package grpcapi

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/google/uuid"
	"github.com/lib/pq"
	db "github.com/ridhwan90/Golang-grpc/db/sqlc"
	"github.com/ridhwan90/Golang-grpc/pb"
	"github.com/ridhwan90/Golang-grpc/util"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (server *Server) CreateUser(ctx context.Context, req *pb.CreateUserRequest) (*pb.CreateUserResponse, error) {
	// var dbFirstName sql.NullString
	// firstName := req.UserInfo.GetFirstName()
	// if firstName == "" {
	// 	dbFirstName.String = firstName
	// 	dbFirstName.Valid = false
	// } else {
	// 	dbFirstName.String = firstName
	// 	dbFirstName.Valid = true
	// }

	// var dbLastName sql.NullString
	// lastName := req.UserInfo.GetLastName()
	// if lastName == "" {
	// 	dbLastName.String = lastName
	// 	dbLastName.Valid = false
	// } else {
	// 	dbLastName.String = lastName
	// 	dbLastName.Valid = true
	// }

	// var dbPhoneNumber sql.NullString
	// phoneNumber := req.UserInfo.GetPhoneNumber()
	// if phoneNumber == "" {
	// 	dbPhoneNumber.String = phoneNumber
	// 	dbPhoneNumber.Valid = false
	// } else {
	// 	dbPhoneNumber.String = phoneNumber
	// 	dbPhoneNumber.Valid = true
	// }

	// var dbEmail sql.NullString
	// email := req.UserInfo.GetEmail()
	// if email == "" {
	// 	dbEmail.String = email
	// 	dbEmail.Valid = false
	// } else {
	// 	dbEmail.String = email
	// 	dbEmail.Valid = true
	// }

	arg := db.CreateUserInfoParams{
		ID:          uuid.New(),
		FirstName:   sql.NullString{String: req.UserInfo.GetFirstName(), Valid: true},
		LastName:    sql.NullString{String: req.UserInfo.GetLastName(), Valid: true},
		PhoneNumber: sql.NullString{String: req.UserInfo.GetPhoneNumber(), Valid: true},
		Email:       sql.NullString{String: req.UserInfo.GetEmail(), Valid: true},
	}

	userInfo, err := server.db_query.CreateUserInfo(ctx, arg)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "failed to insert userInfo: %s", err)
	}

	hashedPassword, err := util.HashPassword(req.GetPassword())
	if err != nil {
		return nil, status.Errorf(codes.Internal, "failed to hashed password: %s", err)
	}

	arg2 := db.CreateUserParams{
		ID:             uuid.New(),
		Username:       req.GetUsername(),
		HashedPassword: hashedPassword,
		UserinfoUuid:   uuid.NullUUID{UUID: userInfo.ID, Valid: true},
	}

	user, err := server.db_query.CreateUser(ctx, arg2)
	if err != nil {
		if pqErr, ok := err.(*pq.Error); ok {
			fmt.Println(pqErr.Code.Name())
			switch pqErr.Code.Name() {
			case "unique_violation":
				return nil, status.Errorf(codes.AlreadyExists, "username already exist: %s", err)
			}
		}
		return nil, status.Errorf(codes.Internal, "failed to create user: %s", err)
	}

	var resFirstName string
	if userInfo.FirstName.Valid {
		resFirstName = userInfo.FirstName.String
	} else {
		resFirstName = ""
	}

	var resLastName string
	if userInfo.FirstName.Valid {
		resLastName = userInfo.LastName.String
	} else {
		resLastName = ""
	}

	var resPhoneNumber string
	if userInfo.FirstName.Valid {
		resPhoneNumber = userInfo.PhoneNumber.String
	} else {
		resPhoneNumber = ""
	}

	var resEmail string
	if userInfo.FirstName.Valid {
		resEmail = userInfo.Email.String
	} else {
		resEmail = ""
	}

	rsp := pb.CreateUserResponse{
		Id:       user.ID.String(),
		Username: user.Username,
		UserInfo: &pb.UserInfo{
			FirstName:   resFirstName,
			LastName:    resLastName,
			PhoneNumber: resPhoneNumber,
			Email:       resEmail,
		},
	}

	return &rsp, nil
}

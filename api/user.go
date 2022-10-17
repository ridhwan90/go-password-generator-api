package api

import (
	"database/sql"
	"fmt"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/lib/pq"
	db "github.com/ridhwan90/Golang-grpc/db/sqlc"
	"github.com/ridhwan90/Golang-grpc/util"
)

type userInfoAPI struct {
	FirstName   *string `json:"first_name"`
	LastName    *string `json:"last_name"`
	PhoneNumber *string `json:"phone_number"`
	Email       *string `json:"email"`
}

type createUserRequest struct {
	Username string      `json:"username" binding:"required"`
	Password string      `json:"password" binding:"required,min=8"`
	UserInfo userInfoAPI `json: "userinfo"`
}

type UserResponse struct {
	ID       uuid.UUID   `json:"id"`
	Username string      `json:"username"`
	UserInfo userInfoAPI `json: "userinfo"`
}

func (server *Server) createUser(ctx *gin.Context) {
	var req createUserRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	var firstName sql.NullString
	if req.UserInfo.FirstName == nil {
		firstName = sql.NullString{
			String: "",
			Valid:  false,
		}
	} else {
		firstName = sql.NullString{
			String: *req.UserInfo.FirstName,
			Valid:  true,
		}
	}

	var lastName sql.NullString
	if req.UserInfo.FirstName == nil {
		lastName = sql.NullString{
			String: "",
			Valid:  false,
		}
	} else {
		lastName = sql.NullString{
			String: *req.UserInfo.LastName,
			Valid:  true,
		}
	}

	var phoneNumber sql.NullString
	if req.UserInfo.FirstName == nil {
		phoneNumber = sql.NullString{
			String: "",
			Valid:  false,
		}
	} else {
		phoneNumber = sql.NullString{
			String: *req.UserInfo.PhoneNumber,
			Valid:  true,
		}
	}

	var email sql.NullString
	if req.UserInfo.FirstName == nil {
		email = sql.NullString{
			String: "",
			Valid:  false,
		}
	} else {
		email = sql.NullString{
			String: *req.UserInfo.Email,
			Valid:  true,
		}
	}

	arg := db.CreateUserInfoParams{
		ID:          uuid.New(),
		FirstName:   firstName,
		LastName:    lastName,
		PhoneNumber: phoneNumber,
		Email:       email,
	}

	userInfo, err := server.db_query.CreateUserInfo(ctx, arg)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
	}

	hashedPassword, err := util.HashPassword(req.Password)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
	}

	arg2 := db.CreateUserParams{
		ID:             uuid.New(),
		Username:       req.Username,
		HashedPassword: hashedPassword,
		UserinfoUuid:   uuid.NullUUID{UUID: userInfo.ID, Valid: true},
	}

	user, err := server.db_query.CreateUser(ctx, arg2)
	if err != nil {
		if pqErr, ok := err.(*pq.Error); ok {
			fmt.Println(pqErr.Code.Name())
			switch pqErr.Code.Name() {
			case "unique_violation":
				ctx.JSON(http.StatusForbidden, errorResponse(err))
				return
			}
		}
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	var resFirstName *string
	if userInfo.FirstName.Valid {
		resFirstName = &userInfo.FirstName.String
	} else {
		resFirstName = nil
	}

	var resLastName *string
	if userInfo.FirstName.Valid {
		resLastName = &userInfo.LastName.String
	} else {
		resLastName = nil
	}

	var resPhoneNumber *string
	if userInfo.FirstName.Valid {
		resPhoneNumber = &userInfo.PhoneNumber.String
	} else {
		resPhoneNumber = nil
	}

	var resEmail *string
	if userInfo.FirstName.Valid {
		resEmail = &userInfo.Email.String
	} else {
		resEmail = nil
	}

	response := UserResponse{
		ID:       user.ID,
		Username: user.Username,
		UserInfo: userInfoAPI{
			FirstName:   resFirstName,
			LastName:    resLastName,
			PhoneNumber: resPhoneNumber,
			Email:       resEmail,
		},
	}

	ctx.JSON(http.StatusCreated, response)
}

type loginUserRequest struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required,min=8"`
}

type loginUserResponse struct {
	AccessToken string `json:"acces_token"`
	Username    string `json:"username"`
}

func (server *Server) loginUser(ctx *gin.Context) {
	var req loginUserRequest
	var duration time.Duration = 15 * time.Minute
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	user, err := server.db_query.GetUserbyUsername(ctx, req.Username)
	if err != nil {
		if err == sql.ErrNoRows {
			ctx.JSON(http.StatusNotFound, errorResponse(err))
			return
		}
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	err = util.CheckPassword(req.Password, user.HashedPassword)
	if err != nil {
		ctx.JSON(http.StatusUnauthorized, errorResponse(err))
		return
	}

	accessToken, err := server.tokenMaker.CreateToken(user.Username, duration)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
	}

	rsp := loginUserResponse{
		AccessToken: accessToken,
		Username:    user.Username,
	}

	ctx.JSON(http.StatusOK, rsp)
}

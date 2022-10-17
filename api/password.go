package api

import (
	"database/sql"
	"math/rand"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	db "github.com/ridhwan90/Golang-grpc/db/sqlc"
	"github.com/ridhwan90/Golang-grpc/util"
)

type createPasswordRequest struct {
	Username     string  `json:"username" binding:"required"`
	Site         string  `json:"site" binding:"required"`
	SiteUsername *string `json:"site_username"`
	SiteEmail    *string `json:"site_email"`
}

type createPasswordResponse struct {
	Username          string  `json:"username"`
	Site              string  `json:"site"`
	SiteUsername      *string `json:"site_username"`
	SiteEmail         *string `json:"site_email"`
	GeneratedPassword string  `json:"generated_password"`
}

func (server *Server) CreatePassword(ctx *gin.Context) {
	var req createPasswordRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	// fmt.Println(*req.SiteEmail, *req.SiteUsername)

	user, err := server.db_query.GetUserbyUsername(ctx, req.Username)
	if err != nil {
		if err == sql.ErrNoRows {
			ctx.JSON(http.StatusNotFound, errorResponse(err))
			return
		}
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}
	rand.Seed(time.Now().Unix())
	generatedPassword := util.GeneratePassword(20, 2, 2, 2)

	var siteUsername sql.NullString
	if req.SiteUsername == nil {
		siteUsername.String = ""
		siteUsername.Valid = false
	} else {
		siteUsername.String = *req.SiteUsername
		siteUsername.Valid = true
	}

	var siteEmail sql.NullString
	if req.SiteEmail == nil {
		siteEmail.String = ""
		siteEmail.Valid = false
	} else {
		siteEmail.String = *req.SiteEmail
		siteEmail.Valid = true
	}

	arg := db.CreatePasswordParams{
		ID:                uuid.New(),
		UserUuid:          uuid.NullUUID{UUID: user.ID, Valid: true},
		Site:              req.Site,
		SiteUsername:      siteUsername,
		SiteEmail:         siteEmail,
		GeneratedPassword: generatedPassword,
	}

	userPassword, err := server.db_query.CreatePassword(ctx, arg)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	var resSiteUsername *string
	if userPassword.SiteUsername.Valid {
		resSiteUsername = &userPassword.SiteUsername.String
	} else {
		resSiteUsername = nil
	}

	var resSiteEmail *string
	if userPassword.SiteEmail.Valid {
		resSiteEmail = &userPassword.SiteEmail.String
	} else {
		resSiteEmail = nil
	}

	rsp := createPasswordResponse{
		Username:          user.Username,
		Site:              userPassword.Site,
		SiteUsername:      resSiteUsername,
		SiteEmail:         resSiteEmail,
		GeneratedPassword: userPassword.GeneratedPassword,
	}

	ctx.JSON(http.StatusCreated, rsp)
}

func (server *Server) listPasswordbyUser(ctx *gin.Context) {
	username, ok := ctx.GetQuery("username")
	if ok == false {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "username query parameter not found"})
		return
	}

	user, err := server.db_query.GetUserbyUsername(ctx, username)
	if err != nil {
		if err == sql.ErrNoRows {
			ctx.JSON(http.StatusNotFound, errorResponse(err))
			return
		}
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	listUserPassword, err := server.db_query.ListPasswordbyUser(ctx, uuid.NullUUID{UUID: user.ID, Valid: true})
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, listUserPassword)

}

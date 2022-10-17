package api

import (
	"fmt"

	"github.com/ridhwan90/Golang-grpc/token"

	"github.com/gin-gonic/gin"
	db "github.com/ridhwan90/Golang-grpc/db/sqlc"
)

type Server struct {
	db_query   *db.Queries
	tokenMaker token.Maker
	router     *gin.Engine
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
	router := gin.Default()

	router.POST("/user", server.createUser)
	router.POST("/user/login", server.loginUser)

	authRoutes := router.Group("/").Use(authMiddleware(server.tokenMaker))

	authRoutes.POST("/user/generate-password", server.CreatePassword)
	authRoutes.GET("/user/list-user-password", server.listPasswordbyUser)

	server.router = router
	return server, nil
}

func (server *Server) Start(address string) error {
	return server.router.Run(address)
}

func errorResponse(err error) gin.H {
	return gin.H{"error": err.Error()}
}

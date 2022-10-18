package main

import (
	"database/sql"
	"log"
	"net"
	"os"
	"sync"

	_ "github.com/lib/pq"
	"github.com/ridhwan90/Golang-grpc/api"
	db "github.com/ridhwan90/Golang-grpc/db/sqlc"
	"github.com/ridhwan90/Golang-grpc/grpcapi"
	"github.com/ridhwan90/Golang-grpc/pb"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

const (
	dbDriver              = "postgres"
	dbSource              = "postgresql://postgres:postgres@postgres:5432/postgres?sslmode=disable"
	dbDriverProd          = "postgres"
	serverAddress         = "0.0.0.0:"
	GrpcServerAddress     = "0.0.0.0:9090"
	GrpcProdServerAddress = "0.0.0.0:"
)

func main() {
	var conn *sql.DB
	var err error
	if os.Getenv("ENV") == "PROD" {
		conn, err = sql.Open(dbDriverProd, os.Getenv("DATABASE_URL"))
		if err != nil {
			log.Fatal("cannot connect to db:", err)
		}
	} else {
		conn, err = sql.Open(dbDriver, dbSource)
		if err != nil {
			log.Fatal("cannot connect to db:", err)
		}
	}

	db_query := db.New(conn)
	wg := new(sync.WaitGroup)
	wg.Add(2)
	go runGrpcServer(db_query, wg)
	go runGinServer(db_query, wg)
	wg.Wait()
}

func runGrpcServer(db_query *db.Queries, wg *sync.WaitGroup) {
	server, err := grpcapi.NewServer(db_query)
	if err != nil {
		log.Fatal("cannot create server:", err)
	}

	grpcServer := grpc.NewServer()
	pb.RegisterGeneratePasswordServer(grpcServer, server)
	reflection.Register(grpcServer)

	var listner net.Listener

	if os.Getenv("ENV") == "PROD" {
		var port string = os.Getenv("PORT")
		listner, err = net.Listen("tcp", GrpcProdServerAddress+port)
		if err != nil {
			log.Fatal("cannot create listener", err)
		}

	} else {
		listner, err = net.Listen("tcp", GrpcServerAddress)
		if err != nil {
			log.Fatal("cannot create listener", err)
		}
	}

	log.Printf("start gRPC server at %s", listner.Addr().String())
	err = grpcServer.Serve(listner)
	if err != nil {
		log.Fatal("cannot start gRPC server")
	}
	wg.Done()
}

func runGinServer(db_query *db.Queries, wg *sync.WaitGroup) {
	server, err := api.NewServer(db_query)
	if err != nil {
		log.Fatal("cannot create server:", err)
	}
	if os.Getenv("ENV") == "PROD" {
		var port string = os.Getenv("PORT")
		err = server.Start(serverAddress + port)
		if err != nil {
			log.Fatal("cannot start server:", err)
		}
	} else {
		err = server.Start("0.0.0.0:8080")
		if err != nil {
			log.Fatal("cannot start server:", err)
		}
	}
	wg.Done()
}

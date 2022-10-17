package main

import (
	"database/sql"
	"log"
	"os"

	_ "github.com/lib/pq"
	"github.com/ridhwan90/Golang-grpc/api"
	db "github.com/ridhwan90/Golang-grpc/db/sqlc"
)

const (
	dbDriver      = "postgres"
	dbSource      = "postgresql://postgres:postgres@db:5432/postgres?sslmode=disable"
	dbDriverProd  = "postgres"
	serverAddress = "0.0.0.0:8080"
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
	server, err := api.NewServer(db_query)
	if err != nil {
		log.Fatal("cannot create server:", err)
	}

	err = server.Start(serverAddress)
	if err != nil {
		log.Fatal("cannot start server:", err)
	}
}

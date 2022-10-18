migrateup:
	migrate -path db/migration -database "postgresql://postgres:postgres@db:5432/postgres?sslmode=disable" -verbose up

migratedown:
	migrate -path db/migration -database "postgresql://postgres:postgres@db:5432/postgres?sslmode=disable" -verbose down

sqlc:
	sqlc generate

proto:
	rm -f pb/*.go
	protoc --proto_path=proto --go_out=pb --go_opt=paths=source_relative \
    --go-grpc_out=pb --go-grpc_opt=paths=source_relative \
    proto/*.proto

.PHONY: migrateup migratedown proto
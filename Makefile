migrateup:
	migrate -path db/migration -database "postgresql://postgres:postgres@db:5432/postgres?sslmode=disable" -verbose up

migratedown:
	migrate -path db/migration -database "postgresql://postgres:postgres@db:5432/postgres?sslmode=disable" -verbose down

sqlc:
	sqlc generate

.PHONY: migrateup migratedown
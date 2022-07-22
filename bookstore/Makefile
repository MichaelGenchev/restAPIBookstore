createdb:
	 docker exec -it postgres13 createdb --username=root --owner=root bookshelf

dropdb:
	 docker exec -it postgres13 dropdb bookshelf

postgres:
	docker run --name postgres13 -p 5432:5432  -e POSTGRES_USER=root -e POSTGRES_PASSWORD=secret -d postgres:13-alpine

migrateup:
	migrate -path db/migration -database "postgresql://root:secret@localhost:5432/bookshelf?sslmode=disable" -verbose up

migratedown:
	 migrate -path db/migration -database "postgresql://root:secret@localhost:5432/bookshelf?sslmode=disable" -verbose down

sqlc:
	sqlc generate

test:
	go test -v -cover ./...
.PHONY: createdb postgres dropdb migrateup migratedown sqlc test
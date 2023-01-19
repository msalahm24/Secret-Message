server:
	go run main.go

test:
	go test -v -cover ./...

migrateup:
	migrate -path db/migration -database "postgres://root:123@localhost:5432/secretMessage?sslmode=disable" -verbose up

migratedown:
	migrate -path db/migration -database "postgres://root:123@localhost:5432/secretMessage?sslmode=disable" -verbose down

createdb:
	docker exec -it postgres createdb --username=root --owner=root secretMessage

dropdb:
	docker exec -it postgres dropdb secretMessage


.PHONY: server test migrateup migratedown createdb dropdb
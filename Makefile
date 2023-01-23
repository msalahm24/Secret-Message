server:
	go run main.go

test:
	go test -v -cover ./...

migrateup:
	migrate -path db/migration -database "postgres://root:secret@localhost:5432/secretMessage?sslmode=disable" -verbose up

migratedown:
	migrate -path db/migration -database "postgres://root:secret@localhost:5432/secretMessage?sslmode=disable" -verbose down

createdb:
	sudo docker exec -it postgres createdb --username=root --owner=root secretMessage

dropdb:
	sudo docker exec -it postgres dropdb secretMessage

sqlc:
	sqlc generate


.PHONY: server test migrateup migratedown createdb dropdb sqlc
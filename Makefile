

postgres:
	docker run --name songs -p 5432:5432 -e POSTGRES_USER=root -e POSTGRES_PASSWORD=secret -d postgres

createdb:
	docker exec -it songs createdb --username=root --owner=root songs

migrate:
	migrate create -ext sql -dir db/migration -seq $(name)
	
migrateup:
	migrate -path db/migration -database 'postgresql://root:secret@localhost:5432/songs?sslmode=disable' -verbose up

migratedown:
	migrate -path db/migration -database "postgresql://root:secret@localhost:5432/songs?sslmode=disable" -verbose down

sqlc:
	sqlc generate

server:
	go run ./cmd/api


.PHONY: server postgres createdb  migrateup migratedown sqlc






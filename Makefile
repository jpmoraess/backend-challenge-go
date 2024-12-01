postgres:
	docker run --name postgresdb -p 5432:5432 -e POSTGRES_USER=postgres -e POSTGRES_PASSWORD=postgres -d postgres:latest

createdb:
	docker exec -it postgresdb createdb --username=postgres --owner=postgres backend_challenge

dropdb:
	docker exec -it postgresdb dropdb backend_challenge

migrateup:
	migrate -path db/migration -database "postgresql://postgres:postgres@localhost:5432/backend_challenge?sslmode=disable" -verbose up

migratedown:
	migrate -path db/migration -database "postgresql://postgres:postgres@localhost:5432/backend_challenge?sslmode=disable" -verbose down


.PHONY: postgres createdb dropdb migrateup migratedown
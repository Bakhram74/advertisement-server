

docker_run:
	docker run --name advertisement -e POSTGRES_USER=root -e POSTGRES_PASSWORD=secret -p 5432:5432 -d postgres

createDB:
	docker exec -it advertisement createdb --username=root --owner=root advertisement

dropDB:
	docker exec -it advertisement dropdb advertisement

docker_exec:
	docker exec -it advertisement psql -U root

migrate:
	migrate create -ext sql -dir db/schema/migration -seq init_schema

migrateUP:
	migrate -path db/schema/migration -database "postgresql://root:secret@localhost:5432/advertisement?sslmode=disable" -verbose up

migrateDown:
	migrate -path db/schema/migration -database "postgresql://root:secret@localhost:5432/advertisement?sslmode=disable" -verbose down

sqlc:
	sqlc generate

PHONY:  docker_run createDB dropDB docker_exec migrate migrateUP migrateDown sqlc
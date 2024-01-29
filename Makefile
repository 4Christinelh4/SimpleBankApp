postgres:
	docker run --name postgres14 -p 5435:5432 -e POSTGRES_USER=root -e POSTGRES_PASSWORD=secret1 -d postgres:14-alpine

createdb:
	docker exec -it postgres14 createdb --username=root --owner=root simple_bank

dropdb:
	docker exec -it postgres14 dropdb simple_bank

migratedbup:
	migrate -path db/migration -database "postgresql://root:secret1@localhost:5435/simple_bank?sslmode=disable" -verbose up

sqlc:
	sqlc generate

.PHONY: postgres creatdb migratedbup sqlc
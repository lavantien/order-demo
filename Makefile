makeFileDir := $(dir $(abspath $(lastword $(MAKEFILE_LIST))))

postgres:
	- docker run --name postgres -p 5432:5432 -e POSTGRES_USER=root -e POSTGRES_PASSWORD=secret -d postgres:alpine
	- docker start postgres

createdb:
	docker exec -it postgres createdb --username=root --owner=root order_demo

dropdb:
	docker exec -it postgres dropdb order_demo

migrateup:
	migrate -path db/migration -database "postgresql://root:secret@localhost:5432/order_demo?sslmode=disable" -verbose up

migratedown:
	migrate -path db/migration -database "postgresql://root:secret@localhost:5432/order_demo?sslmode=disable" -verbose down

sqlc:
	docker run --rm -v $(makeFileDir):/src -w /src kjconroy/sqlc generate

test:
	go test -v -cover ./...

server:
	go run main.go

mock:
	mockgen -package mockdb -destination db/mock/store.go order-demo/db/sqlc Store

.PHONY: postgres createdb dropdb migrateup migratedown sqlc test server mock

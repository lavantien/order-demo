makeFileDir := $(dir $(abspath $(lastword $(MAKEFILE_LIST))))

network:
	docker network create order-network

postgres:
	- docker run --name postgres --network order-network -p 5432:5432 -e POSTGRES_USER=root -e POSTGRES_PASSWORD=secret -d postgres:alpine
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
	go test -v -count=1 -race -cover ./...

server:
	go run main.go

mock:
	mockgen -package mockdb -destination db/mock/store.go order-demo/db/sqlc Store

build:
	docker build -t order-demo:latest .

clean:
	- docker stop order-demo && docker rm order-demo
	- docker stop postgres && docker rm postgres
	- docker rmi order-demo_api

order-demo:
	- docker run --name order-demo --network order-network -p 8080:8080 -e GIN_MODE=release -e DB_SOURCE="postgresql://root:secret@postgres:5432/order_demo?sslmode=disable" -d order-demo:latest
	- docker start order-demo

.PHONY: postgres createdb dropdb migrateup migratedown sqlc test server mock build clean order-demo all

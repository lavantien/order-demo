# Production-ready Simple Order Service Demo

![ci-test](https://github.com/lavantien/order-demo/actions/workflows/ci.yml/badge.svg?branch=main)

## 4 Requirements and 4 Bonuses

1. [X] View list of products
2. [X] Add/Remove product from cart
3. [X] Create new order with payment
4. [X] Users can login, sign up

----

1. [X] Cleanly structured and CI integration
2. [X] Well-documented
3. [X] Well-tested
4. [X] Containerized

### **Authorization Rules**

```go
// From the server's routes, add authentication middleware, using Paseto Token with local symmetric encryption
tokenMaker, _ := token.NewPasetoMaker(config.TokenSymmetricKey)
authRoutes := router.Group("/").Use(authMiddleware(tokenMaker))
// These 6 endpoints need authorization, implements in their handlers repsectively
authRoutes.GET("/users", server.listUsers)
authRoutes.POST("/products", server.createProduct)
authRoutes.POST("/products/cart/add", server.addToCart)
authRoutes.POST("/products/cart/remove", server.removeFromCart)
authRoutes.GET("/orders", server.listOrders)
authRoutes.POST("/orders", server.createOrder)
```

0. [X] An admin user `{"username":"admin",password:"secret"}` have all the powers, is created upon migrate up, log in as admin to test all of the endpoints
1. [X] A logged-in user can only view their own user's details
2. [X] A logged-in user can create a product
3. [X] A logged-in user can add a product to cart
4. [X] A logged-in user can remove a product from card
5. [X] A logged-in user can create an order
6. [X] A logged-in user can only view a the orders that they've made

### **API Endpoints**

<details>
	<summary>See details</summary>

See **Booting Up** running and testing instructions in the section below first, and then continue:

![Server running](/resources/readme/server-running.png "Server running")

### Rqm4.1: Create user via endpoint

```bash
curl "http://localhost:8080/users" -H "Content-Type: application/json" -d '{"username":"tien1","full_name":"Tien La","email":"tien@email.com","password":"secret"}' | jq
# Should return
{
  "username": "tien1",
  "full_name": "Tien La",
  "email": "tien@email.com",
  "password_change_at": "0001-01-01T00:00:00Z",
  "created_at": "2021-12-26T18:24:12.73219Z"
}
```

### Rqm4.2: User login with wrong password

```bash
curl "http://localhost:8080/users/login" -H "Content-Type: application/json" -d '{"username":"tien1","password":"abc123"}' | jq
# Should return (401)
{
  "error": "crypto/bcrypt: hashedPassword is not the hash of the given password"
}
```

### Rqm4.3: User log in

After logged-in, copy the `access_token` to the `TOKEN` variable to be use in appropriated endpoints' `-H 'Authorization: Bearer ...'`. Or use Postman/Insomnia/vscode-rest, ...

If having the error `token has expired`, log in again

```bash
curl "http://localhost:8080/users/login" -H "Content-Type: application/json" -d '{"username":"tien1","password":"secret"}' | jq
# Should return
{
  "access_token": "v2.local.iMAQ5gAOXIWxvl446dWq_Z7D7tV_J9MzRQov7HXEi0cbXFU0ZBhsR2GsHlhAeyMbpKMXH8ie-XTW6aKnIFgEfxZNnWXpsUl_QVTsuum1X2H_97UA0iqyP4NEG4JvWdqtrQ30HFN-BdvvXle98eUnKbCFn-28ot60kMGotwRySXJvI-LKCl04crKV31C6yjmKsj-2kPQ14d7eWM7bW8TyDm2DkPy5ZyrmrUTptk3LPLKZSCHPFDa9nfVwO_u4DcG-XZh_Nt6QB3NRTvSwVw.bnVsbA",
  "user": {
    "username": "tien1",
    "full_name": "Tien La",
    "email": "tien@email.com",
    "password_change_at": "0001-01-01T00:00:00Z",
    "created_at": "2021-12-26T18:24:12.73219Z"
  }
}

# Set TOKEN variable
TOKEN='v2.local.iMAQ5gAOXIWxvl446dWq_Z7D7tV_J9MzRQov7HXEi0cbXFU0ZBhsR2GsHlhAeyMbpKMXH8ie-XTW6aKnIFgEfxZNnWXpsUl_QVTsuum1X2H_97UA0iqyP4NEG4JvWdqtrQ30HFN-BdvvXle98eUnKbCFn-28ot60kMGotwRySXJvI-LKCl04crKV31C6yjmKsj-2kPQ14d7eWM7bW8TyDm2DkPy5ZyrmrUTptk3LPLKZSCHPFDa9nfVwO_u4DcG-XZh_Nt6QB3NRTvSwVw.bnVsbA'
```

### Rqm4.4: List users

```bash
curl "http://localhost:8080/users?page_id=1&page_size=5" -H "Authorization: Bearer $TOKEN" | jq
# Should return
[
  {
    "username": "tien1",
    "full_name": "Tien La",
    "email": "tien@email.com",
    "password_change_at": "0001-01-01T00:00:00Z",
    "created_at": "2021-12-26T18:24:12.73219Z"
  }
]
```

### Rqm1: View list of products

```bash
curl "http://localhost:8080/products?page_id=1&page_size=3" | jq
# Should return
[
  {
    "id": 1,
    "name": "ndomrf",
    "cost": 789,
    "quantity": 4,
    "created_at": "2021-12-26T18:20:27.991534Z"
  },
  {
    "id": 2,
    "name": "qsuwja",
    "cost": 913,
    "quantity": 5,
    "created_at": "2021-12-26T18:20:28.05339Z"
  },
  {
    "id": 3,
    "name": "jesmsw",
    "cost": 754,
    "quantity": 9,
    "created_at": "2021-12-26T18:20:28.11771Z"
  }
]
```

### Rqm2.1: Add product to cart

```bash
curl "http://localhost:8080/products/cart/add" -H "Authorization: Bearer $TOKEN" -H 'Content-Type: application/json' -d '{"product_id":1,"quantity":2}' | jq
# Should return
{
  "product": {
    "id": 1,
    "name": "ndomrf",
    "cost": 789,
    "quantity": 2,
    "created_at": "2021-12-26T18:20:27.991534Z"
  }
}
```

### Rqm2.2: Remove product from cart

```bash
curl "http://localhost:8080/products/cart/remove" -H "Authorization: Bearer $TOKEN" -H 'Content-Type: application/json' -d '{"product_id":1,"quantity":2}' | jq
# Should return
{
  "product": {
    "id": 1,
    "name": "ndomrf",
    "cost": 789,
    "quantity": 6,
    "created_at": "2021-12-26T18:20:27.991534Z"
  }
}
```

### Rqm3.1: Create new order with payment

```bash
curl "http://localhost:8080/orders" -H "Authorization: Bearer $TOKEN" -H 'Content-Type: application/json' -d '{"user_id":1,"product_id":1,"quantity":2}' | jq
# Should return
{
  "user": {
    "username": "tien1",
    "full_name": "Tien La",
    "email": "tien@email.com",
    "password_change_at": "0001-01-01T00:00:00Z",
    "created_at": "2021-12-26T18:24:12.73219Z"
  },
  "product": {
    "id": 1,
    "name": "ndomrf",
    "cost": 789,
    "quantity": 2,
    "created_at": "2021-12-26T18:20:27.991534Z"
  },
  "order": {
    "id": 24,
    "owner": "tien1",
    "product_id": 1,
    "quantity": 2,
    "price": 1578,
    "created_at": "2021-12-26T18:26:26.003027Z"
  }
}
```

### Rqm3.2: Check the result

```bash
curl "http://localhost:8080/orders?page_id=1&page_size=5" -H "Authorization: Bearer $TOKEN" | jq
# Should return
[
  {
    "id": 24,
    "owner": "tien1",
    "product_id": 1,
    "quantity": 2,
    "price": 1578,
    "created_at": "2021-12-26T18:26:26.003027Z"
  }
]
```

</details>

## Database UML

![Database UML](/resources/readme/order-demo.png "Database UML")

## Technology Stack

- **Go 1.17**: *Leverage the standard libraries as much as possible*
- **SQLc**: *Generates efficient native SQL CRUD code*
- **PostgreSQL**: *RDBMS of choice because of faster read due to its indexing model and safer transaction with better isolation levels handling*
- **Gin**: *Fast and have respect for native net/http API*
- **Paseto Token**: *Better choice than JWT because of enforcing better cryptographic standards and debloated of useless information*
- **JWT Token**: *Also implemented to demonstrate the decoupility*
- **Golang-Migrate**: *Efficient schema generating, up/down migrating*
- **GoMock**: *Generates mocks of about anything*
- **Docker** + **Docker-Compose**: *Containerization, what else to say ...*
- **Github Actions CI**: *Make sure we don't push trash code into the codebase*
- **Viper**: *Add robustness to configurations*

## Philosophy and Architecture

- **Adaptive Minimalism**: *I always keep it as simple as possible, but with a highly decoupled structure we ensure high adaptivity and extensibility, on top of that minimal solid head start. Things are implement only when they're absolutely needed*

## Booting Up

### Non-docker way

- Spin up a PostgreSQL instance, run createdb and migrateup:

```bash
make network

make postgres

make createdb

make migrateup
```

- Run test:

```bash
# go get github.com/golang/mock/mockgen/model

# make mock

make test
```

- Run server:

```bash
make server

# make migratedown
```

### Docker way

```bash
make network

make postgres

make createdb

make migrateup

make test

make order-demo

# Rebuild server image
# make build

# make clean
```

### Docker Compose way

```bash
docker compose up -d

# docker-compose down

# docker rmi order-demo_api
```

## Development Infrastructure Setup

### Helpful Commands

```bash
# Update Go toolings
go get -u
go mod tidy

# Spin up a container for local development, for example postgres
# The default database will be root, the same name as POSTGRES_USER
docker run --name postgres -p 5432:5432 -e POSTGRES_USER=root -e POSTGRES_PASSWORD=secret -d postgres:alpine

# To access postgres psql
docker exec -it postgres psql -U root

# To access container console and run postgres commands
docker exec -it postgres /bin/sh

# To quit the console
\q

# To view its logs
docker logs postgres

# To stop it
docker stop postgres

# To just run it again
docker start postgres

# To remove it completely
docker rm postgres
```

### Tooling Installation Guide

<details>
    <summary>Expand</summary>

- [**Golang**](https://go.dev/doc/install):

```bash
# Go to go.dev/dl and download a binary, in this example it's version 1.17.5

sudo rm -rf /usr/local/go && sudo tar -C /usr/local -xzf go1.17.5.linux-amd64.tar.gz

# Add these below to your .bashrc or .zshrc
export GOPATH=/home/<username>/go
export GOBIN=/home/<username>/go/bin
export PATH=$PATH:/usr/local/go/bin
export PATH=$PATH:$GOBIN
```

- [**Docker**](https://docs.docker.com/engine/install/ubuntu/):

```bash
sudo apt remove docker docker-engine docker.io containerd runc

sudo apt update

sudo apt install apt-transport-https ca-certificates curl gnupg lsb-release software-properties-common

curl -fsSL https://download.docker.com/linux/ubuntu/gpg | sudo gpg --dearmor -o /usr/share/keyrings/docker-archive-keyring.gpg

echo \
  "deb [arch=amd64 signed-by=/usr/share/keyrings/docker-archive-keyring.gpg] https://download.docker.com/linux/ubuntu \
  $(lsb_release -cs) stable" | sudo tee /etc/apt/sources.list.d/docker.list > /dev/null

sudo apt update

apt-cache policy docker-ce

sudo apt install docker-ce docker-ce-cli containerd.io

sudo usermod -aG docker $USER

newgrp docker

# Restart the machine then test the installation

docker run hello-world

# On older system you also need to activate the services

sudo systemctl enable docker.service

sudo systemctl enable containerd.service
```

- [**Docker-Compose**](https://docs.docker.com/compose/install/):

```bash
# Check their github repo for latest version number
sudo curl -L "https://github.com/docker/compose/releases/download/v2.2.2/docker-compose-linux-x86_64" -o /usr/local/bin/docker-compose && sudo chmod +x /usr/local/bin/docker-compose

# To self-update docker-compose
docker-compose migrate-to-labels
```

- [**Golang-Migrate**](https://github.com/golang-migrate/migrate/tree/master/cmd/migrate):

```bash
go install -tags 'postgres' github.com/golang-migrate/migrate/v4/cmd/migrate@latest
```

- [**SQLc**](https://docs.sqlc.dev/en/latest/overview/install.html):

```bash
go install github.com/kyleconroy/sqlc/cmd/sqlc@latest
```

- [**GoMock**](https://github.com/golang/mock):

```bash
go install github.com/golang/mock/mockgen@latest
```

- [**Viper**](https://github.com/spf13/viper):

```bash
go install https://github.com/spf13/viper@latest
```

- [**Gin**](https://github.com/gin-gonic/gin#installation):

```bash
go install github.com/gin-gonic/gin@latest

go get -u github.com/gin-gonic/gin
```

- [**Paseto**](https://github.com/o1egl/paseto):

```bash
go get -u github.com/o1egl/paseto
```

- [**JWT**](https://github.com/golang-jwt/jwt):

```bash
go get -u https://github.com/golang-jwt/jwt
```

- [**CURL**](https://curl.se/download.html) + [**JQ**](https://stedolan.github.io/jq/) + [**Chocolatery**](https://docs.chocolatey.org/en-us/choco/setup) + [**Make**](https://community.chocolatey.org/packages/make):

```bash
sudo apt install curl jq

# These tools are needed only for Windows users

# Run this in an Admin cmd to install Chocolatery first
@"%SystemRoot%\System32\WindowsPowerShell\v1.0\powershell.exe" -NoProfile -InputFormat None -ExecutionPolicy Bypass -Command "[System.Net.ServicePointManager]::SecurityProtocol = 3072; iex ((New-Object System.Net.WebClient).DownloadString('https://community.chocolatey.org/install.ps1'))" && SET "PATH=%PATH%;%ALLUSERSPROFILE%\chocolatey\bin"

# Then install GNU-Make, cURL, and jq via Chocolatery in Admin pwsh
choco install make curl jq
```

### Infrastructure

- Create order-demo-network

```bash
make network
```

- Start postgres container:

```bash
make postgres
```

- Create order_demo database

```bash
make createdb
```

- Run DB migrate up for all versions:

```bash
make migrateup
```

- Run DB migrate down for all versions:

```bash
make migratedown
```

### Code Generation

- Generate SQL CRUD via SQLc:

```bash
make sqlc
```

- Generate DB mock via GoMock:

```bash
make mock
```

- Start postgres container:

```bash
migrate create -ext sql -dir db/migration -seq <migration_name>
```

</details>

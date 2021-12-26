# Production-ready Simple Order Service Demo

![ci-test](https://github.com/lavantien/order-demo/actions/workflows/ci.yml/badge.svg?branch=main)

## Requirements

1. [X] View list of products
2. [X] Add/Remove product from cart
3. [X] Create new order with payment
4. [ ] Users can login, sign up

### **API Endpoints**

<details>
 <summary>See details</summary>

See Booting Up running and testing instructions in the section below first, and then

### Rqm1: View list of products

```bash
curl http://localhost:8080/products?page_id=1&page_size=5 | jq
# Should return
[
  {
    "id": 1,
    "name": "lxhoak",
    "cost": 445,
    "quantity": 6,
    "created_at": "2021-12-20T19:25:16.15668Z"
  },
  {
    "id": 2,
    "name": "yyfxbi",
    "cost": 777,
    "quantity": 10,
    "created_at": "2021-12-20T19:25:16.159625Z"
  },
  {
    "id": 3,
    "name": "vyloqc",
    "cost": 975,
    "quantity": 1,
    "created_at": "2021-12-20T19:25:16.162256Z"
  },
  {
    "id": 4,
    "name": "csibko",
    "cost": 271,
    "quantity": 6,
    "created_at": "2021-12-20T19:25:16.163474Z"
  },
  {
    "id": 5,
    "name": "aymlpf",
    "cost": 93,
    "quantity": 3,
    "created_at": "2021-12-20T19:25:16.164919Z"
  }
]
```

### Rqm2.1: Add product to cart

```bash
curl http://localhost:8080/products/cart/add -H 'Content-Type: application/json' -d '{"product_id":1,"quantity":2}' | jq
# Should return
{
  "product": {
    "id": 1,
    "name": "lxhoak",
    "cost": 445,
    "quantity": 4,
    "created_at": "2021-12-20T19:25:16.15668Z"
  }
}
```

### Rqm2.2: Remove product from cart

```bash
curl http://localhost:8080/products/cart/remove -H 'Content-Type: application/json' -d '{"product_id":1,"quantity":2}' | jq
# Should return
{
  "product": {
    "id": 1,
    "name": "lxhoak",
    "cost": 445,
    "quantity": 8,
    "created_at": "2021-12-20T19:25:16.15668Z"
  }
}
```

### Rqm3.1: Create new order with payment

```bash
curl http://localhost:8080/orders -H 'Content-Type: application/json' -d '{"user_id":1,"product_id":1,"quantity":2}' | jq
# Should return
{
  "user": {
    "id": 1,
    "email": "dhksfo@email.com",
    "hashed_password": "lyxceaqfnueo",
    "created_at": "2021-12-20T19:32:33.096859Z"
  },
  "product": {
    "id": 1,
    "name": "lxhoak",
    "cost": 445,
    "quantity": 4,
    "created_at": "2021-12-20T19:25:16.15668Z"
  },
  "order": {
    "id": 96,
    "user_id": 1,
    "product_id": 1,
    "quantity": 2,
    "price": 890,
    "created_at": "2021-12-22T19:31:52.272728Z"
  }
}
```

### Rqm3.2: Check the result

```bash
curl http://localhost:8080/orders?page_id=1&page_size=5 | jq
# Should return
[
  {
    "id": 1,
    "user_id": 13,
    "product_id": 37,
    "quantity": 9,
    "price": 1962,
    "created_at": "2021-12-20T19:42:26.68327Z"
  },
  {
    "id": 2,
    "user_id": 14,
    "product_id": 38,
    "quantity": 4,
    "price": 1124,
    "created_at": "2021-12-20T19:42:26.688983Z"
  },
  {
    "id": 3,
    "user_id": 15,
    "product_id": 39,
    "quantity": 0,
    "price": 0,
    "created_at": "2021-12-20T19:42:26.693553Z"
  },
  {
    "id": 4,
    "user_id": 16,
    "product_id": 40,
    "quantity": 3,
    "price": 1578,
    "created_at": "2021-12-20T19:42:26.697026Z"
  },
  {
    "id": 5,
    "user_id": 17,
    "product_id": 41,
    "quantity": 5,
    "price": 3525,
    "created_at": "2021-12-20T19:42:26.701013Z"
  }
]
```

### Rqm4.1: Create user via endpoint

```bash
curl http://localhost:8080/users -H "Content-Type: application/json" -d '{"username":"tien1","full_name":"Tien La","email":"tien@email.com","password":"secret"}' | jq
# Should return
{
  "username": "tien1",
  "full_name": "Tien La",
  "email": "tien@email.com",
  "password_change_at": "0001-01-01T00:00:00Z",
  "created_at": "2021-12-25T12:23:04.82391Z"
}
```

### Rqm4.2: List users

```bash
curl http://localhost:8080/users?page_id=1&page_size=5 | jq
# Should return
[
  {
    "username": "adumqa",
    "hashed_password": "$2a$10$uLLL/MJ5v2QhJFNrnviUDuPjZrj9JBBIZvGhUSHMQN6gV7Ao0DNke",
    "full_name": "fglkbw",
    "email": "fldwtr@email.com",
    "password_change_at": "0001-01-01T00:00:00Z",
    "created_at": "2021-12-25T22:41:26.314508Z"
  },
  {
    "username": "bjqyze",
    "hashed_password": "$2a$10$VSKxzf8tRl5Rfdz2.QOuuetIWdFvLhb5CfsJpMjf3MVm3cI0bcdG.",
    "full_name": "xadprw",
    "email": "qjaxmw@email.com",
    "password_change_at": "0001-01-01T00:00:00Z",
    "created_at": "2021-12-25T22:41:27.330422Z"
  },
  {
    "username": "dhmxjk",
    "hashed_password": "$2a$10$Ep.SYe6VSshu8gsN41evj.jvjBkqwB1Y4SqY/SpCKUyf/JL4D633a",
    "full_name": "unfbtr",
    "email": "duozjr@email.com",
    "password_change_at": "0001-01-01T00:00:00Z",
    "created_at": "2021-12-25T22:41:27.019427Z"
  },
  {
    "username": "dpcelz",
    "hashed_password": "$2a$10$sR7wzk9NEuBrNSNEXX4S/ekS7nqLI3afDdpbEbaUrO9QiFF3nDjgW",
    "full_name": "isnjcf",
    "email": "gjmgxi@email.com",
    "password_change_at": "0001-01-01T00:00:00Z",
    "created_at": "2021-12-25T22:41:26.862502Z"
  },
  {
    "username": "eyzcfe",
    "hashed_password": "$2a$10$UYy4Bx6/FGCUs4nROOKwYOtPWzYK0CtbJQZA208g/ll8NitVQ.6Dy",
    "full_name": "baewxb",
    "email": "pvpwyw@email.com",
    "password_change_at": "0001-01-01T00:00:00Z",
    "created_at": "2021-12-25T22:41:27.17524Z"
  }
]
```

### Rqm4.3: User login

```bash
curl http://localhost:8080/users/login -H "Content-Type: application/json" -d '{"username":"tien1","password":"secret"}' | jq
# Should return
{
  "access_token": "v2.local.fXVcYzYD0Zhm9UxoUtxlxER7NtkuJ8RY0_-NxvrUY8lizSR-yIbX1KWohCEOa0QMrxexozDgLV9LeNL681RIOkcz6LJZZivn4cBbWMw4TY0MEmPs_k0wbe2o4Z5COlLeE0SKLR3UcEudAyoXUt0RYbmeeA80SEv2OzkyPcd3N83czo2qe3y1Lu45uJdACOG2hwgn0_wKmq82R1pUJz5SnCUnE_6h3G1IKnt3Ic-BfkFSx2Q89Z1jKFH1jwPxyml3gUamvwIydcBszrF6FQ.bnVsbA",
  "user": {
    "username": "tien1",
    "full_name": "Tien La",
    "email": "tien@email.com",
    "password_change_at": "0001-01-01T00:00:00Z",
    "created_at": "2021-12-26T13:08:51.243974Z"
  }
}
```

### Rqm4.4: User login with wrong password

```bash
curl http://localhost:8080/users/login -H "Content-Type: application/json" -d '{"username":"tien1","password":"abc123"}' | jq
# Should return (401)
{
  "error": "crypto/bcrypt: hashedPassword is not the hash of the given password"
}
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

- Run createdb and migrateup:

```bash
make createdb

make migrateup
```

- Run test:

```bash
go get github.com/golang/mock/mockgen/model

make mock

make test
```

- Run server:

```bash
make server
```

## Development Infrastructure Setup

### Helpful Commands

```bash
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

### Toolings Installation Guide

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

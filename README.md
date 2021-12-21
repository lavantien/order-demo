# Production-ready Simple Order Service Demo

![ci-test](https://github.com/lavantien/order-demo/actions/workflows/ci.yml/badge.svg?branch=main)

## Requirements

- [ ] View list of products
- [ ] Add/Remove products from cart
- [ ] Create new order with payment
- [ ] Users can login, sign up

## Database UML

![Database UML](/resources/readme/order-demo.png "Database UML")

## Technology Stack

- **Go 1.17**: *Leverage the standard libraries as much as possible*
- **SQLc**: *Generates efficient native SQL CRUD code*
- **PostgreSQL**: *RDBMS of choice because of faster read due to its indexing model and safer transaction with better isolation levels handling*
- **Gin**: *Fast and have respect for native net/http API*
- **Paseto Token**: *Better choice than JWT because of enforcing better cryptographic standards and debloated of useless information*
- **Golang-Migrate**: *Efficient schema generating, up/down migrating*
- **GoMock**: *Generates mocks of about anything*
- **Docker** + **Docker-Compose**: *Containerization, what else to say ...*
- **Github Actions CI**: *Make sure we don't push trash code into the codebase*
- **Viper**: *Add robustness to configurations*

## Philosophy and Architecture

- **Adaptive Minimalist**: *I always keep it as simple as possible, but with a highly decoupled structure we ensure high adaptivity and extensibility, on top of that minimal solid head start. Things are implement only when they're absolutely needed*

## Usage

- Run createdb and migrateup:

```bash
make createdb

make migrateup
```

- Run test:

```bash
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

- [**Chocolatery**](https://docs.chocolatey.org/en-us/choco/setup) + [**Make**](https://community.chocolatey.org/packages/make):

```bash
# These tools are needed only for Windows users

# Run this in an Admin cmd to install Chocolatery first
@"%SystemRoot%\System32\WindowsPowerShell\v1.0\powershell.exe" -NoProfile -InputFormat None -ExecutionPolicy Bypass -Command "[System.Net.ServicePointManager]::SecurityProtocol = 3072; iex ((New-Object System.Net.WebClient).DownloadString('https://community.chocolatey.org/install.ps1'))" && SET "PATH=%PATH%;%ALLUSERSPROFILE%\chocolatey\bin"

# Then install GNU-Make via Chocolatery
choco install make
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

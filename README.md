# Elastic with go

## Purpose
This project is the POC for [CU Get Reg](https://cugetreg.com) and [CU Newbie](https://cunewbie.com) search feature

## Features

- Create Index
  - [x] CLI
  - [ ] API
- Reindex
  - [ ] CLI
  - [ ] API
- Insert Data
  - [ ] Bulk API
  - [x] CQRS
- Search
  - [x] Search with basic query string
  - [ ] Advance search with query string (specific fields)
  - [ ] Pagination

## Stacks
- golang
- gRPC
- elasticsearch
- redis
- rabbitmq

## Getting Start
These instructions will get you a copy of the project up and running on your local machine for development and testing purposes.

### Prerequisites
- golang 1.19 or [later](https://go.dev)
- docker
- makefile

### Installing
1. Clone the project from [Elastic with go](https://github.com/samithiwat/elastic-with-go)
2. Import project
3. Copy `app.example.yaml` in `config` and paste it in the same location then remove `.example` from its name.
3. Copy `elasticsearch.example.yaml` in `config` and paste it in the same location then remove `.example` from its name.
3. Copy `redis.example.yaml` in `config` and paste it in the same location then remove `.example` from its name.
4. Download dependencies by `go mod download`

### Testing
1. Run `	go test  -v -coverpkg ./src/internal/... -coverprofile coverage.out -covermode count ./src/internal/...` or `make test`

### Running
1. Run `docker-compose up -d` or `make compose-up`
2. Run `go run ./src/.` or `make server`

### Compile proto file
1. Run `make proto`

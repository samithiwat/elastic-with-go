proto:
	protoc --proto_path=src/proto --go_out=. --go-grpc_out=require_unimplemented_servers=false:. common.proto
	protoc --proto_path=src/proto --go_out=. --go-grpc_out=require_unimplemented_servers=false:. course-search.proto
	protoc --proto_path=src/proto --go_out=. --go-grpc_out=require_unimplemented_servers=false:. course-suggest.proto

test:
	go vet ./...
	go test  -v -coverpkg ./src/internal/... -coverprofile coverage.out -covermode count ./src/internal/...
	go tool cover -func=coverage.out
	go tool cover -html=coverage.out -o coverage.html

server:
	go run ./src/cmd/main.go

compose-up:
	docker-compose up -d

compose-down:
	docker-compose down

seed:
	go run ./src/. seed
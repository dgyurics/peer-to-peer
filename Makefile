PROJECT_NAME = peer-to-peer

.DEFAULT_GOAL := run

run:
	go run cmd/main.go

run-with-port:
	go run cmd/main.go -port=$(port)

test:
	go test -v -race ./...

compile:
	protoc api/*.proto \
		--go_out=. \
		--go_opt=paths=source_relative \
		--go-grpc_out=. \
		--go-grpc_opt=paths=source_relative \
		--proto_path=.

all: clean generate_grpc_code build

clean:
	go clean

generate_grpc_code:
	protoc --go_out=plugins=grpc:. api/api.proto

build:
	go build -o bin/server

run:
	bin/server

# Testing
.PHONY: run test integration
test:
	go test -v -count=1 ./...

integration:
	go test -v -count=1 -tags integration ./...

# Database
create_db:
	docker run --name grpc-redis-data -p ${GRPC_DB_PORT}:6379 \
	-d redis redis-server --appendonly no --save "" --requirepass ${GRPC_DB_REQ_PSWD}

start_db:
	docker start grpc-redis-data

stop_db:
	docker stop grpc-redis-data

remove_db:
	docker rm -f grpc-redis-data

cli:
	docker run --rm -it --net host redis redis-cli -p ${GRPC_DB_PORT}

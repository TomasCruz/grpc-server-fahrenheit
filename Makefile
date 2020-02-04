all: clean generate build

clean:
	go clean

generate:
	protoc --go_out=plugins=grpc:. presenter/api.proto

build:
	go build -o bin/server

run:
	bin/server

.PHONY: run test integration
test:
	go test -v -count=1 ./...

create_db:
	docker run --name grpc-data -e POSTGRES_USER=grpc -e POSTGRES_PASSWORD=secret -p 5432:5432 -d postgres

start_db:
	docker start grpc-data

stop_db:
	docker stop grpc-data

remove_db:
	docker rm -f grpc-data

psql:
	docker run -it --rm --link grpc-data:postgres postgres psql postgresql://grpc:secret@postgres:5432

integration:
	go test -v -count=1 -tags integration ./...

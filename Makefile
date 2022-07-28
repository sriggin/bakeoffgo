#!/usr/bin/make

proto:
	protoc --go_out=. --go_opt=paths=source_relative \
		--go-grpc_out=. \
		--go-grpc_opt=paths=source_relative \
		trafficscanner/service.proto

build: proto
	go build -o bin/bakeoffgo main.go

run:
	go run main.go

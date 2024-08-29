CURRENT_DIR=$(shell pwd)
APP=template
APP_CMD_DIR=./cmd

DB_URL="postgres://postgres:feruza1727@database-1.c7cqkqsa66fu.us-east-2.rds.amazonaws.com:5432/auth_service?sslmode=disable"

run:
	go run cmd/main.go
init:
	# go mod init
	go mod tidy 
	go mod vendor

proto-gen:
	./internal/scripts/gen-proto.sh ${CURRENT_DIR}

migrate_up:
	migrate -path migrations -database ${DB_URL} -verbose up

migrate_down:
	migrate -path migrations -database ${DB_URL} -verbose down

migrate_force:
	migrate -path migrations -database ${DB_URL} -verbose force 1

migrate_file:
	migrate create -ext sql -dir migrations -seq create_table

build:
	CGO_ENABLED=0 GOOS=darwin go build -mod=vendor -a -installsuffix cgo -o ${CURRENT_DIR}/bin/${APP} ${APP_CMD_DIR}/main.go

swag-gen:
	~/go/bin/swag init -g ./api/api.go -o api/docs force 1

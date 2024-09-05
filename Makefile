run:
	go mod tidy
	clear
	go run main.go

proto-gen:
	protoc --go_out=./ --go-grpc_out=./ submodule/protos/auth_service/*.proto

migrate_up:
	migrate -path migrations -database postgres://postgres:postgres@localhost:5432/project_control?sslmode=disable -verbose up

migrate_down:
	migrate -path migrations -database postgres://postgres:postgres@localhost:5432/project_control?sslmode=disable -verbose down

migrate_force:
	migrate -path migrations -database postgres://postgres:postgres@localhost:5432/project_control?sslmode=disable -verbose force 1

migrate_file:
	migrate create -ext sql -dir migrations -seq create_table

swag-gen:
	~/go/bin/swag init -g ./api/api.go -o api/docs force 1
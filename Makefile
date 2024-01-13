
go.install:
	go install -tags 'postgres' github.com/golang-migrate/migrate/v4/cmd/migrate@v4.17.0
	go install github.com/sqlc-dev/sqlc/cmd/sqlc@v1.25.0
	go install github.com/golangci/golangci-lint/cmd/golangci-lint@v1.55.2
	go install github.com/google/wire/cmd/wire@v0.5.0

migrate.create:
	migrate create -ext sql -dir db/migrations -seq create_users_table

sqlc.gen:
	sqlc generate

wire.gen:
	wire ./...

go.lint:
	golangci-lint run ./...

go.install:
	go install -tags 'postgres' github.com/golang-migrate/migrate/v4/cmd/migrate@v4.17.0

migrate.create:
	migrate create -ext sql -dir db/migrations -seq create_users_table
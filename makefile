run:
	go run cmd/main.go

migrate-install:
	go install -tags 'postgres' github.com/golang-migrate/migrate/v4/cmd/migrate@latest

create_migration: migrate-install
	migrate create -ext sql -dir ./migrations $(name)

migrate-up:
	go run cmd/cli migrate up

migrate-down:
	go run cmd/cli.go migrate down

migrate-test-up:
	GO_ENVIRONMENT=test go run cmd/cli.go migrate up

migrate-test-down:
	GO_ENVIRONMENT=test go run cmd/cli.go migrate down

compose-up:

compose-up-integration-test:

get-version:
	git describe --tags

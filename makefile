GO_ENVIRONMENT=dev

run:
	go run -ldflags "-X build.Version=$(git describe --tags) -X build.HashCommit=$(git rev-parse HEAD)" app/main.go

migrate-install:
	go install -tags 'postgres' github.com/golang-migrate/migrate/v4/cmd/migrate@latest

create_migration: migrate-install
	migrate create -ext sql -dir ./migrations $(name)

migrate-up:
	go run cmd/*.go migrate up

migrate-down:
	go run cmd/*.go migrate down -v

migrate-test-up:
	GO_ENVIRONMENT=test go run cmd/*.go migrate up

migrate-test-down:
	GO_ENVIRONMENT=test go run cmd/*.go migrate down

test:
	go clean -testcache
	GO_ENVIRONMENT=test CONFIG_FOLDER=../../../config go test -v ./internal/tests/integration/...

compose-up:
	docker-compose -f ./deployment/docker-compose-local.yaml -p four_in_a_row_game up

compose-down:
	docker-compose -f ./deployment/docker-compose-local.yaml -p four_in_a_row_game down

docker:
	docker build . --tag 'game' -f deployment/Dockerfile
	docker run game

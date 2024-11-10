lint:
	golangci-lint run

gen-sql:
	sqlc generate

pkg-update:
	go get -u
	go mod tidy

cache-clean:
	go clean -modcache

run:
	go build && ./iggy

int-test:
	k6 run integration_tests/main.js

deps:
	go install github.com/sqlc-dev/sqlc/cmd/sqlc@latest
	curl -sSfL https://raw.githubusercontent.com/golangci/golangci-lint/master/install.sh | sh -s -- -b $(go env GOPATH)/bin v1.61.0
	go install -tags 'postgres' github.com/golang-migrate/migrate/v4/cmd/migrate
	go install golang.org/x/vuln/cmd/govulncheck@latest
	go install mvdan.cc/gofumpt@latest
	go install github.com/daixiang0/gci@latest

fmt:
	gofumpt -l -w .

migrate:
	migrate -source file://migration/migration_data -database "postgres://${IGGY_DB_USER}:${IGGY_DB_PASSWORD}@${IGGY_DB_HOST}:${IGGY_DB_PORT}/${IGGY_DB_DATABASE}?sslmode=disable&x-migrations-table=\"public\".\"iggy_schema_migrations\"&x-migrations-table-quoted=1" up

rollback:
	migrate -source file://migration/migration_data -database "postgres://${IGGY_DB_USER}:${IGGY_DB_PASSWORD}@${IGGY_DB_HOST}:${IGGY_DB_PORT}/${IGGY_DB_DATABASE}?sslmode=disable&x-migrations-table=\"public\".\"iggy_schema_migrations\"&x-migrations-table-quoted=1" down
psql:
	PGPASSWORD=${IGGY_DB_PASSWORD} psql -h ${IGGY_DB_HOST} -U ${IGGY_DB_USER} ${IGGY_DB_DATABASE}

vuln:
	govulncheck ./...

gci:
	gci write --skip-generated -s standard -s default .


full-lint: gci fmt lint vuln
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
	# go install github.com/golang-migrate/migrate/v4/cmd/migrate

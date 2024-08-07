run:
	@ go run cmd/cli/main.go

lint:
	@ golangci-lint run

it:
	@ go test -count=1 -v ./internal/test

build:
	@ GOOS=linux GOARCH=amd64 go build -o tree cmd/cli/main.go
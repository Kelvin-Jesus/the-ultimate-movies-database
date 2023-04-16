all: test vet lint build

test:
	go test ./...

vet:
	go vet ./...

lint:
	go list ./... | grep -v /vendor/ | xargs -L1 golint -set_exit_status

build:
	go build -o bin/db-populate ./cmd/main.go
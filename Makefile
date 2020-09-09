run:
	go run main.go

build:
	go build -o bin/ .

test:
	go test -v -cover -race ./...
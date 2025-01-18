build:
	go build .

run:
	go run .

lint:
	golangci-lint run -c .golangci.yml

unit-test:
	go test ./... -short

generate-mocks:
	mockgen -source=./service.go -destination=./mock_service.go -package=main
	mockgen -source=./handler.go -destination=./mock_handler.go -package=main
	mockgen -source=./store.go -destination=./mock_store.go -package=main

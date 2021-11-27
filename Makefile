pre-requisites:
	go mod tidy
	go get
	npm install -g newman
run:
	go run ./cmd/aluraflix-api/main.go 
test:
	go test -v ./... -cover
integration-test:
	docker-compose up -d
	clear
	newman run Aluraflix.postman_collection.json
	docker-compose down
all-tests:
	make test
	make integration-test
lint:
	go fmt ./... 
	clear
	go vet ./...
	clear
	golangci-lint run ./internal/... ./cmd/... --skip-files wire.go
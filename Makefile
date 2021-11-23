pre-requisites:
	go mod tidy
	go get
	npm install -g newman
run:
	go run main.go wire_gen.go app.go
test:
	go test -v ./... -cover
integration-test:
	newman run Aluraflix.postman_collection.json
all-tests:
	make test
	make integration-test
lint:
	go fmt ./... 
	clear
	go vet ./...
	clear
	golangci-lint run ./internal/... ./cmd/...
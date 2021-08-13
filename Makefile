pre-requisites:
	go mod tidy
	go get
	npm install -g newman
run:
	go run main.go app.go
test:
	go test -v ./... -cover
	make integration
integration:
	newman run Aluraflix.postman_collection.json

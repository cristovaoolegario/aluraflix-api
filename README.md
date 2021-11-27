# Aluraflix API

![GitHub Workflow Status](https://img.shields.io/github/workflow/status/cristovaoolegario/aluraflix-api/CI)
![Codecov](https://img.shields.io/codecov/c/gh/cristovaoolegario/aluraflix-api)
[![Go Report Card](https://goreportcard.com/badge/github.com/cristovaoolegario/aluraflix-api)](https://goreportcard.com/report/github.com/cristovaoolegario/aluraflix-api)

![GitHub tag (latest SemVer)](https://img.shields.io/github/v/tag/cristovaoolegario/aluraflix-api)
![GitHub release (latest SemVer)](https://img.shields.io/github/v/release/cristovaoolegario/aluraflix-api)

![Badge AluraFlix reduzido - Sharer](https://user-images.githubusercontent.com/79534537/130669222-e3e649dd-565b-4bb3-85a7-54bdc4f02dcb.png)

## About the project

This is a challenge from [Alura Backend Challenges](https://github.com/alura-challenges/challenge-back-end).

This application was developed with:

- Golang
- Mongodb
- Docker-Compose

This project image is being published to
this [Docker hub repository](https://hub.docker.com/repository/docker/cristovaoolegario/aluraflix-api).

## Pre-requisites

You need to have [Go](https://golang.org/dl/), [Docker](https://www.docker.com/products/docker-desktop)
and [NPM](https://www.npmjs.com/) (for integration tests on command line) installed on your machine.

`make pre-requisites`

## Running tests

To run unit tests you can use:

`make test`

To run integration tests you can use:

`make integration-test`

To run all the tests (unit and integration test) in this repo you can use:

`make all-tests`

## Run instructions

### Locally

To run it locally, you will need to set up a .env file like below:

- ```shell
  PORT=
  ENV=
  AUD=
  ISS=
  APP_DB_USERNAME=
  APP_DB_PASSWORD=
  APP_DB_HOST=
  APP_DB_NAME=
  ```

- Then run `go run ./cmd/aluraflix-api/main.go`

### Docker container

`docker-compose up -d`

## Import Postman Collection (API's)

Download [Postman](https://www.getpostman.com/) -> Import -> Import from link

Paste the link
to : [Aluraflix.postman_collection.json](https://raw.githubusercontent.com/cristovaoolegario/aluraflix-api/main/Aluraflix.postman_collection.json)

Includes the following:

- Auth
  - Credentials
  - Testing all endpoints with invalid token
- Categories
  - Create Category
  - Get Categories with filters
  - Get Category By Id
  - Update Category
  - Delete Category
- Videos
  - Create Video
  - Get Videos with filters
  - Get Video By Id
  - Update Video
  - Delete Video

## License

[MIT](https://choosealicense.com/licenses/mit/)

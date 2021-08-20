# Aluraflix API

## About the project
This is a challenge from [Alura Backend Challenges](https://github.com/alura-challenges/challenge-back-end).

This application was developed with:

- Golang
- Mongodb
- Docker-Compose

This project image is being published to this [Docker hub repository](https://hub.docker.com/repository/docker/cristovaoolegario/aluraflix-api).

## Pre-requisites

You need to have [Go](https://golang.org/dl/), [Docker](https://www.docker.com/products/docker-desktop) and [NPM](https://www.npmjs.com/) (for integration tests on command line) installed on your machine.

```make pre-requisites```

## Running tests

To run unit tests you can use:

```make test```

To run integration tests you can use:

```make integration-test```

To run all the tests (unit and integration test) in this repo you can use:

```make all-tests```


## Run instructions:

``` docker-compose up -d ```

## License
[MIT](https://choosealicense.com/licenses/mit/)

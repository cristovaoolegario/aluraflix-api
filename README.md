# Aluraflix API

## Pre-requisites

You need to have Go, Docker and NPM (for integration tests on command line) installed on your machine.

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

``` go mod download ```

``` go run main.go app.go ```

## License
[MIT](https://choosealicense.com/licenses/mit/)

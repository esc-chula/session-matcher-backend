# Session Matcher Backend

## Contents

- [Stack](#stack)
- [Env](#env)
- [Run](#run)
- [Endpoint](#endpoint)

<div id="stack"></div>

## Stack

- [Golang](https://go.dev/)
- [Gin](https://github.com/gin-gonic/)
- [Docker](https://www.docker.com/)

<div id="env"></div>

## Env

| name     | optional | default | description             |   |
|----------|----------|---------|-------------------------|---|
| PORT     |    yes   | 8080    | expose port             |   |
| PREFIX   |    yes   |         | endpoint prefix         |   |
| DATA_URL |          |         | url to request sections |   |
<div id="run"></div>

## Run

### Locally

1. Setup .env
2. Run `go run src/main`

### Docker

1. Build the image with `docker build -t <image-name> .`
2. Run `docker run -p 8080:8080 -e DATA_URL=<your-data-url> go/match`

<div id="endpoint"></div>

## Endpoint

Get /PREFIX/get-match-sections <br>
query ids eg. 65XXXXXX21,65XXXXXX21,...
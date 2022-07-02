FROM golang:1.18-alpine AS build
WORKDIR /app
COPY go.mod .
COPY go.sum .
RUN go mod download
COPY . .
RUN go build -o main src/main.go

FROM alpine
ENV ENV production
WORKDIR /app
COPY --from=build /app/main .
CMD [ "./main" ]
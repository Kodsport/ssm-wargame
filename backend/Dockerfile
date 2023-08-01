FROM golang:1.20-alpine AS build

WORKDIR /app/
COPY go.mod go.sum /app/

RUN go mod download

COPY . /app/
RUN go build -o /bin/server ./cmd/api/main.go

FROM alpine:3.18.2
 
COPY --from=build /bin/server /bin/server

ENTRYPOINT /bin/server
FROM golang:1.18
  WORKDIR /usr/src/app
  COPY go.mod go.sum ./
  RUN go mod download && go mod verify

  COPY . .
  COPY ./src/cmd/schema.sql /docker-entrypoint-initdb.d/schema.sql
  RUN go build -v -o . ./...

  CMD ["./Q2Bank"]
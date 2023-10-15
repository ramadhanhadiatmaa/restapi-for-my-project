FROM golang:1.21.1 AS builder

WORKDIR /app

COPY go.mod ./

COPY go.sum ./

COPY . .

RUN go mod tidy
 
RUN go mod download

RUN go build -o apicsmfib

FROM alpine:latest

COPY --from=builder /app/apicsmfib .

EXPOSE 8080

CMD ./apicsmfib


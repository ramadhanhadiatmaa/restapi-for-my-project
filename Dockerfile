FROM golang:1.21.1

WORKDIR /app

COPY go.mod .
COPY go.sum .
COPY . .

RUN go mod download


RUN go build -o /apicsmfib

EXPOSE 8080

CMD ["/apicsmfib"]

# ENTRYPOINT [ "/app/apicsmfib" ]
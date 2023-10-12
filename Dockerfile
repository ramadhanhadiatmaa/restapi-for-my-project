FROM golang:alpine

RUN apk update && apk add --no-cache git

WORKDIR /app

COPY go.mod .

COPY go.sum .

RUN go mod download

COPY . .

RUN go mod tidy

RUN go build -o apicsmfib .

EXPOSE 8000

CMD [ "./apicsmfib" ]
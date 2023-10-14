FROM golang:1.21.1

WORKDIR /app

COPY . .

# COPY go.mod .
# COPY go.sum .

# COPY apicsmfib /app/

# RUN go mod download

RUN go build -o apicsmfib

EXPOSE 8080

CMD ["./apicsmfib"]

# ENTRYPOINT [ "/app/apicsmfib" ]
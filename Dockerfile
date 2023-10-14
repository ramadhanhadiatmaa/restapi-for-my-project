FROM golang:1.21.1

WORKDIR /app

COPY go.mod ./

COPY go.sum ./

COPY . .

RUN go mod tidy
 
RUN go mod download


# RUN go build -o apicsmfib
EXPOSE 8080

RUN go build -o apicsmfib


CMD ./apicsmfib
# CMD ["./apicsmfib"]

# COPY apicsmfib /app/
# ENTRYPOINT [ "/app/apicsmfib" ]
FROM golang:1.21.1

WORKDIR /app

COPY go.mod ./

COPY go.sum ./

COPY . .

# RUN go build -o apicsmfib
EXPOSE 8080

RUN CGO_ENABLED=1 go build -o /apicsmfib cmd/apicsmfib/main.go


CMD ./apicsmfib
# CMD ["./apicsmfib"]

# COPY apicsmfib /app/
# ENTRYPOINT [ "/app/apicsmfib" ]
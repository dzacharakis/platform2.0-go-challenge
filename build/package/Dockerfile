FROM golang:1.18

WORKDIR /tmp/go-webservice

# pre-copy/cache go.mod for pre-downloading dependencies and only redownloading them in subsequent builds if they change
COPY go.mod go.sum ./
RUN go mod download && go mod verify

COPY . .
RUN go build -v -o cmd/api/main .

EXPOSE 8080

CMD ["./cmd/api/main"]
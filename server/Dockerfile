FROM golang:1.26.2

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build -o /gogamev3 ./cmd/server

EXPOSE 8080

CMD ["/gogamev3"]
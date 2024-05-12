FROM golang:1.22

WORKDIR /app
ADD go.mod .
ADD go.sum .
RUN go mod download

COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build -o /docker-gs-ping cmd/main.go
CMD ["/docker-gs-ping"]
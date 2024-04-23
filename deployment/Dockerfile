FROM golang:1.22
WORKDIR /app
COPY ../ ./
RUN go mod download
WORKDIR /cmd
RUN CGO_ENABLED=0 GOOS=linux go build -o /docker-gs-ping
EXPOSE 8080
CMD ["/docker-gs-ping"]
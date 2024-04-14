FROM golang:1.22.2
WORKDIR /app
COPY .. .
RUN go mod download
WORKDIR /cmd
RUN go build main.go
EXPOSE 8080
CMD ["./main"]
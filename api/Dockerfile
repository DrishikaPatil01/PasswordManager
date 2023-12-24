FROM golang:1.21.5

WORKDIR /app

COPY /password-manager-service/go.mod /password-manager-service/go.sum ./

RUN go mod download

COPY ./password-manager-service ./

RUN CGO_ENABLED=0 GOOS=linux go build -o /docker-gs-ping

EXPOSE 8080

# Run
CMD ["/docker-gs-ping"]
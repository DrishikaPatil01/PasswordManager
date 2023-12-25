FROM golang:1.21.5

WORKDIR /app

EXPOSE 8080

COPY /password-manager-service/go.mod /password-manager-service/go.sum ./

RUN go mod download

COPY ./password-manager-service ./

RUN go build -o /password-manager-service

CMD [ "/password-manager-service" ]
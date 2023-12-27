FROM golang:1.21.5

WORKDIR /app

EXPOSE 8080

COPY ./password-manager-api ./

RUN go mod download

RUN go build -o /password-manager-api

CMD [ "/password-manager-api" ]
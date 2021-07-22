FROM golang:1.16-alpine

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download


COPY . ./


RUN go build -o /my-docker-app
EXPOSE 1323
CMD [ "/my-docker-app" ]
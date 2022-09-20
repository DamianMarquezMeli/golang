FROM golang:1.19.1-alpine3.16

EXPOSE 8000
EXPOSE 9000

RUN apk update \
    && apk add --no-cache \ 
    mysql-client \
    build-base
  
RUN mkdir /app
WORKDIR /app

COPY go.mod .
COPY go.sum .
RUN go mod download
COPY . .


RUN go build cmd/main.go
RUN mv main /usr/local/bin/

CMD ["main"]

  
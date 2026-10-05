FROM golang:1.27

WORKDIR /usr/src/app

COPY ./app .
RUN go mod download
RUN go build -v -o /usr/local/bin/app

EXPOSE 3000

CMD ["app"]
FROM golang:latest

RUN go version

COPY . /go/src/app
ENV GOPATH=/go
WORKDIR /go/src/app

RUN apt-get update
RUN apt-get -y install postgresql-client
RUN apt-get -y install make

RUN apt-get -y install curl
RUN curl -sSL https://packagecloud.io/golang-migrate/migrate/gpgkey | apt-key add -
RUN echo "deb https://packagecloud.io/golang-migrate/migrate/ubuntu/ bionic main" > /etc/apt/sources.list.d/migrate.list
RUN apt-get update && \
    apt-get install -y migrate

RUN migrate -version


RUN chmod +x ./cmd/wait-for-postgres.sh
RUN chmod +x ./migrate.sh

RUN go build -o todo-app ./cmd/main.go

CMD [ "./todo-app" ]

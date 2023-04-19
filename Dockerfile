FROM golang:latest

WORKDIR /whisprite
COPY . .

RUN apt-get update
RUN apt-get install -y libsqlite3-dev sqlite3

RUN go build

CMD ["./whisprite"]

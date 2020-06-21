FROM golang AS build
RUN apt-get update
RUN apt-get install sqlite3

WORKDIR /source
COPY ./go.mod ./go.sum ./
RUN go mod download

COPY . .
RUN go build -v .

EXPOSE 80
CMD ["go", "run", "api.go"]
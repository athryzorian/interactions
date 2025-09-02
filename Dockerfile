##
## Build
##

FROM golang:1.24.6 AS build

WORKDIR /app

COPY go.mod .
COPY go.sum .

RUN go mod download

# Copy everything from this project into the filesystem of the container.
COPY *.go ./

# Obtain the package needed to run code. Alternatively use GO modules.
RUN go get -u github.com/lib/pq

RUN go build -o /interactions

EXPOSE 8080

ENTRYPOINT ["/interactions"]

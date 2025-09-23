##
## Build
##

FROM golang:1.24.6 AS build

WORKDIR /app

COPY go.mod .
COPY go.sum .

RUN go mod download

# Copy everything from this project into the filesystem of the container.
COPY *.go /app/
COPY ./dal/ /app/

# Obtain the package needed to run code. Alternatively use GO modules.
RUN go get -u github.com/lib/pq
RUN go get github.com/athryzorian/interactions/backend/dal/operations
RUN go get -u github.com/labstack/echo/v4
RUN go get -u github.com/labstack/echo/v4/middleware

# Build the Go application.
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o /interactions


EXPOSE 8080

ENTRYPOINT ["/interactions"]

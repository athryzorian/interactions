##
## Build
##

FROM golang:1.24.6 AS build

WORKDIR /app

COPY go.mod .
COPY go.sum .

RUN go mod download

# Copy everything from this project into the filesystem of the container.
COPY *.go .
COPY . .

# Obtain the package needed to run code. Alternatively use GO modules.


# Build the Go application.
#RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o /interactions
RUN go build -o /interactions


EXPOSE 8080

ENTRYPOINT ["/interactions"]

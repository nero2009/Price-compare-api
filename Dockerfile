
FROM golang:1.22.2

# syntax=docker/dockerfile:1


# Set destination for COPY
WORKDIR /app

# Download Go modules
COPY go.mod go.sum ./
RUN go mod download

# Copy the source code. Note the slash at the end, as explained in
# https://docs.docker.com/reference/dockerfile/#copy
COPY *.go ./
COPY ./cmd/api/*.go ./cmd/api/
COPY . ./

# Build
RUN  go build ./cmd/api/main.go 

# Optional:
# To bind to a TCP port, runtime parameters must be supplied to the docker command.
# But we can document in the Dockerfile what ports
# the application is going to listen on by default.
# https://docs.docker.com/reference/dockerfile/#expose
EXPOSE 8090

# Run
CMD ["./main"]
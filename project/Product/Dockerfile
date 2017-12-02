# golang image where workspace (GOPATH) configured at /go.
FROM golang:1.8

# Install dependencies
RUN go get gopkg.in/mgo.v2
RUN go get goji.io
RUN go get github.com/gorilla/mux

# copy the local package files to the container workspace
ADD . /go/src/github.com/hoang1127/project/products

# Setting up working directory
WORKDIR /go/src/github.com/hoang1127/project/products

# Build the users command inside the container.
RUN go install github.com/hoang1127/project/products

# Run the users microservice when the container starts.
ENTRYPOINT /go/bin/products

# Service listens on port 8080.
EXPOSE 8080

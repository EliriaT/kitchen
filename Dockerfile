# first stage - builds the binary from sources
FROM golang:alpine as build

# using build as current directory
WORKDIR /build

# Add the source code:
COPY . .

## install build deps
#RUN apk --update --no-cache add git
#COPY go.mod go.sum ./
RUN go mod download && go mod verify

# downloading dependencies and
## building server binary
#RUN go get github.com/gorilla/mux && \
#  go build -o server .
RUN go build  -o kitchen .

# second stage - using minimal image to run the server
FROM alpine:latest

# using /app as current directory
WORKDIR /app

# copy server binary from `build` layer
COPY --from=build /build/kitchen kitchen

# binary to run
CMD "/app/kitchen"

EXPOSE 8080
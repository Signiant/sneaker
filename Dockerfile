##
## Build
##

FROM golang:alpine AS build

# Git is required for fetching dependencies
# build-base is required for cgo in go test
RUN apk update && apk add --no-cache build-base git

WORKDIR /app
COPY . .
RUN go test && cd cmd/sneaker && go build -o /sneaker

##
## Deploy
##
FROM alpine:latest

WORKDIR /
COPY --from=build /sneaker /sneaker

RUN adduser -D nonroot

USER nonroot
ENTRYPOINT ["/sneaker"]
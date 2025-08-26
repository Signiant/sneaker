##
## Build
##
FROM golang:1.25 AS build

WORKDIR /app
COPY . .
RUN go test && CGO_ENABLED=0 go build -ldflags="-s -w" -o /sneaker ./cmd/sneaker

##
## Deploy
##
FROM gcr.io/distroless/static:nonroot
COPY --from=build /sneaker /sneaker
ENTRYPOINT ["/sneaker"]
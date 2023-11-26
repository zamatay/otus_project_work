FROM golang:alpine AS build
RUN apk add --no-cache git
WORKDIR /app
COPY ./server/go.mod .
COPY ./server/go.sum .
RUN go mod download
COPY ./server .
RUN go build -o server ./cmd/main.go

FROM alpine:latest
RUN adduser -DH appuser
WORKDIR /app
COPY --from=build /app/server .
USER appuser
ENTRYPOINT ["/app/server"]
FROM golang:alpine AS builder
RUN mkdir /go/src/app
ADD main.go /go/src/app/
WORKDIR /go/src/app
RUN CGO_ENABLED=0 go build -o app .

FROM alpine
WORKDIR /app
COPY --from=builder /go/src/app/app .
ENTRYPOINT ./app

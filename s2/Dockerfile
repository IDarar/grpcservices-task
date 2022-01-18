FROM golang:1.17-alpine3.13 AS builder
WORKDIR /app
ADD . .

RUN ls

RUN go build -o main s2/cmd/main.go

FROM alpine:latest

WORKDIR /app/

COPY --from=builder /app/main .

EXPOSE 7002
CMD [ "/app/main" ]
FROM golang:1.17-alpine3.13 AS builder
WORKDIR /app
ADD . .

RUN ls

RUN go build -o main s1/cmd/main.go

FROM alpine:latest

WORKDIR /app/

COPY --from=builder /app/main .

EXPOSE 7001
CMD [ "/app/main" ]
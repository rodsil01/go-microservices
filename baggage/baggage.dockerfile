FROM golang:latest as builder

RUN mkdir /app
COPY . /app

WORKDIR /app/baggage

RUN CGO_ENABLED=0 go build -o baggage .

FROM alpine:latest

RUN mkdir /app

COPY --from=builder /app/baggage /app

CMD ["/app/baggage"]
FROM golang:latest as builder

RUN mkdir /app
COPY . /app

WORKDIR /app/reservations

RUN CGO_ENABLED=0 go build -o reservations .

FROM alpine:latest

RUN mkdir /app

COPY --from=builder /app/reservations /app

CMD ["/app/reservations"]
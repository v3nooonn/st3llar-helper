FROM golang:1.22.2-alpine3.19 AS builder

WORKDIR /builder

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN go build -a -o /builder/supplier cmd/supplier/supplier.go

FROM alpine:latest AS final

WORKDIR /root

COPY --from=builder /builder/supplier .

EXPOSE 8080

CMD ["./supplier"]
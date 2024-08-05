FROM golang:1.22.1-alpine AS builder

WORKDIR /app
COPY main.go .
COPY go.mod .

RUN <<EOF
go mod tidy 
go build -o tiny-service
EOF

FROM scratch
#FROM alpine:3.17.2 AS final

WORKDIR /app

COPY --from=builder /app/tiny-service .
COPY public ./public 
CMD ["./tiny-service"]
#CMD ["sleep", "5m"]


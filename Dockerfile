FROM golang:1.22.1-alpine AS builder

WORKDIR /app

COPY main.go go.mod ./

RUN go mod tidy

RUN go build -o tiny-service .

# Stage 2: Create the Final Image
FROM scratch

WORKDIR /app

COPY --from=builder /app/tiny-service /app/tiny-service

COPY public ./public

CMD ["/app/tiny-service"]
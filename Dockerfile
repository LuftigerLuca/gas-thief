FROM golang:1.25-alpine AS builder

WORKDIR /src

COPY go.mod go.sum ./
RUN go mod download

COPY . .
RUN CGO_ENABLED=0 go build -ldflags="-s -w" -o gas-thief ./app

FROM alpine:3.20

RUN apk --no-cache add ca-certificates \
    && adduser -D -H appuser

USER appuser
WORKDIR /app

COPY --from=builder /src/gas-thief .

ENTRYPOINT ["./gas-thief"]

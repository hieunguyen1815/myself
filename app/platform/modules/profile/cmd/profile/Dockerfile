# Build context must be the platform root (app/platform/) so go.mod is accessible.
# Example: docker build -f modules/profile/cmd/profile/Dockerfile -t profile .

# ---- build stage ----
FROM golang:1.26-alpine AS builder

RUN apk add --no-cache make

WORKDIR /app

COPY go.mod go.sum* ./
RUN go mod download

COPY . .

RUN make build-profile

# ---- run stage ----
FROM alpine:3.21

RUN apk add --no-cache ca-certificates

WORKDIR /app

COPY --from=builder /app/bin/profile .
COPY --from=builder /app/assets ./assets

EXPOSE 8080

ENTRYPOINT ["./profile"]

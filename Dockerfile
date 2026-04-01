FROM golang:1.22-alpine AS builder
WORKDIR /app
COPY . .
RUN CGO_ENABLED=0 go mod download && CGO_ENABLED=0 go build -o drifter ./cmd/drifter/

FROM alpine:3.19
RUN apk add --no-cache ca-certificates
WORKDIR /app
COPY --from=builder /app/drifter .
ENV PORT=9130 DATA_DIR=/data
EXPOSE 9130
CMD ["./drifter"]

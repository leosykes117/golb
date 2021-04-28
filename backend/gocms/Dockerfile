FROM golang:1.16 AS builder
WORKDIR /src
# statically build the binary
ENV CGO_ENABLED=0
# Copy go.mod and go.sum files
COPY go.* .
# Download all dependencies. Dependencies will be cached if the go.mod and the go.sum files are not changed 
RUN go mod download
# Copy repository
COPY . .
# Build
RUN go build -o /out/golb cmd/api/main.go

# Second Stage
FROM alpine:3.13.5
RUN apk add --no-cache ca-certificates apache2-utils
WORKDIR /root/
COPY --from=builder /out/golb .
CMD ["./golb"]

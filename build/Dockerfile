FROM golang:1.16-alpine as build
WORKDIR /build
ARG VERSION="0.0.0"
ARG BUILD_INFO="Not set"

# Install system dependencies
RUN apk update && apk add make
#git gcc musl-dev

# Fetch and cache Go modules
COPY go.mod .
COPY go.sum .
RUN go mod download

# Copy in Go source files
COPY cmd/ ./cmd
COPY makefile .

# Now run the build
RUN make build

# ==================================================

FROM scratch

ENV METRICS_PORT 2112

WORKDIR /app
COPY --from=build /build/bin/k6-prometheus-exporter .

ENTRYPOINT [ "/app/k6-prometheus-exporter" ]

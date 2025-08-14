#syntax=docker/dockerfile:1

FROM --platform=$BUILDPLATFORM tonistiigi/xx:1.6.1 AS xx

FROM --platform=$BUILDPLATFORM golang:1.25.0-alpine AS build
WORKDIR /app

COPY --from=xx / /

COPY go.mod go.sum ./
RUN go mod download

COPY . .

# Set Golang build envs based on Docker platform string
ARG TARGETPLATFORM
RUN --mount=type=cache,target=/root/.cache \
  CGO_ENABLED=0 xx-go build -ldflags='-w -s' -trimpath -o pwgen


FROM gcr.io/distroless/static:nonroot
COPY config_example.toml .config/pwgen-go/config.toml
WORKDIR /
COPY --from=build /app/pwgen /
ENTRYPOINT ["/pwgen"]

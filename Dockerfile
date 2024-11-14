#syntax=docker/dockerfile:1.10

FROM --platform=$BUILDPLATFORM golang:1.23.3-alpine AS build
WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

# Set Golang build envs based on Docker platform string
ARG TARGETPLATFORM
RUN --mount=type=cache,target=/root/.cache <<EOT
  set -eux
  case "$TARGETPLATFORM" in
    'linux/amd64') export GOARCH=amd64 ;;
    'linux/arm/v6') export GOARCH=arm GOARM=6 ;;
    'linux/arm/v7') export GOARCH=arm GOARM=7 ;;
    'linux/arm64') export GOARCH=arm64 ;;
    *) echo "Unsupported target: $TARGETPLATFORM" && exit 1 ;;
  esac
  go build -ldflags='-w -s' -trimpath -o pwgen
  ./pwgen &>/dev/null
EOT


FROM gcr.io/distroless/static:nonroot
COPY --from=build /root/.config/pwgen-go .config/pwgen-go
WORKDIR /
COPY --from=build /app/pwgen /
ENTRYPOINT ["/pwgen"]

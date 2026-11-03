ARG GO_VERSION=1.27

FROM golang:${GO_VERSION} AS build
WORKDIR /src
COPY go.mod go.sum* ./
RUN --mount=type=cache,target=/go/pkg/mod go mod download
COPY . .
ARG VERSION=dev
RUN --mount=type=cache,target=/go/pkg/mod --mount=type=cache,target=/root/.cache/go-build \
    CGO_ENABLED=0 go build -trimpath -ldflags="-s -w -X main.version=${VERSION}" -o /out/go-template ./cmd/go-template

FROM gcr.io/distroless/static-debian13:nonroot
COPY --from=build /out/go-template /usr/local/bin/go-template
ENTRYPOINT ["go-template"]

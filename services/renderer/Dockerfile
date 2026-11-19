FROM --platform=$BUILDPLATFORM golang:1.27.4@sha256:205c7636d74e0ffbe3d007162df7d74f859f5a8c8322dfb9f907977686c126cf AS build
WORKDIR /src
COPY go.mod go.sum ./
RUN go mod download
COPY main.go render.go ./
ARG TARGETOS TARGETARCH
RUN CGO_ENABLED=0 GOOS=$TARGETOS GOARCH=$TARGETARCH go build -trimpath -ldflags="-s -w" -o /renderer .

FROM scratch
LABEL org.opencontainers.image.source="https://github.com/cagojeiger/notehub" \
      org.opencontainers.image.licenses="Apache-2.0"
COPY --from=build /renderer /renderer
COPY LICENSE THIRD_PARTY_NOTICES.md /
ENV LISTEN_ADDR=:8080
USER 10001:10001
EXPOSE 8080
ENTRYPOINT ["/renderer"]

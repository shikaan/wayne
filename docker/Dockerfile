# Install dependencies
FROM golang:1.11-alpine AS build_base

RUN apk add bash ca-certificates git gcc g++ libc-dev
WORKDIR /go/src/shikaan.io/wayne

ENV GO111MODULE=on

COPY go.mod .
COPY go.sum .

RUN go mod download

# Build binary
FROM build_base AS server_builder

COPY . .

RUN CGO_ENABLED=1 GOOS=linux GOARCH=amd64 go install -a -tags netgo -ldflags '-w -extldflags "-static"' ./cmd/main.go

# Run application
FROM alpine AS wayne

ARG FIREBASE_AUTH_JSON_SOURCE
ARG ACTIVEGATE_ADDRESS

RUN apk add ca-certificates

ENV ACTIVEGATE_ADDRESS ${ACTIVEGATE_ADDRESS}

COPY --from=server_builder /go/bin/main /bin/wayne

COPY --from=kjd61284.live.dynatrace.com/linux/oneagent-codemodules-musl:all / /
ENV LD_PRELOAD /opt/dynatrace/oneagent/agent/lib64/liboneagentproc.so

EXPOSE 8080

ENTRYPOINT ["/bin/wayne"]
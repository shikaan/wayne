# Install dependencies
FROM golang:1.11-stretch AS dependencies

RUN apt install bash ca-certificates git
WORKDIR /go/src/shikaan.io/wayne

ENV GO111MODULE=on

COPY go.mod .
COPY go.sum .

RUN go mod download

# Build binary
FROM dependencies AS builder

COPY . .

RUN CGO_ENABLED=1 GOOS=linux GOARCH=amd64 go install -a -tags netgo -ldflags '-w -extldflags "-static"' ./cmd/main.go

# Run application
FROM buildpack-deps:stretch-scm AS wayne

ARG DT_API_URL
ARG DT_API_TOKEN

RUN apt update && apt install -y ca-certificates fastjar

ARG DT_ONEAGENT_OPTIONS="flavor=default&include=go"
ENV DT_HOME="/opt/dynatrace"

RUN mkdir -p "$DT_HOME" && \
    wget -O "$DT_HOME/oneagent" "$DT_API_URL/v1/deployment/installer/agent/unix/paas/latest?Api-Token=$DT_API_TOKEN&$DT_ONEAGENT_OPTIONS"

ENTRYPOINT [ "$DT_HOME/oneagent/dynatrace-agent64.sh" ]

COPY --from=builder /go/bin/main /bin/wayne

EXPOSE 8080

ENTRYPOINT ["/bin/wayne"]
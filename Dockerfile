FROM public.ecr.aws/docker/library/golang:1.21-alpine AS builder

ARG GIT_TOKEN

ENV GO111MODULE=on \
    CGO_ENABLED=1 \
    GOOS=linux \
    GOARCH=amd64 \
    GIT_TERMINAL_PROMPT=1

RUN apk add --no-cache build-base git bash linux-headers eudev-dev curl ca-certificates

WORKDIR /build
COPY . .

# RUN echo "https://reneecok:${GIT_TOKEN}@github.com" > ~/.git-credentials \
#     && git config --global credential.helper store

RUN go mod tidy
RUN go build -o .build/staking-indexer ./

FROM public.ecr.aws/docker/library/alpine:latest

RUN apk add --no-cache build-base bash vim curl busybox-extras

WORKDIR /opt/app

COPY --from=builder /build/.build/staking-indexer /opt/app/

ENTRYPOINT /opt/app/staking-indexer -config /opt/app/configs/config.toml

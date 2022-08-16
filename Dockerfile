FROM golang:1.18-buster AS builder
WORKDIR /app
COPY . /app
RUN make build && strip ./build/pongo2-runner

FROM alpine:3.16
COPY --from=builder /app/build/pongo2-runner /usr/bin/pongo2-runner

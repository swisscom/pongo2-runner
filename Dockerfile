FROM golang:1.16
COPY . $GOPATH/src/github.com/swisscom/pongo2-runner
WORKDIR $GOPATH/src/github.com/swisscom/pongo2-runner
RUN CGO_ENABLED=0 go build -o ./pongo2-runner ./cmd && strip ./pongo2-runner

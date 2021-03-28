FROM golang:1.16-alpine

RUN apk update && \
    apk upgrade && \
    apk add --no-cache git make

RUN go get github.com/uudashr/gopkgs/v2/cmd/gopkgs@latest\
    github.com/ramya-rao-a/go-outline \
    github.com/go-delve/delve/cmd/dlv \
    golang.org/x/lint/golint \
    github.com/josharian/impl \
    honnef.co/go/tools/cmd/staticcheck \
    golang.org/x/tools/...

RUN GO111MODULE=on go get golang.org/x/tools/gopls@master golang.org/x/tools@master
ENV CGO_ENABLED=0

COPY . /workspaces/learn-go-with-tests
WORKDIR /workspaces/learn-go-with-tests

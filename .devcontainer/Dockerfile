FROM golang:alpine

RUN apk add --update --no-cache git gcc
RUN go get github.com/uudashr/gopkgs/v2/cmd/gopkgs \
    && go get github.com/uudashr/gopkgs/v2/cmd/gopkgs \
    && go get github.com/ramya-rao-a/go-outline \
    && go get github.com/cweill/gotests/... \
    && go get github.com/fatih/gomodifytags \
    && go get github.com/josharian/impl \
    && go get github.com/haya14busa/goplay/cmd/goplay \
    && go get github.com/go-delve/delve/cmd/dlv \
    && go get golang.org/x/lint/golint \
    && go get golang.org/x/tools/gopls

CMD ["tail", "-f", "/dev/null"]

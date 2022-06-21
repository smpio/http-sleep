FROM golang:1.18 as builder

WORKDIR /go/src/github.com/smpio/http-sleep/

COPY . .
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -a -ldflags "-s -w"


FROM scratch
COPY --from=builder /go/src/github.com/smpio/http-sleep/http-sleep /
ENTRYPOINT ["/http-sleep"]

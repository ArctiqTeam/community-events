FROM golang:1.12.0-alpine3.9
RUN go version

COPY . /go/src/github.com/alyarctiq/actions-go-mod
WORKDIR /go/src/github.com/alyarctiq/actions-go-mod
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -a -o app .

CMD ["/go/src/github.com/alyarctiq/actions-go-mod/app"]

FROM golang:1.9 AS builder

RUN go version

COPY . /go/src/myapp
WORKDIR /go/src/myapp

# #RUN go get -v -t  .
# RUN set -x && \
#     #go get github.com/2tvenom/go-test-teamcity && \  
#     go get github.com/golang/dep/cmd/dep && \
#     dep ensure -v

# RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build  -o /your-app

FROM scratch

COPY --from=builder /go/src/myapp .

EXPOSE 8080

CMD ["/go/src/myapp"]
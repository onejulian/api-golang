FROM golang:1.19-alpine3.17 as builder

#RUN mkdir -p /go/src/api-aliados/src##
#WORKDIR /go/src/api-aliados/src
WORKDIR /app
COPY . .

RUN go mod download
#COPY . .
#RUN go build -o main main.go
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o main main.go
FROM alpine:latest
WORKDIR /app  
RUN apk --no-cache add ca-certificates

COPY --from=builder /app/main .

EXPOSE 8080
CMD ["/app/main"]

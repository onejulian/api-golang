FROM golang:1.17.5-alpine as builder

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

EXPOSE 3025
CMD ["/app/main"]

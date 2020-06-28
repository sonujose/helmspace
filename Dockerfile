FROM golang:1.13.12-alpine3.12 AS builder

WORKDIR /go/src/app
 
COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o helmspace .

FROM alpine:3.12 
RUN apk --no-cache add ca-certificates

WORKDIR /usr/app/

COPY --from=builder /go/src/app .

EXPOSE 5000

ENTRYPOINT ["./helmspace"]
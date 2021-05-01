FROM golang:1.7.3 AS builder
WORKDIR /src

RUN go get ./...

COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o .


FROM alpine:latest  
RUN apk --no-cache add ca-certificates
WORKDIR /root/
COPY --from=builder . .
CMD ["./metric-app"] 

FROM golang:1.16-alpine as build
WORKDIR /src

COPY . .


#RUN go mod init github.com/naren4b/metric-simulator

RUN go build -o releases/metric-simulator


FROM alpine:latest
RUN apk --no-cache add ca-certificates

# RUN addgroup -g 1001 appgroup && \
#   adduser -H -D -s /bin/false -G appgroup -u 1001 appuser

# USER 1001:1001  

COPY --from=build /src/releases/metric-simulator /bin/metric-simulator
COPY --from=build /src/public/*.html /var/lib/public/

ENTRYPOINT [ "/bin/metric-simulator" ]

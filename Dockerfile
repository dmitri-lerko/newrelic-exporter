FROM golang:1.8
RUN go get github.com/mitchellh/gox
ADD workspace /go
RUN make build

FROM alpine:latest
RUN apk --no-cache add ca-certificates
COPY --from=0 /go/bin/newrelic-exporter_linux_amd64 /usr/local/bin/newrelic-exporter
CMD ["newrelic-exporter"]

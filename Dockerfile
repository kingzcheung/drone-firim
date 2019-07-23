FROM alpine:latest

RUN apk update && \
  apk add \
    ca-certificates && \
  rm -rf /var/cache/apk/*

ADD firim/firim /bin/
ENTRYPOINT ["/bin/firim"]
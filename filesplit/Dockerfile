FROM golang:1.20-alpine3.18

# disable cgo to avoid gcc requirement bug
ENV CGO_ENABLED=0

RUN apk --no-cache add git entr tini staticcheck

WORKDIR /app

ENTRYPOINT ["/sbin/tini", "-g", "--"]
CMD ["/app/bin/boot.sh"]

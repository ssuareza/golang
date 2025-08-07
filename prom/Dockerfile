FROM golang:1.23.11-alpine3.22

# disable cgo to avoid gcc requirement bug
ENV CGO_ENABLED=0

RUN apk --no-cache add git entr tini make

WORKDIR /app

ENTRYPOINT ["/sbin/tini", "-g", "--"]
CMD ["/app/bin/boot.sh"]

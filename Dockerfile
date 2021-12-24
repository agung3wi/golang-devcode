# lightweight container for go
FROM golang:alpine AS app-builder

# update container's packages and install git
RUN apk update && apk add --no-cache git

# set /app to be the active directory
WORKDIR /app

# copy all files from outside container, into the container
COPY . .

# build binary
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -a -tags netgo -ldflags '-w' -o binary *.go

FROM alpine:latest

WORKDIR /app
# RUN apk update && apk add --no-cache supervisor
# COPY supervisord.conf /etc/supervisord.conf


COPY --from=app-builder /app/binary /app/binary

EXPOSE 3030

# CMD ["/usr/bin/supervisord", "-c", "/etc/supervisord.conf"]
CMD [ "/app/binary" ]


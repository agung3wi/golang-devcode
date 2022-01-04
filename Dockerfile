FROM golang:alpine AS app-builder

RUN apk update

WORKDIR /app

COPY . .

RUN CGO_ENABLED=0 go build -gcflags '-S -N' -o binary *.go

FROM golang:alpine AS app-builder2

RUN apk update

WORKDIR /app

COPY ./hardcode/main.go ./main.go

RUN CGO_ENABLED=0 go build -gcflags '-S -N' -o main *.go

FROM alpine:latest

RUN apk update && apk add --no-cache supervisor

# COPY ./default.conf /etc/nginx/conf.d/default.conf

COPY ./supervisord.conf /etc/supervisord.conf

COPY --from=app-builder /app/binary /app/binary
COPY --from=app-builder2 /app/main /app/main


EXPOSE 3030

CMD ["/usr/bin/supervisord", "-c", "/etc/supervisord.conf"]
# CMD [ "/app/binary" ]

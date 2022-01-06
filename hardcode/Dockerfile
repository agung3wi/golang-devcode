FROM golang:alpine AS app-builder

WORKDIR /app

COPY ./main.go .
COPY ./go.mod .

RUN go mod tidy

# build binary
RUN CGO_ENABLED=0 go build -gcflags '-S -N' -o binary *.go

# FROM golang:alpine AS migration


# WORKDIR /app

# COPY ./migration .

# RUN go mod tidy

# # build binary
# RUN CGO_ENABLED=0 go build -gcflags '-S -N' -o migration *.go

FROM alpine:latest

COPY --from=app-builder /app/binary /app/binary
# COPY --from=migration /app/migration /app/migration

EXPOSE 3030

# CMD ["/usr/bin/supervisord", "-c", "/etc/supervisord.conf"]
CMD /app/binary

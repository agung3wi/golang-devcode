FROM nginx:alpine

RUN apk update && apk add --no-cache supervisor

COPY ./default.conf /etc/nginx/conf.d/default.conf
COPY ./supervisord.conf /etc/supervisord.conf
EXPOSE 3030
CMD ["/usr/bin/supervisord", "-c", "/etc/supervisord.conf"]

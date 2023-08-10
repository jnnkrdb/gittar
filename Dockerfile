# Base Image
FROM alpine:3.10
LABEL org.opencontainers.image.source=https://github.com/jnnkrdb/gittar
RUN apk update && apk upgrade --no-cache
RUN apk add tzdata
RUN cp /usr/share/zoneinfo/Europe/Berlin /etc/localtime
RUN echo "Europe/Berlin" > /etc/timezone
RUN apk del tzdata
WORKDIR /
ENTRYPOINT ["/bin/sh"]

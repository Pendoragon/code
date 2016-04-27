FROM index.caicloud.io/caicloud/debian:jessie

RUN mkdir /app
WORKDIR /app

ADD fsnotify /app
RUN touch /tmp/foo

CMD ["./fsnotify"]

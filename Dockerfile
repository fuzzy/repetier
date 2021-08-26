FROM alpine:latest
RUN mkdir -p /drone/src
VOLUME /drone/src
RUN apk update
RUN apk add make go
RUN pwd
RUN ls `pwd`
RUN ls /drone/src

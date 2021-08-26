FROM alpine:latest
RUN apk update
RUN apk add make go
RUN pwd
RUN ls `pwd`
RUN ls /drone/src

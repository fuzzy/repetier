FROM alpine:latest
RUN apk update
RUN apk add make go
RUN ls /drone/src

# FROM golang
FROM alpine
# FROM golang:1.15

COPY ./red /tmp/red

WORKDIR /tmp/

RUN chmod +x red
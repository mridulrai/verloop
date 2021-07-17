# Dockerfile to be used in CI/CD

# Stage 1
# Setting official golang image tag 1.16.5 as baseimage
FROM golang:1.16.5 AS buildenv
ENV GOOS=linux GOARCH=amd64 CGO_ENABLED=0
ADD . /app
WORKDIR /app
RUN go build -o bin/payment main.go

#################################################################

# Stage 2
# Setting official alpine linux image tag latest as baseimage
FROM alpine:latest
RUN apk update \
  && apk add --no-cache --purge ca-certificates tzdata\
  && rm -rf /var/cache/apk/* /tmp/*
COPY --from=buildenv /app/bin/payment /usr/local/bin/
CMD ["payment"]


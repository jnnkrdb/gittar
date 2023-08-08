# Build the manager binary
FROM golang:1.20 as builder
RUN mkdir /build
COPY go /build
WORKDIR /build
RUN go mod download
ENV CGO_ENABLED=0
ENV GOARCH=amd64
ENV GOOS=linux
# START BUILD
RUN go mod download
RUN go build -o /gittar .

# Base Image
FROM alpine:3.10
LABEL org.opencontainers.image.source=https://github.com/jnnkrdb/gittar

RUN apk update && apk upgrade --no-cache
RUN apk add tzdata
RUN cp /usr/share/zoneinfo/Europe/Berlin /etc/localtime
RUN echo "Europe/Berlin" > /etc/timezone
RUN apk del tzdata

WORKDIR /
COPY --from=builder /gittar /gittar
RUN chmod a+x /gittar
ENTRYPOINT ["/gittar"]

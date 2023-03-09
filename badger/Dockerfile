FROM $BUILD_FROM

FROM golang:1-alpine AS builder
ARG BUILD_ARCH

RUN mkdir -p WORKDIR /usr/src/app
WORKDIR /usr/src/app
COPY src/go.mod go.sum ./
RUN go mod download && go mod verify

COPY . /usr/src/app/

RUN GOOS=linux GOARCH=${BUILD_ARCH} go build -v -o /usr/local/bin/builder ./...

# https://developers.home-assistant.io/docs/add-ons/configuration#add-on-dockerfile
ARG BUILD_FROM
ARG TEMPIO_VERSION BUILD_ARCH

# Execute during the build of the image
RUN \
    curl -sSLf -o /usr/bin/tempio \
    "https://github.com/home-assistant/tempio/releases/download/${TEMPIO_VERSION}/tempio_${BUILD_ARCH}"

# Copy root filesystem
COPY rootfs /
COPY builder:/usr/local/bin/builder /usr/bin/badger
# Fix perms
RUN chown root:root /etc/services.d/badger/* /usr/bin/badger
RUN chmod 755 /etc/services.d/badger/* /usr/bin/badger

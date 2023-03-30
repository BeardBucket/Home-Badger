ARG GOLANG_IMAGE="golang:1-alpine"
ARG BUILD_FROM="xxxxx"

FROM $GOLANG_IMAGE AS bldr
ARG BUILD_ARCH
ENV GOLANG_ARCH ""

RUN mkdir -p WORKDIR /usr/src/app
WORKDIR /usr/src/app
COPY src/go.mod src/go.sum ./
RUN go mod download && go mod verify

COPY src/ /usr/src/app

WORKDIR /usr/src/app

RUN GOOS=linux GOARCH=${GOLANG_ARCH} go build --ldflags '-extldflags "-static"' -v -o /usr/local/bin/builder ./main.go

## Debug
#RUN apk add file
#RUN file /usr/local/bin/builder
#RUN /usr/local/bin/builder --help

# https://developers.home-assistant.io/docs/add-ons/configuration#add-on-dockerfile
FROM $BUILD_FROM
RUN apk add --no-cache --virtual .build-deps file .build-deps bash gcc musl-dev openssl go
ARG TEMPIO_VERSION
ARG BUILD_ARCH

# Execute during the build of the image
RUN curl -sSLf -o /usr/bin/tempio \
    "https://github.com/home-assistant/tempio/releases/download/${TEMPIO_VERSION}/tempio_${BUILD_ARCH}"

# Copy root filesystem
COPY rootfs /
COPY --from=bldr /usr/local/bin/builder /usr/bin/badger
# Fix perms
RUN chown root:root /etc/s6-overlay/s6-rc.d/badger/* /usr/bin/badger
RUN chmod 755 /etc/s6-overlay/s6-rc.d/badger/* /usr/bin/badger

# Debug
RUN file /usr/bin/badger
RUN /usr/bin/badger --help

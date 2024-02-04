ARG BASE_NAME=wayland-base
ARG DOCKER_REGISTRY=docker.io
ARG IMAGE_TAG=4-rc

FROM $DOCKER_REGISTRY/torizon/$BASE_NAME:$IMAGE_TAG

LABEL org.toradex.image.base.registry="docker.io" \
      org.toradex.image.base.namespace="torizon" \
      org.toradex.image.base.name="wayland-base-am62" \
      org.toradex.image.base.tag.major="rc" \
      org.toradex.image.base.tag.minor="" \
      org.toradex.image.base.tag.patch="" \
      org.toradex.image.variant="" \
      org.toradex.image.registry="docker.io" \
      org.toradex.image.namespace="torizon" \
      org.toradex.image.name="chromium-am62" \
      org.toradex.image.base.tag.major="rc" \
      org.toradex.image.base.tag.minor="" \
      org.toradex.image.base.tag.patch="" \
      org.toradex.image.license="MIT" \
      org.toradex.image.arch="linux/arm64/v8"

RUN apt-get -y update && \
    apt-get install -y --no-install-recommends chromium chromium-sandbox && \
    apt-get clean && apt-get autoremove && \
    rm -rf /var/lib/apt/lists/*

ADD chrome-virtual-keyboard.tar.gz /chrome-extensions

COPY start-chromium.sh /usr/bin/start-browser

# FIXME: running with root user isn't the best.
# [7:7:0929/164631.238808:ERROR:wayland_connection.cc(198)] Failed to connect to Wayland display: Permission denied (13)
# USER torizon

ENTRYPOINT ["/usr/bin/start-browser"]


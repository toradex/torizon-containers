ARG BASE_NAME=wayland-base
ARG DOCKER_REGISTRY=docker.io
ARG IMAGE_TAG=4-rc

FROM $DOCKER_REGISTRY/torizon/$BASE_NAME:$IMAGE_TAG

RUN apt-get -y update && \
    apt-get install -y --no-install-recommends chromium chromium-sandbox && \
    apt-get clean && apt-get autoremove && \
    rm -rf /var/lib/apt/lists/*

ADD chrome-virtual-keyboard.tar.gz /chrome-extensions

COPY start-chromium.sh /usr/bin/start-browser

# FIXME: running with root user isn't the best.
# [7:7:0929/164631.238808:ERROR:wayland_connection.cc(198)] Failed to connect to Wayland display: Permission denied (13)
# USER torizon

ENV WAYLAND_DISPLAY=wayland-1
ENV XDG_RUNTIME_DIR=/tmp/1000-runtime-dir/
ENV DISPLAY=:0

ENTRYPOINT ["/usr/bin/start-browser"]


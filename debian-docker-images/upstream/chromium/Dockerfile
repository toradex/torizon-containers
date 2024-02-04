# Copyright (c) 2019-2023 Toradex AG
# SPDX-License-Identifier: MIT

ARG BASE_IMAGE_REGISTRY
ARG BASE_IMAGE_NAMESPACE
ARG BASE_IMAGE_NAME
ARG BASE_IMAGE_TAG

FROM $BASE_IMAGE_REGISTRY/$BASE_IMAGE_NAMESPACE/$BASE_IMAGE_NAME:$BASE_IMAGE_TAG

LABEL org.toradex.image.base.registry="docker.io" \
      org.toradex.image.base.namespace="torizon" \
      org.toradex.image.base.name="wayland-base" \
      org.toradex.image.base.tag.major="rc" \
      org.toradex.image.base.tag.minor="" \
      org.toradex.image.base.tag.patch="" \
      org.toradex.image.base.tag.variant="" \
      org.toradex.image.registry="docker.io" \
      org.toradex.image.namespace="torizon" \
      org.toradex.image.name="chromium" \
      org.toradex.image.tag.major="rc" \
      org.toradex.image.tag.minor="" \
      org.toradex.image.tag.patch="" \
      org.toradex.image.tag.variant="" \
      org.toradex.image.license="MIT" \
      org.toradex.image.arch="linux/amd64,linux/arm64/v8,linux/arm/v7"

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


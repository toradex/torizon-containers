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
      org.toradex.image.base.tag.variant="bookworm" \
      org.toradex.image.registry="docker.io" \
      org.toradex.image.namespace="torizon" \
      org.toradex.image.name="weston" \
      org.toradex.image.tag.major="rc" \
      org.toradex.image.tag.minor="" \
      org.toradex.image.tag.patch="" \
      org.toradex.image.tag.variant="" \
      org.toradex.image.license="MIT" \
      org.toradex.image.arch="linux/amd64,linux/arm64/v8,linux/arm/v7"

# Install the weston compositor.
RUN apt-get -y update && apt-get install -y --no-install-recommends \
    weston \
    xwayland \
    kbd \
    dos2unix \
    seatd \
    && apt-get clean && apt-get autoremove && rm -rf /var/lib/apt/lists/*

# Container entry point and support files.
COPY entry.sh /usr/bin/entry.sh
COPY switchvtmode.pl /usr/bin/switchvtmode.pl
COPY weston.ini /etc/xdg/weston/weston.ini
COPY weston-dev.ini /etc/xdg/weston-dev/weston.ini

# The compositor needs access to input devices
RUN usermod -a -G input torizon

# SUID seatd-launch
# seatd-launch requires root privileges to run
# This allows 'torizon' user to launch weston using seatd-launch
# See https://github.com/kennylevinsen/seatd/blob/a803ba0502cccf147eec7fbcacd11c5b8643c0e0/man/seatd-launch.1.scd
RUN chmod u+s /usr/bin/seatd-launch

HEALTHCHECK --interval=1s --timeout=5s --start-period=0s --retries=3 \
    CMD find /tmp/1000-runtime-dir -maxdepth 1 -type s -name 'wayland-*' -print -quit | grep -q .

WORKDIR /home/torizon

ENTRYPOINT ["/usr/bin/entry.sh"]

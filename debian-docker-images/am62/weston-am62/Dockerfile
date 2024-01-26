# Copyright (c) 2019-2024 Toradex AG
# SPDX-License-Identifier: MIT

ARG BASE_NAME=wayland-base
ARG IMAGE_TAG=3
ARG DOCKER_REGISTRY=docker.io
FROM $DOCKER_REGISTRY/torizon/$BASE_NAME:$IMAGE_TAG AS base

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

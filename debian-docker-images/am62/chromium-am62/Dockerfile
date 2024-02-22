# Copyright (c) 2019-2024 Toradex AG
# SPDX-License-Identifier: MIT

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

USER torizon

ENTRYPOINT ["/usr/bin/start-browser"]


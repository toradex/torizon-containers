# Copyright (c) 2019-2023 Toradex AG
# SPDX-License-Identifier: MIT

ARG IMAGE_ARCH=linux/arm/v7
# For arm64 use:
# ARG IMAGE_ARCH=linux/arm64/v8
ARG BASE_NAME=weston
# For arm64v8 use:
# ARG BASE_NAME=weston-vivante
ARG IMAGE_TAG=3
ARG DOCKER_REGISTRY=torizon
FROM --platform=$IMAGE_ARCH $DOCKER_REGISTRY/$BASE_NAME:$IMAGE_TAG AS base

COPY entry.sh /usr/bin/
RUN mkdir -p /etc/xdg/weston/
COPY weston.ini /etc/xdg/weston/
COPY save-calibration.sh /usr/bin/

USER root

ENTRYPOINT ["/usr/bin/entry.sh"]

# Copyright (c) 2019-2023 Toradex AG
# SPDX-License-Identifier: MIT

ARG BASE_NAME=qt5-wayland
ARG IMAGE_TAG=3
ARG DOCKER_REGISTRY

FROM $DOCKER_REGISTRY/torizon/$BASE_NAME:$IMAGE_TAG AS base

LABEL org.toradex.image.base.registry="docker.io" \
      org.toradex.image.base.namespace="torizon" \
      org.toradex.image.base.name="qt6-wayland-am62" \
      org.toradex.image.base.tag.major="rc" \
      org.toradex.image.base.tag.minor="" \
      org.toradex.image.base.tag.patch="" \
      org.toradex.image.variant="" \
      org.toradex.image.registry="docker.io" \
      org.toradex.image.namespace="torizon" \
      org.toradex.image.name="qt6-wayland-examples-am62" \
      org.toradex.image.base.tag.major="rc" \
      org.toradex.image.base.tag.minor="" \
      org.toradex.image.base.tag.patch="" \
      org.toradex.image.license="MIT" \
      org.toradex.image.arch="linux/arm64/v8"

RUN apt-get -y update && apt-get install -y --no-install-recommends \
    qt6-base-examples \
    && apt-get clean && apt-get autoremove && rm -rf /var/lib/apt/lists/*

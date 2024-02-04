# Copyright (c) 2019-2023 Toradex AG
# SPDX-License-Identifier: MIT

ARG BASE_NAME=qt5-wayland
ARG IMAGE_TAG=3
ARG DOCKER_REGISTRY=torizon

ARG QT5_EXAMPLES_DEBIAN
# Uncomment the line below if building outside of the context of GitLab CI
# ARG QT5_EXAMPLES_DEBIAN=1
ARG QT5_EXAMPLES_DEBIAN_INSTALL=${QT5_EXAMPLES_DEBIAN:+"qtbase5-examples qtdeclarative5-examples"}

FROM $DOCKER_REGISTRY/torizon/$BASE_NAME:$IMAGE_TAG AS base

LABEL org.toradex.image.base.registry="docker.io" \
      org.toradex.image.base.namespace="torizon" \
      org.toradex.image.base.name="qt5-wayland-am62" \
      org.toradex.image.base.tag.major="rc" \
      org.toradex.image.base.tag.minor="" \
      org.toradex.image.base.tag.patch="" \
      org.toradex.image.variant="" \
      org.toradex.image.registry="docker.io" \
      org.toradex.image.namespace="torizon" \
      org.toradex.image.name="qt5-wayland-examples-am62" \
      org.toradex.image.base.tag.major="rc" \
      org.toradex.image.base.tag.minor="" \
      org.toradex.image.base.tag.patch="" \
      org.toradex.image.license="MIT" \
      org.toradex.image.arch="linux/arm64/v8"

ARG QT5_EXAMPLES_DEBIAN_INSTALL

RUN apt-get -y update && apt-get install -y --no-install-recommends \
    ${QT5_EXAMPLES_DEBIAN_INSTALL} \
    && apt-get clean && apt-get autoremove && rm -rf /var/lib/apt/lists/*

# Copyright (c) 2019-2023 Toradex AG
# SPDX-License-Identifier: MIT

ARG BASE_IMAGE_REGISTRY
ARG BASE_IMAGE_NAMESPACE
ARG BASE_IMAGE_NAME
ARG BASE_IMAGE_TAG

FROM $BASE_IMAGE_REGISTRY/$BASE_IMAGE_NAMESPACE/$BASE_IMAGE_NAME:$BASE_IMAGE_TAG

LABEL org.toradex.image.base.namespace="torizon" \
      org.toradex.image.base.name="qt5-wayland-imx8" \
      org.toradex.image.base.tag="rc" \
      org.toradex.image.namespace="torizon" \
      org.toradex.image.name="qt5-wayland-examples-imx8" \
      org.toradex.image.tag="rc" \
      org.toradex.image.license="MIT" \
      org.toradex.image.arch="linux/arm64/v8"

RUN apt-get -y update && apt-get install -y --no-install-recommends \
    qtbase5-examples qtdeclarative5-examples \
    && apt-get clean && apt-get autoremove && rm -rf /var/lib/apt/lists/*

# Copyright (c) 2019-2023 Toradex AG
# SPDX-License-Identifier: MIT

ARG BASE_NAME=wayland-base-vivante
ARG IMAGE_TAG=3
ARG DOCKER_REGISTRY=torizon

FROM $DOCKER_REGISTRY/torizon/$BASE_NAME:$IMAGE_TAG AS base

LABEL org.toradex.image.base.namespace="torizon" \
      org.toradex.image.base.name="wayland-base-imx8" \
      org.toradex.image.base.tag="rc" \
      org.toradex.image.namespace="torizon" \
      org.toradex.image.name="graphics-tests-imx8" \
      org.toradex.image.tag="rc" \
      org.toradex.image.license="MIT" \
      org.toradex.image.arch="linux/arm64/v8"

RUN apt-get -y update &&  apt-get install -y --no-install-recommends \
        clinfo \
        glmark2-es2-wayland \
        kmscube \
        libdrm-tests \
        imx-gpu-viv-demos \
    && apt-get clean && apt-get autoremove && rm -rf /var/lib/apt/lists/*

WORKDIR /home/torizon

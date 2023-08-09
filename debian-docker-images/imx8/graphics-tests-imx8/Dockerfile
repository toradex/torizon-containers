# Copyright (c) 2019-2023 Toradex AG
# SPDX-License-Identifier: MIT

ARG IMAGE_ARCH=linux/arm64/v8
ARG BASE_NAME=wayland-base-vivante
ARG IMAGE_TAG=3
ARG DOCKER_REGISTRY=torizon

FROM --platform=$IMAGE_ARCH $DOCKER_REGISTRY/$BASE_NAME:$IMAGE_TAG AS base

RUN apt-get -y update &&  apt-get install -y --no-install-recommends \
        clinfo \
        glmark2-es2-wayland \
        kmscube \
        libdrm-tests \
        imx-gpu-viv-demos \
    && apt-get clean && apt-get autoremove && rm -rf /var/lib/apt/lists/*

WORKDIR /home/torizon

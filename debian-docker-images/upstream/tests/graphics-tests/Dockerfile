# Copyright (c) 2019-2024 Toradex AG
# SPDX-License-Identifier: MIT

ARG BASE_NAME=wayland-base
ARG DOCKER_REGISTRY=docker.io
ARG IMAGE_TAG=rc

FROM $DOCKER_REGISTRY/torizon/$BASE_NAME:$IMAGE_TAG AS base

RUN apt-get -y update \
    && apt-get install -y --no-install-recommends \
        clinfo \
        glmark2-es2-wayland \
        kmscube \
        libdrm-tests \
        openimageio-tools \
    && apt-get clean \
    && apt-get autoremove \
    && rm -rf /var/lib/apt/lists/*

WORKDIR /home/torizon

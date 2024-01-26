# Copyright (c) 2019-2024 Toradex AG
# SPDX-License-Identifier: MIT

ARG BASE_NAME=qt5-wayland
ARG IMAGE_TAG=3
ARG DOCKER_REGISTRY

FROM $DOCKER_REGISTRY/torizon/$BASE_NAME:$IMAGE_TAG AS base

RUN apt-get -y update && apt-get install -y --no-install-recommends \
    qt6-base-examples \
    && apt-get clean && apt-get autoremove && rm -rf /var/lib/apt/lists/*

# Copyright (c) 2019-2024 Toradex AG
# SPDX-License-Identifier: MIT

ARG BASE_NAME=qt5-wayland
ARG IMAGE_TAG=3
ARG DOCKER_REGISTRY=torizon

ARG QT5_EXAMPLES_DEBIAN
# Uncomment the line below if building outside of the context of GitLab CI
# ARG QT5_EXAMPLES_DEBIAN=1
ARG QT5_EXAMPLES_DEBIAN_INSTALL=${QT5_EXAMPLES_DEBIAN:+"qtbase5-examples qtdeclarative5-examples"}

FROM $DOCKER_REGISTRY/torizon/$BASE_NAME:$IMAGE_TAG AS base
ARG QT5_EXAMPLES_DEBIAN_INSTALL

RUN apt-get -y update && apt-get install -y --no-install-recommends \
    ${QT5_EXAMPLES_DEBIAN_INSTALL} \
    && apt-get clean && apt-get autoremove && rm -rf /var/lib/apt/lists/*

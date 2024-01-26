# Copyright (c) 2019-2024 Toradex AG
# SPDX-License-Identifier: MIT

ARG BASE_NAME=wayland-base
ARG DOCKER_REGISTRY=docker.io
ARG IMAGE_TAG=rc

FROM $DOCKER_REGISTRY/torizon/$BASE_NAME:$IMAGE_TAG AS base

# Install remaining dependencies required to run qtbase
RUN apt-get -y update \
    && apt-get install -y --no-install-recommends \
        libfontconfig1-dev \
        libqt6concurrent6 \
        libqt6dbus6 \
        libqt6gui6 \
        libqt6network6 \
        libqt6opengl6 \
        libqt6printsupport6 \
        libqt6qml6 \
        libqt6quicktest6 \
        libqt6quickwidgets6 \
        libqt6sql6 \
        libqt6test6 \
        libqt6widgets6 \
        libqt6xml6 \
        qt6-wayland \
    && apt-get clean \
    && apt-get autoremove \
    && rm -rf /var/lib/apt/lists/*

COPY kms-setup.sh /usr/bin/kms-setup.sh

ENV QT_QPA_PLATFORM="wayland"

# EGLFS configuration
ENV QT_QPA_EGLFS_INTEGRATION="eglfs_kms"
ENV QT_QPA_EGLFS_KMS_ATOMIC="1"
ENV QT_QPA_EGLFS_KMS_CONFIG="/etc/kms.conf"

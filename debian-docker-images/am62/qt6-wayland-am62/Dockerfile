# Copyright (c) 2019-2024 Toradex AG
# SPDX-License-Identifier: MIT

ARG BASE_NAME=wayland-base
ARG IMAGE_TAG=3
ARG DOCKER_REGISTRY
FROM $DOCKER_REGISTRY/torizon/$BASE_NAME:$IMAGE_TAG AS base

COPY kms-setup.sh /usr/bin/kms-setup.sh

RUN apt-get -y update && apt-get install -y --no-install-recommends \
    apt-utils \
    && apt-get -y upgrade \
    && apt-get clean && apt-get autoremove && rm -rf /var/lib/apt/lists/*

# Install remaining dependencies required to run qtbase
RUN apt-get -y update && apt-get install -y --no-install-recommends \
       libqt6gui6 \
       libfontconfig1-dev \
       libqt6concurrent6 \
       libqt6dbus6 \
       libqt6network6 \
       libqt6printsupport6 \
       libqt6sql6 \
       libqt6test6 \
       libqt6widgets6 \
       libqt6xml6 \
       libqt6qml6 \
       libqt6opengl6 \
       libqt6quicktest6 \
       libqt6quickwidgets6 \
       qt6-wayland \
       && apt-get clean && apt-get autoremove && rm -rf /var/lib/apt/lists/*

ENV QT_QPA_PLATFORM="wayland"

# EGLFS configuration
ENV QT_QPA_EGLFS_INTEGRATION="eglfs_kms"
ENV QT_QPA_EGLFS_KMS_ATOMIC="1"
ENV QT_QPA_EGLFS_KMS_CONFIG="/etc/kms.conf"

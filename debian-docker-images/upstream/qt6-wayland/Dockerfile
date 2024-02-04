# Copyright (c) 2019-2023 Toradex AG
# SPDX-License-Identifier: MIT

ARG BASE_IMAGE_REGISTRY
ARG BASE_IMAGE_NAMESPACE
ARG BASE_IMAGE_NAME
ARG BASE_IMAGE_TAG

FROM $BASE_IMAGE_REGISTRY/$BASE_IMAGE_NAMESPACE/$BASE_IMAGE_NAME:$BASE_IMAGE_TAG

LABEL org.toradex.image.base.registry="docker.io" \
      org.toradex.image.base.namespace="torizon" \
      org.toradex.image.base.name="wayland-base" \
      org.toradex.image.base.tag.major="rc" \
      org.toradex.image.base.tag.minor="" \
      org.toradex.image.base.tag.patch="" \
      org.toradex.image.base.tag.variant="" \
      org.toradex.image.registry="docker.io" \
      org.toradex.image.namespace="torizon" \
      org.toradex.image.name="qt6-wayland" \
      org.toradex.image.tag.major="rc" \
      org.toradex.image.tag.minor="" \
      org.toradex.image.tag.patch="" \
      org.toradex.image.tag.variant="" \
      org.toradex.image.license="MIT" \
      org.toradex.image.arch="linux/arm64/v8"

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

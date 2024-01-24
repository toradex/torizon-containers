# Copyright (c) 2019-2023 Toradex AG
# SPDX-License-Identifier: MIT

ARG BASE_NAME=wayland-base
ARG IMAGE_TAG=3
ARG DOCKER_REGISTRY
FROM $DOCKER_REGISTRY/torizon/$BASE_NAME:$IMAGE_TAG AS base

RUN apt-get -y update && apt-get install -y --no-install-recommends \
    apt-utils \
    && apt-get -y upgrade \
    && apt-get clean && apt-get autoremove && rm -rf /var/lib/apt/lists/*

# Install libqt5gui5-gles before libqt5opengl5 which also has an alternate dependency on libqt5gui5(non-gles)
RUN apt-get -y update  && apt-get install -y --no-install-recommends libqt5gui5-gles \
    && apt-get clean && apt-get autoremove && rm -rf /var/lib/apt/lists/*

# Install libqt5opengl5.
# Under Bookworm, libqt5opengl5 is currently at version 5.15.8+dfsg-11 and contains a redundant dependency for libqt5gui5 (>= 5.1.0):
#
# This forbids installing qtbase5-examples with libqt5gui5-gles.
# Workaround the issue by mangling the package file and remove the leftover dependency for each of the architectures
RUN apt-get -y update \
    && apt-get -y install --no-install-recommends binutils xz-utils \
    && WORK_DIR=$(mktemp -d) \
    && cd $WORK_DIR \
    && ARCH=$(dpkg --print-architecture) apt-get download libqt5opengl5:$ARCH \
    && ar x libqt5opengl5*.deb \
    && tar -xJf control.tar.xz \
    && sed -i '/^Depends:/s/, libqt5gui5 (>= 5.1.0)//' control \
    && tar -cJf control.tar.xz control md5sums shlibs symbols triggers \
    && ar rcs libqt5opengl5.deb debian-binary control.tar.xz data.tar.xz \
    && apt-get -y install --no-install-recommends ./libqt5opengl5.deb \
    && cd ~ \
    && rm -rf $WORK_DIR \
    && apt-get -y remove binutils xz-utils \
    && apt-mark hold libqt5opengl5 \
    && apt-get clean && apt-get autoremove && rm -rf /var/lib/apt/lists/*

# Install remaining dependencies required to run qtbase and qtdeclarative examples
RUN apt-get -y update && apt-get install -y --no-install-recommends \
        libfontconfig1-dev \
        libqt5quick5-gles \
        libqt5quickparticles5-gles \
        libqt5concurrent5 \
        libqt5dbus5 \
        libqt5network5 \
        libqt5printsupport5 \
        libqt5sql5 \
        libqt5test5 \
        libqt5widgets5 \
        libqt5xml5 \
        libqt5qml5 \
        libqt5quicktest5 \
        libqt5quickwidgets5 \
        qml-module-qt-labs-qmlmodels \
        qml-module-qtqml-models2 \
        qml-module-qtquick-layouts \
        qml-module-qtquick-localstorage \
        qml-module-qtquick-particles2 \
        qml-module-qtquick-shapes \
        qml-module-qttest \
    && apt-get clean && apt-get autoremove && rm -rf /var/lib/apt/lists/*

# Install Wayland Qt module
RUN apt-get -y update && apt-get install -y --no-install-recommends \
    qtwayland5 \
    && apt-get clean && apt-get autoremove && rm -rf /var/lib/apt/lists/*

COPY kms-setup.sh /usr/bin/kms-setup.sh

ENV QT_QPA_PLATFORM="wayland"

# EGLFS configuration
ENV QT_QPA_EGLFS_INTEGRATION="eglfs_kms"
ENV QT_QPA_EGLFS_KMS_ATOMIC="1"
ENV QT_QPA_EGLFS_KMS_CONFIG="/etc/kms.conf"

# Copyright (c) 2019-2023 Toradex AG
# SPDX-License-Identifier: MIT

ARG DEBIAN_BASE_IMAGE=bookworm-slim
ARG DOCKER_REGISTRY

FROM ${DOCKER_REGISTRY}/debian:$DEBIAN_BASE_IMAGE

# For arm64 use:
# ARG CROSS_COMPILER_TARGET_ARCH=arm64
ARG CROSS_COMPILER_TARGET_ARCH=armhf
# For arm64 use:
# ARG CROSS_COMPILE_ARG="aarch64-linux-gnu"
ARG CROSS_COMPILE_ARG="arm-linux-gnueabihf"

ENV CROSS_COMPILE=${CROSS_COMPILE_ARG}-
ENV PKG_CONFIG_PATH=/usr/lib/pkgconfig:/usr/lib/${CROSS_COMPILE_ARG}/pkgconfig

# Make sure we don't get notifications we can't answer during building.
ENV DEBIAN_FRONTEND="noninteractive"

# Enable multiarch
RUN dpkg --add-architecture ${CROSS_COMPILER_TARGET_ARCH}

# Install required packages
RUN apt-get -q -y update \
    && apt-get -q -y install --no-install-recommends \
        ca-certificates \
        gpg \
        wget \
    && apt-get clean \
    && apt-get autoremove -y \
    && rm -rf /var/lib/apt/lists/*

# Upgrade & install required packages

RUN apt-get -q -y update \
    && apt-get -q -y upgrade \
    && apt-get -q -y install --no-install-recommends \
        git \
        openssl \
        wget \
        curl \
        crossbuild-essential-${CROSS_COMPILER_TARGET_ARCH} \
    && apt-get clean \
    && apt-get autoremove -y \
    && rm -rf /var/lib/apt/lists/*

COPY users-groups.sh /users-groups.sh
RUN ./users-groups.sh \
    && rm users-groups.sh

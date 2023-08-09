# Copyright (c) 2019-2023 Toradex AG
# SPDX-License-Identifier: MIT

ARG IMAGE_ARCH=linux/arm/v7
# For arm64v8 use:
# ARG IMAGE_ARCH=linux/arm64/v8
ARG IMAGE_TAG=bookworm-slim
ARG DOCKER_REGISTRY
FROM --platform=$IMAGE_ARCH ${DOCKER_REGISTRY}/debian:$IMAGE_TAG AS base

ARG DEBIAN_FRONTEND=noninteractive
ONBUILD ARG DEBIAN_FRONTEND=noninteractive

# Upgrade & install required packages
RUN apt-get update \
    && apt-get upgrade -y \
    && apt-get install -y --no-install-recommends \
        sudo \
        ca-certificates \
        netbase \
        gpg \
        wget \
    && apt-get clean \
    && apt-get autoremove \
    && rm -rf /var/lib/apt/lists/*

ENV LC_ALL C.UTF-8

# Create 01_nodoc
COPY 01_nodoc /etc/dpkg/dpkg.cfg.d/01_nodoc

# Create 01_buildconfig
RUN echo 'APT::Get::Assume-Yes "true";\n\
    APT::Install-Recommends "0";\n\
    APT::Install-Suggests "0";\n\
    quiet "true";' > /etc/apt/apt.conf.d/01_buildconfig \
    && mkdir -p /usr/share/man/man1

COPY users-groups.sh /users-groups.sh
RUN ./users-groups.sh \
    && rm users-groups.sh

RUN apt-get -y update && apt-get install -y --no-install-recommends \
    neofetch \
    && apt-get clean && apt-get autoremove && rm -rf /var/lib/apt/lists/*

COPY neofetch.conf /root/.config/neofetch/config.conf

CMD ["/bin/bash"]

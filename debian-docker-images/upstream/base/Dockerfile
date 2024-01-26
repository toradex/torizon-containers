# Copyright (c) 2019-2024 Toradex AG
# SPDX-License-Identifier: MIT

ARG DEBIAN_BASE_IMAGE=bookworm-slim
ARG DOCKER_REGISTRY=docker.io

FROM ${DOCKER_REGISTRY}/debian:$DEBIAN_BASE_IMAGE AS base

ARG DEBIAN_FRONTEND=noninteractive
ONBUILD ARG DEBIAN_FRONTEND=noninteractive

# Upgrade & install required packages
RUN apt-get update \
    && apt-get upgrade -y \
    && apt-get install -y --no-install-recommends \
        ca-certificates \
        gpg \
        netbase \
        sudo \
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

RUN apt-get -y update \
    && apt-get install -y --no-install-recommends \
        neofetch \
    && apt-get clean \
    && apt-get autoremove \
    && rm -rf /var/lib/apt/lists/*

COPY neofetch.conf /root/.config/neofetch/config.conf

CMD ["/bin/bash"]

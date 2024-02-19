# Copyright (c) 2019-2023 Toradex AG
# SPDX-License-Identifier: MIT

ARG IMAGE_ARCH=linux/arm/v7
# For arm64v8 use:
# ARG IMAGE_ARCH=linux/arm64/v8
ARG IMAGE_TAG=12.5-slim
ARG TORADEX_SNAPSHOT
ARG USE_TORADEX_SNAPSHOT=1
ARG ADD_TORADEX_REPOSITORY=1
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

FROM base as add_toradex-repository-0

FROM base as add_toradex-repository-1
ARG IMAGE_ARCH
ARG TORADEX_SNAPSHOT
ARG USE_TORADEX_SNAPSHOT
ARG TORADEX_FEED_BASE_URL="https://feeds.toradex.com/debian"

# Add toradex key and package feed
RUN wget -O- ${TORADEX_FEED_BASE_URL}/toradex-debian-repo-19092023.gpg \
    | gpg --dearmor \ 
    | tee /usr/share/keyrings/toradex-debian-repo.gpg

RUN if [ "${USE_TORADEX_SNAPSHOT}" = 1 ]; then \
           TORADEX_FEED_URL="${TORADEX_FEED_BASE_URL}/snapshots/${TORADEX_SNAPSHOT}"; \
       else \
           TORADEX_FEED_URL="${TORADEX_FEED_BASE_URL}"; \
       fi \
    && echo "Types: deb\n\
URIs: https://feeds.toradex.com/debian/snapshots/$TORADEX_SNAPSHOT\n\
Suites: testing\n\
Components: main non-free\n\
Signed-By: /usr/share/keyrings/toradex-debian-repo.gpg" > /etc/apt/sources.list.d/toradex.sources \
    && echo "Package: *\nPin: origin feeds.toradex.com\nPin-Priority: 900" > /etc/apt/preferences.d/toradex-feeds

RUN if [ "${IMAGE_ARCH}" != "linux/arm64/v8" ]; then \
        echo "\
# For armhf, install libdrm-common from Debian stable\n\
Package: libdrm-common\n\
Pin: release a=stable\n\
Pin-Priority: 969"\
       >> /etc/apt/preferences.d/pinned-packages; \
       fi

FROM add_toradex-repository-${ADD_TORADEX_REPOSITORY}

RUN apt-get -y update && apt-get install -y --no-install-recommends \
    neofetch \
    && apt-get clean && apt-get autoremove && rm -rf /var/lib/apt/lists/*

COPY neofetch.conf /root/.config/neofetch/config.conf

CMD ["/bin/bash"]

# Copyright (c) 2019-2023 Toradex AG
# SPDX-License-Identifier: MIT

ARG DEBIAN_BASE_IMAGE=bookworm-slim
ARG TORADEX_SNAPSHOT
ARG USE_TORADEX_SNAPSHOT=1
ARG ADD_TORADEX_REPOSITORY=1
ARG DOCKER_REGISTRY
FROM ${DOCKER_REGISTRY}/debian:$DEBIAN_BASE_IMAGE AS base

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
ARG TORADEX_SNAPSHOT
ARG USE_TORADEX_SNAPSHOT
# Run with executor in network_mode host or privileged
ARG TORADEX_FEED_BASE_URL="http://10.12.1.212:8080/"

# Add toradex key and package feed
RUN wget -O- ${TORADEX_FEED_BASE_URL}/toradex-debian-repo-ti.gpg \
    | gpg --dearmor \ 
    | tee /usr/share/keyrings/toradex-debian-repo-ti.gpg

RUN TORADEX_FEED_URL="20231005T141153Z" && echo "Types: deb\n\
URIs: http://10.12.1.212:8080/snapshots/20231005T141153Z/\n\
Suites: testing\n\
Components: main non-free\n\
Signed-By: /usr/share/keyrings/toradex-debian-repo-ti.gpg" > /etc/apt/sources.list.d/toradex.sources \
    && echo "Package: *\nPin: origin 10.12.1.212\nPin-Priority: 900" > /etc/apt/preferences.d/toradex-feeds

FROM add_toradex-repository-${ADD_TORADEX_REPOSITORY}

RUN apt-get -y update && apt-get install -y --no-install-recommends \
    neofetch \
    && apt-get clean && apt-get autoremove && rm -rf /var/lib/apt/lists/*

COPY neofetch.conf /root/.config/neofetch/config.conf

CMD ["/bin/bash"]

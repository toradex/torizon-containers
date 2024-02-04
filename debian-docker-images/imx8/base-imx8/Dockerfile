# Copyright (c) 2019-2023 Toradex AG
# SPDX-License-Identifier: MIT

ARG BASE_IMAGE_REGISTRY
ARG BASE_IMAGE_NAMESPACE
ARG BASE_IMAGE_NAME
ARG BASE_IMAGE_TAG

FROM $BASE_IMAGE_REGISTRY/$BASE_IMAGE_NAMESPACE/$BASE_IMAGE_NAME:$BASE_IMAGE_TAG

LABEL org.toradex.image.base.registry="docker.io" \
      org.toradex.image.base.namespace="library" \
      org.toradex.image.base.name="debian" \
      org.toradex.image.base.tag.major="12" \
      org.toradex.image.base.tag.minor="4" \
      org.toradex.image.base.tag.patch="" \
      org.toradex.image.variant="slim" \
      org.toradex.image.registry="docker.io" \
      org.toradex.image.namespace="torizon" \
      org.toradex.image.name="debian-imx8" \
      org.toradex.image.base.tag.major="rc" \
      org.toradex.image.base.tag.minor="" \
      org.toradex.image.base.tag.patch="" \
      org.toradex.image.license="MIT" \
      org.toradex.image.arch="linux/arm64/v8"

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

FROM add_toradex-repository-${ADD_TORADEX_REPOSITORY}

RUN apt-get -y update && apt-get install -y --no-install-recommends \
    neofetch \
    && apt-get clean && apt-get autoremove && rm -rf /var/lib/apt/lists/*

COPY neofetch.conf /root/.config/neofetch/config.conf

CMD ["/bin/bash"]

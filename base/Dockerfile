ARG DEBIAN_BASE_IMAGE=bullseye-slim

FROM debian:$DEBIAN_BASE_IMAGE

ARG ARCH_ARG=armhf
# For arm64 use:
# ARG ARCH_ARG=arm64
ARG CROSS_COMPILE_ARG="arm-linux-gnueabihf"
# For arm64 use:
# ARG CROSS_COMPILE_ARG="aarch64-linux-gnu"
ARG TORADEX_SNAPSHOT=20220512T021145Z
ARG USE_TORADEX_SNAPSHOT=1
ARG TORADEX_FEED_BASE_URL="https://feeds.toradex.com/debian"

ENV CROSS_COMPILE=${CROSS_COMPILE_ARG}-
ENV PKG_CONFIG_PATH=/usr/lib/pkgconfig:/usr/lib/${CROSS_COMPILE_ARG}/pkgconfig

# Make sure we don't get notifications we can't answer during building.
ENV DEBIAN_FRONTEND="noninteractive"

# Enable multiarch
RUN dpkg --add-architecture ${ARCH_ARG}

# Enable the Toradex package feed
# (same key is used in https://gitlab.int.toradex.com/rd/torizon-core-containers/debian-docker-image
# if you change the key or feed configuration, please check the other repo!)
RUN apt-get -q -y update \
    && apt-get -q -y install --no-install-recommends \
    ca-certificates \
    && apt-get clean \
    && apt-get autoremove \
    && rm -rf /var/lib/apt/lists/*
ADD ${TORADEX_FEED_BASE_URL}/toradex-debian-repo.gpg /etc/apt/trusted.gpg.d/
RUN chmod 0644 /etc/apt/trusted.gpg.d/toradex-debian-repo.gpg \
    && if [ "${USE_TORADEX_SNAPSHOT}" = 1 ]; then \
            TORADEX_FEED_URL="${TORADEX_FEED_BASE_URL}/snapshots/${TORADEX_SNAPSHOT}"; \
        else \
            TORADEX_FEED_URL="${TORADEX_FEED_BASE_URL}"; \
        fi \
    && echo "deb ${TORADEX_FEED_URL} testing main non-free" >>/etc/apt/sources.list \
    && echo "Package: *\nPin: origin feeds.toradex.com\nPin-Priority: 900" > /etc/apt/preferences.d/toradex-feeds

# Upgrade & install required packages

RUN if [ "${ARCH_ARG}" != "riscv64" ];then \
        apt-get -q -y update \
        && apt-get -q -y upgrade \
        && apt-get -q -y -o Acquire::Http::Dl-limit=1000 install --no-install-recommends \
        git \
        openssl \
        wget \
        curl \
        crossbuild-essential-${ARCH_ARG} \
        && apt-get clean \
        && apt-get autoremove \
        && rm -rf /var/lib/apt/lists/* \
    else \
        apt-get -q -y update \
        && apt-get -q -y upgrade \
        && apt-get -q -y -o Acquire::Http::Dl-limit=1000 install --no-install-recommends \
        git \
        openssl \
        wget \
        curl \
        dpkg-cross \
        gcc-${ARCH_ARG}-linux-gnu \
        g++-${ARCH_ARG}-linux-gnu \
        crossbuild-essential-${ARCH_ARG} \
        && apt-get clean \
        && apt-get autoremove \
        && rm -rf /var/lib/apt/lists/* \
    fi

# Add torizon user to the container
RUN mkdir /home/torizon \
    && addgroup torizon \
    && useradd -s /bin/bash -d /home/torizon -g torizon -G sudo -p $(echo "torizon" | openssl passwd -1 -stdin) torizon \
    && chown -R torizon /home/torizon

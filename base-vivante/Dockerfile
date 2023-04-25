ARG DEBIAN_BASE_IMAGE=bookworm-slim
ARG DOCKER_REGISTRY

FROM ${DOCKER_REGISTRY}debian:$DEBIAN_BASE_IMAGE

ARG ARCH_ARG=armhf
# For arm64 use:
# ARG ARCH_ARG=arm64
ARG CROSS_COMPILE_ARG="arm-linux-gnueabihf"
# For arm64 use:
# ARG CROSS_COMPILE_ARG="aarch64-linux-gnu"
ARG TORADEX_SNAPSHOT=20230328T092716Z
ARG USE_TORADEX_SNAPSHOT=1
ARG TORADEX_FEED_BASE_URL="https://feeds.toradex.com/debian"

ENV CROSS_COMPILE=${CROSS_COMPILE_ARG}-
ENV PKG_CONFIG_PATH=/usr/lib/pkgconfig:/usr/lib/${CROSS_COMPILE_ARG}/pkgconfig

# Make sure we don't get notifications we can't answer during building.
ENV DEBIAN_FRONTEND="noninteractive"

RUN apt-get -y update && apt-get install -y --no-install-recommends \
    apt-utils \
    curl \
    html2text \
    ca-certificates \
    && apt-get -y upgrade \
    && apt-get clean && apt-get autoremove && rm -rf /var/lib/apt/lists/*

ARG LIC_URL=https://www.nxp.com/docs/en/disclaimer/LA_OPT_NXP_SW.html
ARG LIC_FILE_HTML=LA_OPT_NXP_SW.html
ARG LIC_FILE_TXT=LA_OPT_NXP_SW.txt
RUN curl -sSf ${LIC_URL} --output /etc/${LIC_FILE_HTML} \
    && html2text -o /etc/${LIC_FILE_TXT} /etc/${LIC_FILE_HTML} \
    && grep -q "NXP Software License Agreement" /etc/${LIC_FILE_TXT}
ARG ACCEPT_FSL_EULA=0
RUN if [ "${ACCEPT_FSL_EULA}" != "1" ];then \
        echo "Error: This image contains Vivante Packages."; \
        echo "Use --build-arg ACCEPT_FSL_EULA=1 to accept Vivante EULA"; \
        exit 1; \
    else \
        echo "INFO: Vivante EULA has been accepted!"; \
    fi

# Enable multiarch
RUN dpkg --add-architecture ${ARCH_ARG}

# Enable the Toradex package feed
# (same key is used in https://gitlab.int.toradex.com/rd/torizon-core-containers/debian-docker-image
# if you change the key or feed configuration, please check the other repo!)
RUN apt-get -q -y update \
    && apt-get -q -y install --no-install-recommends \
    ca-certificates \
    && apt-get clean \
    && apt-get autoremove -y \
    && rm -rf /var/lib/apt/lists/*
ADD ${TORADEX_FEED_BASE_URL}/toradex-debian-repo.gpg /etc/apt/trusted.gpg.d/
RUN chmod 0644 /etc/apt/trusted.gpg.d/toradex-debian-repo.gpg \
    && if [ "${USE_TORADEX_SNAPSHOT}" = 1 ]; then \
            TORADEX_FEED_URL="${TORADEX_FEED_BASE_URL}/snapshots/${TORADEX_SNAPSHOT}"; \
        else \
            TORADEX_FEED_URL="${TORADEX_FEED_BASE_URL}"; \
        fi \
    && echo "Types: deb\n\
URIs: https://feeds.toradex.com/debian/snapshots/$TORADEX_SNAPSHOT\n\
Suites: testing\n\
Components: main non-free" > /etc/apt/sources.list.d/toradex.sources \
    && echo "Package: *\nPin: origin feeds.toradex.com\nPin-Priority: 900" > /etc/apt/preferences.d/toradex-feeds

# Upgrade & install required packages

RUN if [ "${ARCH_ARG}" != "riscv64" ];then \
        apt-get -q -y update \
        && apt-get -q -y upgrade \
        && apt-get -q -y install --no-install-recommends \
        git \
        openssl \
        wget \
        curl \
        crossbuild-essential-${ARCH_ARG} \
        && apt-get clean \
        && apt-get autoremove -y \
        && rm -rf /var/lib/apt/lists/*; \
    else \
        apt-get -q -y update \
        && apt-get -q -y upgrade \
        && apt-get -q -y install --no-install-recommends \
        git \
        openssl \
        wget \
        curl \
        gnupg \
        dpkg-cross \
        gcc-${ARCH_ARG}-linux-gnu \
        g++-${ARCH_ARG}-linux-gnu \
        debian-ports-archive-keyring && \
        # add the riscv64 ports packages
        echo "Types: deb\n\
URIs: http://deb.debian.org/debian-ports\n\
Suites: sid\n\
Components: main\n\
Architectures: riscv64" > /etc/apt/sources.list.d/debian-ports.sources \
        && cat /etc/apt/sources.list.d/debian-ports.sources \
        && apt-get clean \
        && apt-get autoremove -y \
        && rm -rf /var/lib/apt/lists/*; \
    fi

COPY users-groups.sh /users-groups.sh
RUN ./users-groups.sh \
    && rm users-groups.sh


ARG DEBIAN_BASE_IMAGE=bookworm-slim
ARG DOCKER_REGISTRY

FROM ${DOCKER_REGISTRY}/debian:$DEBIAN_BASE_IMAGE

ARG IMAGE_ARCH=armhf
# For arm64 use:
# ARG IMAGE_ARCH=arm64
ARG CROSS_COMPILE_ARG="arm-linux-gnueabihf"
# For arm64 use:
# ARG CROSS_COMPILE_ARG="aarch64-linux-gnu"
ARG TORADEX_SNAPSHOT
ARG USE_TORADEX_SNAPSHOT=1
ARG TORADEX_FEED_BASE_URL="https://feeds.toradex.com/debian"

ENV CROSS_COMPILE=${CROSS_COMPILE_ARG}-
ENV PKG_CONFIG_PATH=/usr/lib/pkgconfig:/usr/lib/${CROSS_COMPILE_ARG}/pkgconfig

# Make sure we don't get notifications we can't answer during building.
ENV DEBIAN_FRONTEND="noninteractive"

RUN apt-get -y update && apt-get install -y --no-install-recommends \
    apt-utils \
    ca-certificates \
    gpg \
    html2text \
    wget \
    && apt-get -y upgrade \
    && apt-get clean && apt-get autoremove && rm -rf /var/lib/apt/lists/*

ARG LIC_URL=https://www.nxp.com/docs/en/disclaimer/LA_OPT_NXP_SW.html
ARG LIC_FILE_HTML=LA_OPT_NXP_SW.html
ARG LIC_FILE_TXT=LA_OPT_NXP_SW.txt
RUN wget ${LIC_URL} -O /etc/${LIC_FILE_HTML} \
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
RUN dpkg --add-architecture ${IMAGE_ARCH}

# Add toradex key and package feed
RUN wget -O- ${TORADEX_FEED_BASE_URL}/toradex-debian-repo.gpg \
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

# Upgrade & install required packages

RUN apt-get -q -y update \
    && apt-get -q -y upgrade \
    && apt-get -q -y install --no-install-recommends \
    git \
    openssl \
    wget \
    curl \
    crossbuild-essential-${IMAGE_ARCH} \
    && apt-get clean \
    && apt-get autoremove -y \
    && rm -rf /var/lib/apt/lists/*

COPY users-groups.sh /users-groups.sh
RUN ./users-groups.sh \
    && rm users-groups.sh


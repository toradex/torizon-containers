ARG DEBIAN_BASE_IMAGE=bookworm-20230208-slim

FROM debian:$DEBIAN_BASE_IMAGE

ARG ARCH_ARG=armhf
# For arm64 use:
# ARG ARCH_ARG=arm64
ARG CROSS_COMPILE_ARG="arm-linux-gnueabihf"
# For arm64 use:
# ARG CROSS_COMPILE_ARG="aarch64-linux-gnu"
ARG TORADEX_SNAPSHOT=20230224T051016Z
ARG USE_TORADEX_SNAPSHOT=1
ARG TORADEX_FEED_BASE_URL="https://feeds.toradex.com/debian"

ENV CROSS_COMPILE=${CROSS_COMPILE_ARG}-
ENV PKG_CONFIG_PATH=/usr/lib/pkgconfig:/usr/lib/${CROSS_COMPILE_ARG}/pkgconfig

# Make sure we don't get notifications we can't answer during building.
ENV DEBIAN_FRONTEND="noninteractive"

# Debian Bookworm is not yet a stable distribution at the moment of this writing;
# therefore its package list may change in incompatible ways with Torizon software.
# Let's lock Torizon containers to a known snapshot of the Bookworm package list as a workaround.
# In the bookworm-slim images, the snapshot is commented out and the main feed
# is used as the URI. With the following sed command, we reverse this and use
# the snapshot as the URI.
RUN grep -q sid /etc/apt/sources.list.d/debian.sources || sed -i \
    -e 's/^#/URIs:/g' \
    -e '/deb.debian.org/d' \
    -e '/^Signed-By/a check-valid-until: no' \
    /etc/apt/sources.list.d/debian.sources

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


# Debian Bookworm has higher versions of some same packages than the Toradex feed.
# Force install Bookworm versions of these packages in order to keep the container equivalent for 32 and 64 bits architectures.
# This packages are necessary for the Qt6 C++ templates of the IDE Extension V2 (ApolloX)
RUN HOLD_PKGS='libdrm-common libdrm-amdgpu1 libdrm2' \
    && apt-get -y update \
    && for P in $HOLD_PKGS ; do \
        echo ${P}=$(apt-cache show $P | sed -r -e '/^Version:/!d' -e 's/.* //' -e '/toradex/d' -e 'q') ; \
    done | xargs -r apt-get install -y --no-install-recommends \
    && apt-mark hold $HOLD_PKGS \
    && apt-get clean && apt-get autoremove && rm -rf /var/lib/apt/lists/*


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
        && apt-get autoremove -y \
        && rm -rf /var/lib/apt/lists/*; \
    else \
        apt-get -q -y update \
        && apt-get -q -y upgrade \
        && apt-get -q -y -o Acquire::Http::Dl-limit=1000 install --no-install-recommends \
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

# Add torizon user to the container
RUN mkdir /home/torizon \
    && groupadd torizon \
    && useradd -s /bin/bash -d /home/torizon -g torizon -G sudo -p $(echo "torizon" | openssl passwd -1 -stdin) torizon \
    && chown -R torizon /home/torizon

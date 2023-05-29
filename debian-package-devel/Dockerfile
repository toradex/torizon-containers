ARG IMAGE_ARCH=arm64v8
ARG IMAGE_TAG=bookworm-slim
ARG DOCKER_REGISTRY

FROM --platform=$IMAGE_ARCH ${DOCKER_REGISTRY}/debian:$IMAGE_TAG AS base

ARG DEBIAN_FRONTEND=noninteractive

# Install a base set of build tools
RUN apt-get update && apt-get install -y -o Acquire::http::Dl-Limit=1000 --no-install-recommends \
    aptly \
    bash-completion \
    build-essential:native \
    ca-certificates \
    curl \
    debconf \
    debhelper \
    debmake \
    debootstrap \
    devscripts \
    dh-make \
    dh-python \
    dpkg-dev \
    equivs \
    fakeroot \
    git \
    git-buildpackage \
    gnupg \
    libpng16-16 \
    libtool \
    netbase \
    openssh-client \
    pbuilder \
    pkg-config \
    procps \
    python3-debian \
    rsync \
    sudo \
    vim \
    wget \
    && rm -rf /var/lib/apt/lists/*

# Use UID 1000 to build packages.
# Replace to the id of your user in the host.
RUN useradd debian -u 1000 -m -G tty,sudo,dialout,users,plugdev

# Tell sudo not to ask for passwords for users of the sudo group
RUN sed -i '/^%sudo\>/s/) *ALL$/) NOPASSWD: ALL/' /etc/sudoers && visudo -c


# Setup access to the Toradex package feed
RUN echo "Types: deb\n\
URIs: https://feeds.toradex.com/debian\n\
Suites: testing\n\
Components: main non-free" > /etc/apt/sources.list.d/toradex.sources
RUN echo "Package: *\nPin: origin feeds.toradex.com\nPin-Priority: 900" > /etc/apt/preferences.d/toradex-feeds
RUN wget -P /etc/apt/trusted.gpg.d https://feeds.toradex.com/debian/toradex-debian-repo.gpg

# Apply eventual upgrades to installed packages
RUN apt-get update && apt-get upgrade -y && rm -rf /var/lib/apt/lists/*

ENV LC_ALL C.UTF-8

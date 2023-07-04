# Copyright (c) 2019-2023 Toradex AG
# SPDX-License-Identifier: MIT

ARG IMAGE_ARCH=armhf
# For arm64 use:
# ARG IMAGE_ARCH=arm64
ARG IMAGE_TAG=3-bookworm
ARG DOCKER_REGISTRY=torizon

FROM ${DOCKER_REGISTRY}/debian-cross-toolchain-$IMAGE_ARCH:$IMAGE_TAG

# SSH
EXPOSE 22

# Install required packages
RUN apt-get -q -y update \
    && apt-get -q -y install --no-install-recommends \
    openssh-server \
    ssh \
    rsync \
    zip \
    && apt-get clean \
    && apt-get autoremove -y \
    && rm -rf /var/lib/apt/lists/*

# create folders needed for the different components
RUN mkdir /var/run/sshd && mkdir /home/build && mkdir /home/build/sysroot

# configures SSH access to the container and sets environment by default
RUN sed 's@session\s*required\s*pam_loginuid.so@session optional pam_loginuid.so@g' -i /etc/pam.d/sshd \
    #&& useradd -s /bin/bash -d /home/build -g root -G sudo,ssh -p $(echo "build" | openssl passwd -1 -stdin) build \
    && useradd -s /bin/bash -d /home/build -g root -G sudo -p $(echo "build" | openssl passwd -1 -stdin) build \
    && chown -R build /home/build

CMD ["/usr/sbin/sshd", "-D"]

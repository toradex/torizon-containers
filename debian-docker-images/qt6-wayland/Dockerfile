# ARG IMAGE_ARCH=linux/arm
# For arm64v8 use:
ARG IMAGE_ARCH=linux/arm64
# ARG BASE_NAME=wayland-base
# For arm64v8 with Vivante use:
ARG BASE_NAME=wayland-base-vivante
ARG IMAGE_TAG=3
ARG DOCKER_REGISTRY=torizon
FROM --platform=$IMAGE_ARCH $DOCKER_REGISTRY/$BASE_NAME:$IMAGE_TAG AS base

COPY kms-setup.sh /usr/bin/kms-setup.sh

RUN apt-get -y update && apt-get install -y --no-install-recommends \
    apt-utils \
    libqt6gui6-gles \
    && apt-get -y upgrade \
    && apt-get clean && apt-get autoremove && rm -rf /var/lib/apt/lists/*

# Install remaining dependencies required to run qtbase
RUN apt-get -y update && apt-get install -y --no-install-recommends \
       libfontconfig1-dev \
       libqt6concurrent6 \
       libqt6dbus6 \
       libqt6network6 \
       libqt6printsupport6 \
       libqt6sql6 \
       libqt6test6 \
       libqt6widgets6 \
       libqt6xml6 \
       libqt6qml6 \
       libqt6opengl6 \
       libqt6quicktest6 \
       libqt6quickwidgets6 \
       qt6-wayland \
       && apt-get clean && apt-get autoremove && rm -rf /var/lib/apt/lists/*

ENV QT_QPA_PLATFORM="wayland"

# EGLFS configuration
ENV QT_QPA_EGLFS_INTEGRATION="eglfs_kms"
ENV QT_QPA_EGLFS_KMS_ATOMIC="1"
ENV QT_QPA_EGLFS_KMS_CONFIG="/etc/kms.conf"

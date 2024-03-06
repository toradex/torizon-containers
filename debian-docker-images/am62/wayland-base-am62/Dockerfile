# Copyright (c) 2019-2024 Toradex AG
# SPDX-License-Identifier: MIT

ARG BASE_NAME=debian
ARG IMAGE_TAG=3-bookworm
ARG DOCKER_REGISTRY=torizon
FROM $DOCKER_REGISTRY/torizon/$BASE_NAME:$IMAGE_TAG AS base

RUN apt-get -y update && apt-get install -y \
    libd3dadapter9-mesa \
    libd3dadapter9-mesa-dev \
    libegl1-mesa \
    libegl1-mesa-dev \
    libegl-mesa0 \
    libgbm1 \
    libgbm-dev \
    libgl1-mesa-dev \
    libgl1-mesa-dri \
    libgl1-mesa-glx \
    libglapi-mesa \
    libgles2-mesa \
    libgles2-mesa-dev \
    libglx-mesa0 \
    libosmesa6 \
    libosmesa6-dev \
    libpru-pssp-dev \
    libwayland-egl1-mesa \
    libwayland-client0 \
    libwayland-server0 \
    mesa-utils-extra \
    mesa-common-dev \
    mesa-opencl-icd \
    mesa-utils-extra \
    ti-img-rogue-umlibs-am62 \
    && apt-get clean && apt-get autoremove && rm -rf /var/lib/apt/lists/*

# Make sure the user can access DRM and video devices
RUN usermod -a -G video,render torizon

ENV WAYLAND_USER="torizon"
ENV XDG_RUNTIME_DIR="/tmp/1000-runtime-dir"
ENV WAYLAND_DISPLAY="wayland-0"
ENV DISPLAY=":0"

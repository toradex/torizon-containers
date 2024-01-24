# Copyright (c) 2019-2023 Toradex AG
# SPDX-License-Identifier: MIT

ARG BASE_NAME=debian
ARG IMAGE_TAG=3-bookworm
ARG DOCKER_REGISTRY=torizon
FROM $DOCKER_REGISTRY/torizon/$BASE_NAME:$IMAGE_TAG AS base

RUN apt-get -y update && apt-get install -y --no-install-recommends \
    libdrm2 \
    libwayland-client0 \
    libwayland-server0 \
    libglx-mesa0 \
    libegl1 \
    mesa-utils-extra \
    && apt-get clean && apt-get autoremove && rm -rf /var/lib/apt/lists/*

# Make sure the user can access DRM and video devices
RUN usermod -a -G video,render torizon

ENV WAYLAND_USER="torizon"
ENV XDG_RUNTIME_DIR="/tmp/1000-runtime-dir"
ENV WAYLAND_DISPLAY="wayland-0"
ENV DISPLAY=":0"

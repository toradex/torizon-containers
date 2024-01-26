# Copyright (c) 2019-2024 Toradex AG
# SPDX-License-Identifier: MIT

ARG BASE_NAME=weston
ARG DOCKER_REGISTRY=docker.io
ARG IMAGE_TAG=rc

FROM $DOCKER_REGISTRY/torizon/$BASE_NAME:$IMAGE_TAG AS base

COPY entry.sh /usr/bin/
RUN mkdir -p /etc/xdg/weston/
COPY weston.ini /etc/xdg/weston/
COPY save-calibration.sh /usr/bin/

USER root

ENTRYPOINT ["/usr/bin/entry.sh"]

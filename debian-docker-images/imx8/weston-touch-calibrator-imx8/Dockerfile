# Copyright (c) 2019-2023 Toradex AG
# SPDX-License-Identifier: MIT

ARG BASE_NAME=weston-vivante
ARG IMAGE_TAG=3
ARG DOCKER_REGISTRY=torizon
FROM $DOCKER_REGISTRY/torizon/$BASE_NAME:$IMAGE_TAG AS base

LABEL org.toradex.image.base.namespace="torizon" \
      org.toradex.image.base.name="weston-imx8" \
      org.toradex.image.base.tag="rc" \
      org.toradex.image.namespace="torizon" \
      org.toradex.image.name="weston-touch-calibrator-imx8" \
      org.toradex.image.tag="rc" \
      org.toradex.image.license="MIT" \
      org.toradex.image.arch="linux/arm64/v8"

COPY entry.sh /usr/bin/
RUN mkdir -p /etc/xdg/weston/
COPY weston.ini /etc/xdg/weston/
COPY save-calibration.sh /usr/bin/

USER root

ENTRYPOINT ["/usr/bin/entry.sh"]

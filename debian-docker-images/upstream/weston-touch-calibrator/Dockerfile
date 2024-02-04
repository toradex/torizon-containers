# Copyright (c) 2019-2023 Toradex AG
# SPDX-License-Identifier: MIT

ARG BASE_IMAGE_REGISTRY
ARG BASE_IMAGE_NAMESPACE
ARG BASE_IMAGE_NAME
ARG BASE_IMAGE_TAG

FROM $BASE_IMAGE_REGISTRY/$BASE_IMAGE_NAMESPACE/$BASE_IMAGE_NAME:$BASE_IMAGE_TAG

LABEL org.toradex.image.base.registry="docker.io" \
      org.toradex.image.base.namespace="torizon" \
      org.toradex.image.base.name="weston" \
      org.toradex.image.base.tag.major="rc" \
      org.toradex.image.base.tag.minor="" \
      org.toradex.image.base.tag.patch="" \
      org.toradex.image.base.tag.variant="bookworm" \
      org.toradex.image.registry="docker.io" \
      org.toradex.image.namespace="torizon" \
      org.toradex.image.name="weston-touch-calibrator" \
      org.toradex.image.tag.major="rc" \
      org.toradex.image.tag.minor="" \
      org.toradex.image.tag.patch="" \
      org.toradex.image.tag.variant="" \
      org.toradex.image.license="MIT" \
      org.toradex.image.arch="linux/amd64,linux/arm64/v8,linux/arm/v7"

COPY entry.sh /usr/bin/
RUN mkdir -p /etc/xdg/weston/
COPY weston.ini /etc/xdg/weston/
COPY save-calibration.sh /usr/bin/

USER root

ENTRYPOINT ["/usr/bin/entry.sh"]

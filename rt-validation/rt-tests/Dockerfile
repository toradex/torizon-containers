# Copyright (c) 2019-2023 Toradex AG
# SPDX-License-Identifier: MIT

ARG IMAGE_ARCH=linux/arm/v7
# For IMX8 use IMAGE_ARCH=linux/arm64/v8
ARG IMAGE_TAG=3-bookworm
ARG DOCKER_REGISTRY=torizon

FROM --platform=$IMAGE_ARCH $DOCKER_REGISTRY/debian:$IMAGE_TAG

RUN apt update \
    && apt install -y --no-install-recommends \
        gnuplot \
        rt-tests \
    && apt clean \
    && rm -rf /var/lib/apt/lists/*

COPY rt-tests.sh /rt-tests.sh

CMD ["/rt-tests.sh"]

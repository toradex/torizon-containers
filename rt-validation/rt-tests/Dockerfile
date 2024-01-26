# Copyright (c) 2019-2024 Toradex AG
# SPDX-License-Identifier: MIT

ARG DEBIAN_BASE_IMAGE=3-bookworm
ARG DOCKER_REGISTRY=torizon

FROM $DOCKER_REGISTRY/debian:$DEBIAN_BASE_IMAGE

RUN apt update \
    && apt install -y --no-install-recommends \
        gnuplot \
        rt-tests \
    && apt clean \
    && rm -rf /var/lib/apt/lists/*

COPY rt-tests.sh /rt-tests.sh

CMD ["/rt-tests.sh"]

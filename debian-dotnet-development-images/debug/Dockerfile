# Copyright (c) 2019-2023 Toradex AG
# SPDX-License-Identifier: MIT

ARG BASE_NAME
ARG IMAGE_ARCH
ARG DOCKER_REGISTRY
ARG DEBIAN_BASE_IMAGE_TAG=bookworm-slim
ARG DOTNET_BASE_IMAGE_TAG

FROM --platform=${IMAGE_ARCH} ${DOCKER_REGISTRY}/debian:${DEBIAN_BASE_IMAGE_TAG} AS build

RUN apt-get update && apt-get install -y \
    --no-install-recommends --no-install-suggests \
    ca-certificates \
    curl \
    && apt-get clean && apt-get autoremove \
    && rm -rf /var/lib/apt/lists/*

ARG DOTNET_DEBUGGER_RELEASE

RUN curl -sSL https://aka.ms/getvsdbgsh | /bin/sh /dev/stdin -v ${DOTNET_DEBUGGER_RELEASE} -l /vsdbg

FROM --platform=${IMAGE_ARCH} ${DOCKER_REGISTRY}/${BASE_NAME}:${DOTNET_BASE_IMAGE_TAG}

COPY --from=build /vsdbg /vsdbg

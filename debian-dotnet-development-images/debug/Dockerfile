# Copyright (c) 2019-2024 Toradex AG
# SPDX-License-Identifier: MIT

ARG BASE_NAME
ARG DOCKER_REGISTRY
ARG IMAGE_TAG
ARG DOTNET_BASE_IMAGE_TAG

FROM ${DOCKER_REGISTRY}/torizon/debian:${IMAGE_TAG} AS build

RUN apt-get update && apt-get install -y \
    --no-install-recommends --no-install-suggests \
    ca-certificates \
    curl \
    && apt-get clean && apt-get autoremove \
    && rm -rf /var/lib/apt/lists/*

ARG DOTNET_DEBUGGER_RELEASE

RUN curl -sSL https://aka.ms/getvsdbgsh | /bin/sh /dev/stdin -v ${DOTNET_DEBUGGER_RELEASE} -l /vsdbg

FROM ${DOCKER_REGISTRY}/torizon/${BASE_NAME}:${DOTNET_BASE_IMAGE_TAG}

COPY --from=build /vsdbg /vsdbg

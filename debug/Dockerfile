ARG BASE_NAME=
ARG BASE_TAG=
ARG IMAGE_ARCH=

FROM --platform=$IMAGE_ARCH artifactory-horw.int.toradex.com/dockerhub-proxy-horw/debian:bookworm-slim AS build

RUN apt-get update && apt-get install -y \
    --no-install-recommends --no-install-suggests \
    ca-certificates \
    curl \
    && apt-get clean && apt-get autoremove \
    && rm -rf /var/lib/apt/lists/*

ARG DEBUGGER_RELEASE=

RUN curl -sSL https://aka.ms/getvsdbgsh | /bin/sh /dev/stdin -v ${DEBUGGER_RELEASE} -l /vsdbg

FROM torizon/$BASE_NAME:$BASE_TAG

COPY --from=build /vsdbg /vsdbg
ARG BASE_NAME=
ARG BASE_TAG=

FROM artifactory-horw.int.toradex.com/dockerhub-proxy-horw/debian:bookworm-slim AS build

RUN apt-get update && apt-get install -y \
    --no-install-recommends --no-install-suggests \
    ca-certificates \
    curl \
    unzip \
    && apt-get clean && apt-get autoremove \
    && rm -rf /var/lib/apt/lists/*

ARG DEBUGGER_RELEASE=
ARG DEBUGGER_ARCH=

RUN curl -L -o /vsdbg.zip https://vsdebugger.azureedge.net/vsdbg-$DEBUGGER_RELEASE/vsdbg-$DEBUGGER_ARCH.zip \
    && mkdir /vsdbg

WORKDIR /vsdbg

RUN unzip /vsdbg.zip \
    && chmod a+x /vsdbg/vsdbg

FROM torizon/$BASE_NAME:$BASE_TAG

COPY --from=build /vsdbg /vsdbg

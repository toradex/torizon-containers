ARG IMAGE_ARCH=
ARG DOCKER_REGISTRY=torizon
ARG IMAGE_TAG=bookworm-slim

FROM --platform=$IMAGE_ARCH ${DOCKER_REGISTRY}/debian:${IMAGE_TAG} AS build

# define specific .NET version you want to use or "latest"
ARG DOTNET_VERSION=latest
# can be arm or arm64
ARG DOTNET_ARCHITECTURE=
# can be dotnet or aspnetcore
ARG DOTNET_RUNTIME=dotnet

# Make sure we don't get notifications we can't answer during building.
ENV DEBIAN_FRONTEND="noninteractive"

RUN apt-get update && apt-get install -y \
    --no-install-recommends --no-install-suggests \
    curl \
    ca-certificates \
    && apt-get clean && apt-get autoremove \
    && rm -rf /var/lib/apt/lists/*

RUN curl -L -o /dotnet-install.sh https://dot.net/v1/dotnet-install.sh \
    && chmod a+x /dotnet-install.sh \
    && mkdir dotnet

WORKDIR /dotnet

RUN /dotnet-install.sh --install-dir /dotnet --runtime ${DOTNET_RUNTIME} --version ${DOTNET_VERSION} --os linux --architecture ${DOTNET_ARCHITECTURE} --no-path

FROM --platform=$IMAGE_ARCH artifactory-horw.int.toradex.com/dockerhub-proxy-horw/torizon/debian:3-bookworm

ENV DOTNET_ROOT=/dotnet
ENV PATH=$PATH:/dotnet

RUN apt-get update && apt-get install -y \
    --no-install-recommends --no-install-suggests \
    libicu72 \
    libssl3 \
    && apt-get clean && apt-get autoremove \
    && rm -rf /var/lib/apt/lists/*

COPY --from=build /dotnet /dotnet

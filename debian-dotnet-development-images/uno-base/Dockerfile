ARG IMAGE_ARCH=
ARG BASE_NAME=
ARG IMAGE_TAG=
ARG BASE_TAG_WAYLAND=
ARG BASE_NAME_WAYLAND=
ARG DOCKER_REGISTRY=torizon

FROM ${DOCKER_REGISTRY}/$BASE_NAME:$IMAGE_TAG AS dotnet

FROM --platform=$IMAGE_ARCH ${DOCKER_REGISTRY}/${BASE_NAME_WAYLAND}:${BASE_TAG_WAYLAND} AS base

ENV DOTNET_ROOT=/dotnet
ENV PATH=$PATH:/dotnet

COPY --from=dotnet /dotnet /dotnet

RUN apt-get -y update && apt-get install -y --no-install-recommends \
    apt-utils \
    && apt-get -y upgrade \
    && apt-get clean && apt-get autoremove && rm -rf /var/lib/apt/lists/*

# install .net uno platform dependencies
RUN apt-get -y update && apt-get install -y --no-install-recommends \
    libicu72 \
    curl \
    gettext \
    apt-transport-https \
    libx11-6 \
	libunwind-13 \
    icu-devtools \
	libfontconfig1 \
	libgtk-3-0 \
    libgtk-3-bin \
    libgtk-3-common \
	libdrm2 \
	libinput10 \
    libssl3 \
	&& apt-get clean && apt-get autoremove && rm -rf /var/lib/apt/lists/*

# fix mime types
RUN apt-get -y update && apt-get --reinstall install -y --no-install-recommends \
    shared-mime-info \
	&& apt-get clean && apt-get autoremove && rm -rf /var/lib/apt/lists/*

RUN printf "{\n  \"device\": \"/dev/dri/card1\",\n  \"outputs\": [ { \"name\": \"HDMI1\" } ],\n  \"hwcursor\": false\n}" > /etc/kms-imx6.conf \
    && printf "{\n  \"hwcursor\": false\n}" > /etc/kms-imx8.conf \
    && ln -s /etc/kms-imx8.conf /etc/kms-imx7.conf

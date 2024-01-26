# Copyright (c) 2019-2024 Toradex AG
# SPDX-License-Identifier: MIT

ARG BASE_NAME
ARG IMAGE_TAG
ARG DOCKER_REGISTRY

FROM $DOCKER_REGISTRY/torizon/$BASE_NAME:$IMAGE_TAG AS base

RUN apt-get -y update && apt-get install -y --no-install-recommends \
    gputop \
    libgpuperfcnt \
    curl \
    html2text \
    && apt-get -y upgrade \
    && apt-get clean && apt-get autoremove && rm -rf /var/lib/apt/lists/*

ARG LIC_URL=https://www.nxp.com/docs/en/disclaimer/LA_OPT_NXP_SW.html
ARG LIC_FILE_HTML=LA_OPT_NXP_SW.html
ARG LIC_FILE_TXT=LA_OPT_NXP_SW.txt
RUN curl -sSf ${LIC_URL} --output /etc/${LIC_FILE_HTML} \
    && html2text -o /etc/${LIC_FILE_TXT} /etc/${LIC_FILE_HTML} \
    && grep -q "NXP Software License Agreement" /etc/${LIC_FILE_TXT}
ARG ACCEPT_FSL_EULA=0
RUN if [ "${ACCEPT_FSL_EULA}" != "1" ];then \
        echo "Error: This image contains iMX Packages."; \
        echo "Use --build-arg ACCEPT_FSL_EULA=1 to accept Vivante EULA"; \
        exit 1; \
    else \
        echo "INFO: Vivante EULA has been accepted!"; \
    fi

CMD [ "gputop"]

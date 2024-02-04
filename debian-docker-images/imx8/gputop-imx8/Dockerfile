# Copyright (c) 2019-2023 Toradex AG
# SPDX-License-Identifier: MIT

ARG BASE_IMAGE_REGISTRY
ARG BASE_IMAGE_NAMESPACE
ARG BASE_IMAGE_NAME
ARG BASE_IMAGE_TAG

FROM $BASE_IMAGE_REGISTRY/$BASE_IMAGE_NAMESPACE/$BASE_IMAGE_NAME:$BASE_IMAGE_TAG

LABEL org.toradex.image.base.namespace="torizon" \
      org.toradex.image.base.name="debian-imx8" \
      org.toradex.image.base.tag="rc" \
      org.toradex.image.namespace="torizon" \
      org.toradex.image.name="gputop-imx8" \
      org.toradex.image.tag="rc" \
      org.toradex.image.license="MIT" \
      org.toradex.image.arch="linux/arm64/v8"

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

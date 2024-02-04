ARG BASE_NAME=wayland-base
ARG DOCKER_REGISTRY
ARG IMAGE_TAG=3

FROM $DOCKER_REGISTRY/torizon/$BASE_NAME:$IMAGE_TAG

LABEL org.toradex.image.base.registry="docker.io" \
      org.toradex.image.base.namespace="torizon" \
      org.toradex.image.base.name="wayland-base-am62" \
      org.toradex.image.base.tag.major="rc" \
      org.toradex.image.base.tag.minor="" \
      org.toradex.image.base.tag.patch="" \
      org.toradex.image.variant="" \
      org.toradex.image.registry="docker.io" \
      org.toradex.image.namespace="torizon" \
      org.toradex.image.name="cog-am62" \
      org.toradex.image.base.tag.major="rc" \
      org.toradex.image.base.tag.minor="" \
      org.toradex.image.base.tag.patch="" \
      org.toradex.image.license="MIT" \
      org.toradex.image.arch="linux/arm64/v8"

RUN apt-get -y update && \
    apt-get install -y --no-install-recommends cog && \
    apt-get clean && apt-get autoremove && rm -rf /var/lib/apt/lists/*

COPY start-cog.sh /usr/bin/start-browser

USER torizon

ENV DISPLAY=:0

ENTRYPOINT ["/usr/bin/start-browser"]
CMD ["http://www.toradex.com"]

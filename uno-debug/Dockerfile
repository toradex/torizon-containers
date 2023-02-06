ARG BASE_NAME=
ARG BASE_TAG=
ARG BASE_NAME_DEBUG=

FROM torizon/$BASE_NAME_DEBUG:$BASE_TAG AS dotnet

FROM torizon/$BASE_NAME:$BASE_TAG AS base

COPY --from=dotnet /vsdbg /vsdbg

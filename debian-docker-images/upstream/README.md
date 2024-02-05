# Upstream containers

Upstream containers don't use third-party package feeds. All packages is from the upstream Debian project.
Upstream images can be used with any target devices, but may not have support for hardware acceleration.

The support strategy for new target devices (or machines in the Yocto lingo) is to add a new platform folder (like imx8 or am62) and customize it to suit that new device. This step will also most likely involve adding a new debian package feed.

## base

## chromium

## cog

## qt5-wayland

## qt5-wayland-examples

## qt6-wayland

## qt6-wayland-examples

## chromium-tests-am62

## graphics-tests

## wayland-base

## weston

Weston for the upstream containers can use multiple backends

## weston-touch-calibrator

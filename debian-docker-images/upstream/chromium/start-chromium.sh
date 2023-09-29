#!/bin/sh

# default URL
URL="www.toradex.com"

chromium --enable-features=UseOzonePlatform --ozone-platform=wayland --no-sandbox --use-gl=egl --disable-seccomp-filter-sandbox --test-type --allow-insecure-localhost --disable-notifications --kiosk $URL

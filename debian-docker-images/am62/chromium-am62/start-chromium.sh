#!/bin/sh

# default URL
URL="www.toradex.com"

chromium --test-type --allow-insecure-localhost --disable-notifications --check-for-update-interval=315360000 --disable-seccomp-filter-sandbox --use-gl=egl --in-process-gpu --no-sandbox $URL

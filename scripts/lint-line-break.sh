#!/usr/bin/env sh

# shellcheck disable=SC2156
find . -type f -not -path '*/.*' -not -name '*.gz' -exec sh -c 'file "{}" | grep -q "CRLF" && echo "Error: {} has non-Unix line endings"' \;

#!/usr/bin/env sh

find . -type f -not -path '*/.*' -not -name '*.gz' -exec sh -c 'file "{}" | grep -q "CRLF" && echo "Error: {} has non-Unix line endings"' \;

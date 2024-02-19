#!/bin/sh

# default URL
URL="www.toradex.com"

if [ -n "$1" ]; then
  URL=$1
fi

eval exec cog "$URL"

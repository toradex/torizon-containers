#!/bin/sh

# default URL
URL="www.toradex.com"

# Enable GPU support by default
ENABLE_GPU=1

# default parms for kiosk mode
chromium_parms_base="--test-type --allow-insecure-localhost --disable-notifications --check-for-update-interval=315360000 "
chromium_parms="--kiosk "

# Additional params should be stacked chromium_parms_extended="$chromium_parms_extended ..."
chromium_parms_extended=""

# FIXME: TOR-1426: work around seccomp-bpf failures in chromium.
chromium_parms_extended="$chromium_parms_extended --disable-seccomp-filter-sandbox"

for arg in "$@"
do
    case $arg in
    --window-mode)
        chromium_parms="--start-maximized --app="
        shift
        ;;
    --browser-mode)
        chromium_parms="--start-maximized "
        shift
        ;;
    --virtual-keyboard)
        # Load the virtual keyboard
        chromium_parms_extended="$chromium_parms_extended --load-extension=/chrome-extensions/chrome-virtual-keyboard-master"
        shift
        ;;
    --disable-gpu)
        # Disable GPU support completely
        ENABLE_GPU=0
        shift
        ;;
    --disable-gpu-compositing)
        # Disable GPU Compositing only
        chromium_parms_extended="$chromium_parms_extended --disable-gpu-compositing"
        shift
        ;;
    esac
done

# Setup GPU flags
# Use EGL and OpenGL ES instead of GLX and regular OpenGL
chromium_parms_extended="$chromium_parms_extended --use-gl=egl"

# Ozone parameters
chromium_parms_extended="$chromium_parms_extended --enable-features=UseOzonePlatform --ozone-platform=wayland"

# Run the GPU process as a thread in the browser process.
# This is required to use Wayland EGL path
# See: https://github.com/OSSystems/meta-browser/issues/510#issuecomment-854653930
chromium_parms_extended="$chromium_parms_extended --in-process-gpu"

if [ ! -z "$1" ]; then
    URL=$1
fi

exec chromium $chromium_parms_base $chromium_parms_extended $chromium_parms$URL

#!/usr/bin/env bash

# exits immediately if any error has non-zero exit
set -e
# don't allow forward referencing variables
set -u
# error the whole pipeline if a subcommand fails
set -o pipefail

declare -a WESTON_EXTRA_ARGS

# Only use long options to keep consistency.
# `option:` means the option has a non-optional arguments
# `option::` means the option has optional arguments
# `option` means the option has no arguments.
OPTIONS=developer,no-change-tty,tty:

WAYLAND_USER=${WAYLAND_USER:-torizon}
WESTON_ARGS=${WESTON_ARGS:--Bdrm-backend.so --current-mode -S${WAYLAND_DISPLAY}}
IGNORE_X_LOCKS=${IGNORE_X_LOCKS:-0}
IGNORE_VT_SWITCH_BACK=${IGNORE_VT_SWITCH_BACK:-0}

#
# Parse options.
#

# Pass `--options ''` to tell getopt we're not defining any short options
PARSED=$(getopt --options '' \
	--longoptions ${OPTIONS} \
	--name "$0" \
	-- "$@")
if [ $? -ne 0 ]; then
	echo "ERROR: getopt failed when setting up the options"
	exit 1
fi

eval set -- "${PARSED}"

COMMAND_HAS_DOUBLE_DASH=false
DO_NOT_SWITCH_TTY=false
# use tty7 for the graphical server by default
VT="7"
while [[ $# -ne 0 ]]; do
	case "$1" in
	--developer)
		export XDG_CONFIG_HOME=/etc/xdg/weston-dev
		echo "XDG_CONFIG_HOME=/etc/xdg/weston-dev" >>/etc/environment
		;;
	--tty)
		# Populates VT with the /dev/ttyX option that follows the --tty option
		VT=${2:8}
		;;
	--no-change-tty)
		DO_NOT_SWITCH_TTY=true
		;;
	--)
		COMMAND_HAS_DOUBLE_DASH=true
		;;
	*)
		if [ "$COMMAND_HAS_DOUBLE_DASH" = true ]; then
			WESTON_EXTRA_ARGS=("${WESTON_EXTRA_ARGS[@]}" "$1")
		fi
		;;
	esac
	shift
done

function vt_setup() {
	# Some applications may leave old VT in graphics mode which causes
	# applications like openvt and chvt to hang at VT_WAITACTIVE ioctl when they
	# try to switch to a new VT

	# grabs the current active VT, before possibly switching
	OLD_VT=$(cat /sys/class/tty/tty0/active)
	OLD_VT_MODE=$(kbdinfo -C /dev/"${OLD_VT}" getmode)
	if [ "$OLD_VT_MODE" = "graphics" ]; then
		/usr/bin/switchvtmode.pl "${OLD_VT:3}" text
	fi
}

# Change the foreground virtual terminal by default.
# This is specifically done because Plymouth deactivates and releases the drm fd
# if the foreground terminal is changed[0].
# Otherwise we can run into an issue where seatd cannot grab the fd.

# plymouth-quit.service is started by docker.service, so we don't need to
# quit plymouth from inside the container.

# [0] https://cgit.freedesktop.org/plymouth/commit/?id=5ab755153356b3f685afe87c5926969389665bb2

if [ "$DO_NOT_SWITCH_TTY" = false ]; then
	echo "Switching VT $(cat /sys/class/tty/tty0/active) to text mode if currently in graphics mode" && vt_setup
	echo "Switching to VT ${VT}" && chvt "${VT}"
fi

set -- "${WESTON_EXTRA_ARGS[@]}"

#
#
# Set desktop defaults.
#

function init_xdg() {
	if test -z "${XDG_RUNTIME_DIR}"; then
		XDG_RUNTIME_DIR=/tmp/$(id -u "${WAYLAND_USER}")-runtime-dir
		export XDG_RUNTIME_DIR
	fi

	echo "XDG_RUNTIME_DIR=${XDG_RUNTIME_DIR}" >>/etc/environment

	if ! test -d "${XDG_RUNTIME_DIR}"; then
		mkdir -p "${XDG_RUNTIME_DIR}"
	fi

	chown "${WAYLAND_USER}" "${XDG_RUNTIME_DIR}"
	chmod 0700 "${XDG_RUNTIME_DIR}"

	# Create folder for XWayland Unix socket
	export X11_UNIX_SOCKET="/tmp/.X11-unix"
	if ! test -d "${X11_UNIX_SOCKET}"; then
		mkdir -p ${X11_UNIX_SOCKET}
	fi

	chown "${WAYLAND_USER}":video ${X11_UNIX_SOCKET}
}

init_xdg

function cleanup() {
	if [ "$IGNORE_VT_SWITCH_BACK" != "1" ]; then
		# switch back to tty1, otherwise the console screen is not displayed.
		echo "Switching back to vt ${OLD_VT:3}"
		chvt "${OLD_VT:3}"
	fi
        rm -rf /run/seatd.sock
}

trap cleanup EXIT

function init() {
	if CMD=$(command -v "$1" 2>/dev/null); then
		shift
		CMD="${CMD} $*"
		runuser -u "${WAYLAND_USER}" -- sh -c "${CMD}"
	else
		echo "Command not found: $1"
		exit 1
	fi
}

if [ "$IGNORE_X_LOCKS" != "1" ]; then
       echo "Removing previously created '.X*-lock' entries under /tmp before starting Weston. Pass 'IGNORE_X_LOCKS=1' environment variable to Weston container to disable this behavior."
       rm -rf /tmp/.X*-lock
fi

# for every argument after "--", append that argument to WESTON_ARGS
for i in "${!WESTON_EXTRA_ARGS[@]}"; do
	WESTON_ARGS+=" "${WESTON_EXTRA_ARGS[$i]}
done

# FIXME: this should be run as ... init seatd-launch ... to make sure Weston
# is running as 'torizon' user, but it didn't really work.
rm -rf /run/seatd.sock || true && seatd-launch -- weston ${WESTON_ARGS}

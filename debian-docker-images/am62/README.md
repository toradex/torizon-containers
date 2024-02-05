# AM62 containers

AM62 containers contain platform-specific packages from the Toradex feed at https://feeds.toradex.com/debian-am62/.
AM62 containers are meant to be used with iMX6 and iMX7 devices.

## base

```
docker run -it --rm --name=debian torizon/debian-am62:$CT_TAG_DEBIAN
```

## chromium

```
docker run -d --rm --name=chromium \
        -v /tmp:/tmp -v /dev/dri:/dev/dri \
        -v /var/run/dbus:/var/run/dbus --device-cgroup-rule='c 226:* rmw' \
        --security-opt seccomp=unconfined --shm-size 256mb \
        torizon/chromium-am62:$CT_TAG_CHROMIUM \
        --virtual-keyboard https://www.toradex.com
```

## cog

```
docker run -d --rm --name=cog \
        -v /tmp:/tmp -v /var/run/dbus:/var/run/dbus \
        -v /dev/dri:/dev/dri --device-cgroup-rule='c 226:* rmw' \
        torizon/cog-am62:$CT_TAG_COG \
        https://www.toradex.com
```

## qt5-wayland

```
docker run --rm -it --name=qt5 \
        -v /tmp:/tmp \
        -v /dev/dri:/dev/dri --device-cgroup-rule='c 226:* rmw' \
        torizon/qt5-wayland-am62:$CT_TAG_QT5_WAYLAND \
        bash
```

## qt5-wayland-examples

```
docker run --rm -it --name=qt5 \
        -v /tmp:/tmp \
        -v /dev/dri:/dev/dri --device-cgroup-rule='c 226:* rmw' \
        torizon/qt5-wayland-examples-am62:$CT_TAG_QT5_WAYLAND_EXAMPLES \
        bash
```

And then run one of the examples availaible in `/usr/lib/aarch64-linux-gnu/qt5/examples/`:

```
/usr/lib/aarch64-linux-gnu/qt5/examples/widgets/widgets/calculator/calculator
```


## qt6-wayland

```
docker run --rm -it --name=qt6 \
        -v /tmp:/tmp \
        -v /dev/dri:/dev/dri --device-cgroup-rule='c 226:* rmw' \
        torizon/qt6-wayland-am62:$CT_TAG_QT6_WAYLAND \
        bash
```

## qt6-wayland-examples

```
docker run --rm -it --name=qt6 \
        -v /tmp:/tmp \
        -v /dev/dri:/dev/dri --device-cgroup-rule='c 226:* rmw' \
        torizon/qt6-wayland-examples-am62:$CT_TAG_QT6_WAYLAND_EXAMPLES \
        bash
```

And then run one of the examples availaible in `/usr/lib/aarch64-linux-gnu/qt6/examples/`

```
/usr/lib/aarch64-linux-gnu/qt6/examples/widgets/widgets/calculator/calculator
```

## chromium-tests-am62

## graphics-tests

```
docker run -d --rm -it --name=graphics-tests -v /dev:/dev \
        --device-cgroup-rule='c 4:* rmw'  --device-cgroup-rule='c 13:* rmw' \
        --device-cgroup-rule='c 226:* rmw' \
        torizon/graphics-tests-am62:$CT_TAG_GRAPHICS_TESTS
```

## wayland-base

```
docker run --rm -it --name=wayland-base \
        -v /tmp:/tmp \
        -v /dev/dri:/dev/dri --device-cgroup-rule='c 226:* rmw' \
        torizon/wayland-base-am62:$CT_TAG_WAYLAND_BASE \
        bash
```

## weston

```
docker run -d --rm --name=weston --net=host --cap-add CAP_SYS_TTY_CONFIG \
        -v /dev:/dev -v /tmp:/tmp -v /run/udev/:/run/udev/ \
        --device-cgroup-rule='c 4:* rmw' --device-cgroup-rule='c 13:* rmw' \
        --device-cgroup-rule='c 226:* rmw' \
        torizon/weston-am62:$CT_TAG_WESTON --developer --tty=/dev/tty7
```

## weston-touch-calibrator

```
docker run -ti --rm --name=weston-touch-calibrator --privileged \
        -v /dev:/dev -v /run/udev/:/run/udev/ -v /etc/udev/rules.d:/etc/udev/rules.d \
        torizon/weston-touch-calibrator-am62:$CT_TAG_WESTON_TOUCH_CALIBRATOR
```

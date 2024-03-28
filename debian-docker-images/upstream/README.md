# Upstream containers

Upstream containers don't use third-party package feeds. All packages is from the upstream Debian project.
Upstream images can be used with any target devices, but may not have support for hardware acceleration.

The support strategy for new target devices (or machines in the Yocto lingo) is to add a new platform folder (like imx8 or am62) and customize it to suit that new device. This step will also most likely involve adding a new debian package feed.

## base

```
docker run -it --rm --name=debian torizon/debian:rc
```

## chromium

```
docker run -d --rm --name=chromium \
        -v /tmp:/tmp -v /dev/dri:/dev/dri \
        -v /var/run/dbus:/var/run/dbus --device-cgroup-rule='c 226:* rmw' \
        --security-opt seccomp=unconfined --shm-size 256mb \
        torizon/chromium:rc \
        --virtual-keyboard https://www.toradex.com
```

## cog

```
docker run -d --rm --name=cog \
        -v /tmp:/tmp -v /var/run/dbus:/var/run/dbus \
        -v /dev/dri:/dev/dri --device-cgroup-rule='c 226:* rmw' \
        torizon/cog:rc \
        https://www.toradex.com
```

## qt5-wayland

```
docker run --rm -it --name=qt5 \
        -v /tmp:/tmp \
        -v /dev/dri:/dev/dri --device-cgroup-rule='c 226:* rmw' \
        torizon/qt5-wayland:rc \
        bash
```

## qt5-wayland-examples

```
docker run --rm -it --name=qt5 \
        -v /tmp:/tmp \
        -v /dev/dri:/dev/dri --device-cgroup-rule='c 226:* rmw' \
        torizon/qt5-wayland-examples:rc \
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
        torizon/qt6-wayland:rc \
        bash
```

## qt6-wayland-examples

```
docker run --rm -it --name=qt6 \
        -v /tmp:/tmp \
        -v /dev/dri:/dev/dri --device-cgroup-rule='c 226:* rmw' \
        torizon/qt6-wayland-examples:rc \
        bash
```

And then run one of the examples availaible in `/usr/lib/aarch64-linux-gnu/qt6/examples/`

```
/usr/lib/aarch64-linux-gnu/qt6/examples/widgets/widgets/calculator/calculator
```

## chromium-tests

## graphics-tests

```
docker run -d --rm -it --name=graphics-tests -v /dev:/dev \
        --device-cgroup-rule='c 4:* rmw'  --device-cgroup-rule='c 13:* rmw' \
        --device-cgroup-rule='c 226:* rmw' \
        torizon/graphics-tests:rc
```

## wayland-base

```
docker run --rm -it --name=wayland-base \
        -v /tmp:/tmp \
        -v /dev/dri:/dev/dri --device-cgroup-rule='c 226:* rmw' \
        torizon/wayland-base:rc \
        bash
```

## weston

```
docker run -d --rm --name=weston --net=host --cap-add CAP_SYS_TTY_CONFIG \
        -v /dev:/dev -v /tmp:/tmp -v /run/udev/:/run/udev/ \
        --device-cgroup-rule='c 4:* rmw' --device-cgroup-rule='c 13:* rmw' \
        --device-cgroup-rule='c 226:* rmw' \
        torizon/weston:rc --developer
```

## weston-touch-calibrator

```
docker run -ti --rm --name=weston-touch-calibrator --privileged \
        -v /dev:/dev -v /run/udev/:/run/udev/ -v /etc/udev/rules.d:/etc/udev/rules.d \
        torizon/weston-touch-calibrator:rc
```

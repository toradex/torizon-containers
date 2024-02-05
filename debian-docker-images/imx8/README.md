# iMX8 containers

iMX8 containers contain platform-specific packages from the Toradex feed at https://feeds.toradex.com/debian/.

## base

## chromium

## cog

## gputop

While we do provide a gputop container, it's useless to run it outside the process namepace where the graphical applications are running. Essentialy, you need to run `gputop` inside the same container as the graphical applications are running, otherwise you won't see anything.

As an example, if you're running a Qt5 application and you want to monitor it with gputop, run the application as usual adding the `--privileged` flag (`gputop` uses debugfs under the hood):

```
docker run -it --rm --name=qt5-wayland-examples -v /dev/dri:/dev/dri -v /dev/galcore:/dev/galcore -v /tmp:/tmp --device-cgroup-rule='c 199:* rmw' --privileged torizon/qt5-wayland-examples-imx8:rc /usr/lib/aarch64-linux-gnu/qt5/examples/opengl/cube/cube
```

For a graphical application you should have mounted the GPU-related device descriptors already, but `--privileged` will auto-mount them anyway, as well as have the effect of `--device-cgroup-rule='c *:* rmw'`.

Then exec into the running container, update, install and run gputop (note the `cube` application):

```
torizon@verdin-imx8mp-06817296:~$ docker ps
CONTAINER ID   IMAGE                                  COMMAND                  CREATED              STATUS              PORTS     NAMES
cb6510315997   torizon/qt5-wayland-examples-imx8:rc   "bash"                   About a minute ago   Up About a minute             qt5-wayland-examples
c59e4bd75e41   torizon/weston-imx8:rc                 "/usr/bin/entry.sh -…"   About a minute ago   Up About a minute             clever_gauss
torizon@verdin-imx8mp-06817296:~$ docker exec -it cb6510315997 bash
root@4d44ad1358d9:/# apt update && apt install gputop
1b1261761bbe, 1.4
3D:GC7000,Rev:6204 Core: 1000 MHz, Shader: 1000 MHz 
3D:GC8000,Rev:8002 Core: 1000 MHz, Shader: 999 MHz 2D:GC520,Rev:5341 
3D Cores:2,2D Cores:1,VG Cores:0

IMX8_DDR0: axid-read:4.67,axid-write:0.83
IMX8_DDR1: 

     PID   RES(kB)   CONT(kB)   VIRT(kB)  Non-PGD(kB)  Total(kB)              CMD
       1     17505          0          0            0      17505             cube

TOT:         17505          0          0            0      17505
TOT_CON:         -          -          -            -     244638 

```

## qt5-wayland

## qt5-wayland-examples

## qt6-wayland

## qt6-wayland-examples

## graphics-tests

## wayland-base

## weston

Weston should run before any other application, because it also instanciates the Wayland Server. In some cases, graphical applications that require the Wayland socket may start faster than Weston. To prevent this, we implement a container HEALTHCHECK so containers that depend on Weston can wait until the Wayland server is fully up.

To use it, add the following as a dependency in `docker-compose.yml`

```
depends_on:
  weston:
    condition: service_healthy
```

If you have a third service that depends on Weston and some other service, specify a condition for that other service as well, for example

```
depends_on:
  portainer:
    condition: service_started
  weston:
    condition: service_healthy
```

## weston-touch-calibrator

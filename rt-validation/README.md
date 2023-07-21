# rt-validation

Containers to validate real-time (`PREEMPT_RT`) support in TorizonCore.

The `stress-tests` container will run stress tests (CPU, memory, I/O) in the device.

The `rt-tests` container will run `cyclictest` and generate a report of the measured latency.


# Prepare the device for the tests

Make sure you are running a `PREEMPT_RT` version of TorizonCore:
```
$ cat /etc/os-release | grep PREEMPT_RT
NAME="Torizoncore Upstream with PREEMPT_RT"
PRETTY_NAME="Torizoncore Upstream with PREEMPT_RT 5.1.0-devel-20201216+build.152 (dunfell)"
```

Make sure the device is connected to the internet (needed to generate network load):

```
$ ping -c 3 8.8.8.8
PING 8.8.8.8 (8.8.8.8): 56 data bytes
64 bytes from 8.8.8.8: seq=0 ttl=118 time=12.344 ms
64 bytes from 8.8.8.8: seq=1 ttl=118 time=10.730 ms
64 bytes from 8.8.8.8: seq=2 ttl=118 time=12.080 ms
```

Mount a pendrive in `/mnt/pendrive` and make sure it has an archive file called `file.tar` (needed to generate USB read/write):

```
$ sudo mkdir -p /mnt/pendrive
$ sudo mount /dev/sda1 /mnt/pendrive
$ ls /mnt/pendrive/file.tar
-rw-rw-r-- 1 torizon torizon 1011353600 Dec  8 19:30 /mnt/pendrive/file.tar
```


# Run the stress-tests container

Run the following command to execute the `stress-tests` container:

```
$ docker run --rm -d --name stress-tests --privileged -v /dev:/dev -v /tmp:/tmp -v /mnt/pendrive/:/mnt/pendrive torizon/stress-tests:$CT_TAG_STRESS_TESTS
```

Check if the stress tests started successfully:

```
$ docker logs stress-tests
Setting up stress tests...
RT stress tests started successfully!
```


# Run the rt-tests container

Run the following command to execute the `rt-tests` container and start measuring the latency:

```
$ docker run --rm -it --name rt-tests --cap-add=sys_nice --cap-add=ipc_lock --cap-add=sys_rawio --ulimit rtprio=99 --device-cgroup-rule='c 10:* rmw' -v /dev:/dev -v /tmp:/tmp torizon/rt-tests:$CT_TAG_RT_TESTS
```

The tests will run for at most 12 hours, but can be interrupted at any time by
pressing CTRL-C. You can specify a different length of the test run in seconds
by passing the environment variable *DURATION* to the container (with
`-e DURATION=30s`); for other units append `m` for minutes, `h` for hours and `d`
for days.

After the tests are finished, stop the `stress-tests` container:

```
$ docker stop stress-tests
```


# Analysing the results

A summary of latency measurements will be available in `/tmp/latency-summary.log`:

```
$ cat /tmp/latency-summary.log
# /dev/cpu_dma_latency set to 0us
# Histogram
# Total: 000297214 000297199 000297189 000297179
# Min Latencies: 00005 00005 00005 00005
# Avg Latencies: 00023 00026 00023 00021
# Max Latencies: 00081 00085 00077 00076
# Histogram Overflows: 00000 00000 00000 00000
# Histogram Overflow at cycle number:
# Thread 0:
# Thread 1:
# Thread 2:
# Thread 3:
```

A latency plot from cyclictest histogram data will be available in `/tmp/latency-plot.png`.


# Running rt-tests on a non-PREEMPT_RT version of TorizonCore

Running the `rt-tests` container on a non-PREEMPT_RT version of TorizonCore may be useful to compare the results with the PREEMPT_RT version. But if you try to run, the execution might fail:

```
$ docker run --rm -it --name rt-tests --cap-add=sys_nice --cap-add=ipc_lock --cap-add=sys_rawio --ulimit rtprio=99 --device-cgroup-rule='c 10:* rmw' -v /dev:/dev -v /tmp:/tmp torizon/rt-tests:$CT_TAG_RT_TESTS
Unable to change scheduling policy!
Probably missing capabilities, either run as root or increase RLIMIT_RTPRIO limits.
ERROR: cyclictest failed
```

That is because `CONFIG_RT_GROUP_SCHED` may be enabled on non-PREEMPT_RT versions of TorizonCore, and a cgroups configuration is required to run real-time tasks.

So if you want to run the `rt-tests` container on a non-PREEMPT_RT version of TorizonCore, run the following commands:

```
sudo sh -c "echo 950000 > /sys/fs/cgroup/cpu,cpuacct/docker/cpu.rt_runtime_us"
docker run --rm -it --name rt-tests --cpu-rt-runtime=950000 --cap-add=sys_nice --cap-add=ipc_lock --cap-add=sys_rawio --ulimit rtprio=99 --device-cgroup-rule='c 10:* rmw' -v /dev:/dev -v /tmp:/tmp torizon/rt-tests:$CT_TAG_RT_TESTS
```

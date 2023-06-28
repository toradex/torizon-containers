#!/bin/sh

echo "Setting up stress tests..."

# CPU load
for i in $(seq $(nproc)); do dd if=/dev/zero of=/dev/null bs=4M & done

# I/O read load
for dev in /dev/mmcblk? /dev/sd?; do
    test -b "$dev" || continue

    while true; do dd if="$dev" of=/dev/null bs=4M; done &
done

# I/O load (USB read/write)
if mount | grep -q /mnt/pendrive; then
    if [ -e /mnt/pendrive/file.tar ]; then
        cd /mnt/pendrive
        while true; do rm -rf output; mkdir -p output; tar xfv file.tar -C output/ >/dev/null 2>&1; done &
    else
        echo "Warning: Archive file.tar not found in /mnt/pendrive. USB read/write stress test will not run."
    fi
else
    echo "Warning: pendrive not mounted. USB read/write stress test will not run."
fi

# RTC read/write
while true; do hwclock >/dev/null 2>&1; sleep 0.1; done &
while true; do hwclock -w >/dev/null 2>&1; sleep 0.1; done &

# Network load
if ! ping -q -c 1 -W 1 8.8.8.8 >/dev/null 2>&1; then
    echo "Warning: could not ping 8.8.8.8. Internet may be unavailable to generate network load."
fi
while true; do
    SERVERS="bouygues.iperf.fr speedtest.wtnet.de speedtest.iveloz.net.br iperf.scottlinux.com"
    for s in $SERVERS; do
        iperf3 -c $s -P 10 >/dev/null 2>&1
    done
done &

ping -A -f -t 3 -q -s 49152 \
    ${STRESS_SERVER:-$(ip route get 1.1.1.1 |sed '/.* via \([^ ]*\) .*/ s//\1/; q')}

echo "RT stress tests started successfully!"

# Dont exit
while true; do sleep 1; done

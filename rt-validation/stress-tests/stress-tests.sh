#!/bin/sh

echo "Setting up stress tests..."

# CPU load
while true; do hackbench >/dev/null 2>&1; done &

# I/O load (eMMC read)
while true; do du /usr >/dev/null 2>&1; done &
while true; do find / -name a* >/dev/null 2>&1; done &

# I/O load (USB read/write)
if mount | grep /mnt/pendrive >/dev/null; then
    if [ -e /mnt/pendrive/file.tar ]; then
        cd /mnt/pendrive
        tar -cf file.tar -C /var/log --exclude=lost+found -p .
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

echo "RT stress tests started successfully!"

# Dont exit
while true; do sleep 1; done

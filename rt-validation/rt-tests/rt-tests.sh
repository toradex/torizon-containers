#!/bin/bash

# Run cyclictest
if ! cyclictest -S -p99 -i1000 -m -D"${DURATION:-12h}" -h400 -q >output; then
  echo "ERROR: cyclictest failed" >&2
  exit 1
fi

# Get maximum latency
max=$(grep "Max Latencies" output | tr " " "\n" | sort -n | tail -1 | sed s/^0*//)

# Grep data lines, remove empty lines and create a common field separator
grep -v -e "^#" -e "^$" output | tr " " "\t" >histogram

# Create two-column data sets with latency classes and frequency values for each core
cores=$(grep -c ^processor /proc/cpuinfo)
for i in $(seq 1 "$cores"); do
  column=$((i + 1))
  cut -f1,"$column" histogram >histogram"$i"
done

# Create plot command header
echo -n -e "set title \"Latency plot\"\n\
set terminal png\n\
set xlabel \"Latency (us), max $max us\"\n\
set logscale y\n\
set xrange [0:400]\n\
set yrange [0.8:*]\n\
set ylabel \"Number of latency samples\"\n\
set output \"plot.png\"\n\
plot " >plotcmd

# Append plot command data references
for i in $(seq 1 "$cores"); do
  if test "$i" != 1; then
    echo -n ", " >>plotcmd
  fi
  cpuno=$((i - 1))
  if test "$cpuno" -lt 10; then
    title=" CPU$cpuno"
  else
    title="CPU$cpuno"
  fi
  echo -n "\"histogram$i\" using 1:2 title \"$title\" with histeps" >>plotcmd
done

# Execute plot command
gnuplot -persist <plotcmd

# Print summary to /tmp
echo "Generating latency summary to /tmp/latency-summary.log"
grep -e "^#" output >/tmp/latency-summary.log

# Copy plot to /tmp/latency-plot.png
echo "Generating latency plot to /tmp/latency-plot.png"
cp plot.png /tmp/latency-plot.png

# Clean up
rm -f histogram* output plotcmd plot.png

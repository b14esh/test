#!/bin/bash

# Subcommand: cpu
cpu() {
    echo $(top -bn1 | grep "Cpu(s)" | sed "s/.*, *\([0-9.]*\)%* id.*/\1/" | sed "s/^/100 - /" | bc)
}

# Subcommand: memory
memory() {
    echo $(free -m | awk '/Mem/{print $3" "$2}')
}

# Subcommand: network
network() {
    echo $(cat /sys/class/net/[e]*/statistics/{r,t}x_bytes)
}

# Subcommand: storage
storage() {
    #echo $(df -Ph | grep mmcblk0p5 | awk '{print $2" "$3}' | sed 's/G//g')
    echo $(df -Ph | grep /dev/mmcblk0p1 | awk '{print $2" "$3}' | sed 's/G//g')
}

# Subcommand: uptime
uptime() {
    echo $(cut -f1 -d. /proc/uptime)
}

# Main script
cpu_result=$(cpu)
memory_result=$(memory)
network_result=$(network)
storage_result=$(storage)
uptime_result=$(uptime)

echo "CPU: $cpu_result"
echo "Memory: $memory_result"
echo "Network: $network_result"
echo "Storage: $storage_result"
echo "Uptime: $uptime_result"

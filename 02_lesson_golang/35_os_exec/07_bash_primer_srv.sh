#!/bin/bash
# Subcommand: cpu
cpu() {
    top -bn1 | grep "Cpu(s)" | sed "s/.*, *\([0-9.]*\)%* id.*/\1/" | sed "s/^/100 - /" | bc
}

# Subcommand: memory
memory() {
    free -m | awk '/Mem/{print $3" "$2}'
}

# Subcommand: network
network() {
    cat /sys/class/net/[e]*/statistics/{r,t}x_bytes
}

# Subcommand: storage
storage() {
    df -Ph | grep mmcblk0p1 | awk '{print $2" "$3}' | sed 's/G//g'
}

# Subcommand: uptime
uptime() {
    cut -f1 -d. /proc/uptime
}

# TCP Server
while true; do
    { 
        while IFS= read -r cmd; do
            case $cmd in
                "cpu") cpu;;
                "memory") memory;;
                "network") network;;
                "storage") storage;;
                "uptime") uptime;;
                "exit") exit;;
                *) echo "Invalid command";;
            esac
        done
    } | nc -l -p 4444 -q 1
done

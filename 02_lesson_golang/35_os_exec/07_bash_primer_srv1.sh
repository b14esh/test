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
    cat /sys/class/net/[e]*/statistics/{r,t}x_bytes | fmt
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
	#echo -ne "HTTP/1.1 200 OK\r\n\r\n"
	(echo -e "\n 
	cpu  $(cpu)
	memory $(memory)
	network $(network)
	storage $(storage)
	uptime $(uptime)"
		)| nc -l -p 4444 -q 1
done

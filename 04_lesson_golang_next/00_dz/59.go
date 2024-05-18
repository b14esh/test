package main

import (
	"fmt"
	"os"
)

var (
	text                         = ""
	file_name             string = ""
	write_world                  = "[Unit]\nDescription=bh2-service\nAfter=network.target\n\n[Service]\nType=forking\nTimeoutStartSec=10\nExecStart=/usr/bin/bh2-script.sh\nСтрока которую заменить\n\nSyslogIdentifier=bh2-service\nSyslogFacility=daemon\n\n[Install]\nWantedBy=multi-user.target\nСтрока которую заменить\n"
	write_world_file_name string = "test.txt"
)

// func write files
func write_file(text string, file_name string) {
	file, err := os.Create(file_name)
	if err != nil {
		fmt.Println("Unable to create file:", err)
		os.Exit(1)
	}
	defer file.Close()
	file.WriteString(text)

}

func main() {
	write_file(write_world, write_world_file_name)
}

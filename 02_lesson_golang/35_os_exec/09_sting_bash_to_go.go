package main

import (
	"fmt"
	"os/exec"
)

func main() {
	commands := []string{
		"ping -c 2 google.com",
		"ping -c 2 facebook.com",
		"ping -c 2 www.golinuxcloud.com",
	}
	for _, command := range commands {
		cmd := exec.Command("bash", "-c", command)
		out, err := cmd.Output()
		if err != nil {
			fmt.Println(err)
		}
		fmt.Println(string(out))
	}
}

package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

func bh2EncryptionToggle(interfaceName string, psk string, kAlgo int) bool {
	if kAlgo < 0 || kAlgo > 2 {
		log.Fatalf("Invalid kAlgo value. Must be 0, 1, or 2.")
		return false
	}

	if strings.HasPrefix(interfaceName, "wlP2p1s0f") {
		file, err := os.OpenFile(filePath, os.O_RDWR, 0644)
		if err != nil {
			log.Fatalf("An error occurred: %v", err)
			return false
		}
		defer file.Close()

		scanner := bufio.NewScanner(file)
		var lines []string
		for scanner.Scan() {
			line := scanner.Text()
			if strings.HasPrefix(strings.TrimSpace(line), "options bh2 antenna_type=v4") {
				parts := strings.Fields(line)
				parts = filterParts(parts, "bh2_psk=", "bh2_kalgo=")
				if kAlgo == 2 && psk != "" {
					parts = append(parts, fmt.Sprintf("bh2_psk=%s", psk), "bh2_kalgo=2")
				} else if kAlgo == 1 {
					parts = append(parts, "bh2_kalgo=1")
				}
				line = strings.Join(parts, " ")
			}
			lines = append(lines, line)
		}

		if err := scanner.Err(); err != nil {
			log.Fatalf("An error occurred: %v", err)
			return false
		}

		file.Seek(0, 0)
		file.Truncate(0)
		for _, line := range lines {
			fmt.Fprintln(file, line)
		}
	}

	return true
}

func filterParts(parts []string, prefixes ...string) []string {
	var filtered []string
	for _, part := range parts {
		keep := true
		for _, prefix := range prefixes {
			if strings.HasPrefix(part, prefix) {
				keep = false
				break
			}
		}
		if keep {
			filtered = append(filtered, part)
		}
	}
	return filtered
}

func main() {
	bh2EncryptionToggle("wlP2p1s0f", "psk", 1)
}

package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

var (
	//var other / rest var
	argsNeed = os.Args
)

// wx -w1 --encryption key_string alg_number
// wx -w1 --encryption key_string 0 // нафиг шифрование
// wx -w1 --encryption key_string 1 // без ключа
// wx -w1 --encryption key_string 2 // шифрование с ключом
func wifi_encryption() {
	//write_etc_modprobe_bh2_conf_file_name
	key_string := argsNeed[3]
	alg_number := argsNeed[4]

	switch {
	case alg_number != "0" && alg_number != "1" && alg_number != "2":
		fmt.Println("Ich will nicht %s", alg_number)
		return
	case len(key_string) < 1 || len(key_string) > 63:
		fmt.Println("Ich will nicht %s", key_string)
		return
	default:
		filePath := "bh2.conf"
		file, err := os.OpenFile(filePath, os.O_RDWR, 0644)
		if err != nil {
			log.Println("An error occurred: %v", err)
		}
		defer file.Close()

		scanner := bufio.NewScanner(file)
		var lines []string
		for scanner.Scan() {
			line := scanner.Text()
			if strings.HasPrefix(strings.TrimSpace(line), "options bh2 antenna_type=v4") {
				parts := strings.Fields(line)
				parts = filterParts(parts, "bh2_psk=", "bh2_kalgo=")
				if alg_number == "2" && key_string != "" {
					parts = append(parts, fmt.Sprintf("bh2_psk=%s", key_string), "bh2_kalgo=2")
				} else if alg_number == "1" {
					parts = append(parts, "bh2_kalgo=1")
				}
				line = strings.Join(parts, " ")
			}
			lines = append(lines, line)
		}

		if err := scanner.Err(); err != nil {
			log.Fatalf("An error occurred: %v", err)
		}

		file.Seek(0, 0)
		file.Truncate(0)
		for _, line := range lines {
			fmt.Fprintln(file, line)
		}

	}

}

// function for help parsing string
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

func flag() {
	switch {
	case len(argsNeed) == 5 && argsNeed[2] == "--encryption" && argsNeed[4] != "" && argsNeed[3] != "" && argsNeed[1] != "":
		wifi_encryption()
	default:
		fmt.Println("Ich will nicht")
	}
}

func main() {
	// wx -w1 --encryption key_string alg_number
	// wx -w1 --encryption key_string 0 // нафиг шифрование
	// wx -w1 --encryption key_string 1 // без ключа
	// wx -w1 --encryption key_string 2 // шифрование с ключом
	if len(argsNeed) == 5 {
		argsNeed[1] = "-w1"
		argsNeed[2] = "--encryption"
		argsNeed[3] = "key_stringkey_stringkey_string"
		argsNeed[4] = "0"
	}
	flag()

	//go run .\64.go 1 1 1 1
}

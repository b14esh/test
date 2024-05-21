package main

import "fmt"

func main() {
	p := 26008.15 + 116117.40 + 131342.93 + 157530.32 + 173835.97 + 184902.13 + 194530.33 + 202480.72 + 209725.36 + 217904.18 + 222320.87 + 109296.48 + 98439.81
	d := 36369.29 + 133392.36 + 118166.83 + 109410.16 + 86656.55 + 75590.39 + 65962.19 + 58011.80 + 50767.16 + 42588.34 + 38171.65 + 14101.09 + 151069.95
	s := d + p

	fmt.Printf("Проценты: %0.2f \n", p)
	fmt.Printf("Основной долг: %0.2f \n", d)
	fmt.Printf("Всего оплачено на 2024: %0.2f \n", s)

	write_opt_bwt_etc_hostapd_minimal_conf_name := "/opt/bwt/etc/hostapd-minimal.conf"
	network_config_hostapd_wlP2p1s0f0_conf_name := "/opt/bwt/etc/hostapd-wlP2p1s0f0.conf"
	network_config_hostapd_wlP2p1s0f1_conf_name := "/opt/bwt/etc/hostapd-wlP2p1s0f1.conf"

	box7 := []string{
		"cat " + network_config_hostapd_wlP2p1s0f0_conf_name,
		"cat " + network_config_hostapd_wlP2p1s0f1_conf_name,
		"cat " + write_opt_bwt_etc_hostapd_minimal_conf_name,
		"echo \"Usage: wx -channel 1\"",
		"echo \"Usage: wx -channel 2\"",
		"echo \"Usage: wx -channel 3\"",
		"echo \"Usage: wx -channel 4\"",
		"echo \"Usage: wx -channel 5\"",
	}
	for _, val := range box7 {
		fmt.Println(val)
	}
}

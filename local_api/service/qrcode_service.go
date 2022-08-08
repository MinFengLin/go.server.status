package service

import (
	"fmt"
)

type Wifi_info_s struct {
	URL     string
	SSID  	string
	WIFIPWD string
}

var (
	Wifi_info Wifi_info_s
)

func Parser_wifiqrcode(SSID string, ENC string, WIFIPWD string) Wifi_info_s {
	wifi_qrcode_url := "https://tw.piliapp.com/generator/qr-code/iframe/?data=WIFI%3AS%3A" + SSID + "%3BT%3A"

	switch ENC {
	case "WPA", "WPA2":
		wifi_qrcode_url += "WPA"
		fmt.Println("ENC Set at WPA/WPA2")
	case "WEP":
		wifi_qrcode_url += "WEP"
		fmt.Println("ENC Set at WEP")
	case "NONE":
		fmt.Println("ENC Set at none")
	default:
		fmt.Println("ENC only support WPA/WPA2/WEP/NONE")
	}

	wifi_qrcode_url += "%3BP%3A"

	if WIFIPWD != "" {
		wifi_qrcode_url += WIFIPWD + "%3B%3B&size=150" 
	} else {
		wifi_qrcode_url += "%3B%3B&size=150" 
	}

	Wifi_info.URL     = wifi_qrcode_url
	Wifi_info.SSID    = SSID
	Wifi_info.WIFIPWD = WIFIPWD

	fmt.Printf("%+v\n", Wifi_info)

	return Wifi_info
}

package helpers

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net"
)

type IPInfoResponse struct {
	IP       string `json:"ip"`
	City     string `json:"city"`
	Region   string `json:"region"`
	Country  string `json:"country"`
	Loc      string `json:"loc"`
	Org      string `json:"org"`
	Timezone string `json:"timezone"`
	Readme   string `json:"readme"`
}

func GetExternalIP() string {
    resp, err := http.Get("https://ipinfo.io/json")
	if err != nil {
		return ""
	}
	defer resp.Body.Close()

	 if resp.StatusCode != http.StatusOK {
        return ""
    }

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return ""
	}

	var data IPInfoResponse
	if err := json.Unmarshal(body, &data); err != nil {
		return ""
	}

	return fmt.Sprintf("%s (%s,%s)", data.IP, data.City, data.Country)
}


func GetLocalIPs() []string {
	 var adapters []string

	interfaces, err := net.Interfaces()
	if err != nil {
        return adapters
    }

	 for _, iface := range interfaces {
		if iface.Flags&net.FlagLoopback != 0 {
            continue
        }
		addrs, err := iface.Addrs()
        if err != nil {
            continue
        }

		for _, addr := range addrs {
            if ipNet, ok := addr.(*net.IPNet); ok && ipNet.IP.To4() != nil {
                adapters = append(adapters, fmt.Sprintf("%s (%s)",ipNet.IP.String(), iface.Name))
            }
        }

	}

    return adapters
}

package helpers

import (
	"fmt"
	"os"
	"os/user"
	"time"

	"github.com/shirou/gopsutil/v4/host"
)

func GetUsername() string {
	currentUser, err := user.Current()
	if err != nil {
		return ""
	}
	return currentUser.Username
}

func GetHostname() string {
	hostname, err := os.Hostname()
	if err != nil {
		return ""
	}
	return hostname
}

func GetUptime() string {
	uptime, err := host.Uptime()
	if err != nil {
		return ""
	}
	duration := time.Duration(uptime) * time.Second
	days := int(duration.Hours() / 24)
	hours := int(duration.Hours()) % 24
	minutes := int(duration.Minutes()) % 60
	seconds := int(duration.Seconds()) % 60

	return fmt.Sprintf("%dd %dh %dm %ds", days, hours, minutes, seconds)
}

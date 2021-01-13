package util

import (
	"os"
)

func HostName() string {
	host, err := os.Hostname()
	if err != nil {
		Log().Error("failed to retrieve hostname", err)
		return "unknown"
	}
	return host
}

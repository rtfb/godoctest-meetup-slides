package foo

import (
	"strings"
)

func swapPort(hostport, newPort string) string {
	if hostport == "" {
		return "localhost:" + newPort
	}
	colon := strings.Index(hostport, ":")
	if colon == -1 {
		return hostport + ":" + newPort
	}
	return hostport[:colon] + ":" + newPort
}

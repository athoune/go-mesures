package tcp_socket

import (
	"regexp"
	"strconv"
)

func Parse(s string) (string, int64, bool) {
	re, _ := regexp.Compile("(\\w*)\\s*(\\d*)")
	results := re.FindStringSubmatch(s)
	if results == nil {
		return "", 0, false
	}
	value, _ := strconv.ParseInt(results[2], 10, 64)
	return results[1], value, true
}

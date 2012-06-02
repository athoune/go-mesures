package tcp_socket

import (
	"regexp"
	"strconv"
)

var reParse *regexp.Regexp

func init() {
	reParse = regexp.MustCompile("(\\w+)\\s+(\\d+)")
}

func Parse(s string) (string, int64, bool) {
	results := reParse.FindStringSubmatch(s)
	if results == nil {
		return "", 0, false
	}
	value, _ := strconv.ParseInt(results[2], 10, 64)
	return results[1], value, true
}

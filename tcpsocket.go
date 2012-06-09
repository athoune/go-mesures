package main

import (
	"regexp"
	"strconv"
)

type msg struct {
    key string
    value int64
}

var reParse *regexp.Regexp

func init() {
	reParse = regexp.MustCompile("([a-zA-Z._0-9]+)\\s+(\\d+)")
}

func parse_cmd(s string) (msg, bool) {
	results := reParse.FindStringSubmatch(s)
	if results == nil {
		return msg{"", 0}, false
	}
	value, _ := strconv.ParseInt(results[2], 10, 64)
	return msg{results[1], value}, true
}

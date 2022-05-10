package consts

import "regexp"

var (
	HostRe = regexp.MustCompile("\\[(.*)\\]:(.*)@(.*)")
)

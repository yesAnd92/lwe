package utils

import (
	"runtime"
)

type OsPlatform int

const (
	Other OsPlatform = iota
	Win
	Mac
	Linux
)

func OsEnv() OsPlatform {
	// get runtime op platform
	op := runtime.GOOS

	if op == "windows" {
		return Win
	} else if op == "darwin" {
		return Mac
	} else if op == "linux" {
		return Linux
	}
	return Other
}

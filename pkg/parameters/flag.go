package parameters

import (
	"flag"
)

func GetFlag() bool {
	var envFlag bool
	flag.BoolVar(&envFlag, "f", false, "Path")
	flag.Parse()

	if !envFlag {
		return false
	} else {
		return true
	}
}

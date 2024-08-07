package parameters

import (
	"flag"
)

func GetFlag() bool {
	var envFlag bool
	flag.BoolVar(&envFlag, "f", false, "Path")
	flag.Parse()

	return envFlag
}

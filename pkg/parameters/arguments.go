package parameters

import (
	"flag"
)

func GetPath() string {
	if flag.NArg() <= 0 {
		return "."
	}

	return flag.Arg(0)
}

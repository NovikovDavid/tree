package app

import (
	"tree/internal/service"
	"tree/pkg/parameters"
)

func Start() {
	varFlag := parameters.GetFlag()
	varPath := parameters.GetPath()

	if !varFlag {
		service.OnlyDir(varPath, "")
	} else {
		service.DirsAndFiles(varPath, "")
	}
}

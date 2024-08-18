package service

import (
	"fmt"
	"io/fs"
)

func BuildingOnlyDir(entry fs.DirEntry, prefix string, enLast bool) {
	line := prefix
	if enLast {
		line += "└───"
	} else {
		line += "├───"
	}

	line += entry.Name()
	fmt.Println(line)
}

func BuildingDirsAndFiles(entry fs.DirEntry, prefix string, enLast bool) {
	line := prefix
	if enLast {
		line += "└───"
	} else {
		line += "├───"
	}

	line += entry.Name()
	if !entry.IsDir() {
		info, err := entry.Info()
		if err == nil {
			if info.Size() == 0 {
				line += " (empty)"
			} else {
				line += fmt.Sprintf(" (%d)", info.Size())
			}
		} else {
			line += " (unknown size)"
		}
	}
	fmt.Println(line)
}

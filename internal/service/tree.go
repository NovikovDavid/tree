package service

import (
	"fmt"
	"io/fs"
	"log"
	"os"
	"path/filepath"
)

func OnlyDir(path, prefix string) {
	entries, err := os.ReadDir(path)
	if err != nil {
		log.Fatalf("Error: %v", err)
	}

	var dirs []fs.DirEntry
	for _, entry := range entries {
		if entry.Name() != ".DS_Store" && entry.IsDir() {
			dirs = append(dirs, entry)
		}
	}

	for i, entry := range dirs {
		enLast := i == len(dirs)-1
		func(entry fs.DirEntry, prefix string, enLast bool) {
			line := prefix
			if enLast {
				line += "└───"
			} else {
				line += "├───"
			}

			line += entry.Name()
			fmt.Println(line)
		}(entry, prefix, enLast)
		newPrefix := prefix
		if enLast {
			newPrefix += "    "
		} else {
			newPrefix += "│   "
		}
		subDir := filepath.Join(path, entry.Name())
		OnlyDir(subDir, newPrefix)
	}
}

func DirsAndFiles(path, prefix string) {
	entries, err := os.ReadDir(path)
	if err != nil {
		log.Fatalf("Error: %v", err)
	}

	for i, entry := range entries {
		if entry.Name() == ".DS_Store" {
			continue
		}
		enLast := i == len(entries)-1
		func(entry fs.DirEntry, prefix string, enLast bool) {
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
		}(entry, prefix, enLast)
		if entry.IsDir() {
			newPrefix := prefix
			if enLast {
				newPrefix += "    "
			} else {
				newPrefix += "│   "
			}
			subDir := filepath.Join(path, entry.Name())
			DirsAndFiles(subDir, newPrefix)
		}
	}
}

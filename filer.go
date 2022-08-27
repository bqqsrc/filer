package filer

import (
	"os"
	"path"
)

func PathExist(path string) bool {
	_, err := os.Lstat(path)
	return !os.IsNotExist(err)
}

func IsFile(path string) bool {
	file, err := os.Stat(path)
	if err != nil {
		return false
	}
	return !file.IsDir()
}

func JoinPathExist(args ...string) bool {
	path := path.Join(args...)
	return PathExist(path)
}

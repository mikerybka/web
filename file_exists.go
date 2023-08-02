package web

import "os"

func fileExists(filename string) bool {
	fi, err := os.Stat(filename)
	if os.IsNotExist(err) {
		return false
	}
	return !fi.IsDir()
}

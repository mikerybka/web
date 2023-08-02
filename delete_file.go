package web

import "os"

func deleteFile(filename string) error {
	return os.Remove(filename)
}

package web

import "os"

func deleteAllFiles(dir string) error {
	return os.RemoveAll(dir)
}

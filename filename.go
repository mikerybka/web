package web

import "path/filepath"

func filename(path ...string) string {
	return filepath.Join(path...) + ".json"
}

package web

import (
	"path/filepath"
)

type DB struct {
	Dir string
}

func NewDB(dir string) *DB {
	return &DB{Dir: dir}
}

func (db *DB) Users() *UserStore {
	return &UserStore{Dir: filepath.Join(db.Dir, "users")}
}

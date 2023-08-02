package web

import (
	"encoding/json"
	"os"
)

type Index struct {
	File string
}

func (i *Index) Get(key string) (string, bool) {
	b, err := os.ReadFile(i.File)
	if err != nil {
		return "", false
	}
	index := map[string]string{}
	err = json.Unmarshal(b, &index)
	if err != nil {
		panic(err)
	}
	val, ok := index[key]
	return val, ok
}

func (i *Index) Set(key, val string) error {
	b, err := os.ReadFile(i.File)
	if err != nil {
		return err
	}
	index := map[string]string{}
	err = json.Unmarshal(b, &index)
	if err != nil {
		panic(err)
	}
	index[key] = val
	b, err = json.Marshal(index)
	if err != nil {
		panic(err)
	}
	err = os.WriteFile(i.File, b, 0644)
	if err != nil {
		return err
	}
	return nil
}

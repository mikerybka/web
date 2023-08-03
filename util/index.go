package util

import (
	"encoding/json"
	"os"

	"github.com/mikerybka/web/types"
)

type Index struct {
	File string
}

func (i *Index) Get(key string) (types.ID, bool) {
	b, err := os.ReadFile(i.File)
	if err != nil {
		return 0, false
	}
	index := map[string]types.ID{}
	err = json.Unmarshal(b, &index)
	if err != nil {
		panic(err)
	}
	val, ok := index[key]
	return val, ok
}

func (i *Index) Set(key string, val types.ID) error {
	b, err := os.ReadFile(i.File)
	if err != nil {
		return err
	}
	index := map[string]types.ID{}
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

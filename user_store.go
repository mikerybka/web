package web

import (
	"encoding/json"
	"os"
	"path/filepath"
)

type UserStore struct {
	Dir string
}

func (us *UserStore) Get(id ID) (*User, error) {
	b, err := os.ReadFile(filepath.Join(us.Dir, id.String()+".json"))
	if err != nil {
		return nil, err
	}
	var user User
	err = json.Unmarshal(b, &user)
	if err != nil {
		return nil, err
	}
	return &user, nil
}

func (us *UserStore) Put(id ID, user User) error {
	b, err := json.MarshalIndent(user, "", "  ")
	if err != nil {
		return err
	}
	err = os.WriteFile(filepath.Join(us.Dir, id.String()+".json"), b, 0644)
	if err != nil {
		return err
	}
	err = us.EmailIndex().Set(user.Email, id.String())
	if err != nil {
		return err
	}
	err = us.PhoneIndex().Set(user.Phone, id.String())
	if err != nil {
		return err
	}
	return nil
}

func (us *UserStore) EmailIndex() *Index {
	return &Index{File: filepath.Join(us.Dir, "emails.json")}
}

func (us *UserStore) PhoneIndex() *Index {
	return &Index{File: filepath.Join(us.Dir, "phones.json")}
}

func (us *UserStore) NextID() ID {
	b, _ := os.ReadFile(filepath.Join(us.Dir, "nextid.txt"))
	var id ID
	json.Unmarshal(b, &id)
	nextID := id + 1
	b, _ = json.Marshal(nextID)
	err := os.WriteFile(filepath.Join(us.Dir, "nextid.txt"), b, 0644)
	if err != nil {
		panic(err)
	}
	return id
}

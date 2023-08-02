package web

import (
	"crypto/rand"
	"encoding/hex"
	"encoding/json"
	"os"
	"path/filepath"
)

type SessionStore struct {
	Dir string
}

func (store *SessionStore) Create(userID ID) (string, error) {
	randomBytes := make([]byte, 32)
	_, err := rand.Read(randomBytes)
	if err != nil {
		return "", err
	}
	var token string
	for {
		token = hex.EncodeToString(randomBytes)
		filename := filepath.Join(store.Dir, token+".json")
		if !fileExists(filename) {
			break
		}
	}
	session := Session{UserID: userID}
	err = store.Put(token, session)
	if err != nil {
		return "", err
	}
	return token, nil
}

func (store *SessionStore) Put(token string, session Session) error {
	b, err := json.MarshalIndent(session, "", "  ")
	if err != nil {
		return err
	}
	err = os.WriteFile(filepath.Join(store.Dir, token+".json"), b, 0644)
	if err != nil {
		return err
	}
	return nil
}

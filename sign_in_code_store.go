package web

type SignInCodeStore struct {
	Dir string
}

func (store *SignInCodeStore) Create(userID ID) (string, error) {
	code := SignInCode{Code: newSixDigitCode(), UserID: userID}
	for {
		if !fileExists(filename(store.Dir, code.Code)) {
			break
		}
		code.Code = newSixDigitCode()
	}
	err := store.Put(code.Code, code)
	if err != nil {
		return "", err
	}
	return code.Code, nil
}

func (store *SignInCodeStore) Get(code string) (*SignInCode, error) {
	var sc SignInCode
	err := readJSON(filename(store.Dir, code), &sc)
	if err != nil {
		return nil, err
	}
	return &sc, nil
}

func (store *SignInCodeStore) Put(code string, sc SignInCode) error {
	return writeJSON(filename(store.Dir, code), sc)
}

func (store *SignInCodeStore) Delete(code string) error {
	return deleteFile(filename(store.Dir, code))
}

func (store *SignInCodeStore) DeleteAll() error {
	return deleteAllFiles(store.Dir)
}

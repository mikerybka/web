package auth

import (
	"flag"
	"net/http"

	"github.com/mikerybka/web"
)

func SignUp() {
	var firstName, lastName, email, phone string
	flag.StringVar(&firstName, "first-name", "", "")
	flag.StringVar(&lastName, "last-name", "", "")
	flag.StringVar(&email, "email", "", "")
	flag.StringVar(&phone, "phone", "", "")
	flag.Parse()

	db := web.NewDB("/data")

	_, ok := db.Users().EmailIndex().Get(email)
	if ok {
		web.Return(http.StatusBadRequest, "email already registered")
	}

	_, ok = db.Users().PhoneIndex().Get(phone)
	if ok {
		web.Return(http.StatusBadRequest, "phone already registered")
	}

	id := db.Users().NextID()

	user := web.User{
		ID:        id,
		FirstName: firstName,
		LastName:  lastName,
		Email:     email,
		Phone:     phone,
	}

	err := db.Users().Put(id, user)
	if err != nil {
		web.Return(http.StatusInternalServerError, err.Error())
	}
}

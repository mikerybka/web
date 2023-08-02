package auth

import (
	"net/http"

	"github.com/mikerybka/web"
)

func SignUp() {
	r := web.Request("/auth/sign-up", "POST")
	r.ParseForm()
	firstName := r.FormValue("first-name")
	lastName := r.FormValue("last-name")
	email := r.FormValue("email")
	phone := r.FormValue("phone")

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

	web.Redirect("/auth/sign-in")
}

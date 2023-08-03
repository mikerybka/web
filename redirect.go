package web

import (
	"net/http"
)

func Redirect(url string) {
	AddHeader("Location", url)
	exit(http.StatusTemporaryRedirect)
}

package web

import (
	"net/http"
	"os"
)

func Redirect(url string) {
	AddHeader("Location", url)
	os.Exit(http.StatusTemporaryRedirect)
}

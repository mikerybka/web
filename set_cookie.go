package web

import "fmt"

func SetCookie(name, value string) {
	AddHeader("Set-Cookie", fmt.Sprintf("%s=%s", name, value))
}

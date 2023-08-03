package web

import (
	"fmt"
)

func Return(statusCode int, err string) {
	fmt.Println(err)
	exit(statusCode)
}

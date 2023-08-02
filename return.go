package web

import (
	"fmt"
	"os"
)

func Return(statusCode int, err string) {
	fmt.Println(err)
	os.Exit(statusCode)
}

package web

import (
	"fmt"
	"os"
)

func Return(statusCode int, err string) {
	fmt.Println(err)
	exitCode := statusCode
	if exitCode == 200 {
		exitCode = 0
	}
	os.Exit(exitCode)
}

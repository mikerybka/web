package web

import (
	"os"

	"github.com/mikerybka/webmachine/pkg/data"
)

func exit(status int) {
	code := data.ExitCodes[status]
	os.Exit(code)
}

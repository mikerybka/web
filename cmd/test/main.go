package main

import (
	"crypto/rand"
	"fmt"
	"math/big"
)

func main() {
	i := 0
	for {
		i++
		code := newSixDigitCode()
		if code[0] == '0' {
			if code[1] == '0' {
				// if code[2] == '0' {
				fmt.Println(code)
				fmt.Println(i)
				break
				// }
			}
		}
	}
}

func newSixDigitCode() string {
	n, err := rand.Int(rand.Reader, big.NewInt(1_000_000))
	if err != nil {
		panic(err)
	}
	s := n.String()
	for len(s) < 6 {
		s = "0" + s
	}
	return s
}

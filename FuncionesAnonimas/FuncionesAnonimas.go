package main

import (
	"fmt"
	"strings"
)

func main() {
	cadenas := "abcde"
	cadenas = strings.Map(func(r rune) rune {
		return r + 2
	}, cadenas)

	fmt.Println(cadenas)
}

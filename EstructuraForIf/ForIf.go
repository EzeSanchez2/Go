package main

import (
	"fmt"
)

func main() {
	for a := 1; a < 10; a++ {
		fmt.Println(a)
	}

	var b int32 = 52
	if b > 60 {
		fmt.Println("El numero b es mas grande que 60")
	} else if b == 60 {
		fmt.Println("El numero b es igual a 60")
	} else {
		fmt.Println("El numero b es menor")
	}
}

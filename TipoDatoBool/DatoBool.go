package main

import (
	"fmt"
)

func main() {
	var resultado bool
	resultado = 5 < 6
	fmt.Println("Resultado: ", resultado)
	//resultado = 7 > 6
	//fmt.Println("Resultado: ", resultado)

	//resultado = 7 > 6 && 3 > 5
	//fmt.Println("Resultado: ", resultado)

	//resultado = 1 > 5 || 3 < 5
	//fmt.Println("Resultado : ", resultado)

	fmt.Println("Resultado!: ", !resultado)

}

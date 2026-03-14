package main

import (
	"fmt"
)

func main() {

	fmt.Println("***********ADIVINAR NUMERO SECRETO***********")

	//Pido un numero
	NumeroSecreto := 50 // LEE ESTE NUMERO
	for true { // BUCLE INFINITO
		var NumeroAdivinar int // DEFINO UN NUMERO PARA ADIVINAR
		fmt.Println("Ingrese un numero a adivinar: ")
		fmt.Scanln(&NumeroAdivinar)
		if NumeroAdivinar > NumeroSecreto {
			fmt.Println("EL NUMERO QUE PUSISTE ES MAYOR AL NUMERO SECRETO")

		} else if NumeroAdivinar < NumeroSecreto {
			fmt.Println("EL NUMERO QUE PUSISTE ES MENOR AL NUMERO SECRETO")

		} else {
			fmt.Println("GANASTE EL NUMERO SECRETO ES: ", NumeroSecreto)
			break
		}

	}

}

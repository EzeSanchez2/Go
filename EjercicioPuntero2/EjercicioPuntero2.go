package main

import (
	"fmt"
)

func Transferir(origen *int, destino *int, monto int) {
	if *origen >= monto {
		*origen = *origen - monto
		*destino = *destino + monto
		fmt.Printf("Transferencia exitosa de $%d \n", monto)
	} else {
		fmt.Println("Fondos insuficientes")
	}
}

func main() {
	cuentaEze := 5000
	cuentaAmiga := 1000

	fmt.Println("Saldo de Eze: ", cuentaEze)
	fmt.Println("Saldo amiga: ", cuentaAmiga)

	Transferir(&cuentaEze, &cuentaAmiga, 1000)

	fmt.Println("Saldo de Eze: ", cuentaEze)
	fmt.Println("Saldo de amiga: ", cuentaAmiga)
}

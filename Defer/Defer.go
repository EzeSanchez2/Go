package main

import (
	"fmt"
	"os"
)

func main() {
	f, err := os.Open("texto.txt")

	if err != nil {
		panic(err)
	}

	defer f.Close() //Esto es lo ultimo que se ejecuta, si llega a haber un error en el medio esto permita que salga , y no que se cuelgue el sistema

	data := make([]byte, 10)
	e, err := f.Read(data)

	if err != nil {
		panic(err)
	}

	fmt.Printf("LA CANTIDAD DE BYTE: %d , TEXTO: %q , ERRORES: %v", e, data, err)

	f.Close()
}

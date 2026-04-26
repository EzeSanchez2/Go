package main

import (
	"errors"
	"fmt"
)

var (
	ErrorUsuarioInvalido  = errors.New("El usuario no es valido")
	ErrorUsuarioEnProceso = errors.New("El usuario esta en proceso")
	ErrorPorDefecto       = errors.New("Ocurrio un error")
)

func baneado(usuario string) (err error) {
	ban := false
	switch usuario {
	case "Myriam":
		ban = true
	case "Silvia":
		return ErrorUsuarioInvalido
	case "Walter":
		ban = true
	case "Roxi":
		return ErrorUsuarioEnProceso
	default:
		return ErrorPorDefecto
	}

	if !ban {
		fmt.Printf("El usuario: %q esta PERMITIDO", usuario)
	} else {
		fmt.Printf("El usuario: %q esta BANEADO", usuario)
	}

	return nil
}

func chekearError(err error) {
	if err != nil {
		fmt.Println("ERROR: ", err)
	}
}

func main() {
	err := baneado("Silvia")
	chekearError(err)
}

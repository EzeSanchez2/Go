package main

import (
	"fmt"
)

type Credenciales struct {
	usuario  string
	password string
}

type Usuario struct {
	Credenciales
	Nombre string
	Rol    string
}

func (a Usuario) Validar(u string, p string) bool {
	if a.usuario != u && a.password != p {
		fmt.Println("Usuario y contraseña incorrecta")
		return false
	}
	return true
}

func main() {
	miUsuario := Usuario{
		Credenciales{
			usuario:  "eze",
			password: "bruceyzaira",
		},
		"Ezequiel",
		"Informatica",
	}
	miUsuario.usuario = "ezeale"
	fmt.Println(miUsuario.usuario)
}

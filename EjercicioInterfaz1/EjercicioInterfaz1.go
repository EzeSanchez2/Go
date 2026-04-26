package main

import (
	"fmt"
)

type Email struct {
	Direccion string
}

type SMS struct {
	Numero string
}

func (e Email) Enviar(mensaje string) {
	fmt.Printf("ENVIANDO EMAIL A %q: %q \n", e.Direccion, mensaje)
}

func (s SMS) Enviar(mensaje string) {
	fmt.Printf("ENVIANDO SMS AL %q: %q  \n", s.Numero, mensaje)
}

type Notificador interface {
	Enviar(mensaje string)
}

func EjecutarNotificacion(a Notificador, mensaje string) {
	a.Enviar(mensaje)
}

func main() {
	email := Email{
		Direccion: "ezealesan17@gmail.com",
	}

	sms := SMS{
		Numero: "1166553788",
	}

	EjecutarNotificacion(email, "RECIBISTE $500")
	EjecutarNotificacion(sms, "RECIBISTE $500 ")

	fmt.Println("--------------------------------------")
	lista := []Notificador{
		email,
		sms,
	}
	for _, n := range lista {
		n.Enviar("Hola desde el bucle")
	}
}

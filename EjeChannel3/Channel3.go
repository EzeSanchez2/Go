package main

import (
	"fmt"
	"time"
)

func UsuarioEscribe(ch chan string) {
	time.Sleep(2 * time.Second)
	ch <- "Hola , necesito ayuda con Tarasca\n"
	fmt.Println("Envio mensaje 1")

	time.Sleep(3 * time.Second)
	ch <- "¿Hay alguien ahi?\n"
	fmt.Printf("Se envio el mensaje 2\n")

	ch <- "Perdon me olvide de responder.¡Gracias!\n"
	fmt.Printf("Enviado mensaje 3\n")
}
func main() {
	chanMensajes := make(chan string)
	go UsuarioEscribe(chanMensajes)
	for {
		select {
		case dato := <-chanMensajes:
			fmt.Printf("Recibi tu mensaje: %s", dato)

		case <-time.After(7 * time.Second):
			fmt.Println("Chat cerrado por inactividad")
			return
		}
	}
}

package main

import (
	"fmt"
	"time"
)

func ImprimirPing(ch chan string) {
	var contador int
	for {
		contador++
		fmt.Println(<-ch, ": ", contador)
		time.Sleep(time.Second * 1)

	}
}

func EnviarPing(ch chan string) {
	for {
		ch <- "Ping"
	}
}

func main() {
	ch := make(chan string)
	go EnviarPing(ch)
	go ImprimirPing(ch)

	var Input string
	fmt.Scanln(&Input)
	fmt.Println("Fin ...")
}

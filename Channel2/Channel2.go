package main

import (
	"fmt"
	"time"
)

func main() {
	numeros := make(chan int)
	cuadrado := make(chan int)

	go func() {
		for x := 0; x < 5; x++ {
			numeros <- x
		}

		close(numeros)
	}()

	go func() {
		for x := range numeros {
			cuadrado <- x * x
		}
		close(cuadrado)
	}()

	for x := range cuadrado {
		fmt.Println(x)
		time.Sleep(1 * time.Second)
	}

}

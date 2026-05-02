package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

var wg sync.WaitGroup //Espera a que las gorutinas terminen antes de cerrar el programa

func main() {
	wg.Add(2)
	fmt.Println("Iniciamos las gorutinas ...")

	go imprimirCantidad("A")
	go imprimirCantidad("B")

	fmt.Println("Esperando que finalice ...")
	wg.Wait() // Si no pusiera esto el main llegaría al final, imprimiría "Terminando el programa" y se cerraría instantáneamente, matando a las goroutines antes de que lleguen a imprimir nada.
	fmt.Println("\nTerminando el programa")
}

func imprimirCantidad(etiqueta string) {
	//Indicamos que lo gorutina termino una a una
	defer wg.Done()

	for a := 0; a <= 10; a++ {
		sleep := rand.Int63n(1000)
		time.Sleep(time.Duration(sleep) * time.Millisecond)
		fmt.Printf("Cantidad: %d de %s\n", a, etiqueta)
	}
}

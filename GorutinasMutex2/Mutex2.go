package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

var (
	Saldo float64 = 1000
	mu    sync.Mutex
	wg    sync.WaitGroup
)

func HacerTransacciones(monto float64, wg *sync.WaitGroup) {
	defer wg.Done()
	sleep := rand.Int63n(50)
	time.Sleep(time.Duration(sleep) * time.Millisecond)

	mu.Lock()
	Saldo -= monto
	fmt.Println("El monto retirado es de: ", Saldo)
	mu.Unlock()

}

func main() {
	for a := 0; a < 10; a++ {
		wg.Add(1)
		fmt.Println("--Iniciando Transaccion --")
		go HacerTransacciones(100, &wg)
	}

	wg.Wait()
	fmt.Println("Saldo Final: ", Saldo)
	fmt.Println("-- Terminando Transaccion --")
}

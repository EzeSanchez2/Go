package main

import (
	"fmt"
	"sync"
	"time"
)

func ProcesarPagos(ch chan float64, wg *sync.WaitGroup) {
	defer wg.Done()
	for monto := range ch {
		if monto > 0 {
			fmt.Printf("Pago de %f procesado con exito \n", monto)
		}
		if monto < 0 {
			fmt.Printf("Error %f inavlido \n", monto)
		}

		time.Sleep(time.Millisecond * 500)
	}
}

func GenerarPagos(ch chan float64) {
	montos := []float64{150.50, -10.0, 500.0, 0.0, 1200.0}
	for _, m := range montos {
		ch <- m
	}
	close(ch)
}

func main() {
	var wg sync.WaitGroup
	chanPagos := make(chan float64)
	go GenerarPagos(chanPagos)
	wg.Add(1)
	go ProcesarPagos(chanPagos, &wg)

	wg.Wait()
}

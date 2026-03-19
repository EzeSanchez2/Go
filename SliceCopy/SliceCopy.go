package main

import (
	"fmt"
)

func main() {
	origen := []int{1, 2}
	destino := []int{3, 6, 7}
	copy(destino, origen)
	fmt.Println(origen, destino)
}

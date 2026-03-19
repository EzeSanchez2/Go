package main

import (
	"fmt"
)

func main() {

	x := make([]byte, 5, 10)
	fmt.Println(x)

	//Inicializamos x
	x = []byte{'H', 'O', 'L', 'A'}
	fmt.Println(x)

	for i := 0; i < len(x); i++ {
		fmt.Printf("Slice x[%d]: %q \n", i, x[i])
	}
	x = append(x, ' ')
	x = append(x, 'M', 'U', 'N', 'D', 'O')
	fmt.Printf("Slice x: %q  \n", x)
	fmt.Println("Capacidad: ", cap(x))

	fmt.Println("*****************************************************************************")
	var y []int
	for e := 1; e < 15; e++ {
		y = append(y, e)
		fmt.Println("Slice y: ", y)
		fmt.Printf("Len y: [%d] , Cap y: [%d], Elementos: %d \n", len(y), cap(y), e)
	}

}

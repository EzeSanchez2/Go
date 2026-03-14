package main

import (
	"fmt"
)

func main() {
	var f32 float32     // 6 digitos decimales
	var f64 float64     // 15 digitos decimales
	var c64 complex64   // Numeros complejos para float 32
	var c128 complex128 // Numeros complejos para float 64

	fmt.Println("------------------------------")

	fmt.Println("f32 , f64 , c64 , c128 =  ", f32, f64, c64, c128)
	f32 = 0.150
	f64 = 0.150

	fmt.Println("f32 * 15.567421 =  ", f32*15.567421)
	fmt.Println("f64 *15.567421= ", f64*15.567421)

	//Numeros Complejos
	c64 = complex(5, 6)
	println(c64 * 2002.2121)

	c128 = complex(5, 6)
	println(c128 * 2002.2121)
}

package main

import (
	"fmt"
	"unsafe"
)

func main() {

	//Enteros con signos
	//var entero8 int8	 // (-128 a 128)
	//var entero16 int16 // (-32.768 a 32.767)
	//var entero32 int32 // (-2.147.483.648 a 2.147.483.647)
	var entero64 int64 // (-9.22 × 10¹⁸ a 9.22 × 10¹⁸)

	// Entero sin signo
	//var uenteros8 uint8 // (0 a 255)
	//var uenteros16 uint16 // (0 a 65.535)
	//var uenteros32 uint32 // (0 a 4 mil millones)
	//var uenteros64 uint64 // (0 a 18 trillones)

	//Alias
	//var enteroByte byte // alias para uint8
	//var enteroRune rune // alias para int32

	//Tipos dependiente de la plataforma
	//var enteroUint uint //32 a 64 bits
	var enteroint int // 32 a 64 bits
	//var enteroUitnpr uintptr // enteros sin signos lo suficientemente grandes para contener al puntero

	//----------------------------------------//
	//Conversion de datos

	//Suma int32 y int64

	//entero32 = 25
	//entero64 = 60
	//fmt.Println(entero32 + int32(entero64))

	//Suma int32 y rune

	//entero32 = 20
	//fmt.Println(entero32 + enteroRune)

	//Suma int16 y uint16

	//entero16 = 200
	//uenteros16 = 20

	//fmt.Println(entero16 + int16(uenteros16))

	//Unsafe
	fmt.Println(unsafe.Sizeof(enteroint), unsafe.Sizeof(entero64))

}

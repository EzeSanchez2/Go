package main

import "fmt"

func main() {
	var dia int
	fmt.Println("ELIGE UN NUMERO PARA SELECCIONAR EL DIA")
	fmt.Println("1. Lunes")
	fmt.Println("2. Martes")
	fmt.Println("3. Miercoles")
	fmt.Println("4. Jueves")
	fmt.Println("5. Viernes")
	fmt.Println("6. Sabado")
	fmt.Println("7. Domingo")
	fmt.Scanln(&dia)

	switch dia {
	case 1:
		fmt.Println("Elegiste el dia Lunes")
	case 2:
		fmt.Println("Elegiste el dia Martes")
	case 3:
		fmt.Println("Elegiste el dia Miercoles")
	case 4:
		fmt.Println("Elegiste el dia Jueves")
	case 5:
		fmt.Println("Elegiste el dia Viernes")
	case 6:
		fmt.Println("Elegiste el dia Sabado")
	case 7:
		fmt.Println("Elegiste el dia Domingo")
	default:
		fmt.Println("Seleecione el numero correcto")

	}

}

package main

import(
	"fmt"
)
//LIFO : 
func fun1()  {
	fmt.Println("Entrando en f1") //3. ENTRA EN LA FUN1
	fun2() // 4. VA A LA FUN2
	fmt.Println("Saliendo de f1")//10. SE TERMINA DE RESOLVER LA FUN1
}
func fun2()  {
	fmt.Println("Entrando en f2") // 5. EJECUTA LA FUN2
	fun3() //6. PASA A LA FUN3
	fmt.Println("Saliendo de f2") //9. SE TERMINA DE RESOLVER LA FUN2
}
func fun3()  {
	fmt.Println("Saliendo de f3") //7. SE EJECUTA LA FUN3
	fmt.Println("Saliendo de f3") //8. SE TERMINA RESOLVIENDO LA FUN3 (ES LA ULTIMA FUNCION ES LA PRIMERA QUE SE RESUELVE)
}

func main()  {
	fmt.Println("Entrando en el main") // 1 . SE EJECUTA EL MAIN
	fun1() // 2. PASA A LA FUN1
	fmt.Println("Saliendo del main ") //11. POR ULTIMO SE RESUELVE EL MAIN

}
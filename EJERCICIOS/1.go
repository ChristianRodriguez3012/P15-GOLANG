// EJERCICIO 1: Desarrolle un programa en Goland que solicite un número por teclado y
// luego imprima por pantalla el factorial de dicho número (4 puntos).
package main

import "fmt"

//fmt viene de format package y es el paquete que se encargará de darle formato a cualquier tipo de entrada o salida de dato.
func main() {
	var n int
	var fact int = 1
	fmt.Print("DETERMINAR FACTORIAL DE N")
	//input
	fmt.Print("\nDETERMINE UN NÚMERO ENTERO: ")
	fmt.Scan(&n)
	if n < 0 {
		//condicional: si el número ingresado es negativo
		fmt.Print("\nEL FACTORIAL DE UN NÚMERO NEGATIVO NO EXISTE")
	}
	//ciclo for: calcular factorial
	for i := 1; i <= n; i++ {
		fact = fact * i
	}
	//output
	fmt.Printf("\nEL FACTORIAL ES %d", fact)
}

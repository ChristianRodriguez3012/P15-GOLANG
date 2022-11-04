//EJERCICIO 2: Desarrolle un programa en Goland que solicite un número por teclado y luego imprima
//por pantalla si dicho número es primo o no (4 puntos).

package main

import "fmt"

func esPrimo(x int) bool {

	var cont = 0
	for i := 1; i <= x; i++ {
		if x%i == 0 {
			cont++
		}
	}
	if cont == 2 {
		return true
	} else {
		return false
	}
}

func main() {
	var n int
	fmt.Print("NÚMEROS PRIMOS")
	fmt.Print("\nDETERMINE UN NÚMERO ENTERO: ")
	fmt.Scan(&n)
	fmt.Print("\n¿EL NÚMERO INGRESADO ES PRIMO?: ", esPrimo(n))
}

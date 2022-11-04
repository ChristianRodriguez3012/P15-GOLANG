//EJERCICIO 4:El número de Euler, e = 2, 71828
//puede ser representado como la siguiente suma infinita:

//e = (1/0!) + (1/1!) + (1/2!) + (1/3!) + (1/4!) + ...

// Desarrolle un programa que entregue un valor aproximado de e, calculando
//esta suma hasta quela diferencia entre dos sumandos consecutivos sea menor
//que 0,01 (8 puntos).

package main

import "fmt"

var factVal uint64 = 1
var i int = 1

//por si es negativo
func factorial(n int) uint64 {
	if n < 0 {
		fmt.Print("EL FACTORIAL DE N NEGATIVO NO EXISTE.")
	} else {
		for i := 1; i <= n; i++ {
			factVal *= uint64(i)
		}
	}
	return factVal
}

//calculando
func euler(n int) float64 {
	var sum float64
	sum = 0
	var term float64
	for i := 0; i <= n; i++ {
		term = 1 / float64(factorial(i))
		sum = sum + term
	}
	return sum
}

//MAIN: INPUT AND OUTPUT
func main() {
	var n int
	fmt.Print("\nNUMERO DE EULER")
	fmt.Print("\nINGRESE UN NUMERO (DEBE SER ENTERO): ")
	fmt.Scan(&n)
	fmt.Print("\nFACTORIAL: ", euler(n))
}

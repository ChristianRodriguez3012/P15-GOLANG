//EJERCICIO 4:

package main

import "fmt"

var factVal uint64 = 1
var i int = 1

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

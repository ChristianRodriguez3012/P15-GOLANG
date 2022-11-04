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
	fmt.Print("\nDetermine un número (tiene que ser entero): ")
	fmt.Scan(&n)
	fmt.Print("\nEl número ingresado es primo: ", esPrimo(n))
}

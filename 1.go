package main

import "fmt"

//fmt viene de format package y es el paquete que se encargará de darle formato a cualquier tipo de entrada o salida de dato.
func main() {
	var n int
	var fact int = 1
	fmt.Print("DETERMINAR FACTORIAL DE N")
	//input
	fmt.Print("\nIngresar un número: ")
	fmt.Scan(&n)
	if n < 0 {
		//condicional: si el número ingresado es negativo
		fmt.Print("\nEl factorial de un número negativo no existe.")
	}
	//ciclo for: calcular factorial
	for i := 1; i <= n; i++ {
		fact = fact * i
	}
	//output
	fmt.Printf("\nEl factorial es: %d", fact)
}

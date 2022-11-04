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

//EJERCICIO 2: Desarrolle un programa en Goland que solicite un número por teclado y luego imprima
//por pantalla si dicho número es primo o no (4 puntos).

///package main

///import "fmt"

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

//EJERCICIO 3:Desarrolle un programa en Goland que solicite un número por teclado y luego imprima por
//pantalla todos los números que pueden dividir a dicho número (4 puntos).

///package main

///import "fmt"

func main() {
	var n int
	fmt.Print("\nDIVISORES DE UN NÚMERO")
	fmt.Print("\nDETERMINE UN NÚMERO ENTERO: ")
	fmt.Scanf("%d", &n)
	fmt.Println("\nLOS DIVISORES DE", n, " SON: ")
	for i := 1; i <= n; i++ {
		if n%i == 0 {
			fmt.Print(i, "\n")
		}
	}
}

//EJERCICIO 4:El número de Euler, e = 2, 71828
//puede ser representado como la siguiente suma infinita:

//e = (1/0!) + (1/1!) + (1/2!) + (1/3!) + (1/4!) + ...

// Desarrolle un programa que entregue un valor aproximado de e, calculando
//esta suma hasta quela diferencia entre dos sumandos consecutivos sea menor
//que 0,01 (8 puntos).

//package main

//import "fmt"

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

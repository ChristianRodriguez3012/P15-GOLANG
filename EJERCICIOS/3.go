//EJERCICIO 3:Desarrolle un programa en Goland que solicite un número por teclado y luego imprima por
//pantalla todos los números que pueden dividir a dicho número (4 puntos).

package main

import "fmt"

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

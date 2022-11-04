//EJERCICIO 3:

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

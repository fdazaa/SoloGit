package main

import "fmt"

var operacion int

func main() {
	fmt.Println("\nCALCULADORA GO")
	fmt.Println("Digite el número del tipo de operacion a realizar")

	fmt.Println("\n(1) SUMA")
	fmt.Println("(2) RESTA")
	fmt.Println("(3) MULTIPLICACION")
	fmt.Println("(4) DIVISIÓN")

	fmt.Println("__")
	fmt.Print("> ")
	fmt.Scan(&operacion)
	fmt.Println("")

	switch operacion {
	case 1:
		println("SUMA")
	case 2:
		println("RESTA")
	case 3:
		println("MULTIPLICACIÓN")
	case 4:
		println("DIVISIÓN")
	}
}

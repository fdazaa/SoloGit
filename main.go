package main

import "fmt"

var c, d float32
var a, b int

func main() {
	var operacion int
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
		suma()
	case 2:
		println("RESTA")
	case 3:
		println("MULTIPLICACIÓN")
	case 4:
		println("DIVISIÓN")
	}
}

func suma() {
	println("Digite los numeros a sumar")
	print("Numero 1: ")
	fmt.Scan(&a)
	print("Numero 2: ")
	fmt.Scan(&b)
	println("La suma es: ", a+b)
	//Para regresar a la calculadora
	main()
}
func division() {
	println("Digite los numeros a dividir")
	print("Numero 1: ")
	fmt.Scan(&c)
	print("Numero 2: ")
	fmt.Scan(&d)
	fmt.Printf("El cociente es: %f\n", c/d)
}

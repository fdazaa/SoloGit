package main

import "fmt"

var a, b float32

func main() {
	division()
}
func division() {
	println("Escriba los numeros a dividir")
	print("Numero 1: ")
	fmt.Scan(&a)
	print("Numero 2: ")
	fmt.Scan(&b)
	fmt.Printf("El cociente es: %f\n", a/b)
}

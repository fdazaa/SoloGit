package main

import "fmt"

var a, b int

func main() {
	multiplicacion()
}
func multiplicacion() {
	println("Escriba los numeros a multiplicar")
	print("Numero 1: ")
	fmt.Scan(&a)
	print("Numero 2: ")
	fmt.Scan(&b)
	println("El producto es: ", a*b)
}

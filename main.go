package main

import "fmt"

var a, b int

func main() {
	multiplicacion()
}
func multiplicacion() {
	println("Digite los numeros a multiplicar")
	print("Numero 1: ")
	fmt.Scan(&a)
	print("Numero 2: ")
	fmt.Scan(&b)
	println("El resultado de la multiplicacón es: ", a*b)
}

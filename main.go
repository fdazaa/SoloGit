package main

import "fmt"

var a, b int

func main() {
	resta()
}
func resta() {
	println("Escriba los numeros a restar")
	print("Numero 1: ")
	fmt.Scan(&a)
	print("Numero 2: ")
	fmt.Scan(&b)
	println("La resta es: ", a-b)
}

package main

import "fmt"

var a, b int

func main() {

}
func suma() {
	println("Escriba los numeros a sumar")
	print("Numero 1: ")
	fmt.Scan(&a)
	print("Numero 2: ")
	fmt.Scan(&b)
	println("La suma es: ", a+b)
}

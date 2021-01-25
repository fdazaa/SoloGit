package main

import "fmt"

var c, d float32
var a, b int

func main() {

}
func opge() {
	print("Numero 1: ")
	fmt.Scan(&a)
	print("Numero 2: ")
	fmt.Scan(&b)
}

func suma() {
	println("Digite los numeros a sumar")
	opge()
	println("La suma es: ", a+b)
	//volver()
}
func resta() {
	println("Digite los numeros a restar")
	opge()
	println("La resta es: ", a-b)
	//volver()
}
func multiplicacion() {
	println("Digite los numeros a multiplicar")
	opge()
	println("El resultado de la multiplicacón es: ", a*b)
	//volver()
}

func division() {
	println("Digite los numeros a dividir")
	opge()
	fmt.Printf("El cociente es: %f\n", c/d)
	//volver()
}

package main

import "fmt"

func main() {
	porcentaje()
}

var a, b float32 //inicializa las variables en un flotante de 32bits

func opge() {
	print("Numero 1: ")
	fmt.Scan(&a)
	print("Numero 2: ")
	fmt.Scan(&b)
}

func suma() {
	println("Digite los numeros a sumar")
	opge()
	fmt.Printf("La suma es: %.2f\n", a+b)
	//volver()
}
func resta() {
	println("Digite los numeros a restar")
	opge()
	fmt.Printf("La resta es: %.2f\n", a-b)
	//volver()
}
func multiplicacion() {
	println("Digite los numeros a multiplicar")
	opge()
	fmt.Printf("El resultado de la multiplicacón es: %.2f\n", a*b)
	//volver()
}

func division() {
	println("Digite los numeros a dividir")
	opge()
	fmt.Printf("El cociente es: %.2f\n", a/b) //para limitar la cantidad de decimales usar .xf, donde x es el # de decimales
	//volver()
}

func porcentaje() {
	var a, b float32
	println("Digita un numero:")
	fmt.Scan(&a)
	println("Que porcentaje quieres saber?:")
	fmt.Scan(&b)
	fmt.Printf("resultado: %.2f\n", (a*b)/100)
}

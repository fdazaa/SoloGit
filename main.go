package main

import "fmt"

var a, b float32 //inicializa las variables en un flotante de 32bits

func main() {
	var operacion int
	fmt.Println("\nCALCULADORA GO")
	fmt.Println("Digite el número del tipo de operacion a realizar")

	fmt.Println("\n(1) SUMA")
	fmt.Println("(2) RESTA")
	fmt.Println("(3) MULTIPLICACION")
	fmt.Println("(4) DIVISIÓN")
	fmt.Println("(5) PORCENTAJE")

	fmt.Println("__")
	fmt.Print("> ")
	fmt.Scan(&operacion)
	fmt.Println("")

	switch operacion {
	case 1:
		suma()
	case 2:
		resta()
	case 3:
		multiplicacion()
	case 4:
		division()
	case 5:
		porcentaje()
	}
}

// f1 - OPERACIONES
func suma() {
	println("Digite los numeros a sumar")
	opge()
	fmt.Printf("La suma es: %.2f\n", a+b)
	volver()
}

func resta() {
	println("Digite los numeros a restar")
	opge()
	fmt.Printf("La resta es: %.2f\n", a-b)
	volver()
}

func multiplicacion() {
	println("Digite los numeros a multiplicar")
	opge()
	fmt.Printf("El resultado de la multiplicacón es: %.2f\n", a*b)
	volver()
}

func division() {
	println("Digite los numeros a dividir")
	opge()
	fmt.Printf("El cociente es: %.2f\n", a/b) //para limitar la cantidad de decimales usar .xf, donde x es el # de decimales
	volver()
}

func porcentaje() {
	println("Digita un numero:")
	fmt.Scan(&a)
	println("Que porcentaje quieres saber?:")
	fmt.Scan(&b)
	fmt.Printf("resultado: %.2f\n", (a*b)/100)
	volver()
}

func opge() {
	print("Numero 1: ")
	fmt.Scan(&a)
	print("Numero 2: ")
	fmt.Scan(&b)
}

func volver() {
	var volver int
	fmt.Println("\n------------")
	fmt.Println("1 -> VOLVER \n2 -> SALIR")
	fmt.Println("------------")
	fmt.Println("__")
	fmt.Print("> ")
	fmt.Scan(&volver)
	if volver == 1 {
		main()
	} else {
		fmt.Println("\n֎ Bye! ֎")
	}
}

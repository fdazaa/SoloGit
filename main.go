package main

import "fmt"

func main() {

}
func porcentaje() {
	println("Digita un numero:")
	fmt.Scan(&a)
	println("Que porcentaje quieres saber?:")
	fmt.Scan(&b)
	fmt.Printf("resultado: %.2f\n", (a*b)/100)
}

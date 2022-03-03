package main

import "fmt"

var numero int
var tecto string
var status bool = true

func main() {
	//var numero2 int
	//numero3 := 4
	var numero2, numero3, numero4 int
	numero2, numero3, numero4, texto, status := 2, 5, 51, "Esto es mi texto", false

	//var moneda float32 = 0
	fmt.Println(numero2)
	fmt.Println(numero3)
	fmt.Println(numero4)
	fmt.Println(texto)
	fmt.Println(status)

	mostrarStatus()
}

func mostrarStatus() {
	fmt.Println(status)
}

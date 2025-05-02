package main

import "fmt"

func OperadoresAritmeticos() {
	//suma
	x := 10
	y := 50
	result := x + y
	fmt.Println("suma: ", result)
	//resta
	result = y - x
	fmt.Println("resta", result)
	//multiplicacion
	result = y * x
	fmt.Sprintln("Multiplicacion", result)
	//division
	result = y / x
	fmt.Println("Division", result)
	//modulo
	result = y % x
	fmt.Println("modulo", result)
	//incrementar
	x++
	fmt.Println("incrementar", x)
	//decrementar
	x--
	fmt.Println("decremental", x)

	//Area de un trapecio
	var BaseP float64 = 21
	var BaseG float64 = 30
	var Altura float64 = 12
	var constante float64 = 2

	resultado := ((BaseP + BaseG) / constante) * Altura
	fmt.Println("El area de un trapecio es: ", resultado, " cm^2")

}

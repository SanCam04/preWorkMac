package main

import "fmt"

func main() {

	//Declaracion de constantes
	const pi float64 = 3.14
	const pi2 = 3.141516
	fmt.Println("pi", pi)
	fmt.Println("pi2", pi2)
	//Declaracion de variables enteras
	base := 12
	var altura int = 14
	var area int
	fmt.Println(base*altura + area)
	//Zero values
	var a int
	var b float64
	var c string
	var d bool
	fmt.Println(a, b, c, d)
	//area cuadrado
	const baseCuadrado = 10
	areaCuadrado := baseCuadrado * baseCuadrado
	fmt.Println("El area del cuadrado es ", areaCuadrado)
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

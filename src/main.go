package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Println("=== Selecciona una opción ===")
	fmt.Println("1. Operadores Aritméticos")
	fmt.Println("2. Tipos de Datos")
	fmt.Print("Opción: ")

	scanner.Scan()
	opcion := scanner.Text()

	switch opcion {
	case "1":
		OperadoresAritmeticos()
	case "2":
		TipoDatos()
	default:
		fmt.Println("Opción no válida")
	}
}

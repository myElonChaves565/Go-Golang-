package main

import (
	"fmt"
)

func main() {
	var N string
	var I int
	var S float64

	fmt.Print("Informe um nome: ")
	fmt.Scan(&N)

	fmt.Print("Qual é a sua idade: ")
	fmt.Scan(&I)

	fmt.Print("Informe o seu salário: ")
	fmt.Scan(&S)

	if len(N) >= 3 && I <= 150 && S >= 0 {
		fmt.Printf("NOME: %s\n", N)
		fmt.Println("\nVÁLIDO")
		fmt.Printf("IDADE: %d\n", I)
		fmt.Println("\nVÁLIDO")
		fmt.Printf("SALARIO: %.2f\n", S)
		fmt.Println("\nVÁLIDO")
	} else {
		fmt.Println("\nINVÁLIDO")
	}
}

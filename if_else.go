package main

import "fmt"

func main() {
	var num int
	fmt.Print("Digite um número:")
	fmt.Scan(&num)

	if num > 5 {
		fmt.Printf("%d é maior que 5", num)
	} else {
		fmt.Printf("%d é menor ou igual a 5", num)
	}
}

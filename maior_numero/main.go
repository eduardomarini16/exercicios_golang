package main

import "fmt"

func maior(a int, b int) int {
	if a > b {
		return a
	}
	return b
}

func main() {
	var a, b int
	fmt.Println("Digite o primeiro número para comparação")
	fmt.Scan(&a)
	fmt.Println("Digite o segundo número para comparação")
	fmt.Scan(&b)
	resultado := maior(a, b)
	fmt.Printf("O maior número entre %d e %d é %d\n", a, b, resultado)
}

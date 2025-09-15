package main

import "fmt"

func main() {

	var num int

	fmt.Println("Digite a quantidade de números que deseja inserir:")
	fmt.Scan(&num)

	n := make([]int, num)

	fmt.Println("Digite números inteiros, para verificarmos:")

	for i := 0; i < num; i++ {
		fmt.Scan(&n[i])
	}

	n_maior, n_menor := 0, 0

	for i := 0; i < num; i++ {
		if n[i] > n_maior {
			n_maior = n[i]
		}
		if n[i] < n_menor || n_menor == 0 {
			n_menor = n[i]
		}
	}

	fmt.Println(n)
	fmt.Println("O maior número é:", n_maior)
	fmt.Println("O menor número é:", n_menor)

}

package main

import "fmt"

func operacoes(a int, b int) (int, int, int, int, error) {
	soma := a + b
	subtracao := a - b
	mult := a * b
	if b == 0 {
		return 0, 0, 0, 0, fmt.Errorf("divisão por zero não é permitida")
	}
	div := a / b
	return soma, subtracao, mult, div, nil
}

func main() {

	var a, b int
	fmt.Println("Digite o primeiro número para realizar a operação")
	fmt.Scan(&a)
	fmt.Println("Digite o segundo número para realizar a operação")
	fmt.Scan(&b)

	var op string
	fmt.Println("Digite qual operação deseja realizar (soma[+], subtração[-], multiplicação[*], divisão[/])")
	fmt.Scan(&op)

	soma, sub, mult, div, err := operacoes(a, b)

	switch op {
	case "+":
		fmt.Printf("O resultado da soma entre %d e %d é %d\n", a, b, soma)
	case "-":
		fmt.Printf("O resultado da subtração entre %d e %d é %d\n", a, b, sub)
	case "*":
		fmt.Printf("O resultado da multiplicação entre %d e %d é %d\n", a, b, mult)
	case "/":
		if err != nil {
			fmt.Println("Erro", err)
		} else {
			fmt.Printf("O resultado da divisão entre %d e %d é %d\n", a, b, div)
		}
	default:
		fmt.Println("Operação inválida. Por favor, escolha entre +, -, * ou /")
	}
}

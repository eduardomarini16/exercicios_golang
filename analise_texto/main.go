package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func analisaTexto(s string) (int, int) {

	letras := 0
	palavras := 0

	s_semEspaco := strings.ReplaceAll(s, " ", "")

	for _, frase := range s_semEspaco {
		if frase != ' ' {
			letras++
		}
	}

	for _, frase := range s {
		if frase == ' ' {
			palavras++
		}
	}

	return letras, palavras + 1
}

func main() {

	reader := bufio.NewReader(os.Stdin)

	fmt.Println("Digite uma frase para analisar:")
	frase, _ := reader.ReadString('\n')
	frase = strings.TrimSpace(frase)

	letras, palavras := analisaTexto(frase)

	fmt.Printf("A frase possui %d letras e %d palavras\n", letras, palavras)

}

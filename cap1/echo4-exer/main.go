package main

import (
	"fmt"
	"os"
)

func main() {
	// Exercício 1.1: Modique o programa echo para exibir também os.Args[0], que é o nome do comando que o chamou.
	// Exercício 1.2: Modique o programa echo para exibir o índice e o valor de cada um de seus argumentos, um por linha.
	// Exercício 1.3: Experimente medir a diferença de tempo de execução entre nossas versões potencialmente inecientes e a versão que usa strings.Join.
	// (A seção 1.6 mostra parte do pacote time, e a seção 11.4 mostra como escrever testes comparativos para uma avaliação sistemática de desempenho.)
	for indice, valor := range os.Args[0:] {
		// poderíamos usar %d e %s para ter maior controle sobre a formatação se necessário
		// %v é um placeholder genérico
		fmt.Printf("%v:%v\n", indice, valor)
	}
}

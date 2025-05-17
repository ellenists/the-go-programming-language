package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	// Se a quantidade de dados envolvida for grande, os exemplos de echo anteriores podem ser custosos.
	// Uma solução mais simples e eficiente seria usar a função Join:
	fmt.Println(strings.Join(os.Args[1:], " ")) // exemplo de resultado: um dois tres
	// outra forma de fazer é exibir os argumentos de forma "crua", sem formatação:
	fmt.Println(os.Args[1:]) // exemplo de resultado: [um dois tres]
}

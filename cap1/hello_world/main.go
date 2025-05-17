package main

import (
	"fmt"
)

var saudacao = "Hello 世界!"

// saudacao := "Hello 世界!" // ERRO! Não permitido fora de função

func main() { // o que a função main faz é o que o programa faz.
	// Go trata Unicode de modo nativo, portanto é capaz de processar texto em todas as línguas do mundo.
	fmt.Print(saudacao)
	// go build hello_world.go gerou um executável hello_world.exe
}

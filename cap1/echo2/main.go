package main

import (
	"fmt"
	"os"
)

func main() {
	s, separator := "", ""
	// os.Args[1:] significa do índice 1 até o fim do slice
	// os.Args[0] exibiria o nome do executável
	for index, arg := range os.Args[1:] {
		// se não quisermos usar o valor de index
		// podemos usar um identificador vazio, o underscore _
		// for _, arg := range os.Args[1:] {
		fmt.Println(index)
		s += separator + arg
		separator = " "
	}
	fmt.Printf(s)
}

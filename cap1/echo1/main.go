// Por convenção, descrevemos cada pacote em um comentário imediatamente antes de sua declaração de pacote.
package main

import (
	"fmt"
	"os"
)

func main() {
	var s, separator string
	// for inicialização; condição; pós
	for i := 1; i < len(os.Args); i++ {
		s += separator + os.Args[i] // equivalente a: s = s + sep + os.Args[i]
		separator = " "
	}
	fmt.Println(s)

	// inicialização, condição e pós podem ser omitidos.
	a := 0
	for a < 10 {
		fmt.Printf("%v ", a)
		a++
	}

	// um break ou return poderia interromper este looping, mas desta forma ele é infinito.
	for {
		fmt.Println("looping infinito...")
	}

}

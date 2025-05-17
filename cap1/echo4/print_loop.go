package print

import (
	"fmt"
)

func printLoop(lista []string) {
	var s, sep string
	for _, v := range lista[0:] {
		s += sep + v
		sep = " "
	}
	fmt.Print(s)
}

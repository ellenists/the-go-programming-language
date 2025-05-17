package print

import (
	"fmt"
	"strings"
)

func printJoin(lista []string) {
	fmt.Println(strings.Join(lista[0:], " "))
}

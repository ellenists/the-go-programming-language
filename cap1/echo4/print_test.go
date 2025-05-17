// Exercício 1.1: Modique o programa echo para exibir também os.Args[0], que é o nome do comando que o chamou.
// Exercício 1.2: Modique o programa echo para exibir o índice e o valor de cada um de seus argumentos, um por linha.
// Exercício 1.3: Experimente medir a diferença de tempo de execução entre nossas versões potencialmente inecientes e a versão que usa strings.Join.
// (A seção 1.6 mostra parte do pacote time, e a seção 11.4 mostra como escrever testes comparativos para uma avaliação sistemática de desempenho.)
package print

import (
	"testing"
)

var lista = []string{
	"kdjlxqwe", "nmsulvqp", "znbgtuio", "wqpexmvk", "dhxkrylc",
	"lvdmsqti", "yutwznae", "qbamczvu", "hgjomwrt", "xvkjtzle",
	"cbyuqpne", "sfivkram", "odwlcjpt", "rnxgqsuv", "wemaybdk",
	"tplxhvno", "ijqzwkte", "fumxrnby", "azlkgwcp", "koynmvde",
	"qlerihzt", "tfdmnaxw", "ybsoqlkv", "juznrgcb", "shxeqvmo",
	"uracwjyn", "nbqeypzk", "dfzvsmxt", "xamltbio", "velrsdqu",
	"wkgiqhzb", "slbtjvec", "mpoxeylk", "giqztwhb", "buzkmdnr",
	"hxacjvup", "lozsrnbm", "ejwmqitz", "tkvyuhrn", "rydcazke",
	"wqofzvct", "dhlbnyma", "qztmsykw", "kixpjhuv", "tzjbnlre",
	"flomvcds", "neajqwyz", "ugrldxtc", "mxvzsrqu", "jcynwbpk",
}

// to run: go test -bench=Loop
func BenchmarkLoop(b *testing.B) {
	for b.Loop() {
		printLoop(lista)
	}
}

// to run: go test -bench=Join
func BenchmarkJoin(b *testing.B) {
	for b.Loop() {
		printJoin(lista)
	}

}

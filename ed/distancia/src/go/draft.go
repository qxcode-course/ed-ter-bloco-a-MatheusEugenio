package main

import (
	"fmt"
)

func buscaParaEsquerdaRec(indiceLimite, value int, runes []rune, indexAtual int) bool {

	if indexAtual < 0 {
		return false
	}

	if indexAtual < indiceLimite {
		return false
	}

	if runes[indexAtual] == rune('0'+value) {
		return true
	}

	if buscaParaEsquerdaRec(indiceLimite, value, runes, indexAtual-1) {
		return true
	}

	return false
}

func buscaParaDireitaRec(indiceLimite, value int, runes []rune, indexAtual int) bool {

	if indexAtual >= len(runes) {
		return false
	}

	if indexAtual > indiceLimite {
		return false
	}

	if runes[indexAtual] == rune('0'+value) {
		return true
	}

	if buscaParaDireitaRec(indiceLimite, value, runes, indexAtual+1) {
		return true
	}

	return false

}

// se retornar true é pq achou um valor igual ao valor passado em um raio de L
func achouIgual(indexLimiteDir, indexLimiteEsq int, runes []rune, indexAtual int, value int) bool {

	if buscaParaDireitaRec(indexLimiteDir, value, runes, indexAtual) || buscaParaEsquerdaRec(indexLimiteEsq, value, runes, indexAtual) {
		return true
	}

	return false
}

func preencher(valoresL []int, runes []rune, limite int) (bool, string) {

	indexPonto := -1
	for i := range runes {
		if runes[i] == '.' {
			indexPonto = i
		}
	}

	if indexPonto == -1 {
		return true, string(runes)
	}

	limiteDir := indexPonto + limite
	if limiteDir == len(runes) {
		limiteDir = len(runes) - 1
	}

	limiteEsq := indexPonto - limite
	if limiteEsq < 0 {
		limiteEsq = 0
	}

	for _, val := range valoresL {

		if !achouIgual(limiteDir, limiteEsq, runes, indexPonto, val) {

			runes[indexPonto] = rune('0' + val)

			caminhos, _ := preencher(valoresL, runes, limite)
			if caminhos {
				return true, string(runes)
			}

			runes[indexPonto] = '.'
		}
	}

	return false, string(runes)
}

func main() {

	var str string
	fmt.Scan(&str)

	var limite int
	fmt.Scan(&limite)

	var valoresL []int

	for i := 0; i <= limite; i++ {
		valoresL = append(valoresL, i)
	}

	runes := []rune(str)
	_, stringFinal := preencher(valoresL, runes, limite)

	fmt.Println(stringFinal)
}

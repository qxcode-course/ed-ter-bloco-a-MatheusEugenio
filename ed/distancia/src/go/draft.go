package main

import (
	"fmt"
)

func buscaParaEsquerda(limite, value int, runes []rune, indexAtual int) bool {

	cont := 1
	for i := indexAtual; i >= 0; i-- {

		if cont > limite {
			return false
		}

		if runes[i] == rune('0'+value) {
			return true
		}

		cont++
	}

	return false
}

func buscaParaDireita(limite, value int, runes []rune, indexAtual int) bool {

	cont := 1
	for i := indexAtual; i < len(runes); i++ {

		if cont > limite {
			return false
		}

		if runes[i] == rune('0'+value) {
			return true
		}

		cont++
	}

	return false
}

func ehValido(runes []rune, pos, limite int, valores []int) bool {

	if pos == len(runes) {
		return false
	}

	if runes[pos] != '.' {
		return ehValido(runes, pos+1, limite, valores)
	}

	if !buscaParaDireita(limite, valores[i], runes, pos) && !buscaParaEsquerda(limite, valores[i], runes, pos) {
		runes[pos] = rune('0' + valores[i])
	}

	i := 0
	for i < len(valores) {

		if ehValido(runes, pos, limite, valores) {

		}

		i++
	}

}

func preencher(valoresL []int, str string, limite int) string {

	runes := []rune(str)

	for i := range runes {
		if runes[i] == '.' {

			j := 0
			for j < len(valoresL) {

				if !buscaParaDireita(limite, valoresL[j], runes, i) && !buscaParaEsquerda(limite, valoresL[j], runes, i) {

					runes[i] = rune('0' + valoresL[j])
					break
				}
				j++
			}
			//volta e tenta outro caminho
		}
	}

	return string(runes)
}

// func preencher(valoresL []int, str string, limite int) string {

// 	runes := []rune(str)

// 	for i := range runes {
// 		if runes[i] == '.' {

// 			j := 0
// 			for j < len(valoresL) {

// 				if !buscaParaDireita(limite, valoresL[j], runes, i) && !buscaParaEsquerda(limite, valoresL[j], runes, i) {

// 					runes[i] = rune('0' + valoresL[j])
// 					break
// 				}
// 				j++
// 			}
// 			//volta e tenta outro caminho
// 		}
// 	}

// 	return string(runes)
// }

func main() {

	var str string
	fmt.Scan(&str)

	var limite int
	fmt.Scan(&limite)

	var valoresL []int

	for i := 0; i <= limite; i++ {
		valoresL = append(valoresL, i)
	}

	fmt.Println(preencher(valoresL, str, limite))
}

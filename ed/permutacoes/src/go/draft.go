package main

import (
	"fmt"
	"slices"
)

func permuta(runes []rune, pos_inicial int, listaOrdenada []string) []string {

	if pos_inicial == len(runes) {
		listaOrdenada = append(listaOrdenada, string(runes))
		return listaOrdenada
	}

	for i := pos_inicial; i < len(runes); i++ {

		runes[pos_inicial], runes[i] = runes[i], runes[pos_inicial]

		listaOrdenada = permuta(runes, pos_inicial+1, listaOrdenada)

		runes[pos_inicial], runes[i] = runes[i], runes[pos_inicial]
	}

	return listaOrdenada
}

func main() {

	var s string
	fmt.Scan(&s)

	runes := []rune(s)
	listaOrdenada := make([]string, 0)
	listaOrdenada = permuta(runes, 0, listaOrdenada)

	slices.Sort(listaOrdenada)

	forEach := func(lista []string, funcao func(string)) {
		for _, s := range lista {
			funcao(s)
		}
	}

	forEach(listaOrdenada, func(s string) {
		fmt.Println(s)
	})
	// que lambda maluca, prefiro java
}

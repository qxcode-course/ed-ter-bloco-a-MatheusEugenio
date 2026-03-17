package main

import (
	"fmt"
)

func main() {

	var totFigDoAlbum int
	fmt.Scan(&totFigDoAlbum)

	var quantidadeDeFigsBaruelPossui int
	fmt.Scan(&quantidadeDeFigsBaruelPossui)

	figurinhas := make([]int, quantidadeDeFigsBaruelPossui)

	for i := 0; i < quantidadeDeFigsBaruelPossui; i++ {
		fmt.Scan(&figurinhas[i])
	}

	var figRepetidas []int
	var faltantes []int

	albumVirtual := make([]bool, totFigDoAlbum+1)

	for i := 0; i < quantidadeDeFigsBaruelPossui; i++ {

		figAtual := figurinhas[i]

		if i > 0 && figAtual == figurinhas[i-1] {
			figRepetidas = append(figRepetidas, figAtual)
		}

		if figAtual <= totFigDoAlbum {
			albumVirtual[figAtual] = true
		}
	}

	for i := 1; i <= totFigDoAlbum; i++ {

		if albumVirtual[i] == false {
			faltantes = append(faltantes, i)
		}
	}

	if len(figRepetidas) == 0 {
		fmt.Println("N")
	} else {

		for i := 0; i < len(figRepetidas); i++ {
			if i > 0 {
				fmt.Print(" ")
			}

			if i == (len(figRepetidas) - 1) {
				fmt.Println(figRepetidas[i])
				break
			}

			fmt.Print(figRepetidas[i])
		}
	}

	if len(faltantes) == 0 {
		fmt.Println("N")
	} else {

		for i := 0; i < len(faltantes); i++ {
			if i > 0 {
				fmt.Print(" ")
			}

			if i == (len(faltantes) - 1) {
				fmt.Println(faltantes[i])
				break
			}

			fmt.Print(faltantes[i])
		}
	}
}

package main

import "fmt"

func main() {

	var totalFigsAlbum int
	fmt.Scan(&totalFigsAlbum)

	if totalFigsAlbum > 50 || totalFigsAlbum < 1 {
		return
	}

	var totalFigsPossui int
	fmt.Scan(&totalFigsPossui)

	arrayFigurinhas := make([]int, totalFigsPossui)
	for i := range arrayFigurinhas {
		fmt.Scan(&arrayFigurinhas[i])
	}

	figsRepetidas := make([]int, 0)
	figsUnicas := make(map[int]bool)

	for _, val := range arrayFigurinhas {
		_, existe := figsUnicas[val]
		if existe {
			figsRepetidas = append(figsRepetidas, val)
		} else {
			figsUnicas[val] = true
		}

	}

	if len(figsRepetidas) == 0 {
		fmt.Println("N")
	} else {
		for i, val := range figsRepetidas {
			if i == len(figsRepetidas)-1 {
				fmt.Printf("%d", val)
				fmt.Println()
			} else {
				fmt.Printf("%d ", val)
			}
		}
	}

	figsFaltando := make([]int, 0)

	if len(figsRepetidas) == totalFigsAlbum {
		fmt.Println("N")
	} else {
		for i := 1; i <= totalFigsAlbum; i++ {
			_, existe := figsUnicas[i]
			if !existe {
				figsFaltando = append(figsFaltando, i)
			}
		}

		for i, val := range figsFaltando {
			if i == len(figsFaltando)-1 {
				fmt.Printf("%d", val)
				fmt.Println()
			} else {
				fmt.Printf("%d ", val)
			}
		}
	}

}

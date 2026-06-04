package main

import (
	"fmt"
)

func imprimeLinha(vetMap map[int]bool, vet []int) string {

	var res string
	res += "[ "

	for i := 0; i < len(vet); i++ {

		posMapa := vet[i]

		if posMapa < 0 {
			posMapa *= -1
		}

		tem := vetMap[posMapa]
		if tem {
			if vet[i] > 0 {
				res += fmt.Sprintf("%d> ", vet[i])
			} else {
				res += fmt.Sprintf("<%d ", vet[i])
			}
		} else {
			res += fmt.Sprintf("%d ", vet[i])
		}
	}

	res += "]"
	return res
}

func colocaEspada(vetMap map[int]bool, e int) {

	for key := range vetMap {
		delete(vetMap, key)
	}
	vetMap[e] = true
}

func main() {
	var n, e, f int

	fmt.Scan(&n, &e, &f)

	vet := make([]int, n)
	vetMap := make(map[int]bool)

	for i := 0; i < n; i++ {
		num := i + 1

		if f == 1 {
			if i%2 != 0 {
				vet[i] = num * -1
			} else {
				vet[i] = num
			}
		} else {
			if i%2 == 0 {
				vet[i] = num * -1
			} else {
				vet[i] = num
			}
		}

	}

	for len(vet) > 1 {

		colocaEspada(vetMap, e)

		fmt.Println(imprimeLinha(vetMap, vet))

		indexEspada := -1
		for i, val := range vet {

			if val < 0 {
				val *= -1
			}

			if val == e {
				indexEspada = i
				break
			}

		}

		dir := -1
		if vet[indexEspada] > 0 {
			dir = 1
		}

		indexMorto := (indexEspada + dir + len(vet)) % len(vet)

		vet = append(vet[:indexMorto], vet[indexMorto+1:]...)

		novoIndexEspada := 0
        if dir == 1{
            novoIndexEspada = indexMorto % len(vet)
        }else {
            novoIndexEspada = (indexMorto-1 + len(vet)) % len(vet)
        }

		e = vet[novoIndexEspada]
		if e < 0 {
			e *= -1
		}
	}

	colocaEspada(vetMap, e)
	fmt.Println(imprimeLinha(vetMap, vet))

}

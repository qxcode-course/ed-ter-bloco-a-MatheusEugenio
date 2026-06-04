package main

import "fmt"

func runRigth(vet []int) {

	indexAtual := len(vet) - 1
	temp := vet[indexAtual]

	for i := len(vet) - 1; i > 0; i-- {
		vet[i] = vet[i-1]
	}

	vet[0] = temp
}

func stringVet(vet []int) string {

	var res string

	for _, val := range vet {
		res += fmt.Sprintf("%d ", val)
	}

	return "[ " + res + "]"
}

func main() {

	var n, numRotacion int
	fmt.Scan(&n, &numRotacion)

	var vet = make([]int, n)

	for i := 0; i < n; i++ {
		fmt.Scan(&vet[i])
	}

	if len(vet) <= 1 {
		fmt.Println(stringVet(vet))
		return
	}

	for numRotacion > 0 {

		runRigth(vet)
		numRotacion--
	}

	fmt.Println(stringVet(vet))

}

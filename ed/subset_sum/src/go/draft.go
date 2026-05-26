package main

import "fmt"

func somatorio(vetNumbers []int, valSoma int, somaTemp int, index int) bool {

	if somaTemp == valSoma {
		return true
	}

	if somaTemp > valSoma || index >= len(vetNumbers) {
		return false
	}

	if somatorio(vetNumbers, valSoma, somaTemp+vetNumbers[index], index+1) {
		return true
	}

	if somatorio(vetNumbers, valSoma, somaTemp, index+1) {
		return true
	}

	return false
}

func main() {

	var n int
	fmt.Scan(&n)

	var valSoma int
	fmt.Scan(&valSoma)

	vet := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Scan(&vet[i])
	}

	somatemp := 0
	if somatorio(vet, valSoma, somatemp, 0) {
		fmt.Println("true")
	} else {
		fmt.Println("false")
	}

}

package main

import "fmt"

func imprimeVet(vet []int) string {

	var res string

	for i, val := range vet {
		res += fmt.Sprint(val)
		if i != len(vet)-1 {
			res += " "
		}
	}

	return res
}

func matchingStrings(consultas []string, stringsVisitadas []string) []int {

	consultasMap := make(map[string]int)

	for _, val := range consultas {

		consultasMap[val]++
	}

	var vetOcorrencias = make([]int, 0)

	for _, str := range stringsVisitadas {
		val, existe := consultasMap[str]
		if existe {
			vetOcorrencias = append(vetOcorrencias, val)
		} else {
			vetOcorrencias = append(vetOcorrencias, 0)
		}
	}

	return vetOcorrencias
}

func main() {

	var n int
	fmt.Scan(&n)

	var vetConsultas = make([]string, n)

	for i := 0; i < n; i++ {
		fmt.Scan(&vetConsultas[i])
	}

	var tamVetbusca int
	fmt.Scan(&tamVetbusca)

	var vetBusca = make([]string, tamVetbusca)

	for i := 0; i < tamVetbusca; i++ {
		fmt.Scan(&vetBusca[i])
	}

	ocorrencias := matchingStrings(vetConsultas, vetBusca)
	fmt.Println(imprimeVet(ocorrencias))

}

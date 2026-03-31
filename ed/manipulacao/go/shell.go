package main

import (
	"bufio"
	"fmt"
	"os"
	"slices"
	"strconv"
	"strings"
)

func getMen(vet []int) []int {

	copiaVet := make([]int, 0)

	for _, val := range vet {
		if val >= 0 {
			copiaVet = append(copiaVet, val)
		}
	}
	return copiaVet
}

func getCalmWomen(vet []int) []int {

	copiaVetMulheres := make([]int, 0)

	for _, val := range vet {
		if val < 0 && val > -10 {
			copiaVetMulheres = append(copiaVetMulheres, val)
		}
	}
	return copiaVetMulheres
}

func sortVet(vet []int) []int {

	copiaVet := slices.Clone(vet)
	slices.Sort(copiaVet)
	return copiaVet
}

func sortStress(vet []int) []int {

	copiaValorAbsoluto := make([]int, len(vet))
	numeroGuardaSinal := make(map[int]bool) //Key: valor absoluto, Value: se tinha sinal

	for i, val := range vet {
		if val < 0 {
			copiaValorAbsoluto[i] = -val
			numeroGuardaSinal[-val] = true
		} else {
			copiaValorAbsoluto[i] = val
			numeroGuardaSinal[val] = false
		}
	}

	slices.Sort(copiaValorAbsoluto)
	vetRetorno := make([]int, 0)

	for _, val := range copiaValorAbsoluto {
		if numeroGuardaSinal[val] {
			vetRetorno = append(vetRetorno, -val)
		} else {
			vetRetorno = append(vetRetorno, val)
		}
	}

	return vetRetorno
}

func reverse(vet []int) []int {

	copiaVet := make([]int, len(vet))

	for i := len(vet) - 1; i >= 0; i-- {
		copiaVet[len(vet)-1-i] = vet[i]
	}

	return copiaVet
}

func unique(vet []int) []int {

	copiaVet := slices.Clone(vet)

	numerosVistos := make(map[int]bool)
	vetResultado := make([]int, 0)

	for _, val := range copiaVet {
		_, existe := numerosVistos[val]
		if !existe {
			numerosVistos[val] = true
			vetResultado = append(vetResultado, val)
		}
	}

	return vetResultado
}

func repeated(vet []int) []int {

	copiaVet := slices.Clone(vet)

	numerosVistos := make(map[int]bool)
	vetResultadosRepetidos := make([]int, 0)

	for _, val := range copiaVet {
		_, existe := numerosVistos[val]
		if !existe {
			numerosVistos[val] = true
		} else {
			vetResultadosRepetidos = append(vetResultadosRepetidos, val)
		}
	}

	return vetResultadosRepetidos
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	for {
		if !scanner.Scan() {
			break
		}
		fmt.Print("$")
		line := scanner.Text()
		args := strings.Split(line, " ")
		fmt.Println(line)

		switch args[0] {
		case "get_men":
			printVec(getMen(str2vet(args[1])))
		case "get_calm_women":
			printVec(getCalmWomen(str2vet(args[1])))
		case "sort":
			printVec(sortVet(str2vet(args[1])))
		case "sort_stress":
			printVec(sortStress(str2vet(args[1])))
		case "reverse":
			array := str2vet(args[1])
			other := reverse(array)
			printVec(array)
			printVec(other)
		case "unique":
			printVec(unique(str2vet(args[1])))
		case "repeated":
			printVec(repeated(str2vet(args[1])))
		case "end":
			return
		}
	}
}

func printVec(vet []int) {
	fmt.Print("[")
	for i, val := range vet {
		if i > 0 {
			fmt.Print(", ")
		}
		fmt.Print(val)
	}
	fmt.Println("]")
}

func str2vet(s string) []int {
	if s == "[]" {
		return nil
	}
	s = s[1 : len(s)-1]
	parts := strings.Split(s, ",")
	var vet []int
	for _, part := range parts {
		n, _ := strconv.Atoi(part)
		vet = append(vet, n)
	}
	return vet
}

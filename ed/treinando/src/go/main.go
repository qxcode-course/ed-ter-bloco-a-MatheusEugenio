package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func auxiliarRecursivaToStr(vet []int) string {

	if len(vet) == 1 {
		return strconv.Itoa(vet[0])
	}

	return strconv.Itoa(vet[0]) + ", " + auxiliarRecursivaToStr(vet[1:])
}

func auxiliarRecursivaToStrInvertida(vet []int) string {

	if len(vet) == 1 {
		return strconv.Itoa(vet[0])
	}

	return strconv.Itoa(vet[len(vet)-1]) + ", " + auxiliarRecursivaToStrInvertida(vet[:len(vet)-1])
}

func tostr(vet []int) string {

	if len(vet) == 0 {
		return "[]"
	}

	return "[" + auxiliarRecursivaToStr(vet) + "]"
}

func tostrrev(vet []int) string {
	if len(vet) == 0 {
		return "[]"
	}

	return "[" + auxiliarRecursivaToStrInvertida(vet) + "]"
}

// reverse: inverte os elementos do slice
func reverse(vet []int) {

	if len(vet) <= 1 {
		return
	}

	aux := vet[0]
	vet[0] = vet[len(vet)-1]
	vet[len(vet)-1] = aux

	reverse(vet[1 : len(vet)-1])
}

// sum: soma dos elementos do slice
func sum(vet []int) int {

	if len(vet) <= 0 {
		return 0
	}

	ultimaPOS := len(vet) - 1
	return sum(vet[:ultimaPOS]) + vet[ultimaPOS]
}

// mult: produto dos elementos do slice
func mult(vet []int) int {

	if len(vet) <= 0 {
		return 1
	}

	ultimaPOS := len(vet) - 1
	return mult(vet[:ultimaPOS]) * vet[ultimaPOS]
}

// min: retorna o índice e valor do menor valor
// crie uma função recursiva interna do modelo
// var rec func(v []int) (int, int)
// para fazer uma recursão que retorna valor e índice
func min(vet []int) int {

	if len(vet) <= 0 {
		return -1
	}

	if len(vet) == 1 {
		return 0
	}

	var rec func(v []int) (int, int)

	rec = func(v []int) (int, int) {

		if len(v) == 1 {
			return 0, v[0]
		}

		indexMenorResto, valMenorResto := rec(v[1:])

		idxAjustado := indexMenorResto + 1

		if v[0] < valMenorResto {
			return 0, v[0]
		}

		return idxAjustado, valMenorResto
	}

	indiceDoMenor, _ := rec(vet)
	return indiceDoMenor
}

func main() {
	var vet []int
	scanner := bufio.NewScanner(os.Stdin)
	for {
		if !scanner.Scan() {
			break
		}
		line := scanner.Text()
		args := strings.Fields(line)
		fmt.Println("$" + line)

		switch args[0] {
		case "end":
			return
		case "read":
			vet = nil
			for _, arg := range args[1:] {
				if val, err := strconv.Atoi(arg); err == nil {
					vet = append(vet, val)
				}
			}
		case "tostr":
			fmt.Println(tostr(vet))
		case "torev":
			fmt.Println(tostrrev(vet))
		case "reverse":
			reverse(vet)
		case "sum":
			fmt.Println(sum(vet))
		case "mult":
			fmt.Println(mult(vet))
		case "min":
			fmt.Println(min(vet))
		default:
			fmt.Println("fail: comando invalido")
		}
	}
}

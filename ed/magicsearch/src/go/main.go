package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func MagicSearch(slice []int, value int) int {

	tam := len(slice)

	if tam == 0 {
		return 0
	}

	low, high := 0, tam-1

	indiceResultado := -1
	for low <= high {

		meio := (low + high) / 2

		if slice[meio] < value {
			low = meio + 1
		} else if slice[meio] > value {
			high = meio - 1
		} else {
			indiceResultado = meio
			low = meio + 1
		}
	}

	if indiceResultado != -1 {
		return indiceResultado
	}

	//retorna a posição devida do número não encontrado
	return low
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	parts := strings.Fields(scanner.Text())
	slice := make([]int, 0, 1)
	for _, elem := range parts[1 : len(parts)-1] {
		value, _ := strconv.Atoi(elem)
		slice = append(slice, value)
	}

	scanner.Scan()
	value, _ := strconv.Atoi(scanner.Text())
	result := MagicSearch(slice, value)
	fmt.Println(result)
}

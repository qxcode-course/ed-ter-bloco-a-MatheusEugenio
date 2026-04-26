package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func BetterSearch(slice []int, value int) (bool, int) {

	tam := len(slice)

	if tam == 0 {
		return false, 0
	}

	low, high := 0, tam-1

	for low <= high {

		meio := (low + high) / 2

		if slice[meio] < value {
			low = meio + 1
		} else if slice[meio] > value {
			high = meio - 1
		} else {
			return true, meio
		} 
	}

	return false, low
}

func main() {

	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	parts := strings.Split(scanner.Text(), " ")
	slice := []int{}
	for _, elem := range parts[1 : len(parts)-1] {
		value, _ := strconv.Atoi(elem)
		slice = append(slice, value)
	}

	scanner.Scan()
	value, _ := strconv.Atoi(scanner.Text())
	found, result := BetterSearch(slice, value)
	if found {
		fmt.Println("V", result)
	} else {
		fmt.Println("F", result)
	}
}

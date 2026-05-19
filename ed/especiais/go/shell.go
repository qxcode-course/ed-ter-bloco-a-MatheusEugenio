package main

import (
	"bufio"
	"fmt"
	"os"
	"slices"
	"strconv"
	"strings"
)

type Pair struct {
	One int
	Two int
}

func occurr(vet []int) []Pair {

	contagemPares := make(map[int]int)

	for _, val := range vet {

		if val < 0 {
			val = val * -1
		}

		_, exists := contagemPares[val]
		if !exists {
			contagemPares[val] = 1
		} else {
			contagemPares[val]++
		}
	}

	pares := make([]Pair, 0)

	for val, repeticao := range contagemPares {
		pares = append(pares, Pair{One: val, Two: repeticao})
	}

	slices.SortFunc(pares, func(a Pair, b Pair) int {
		if a.One < b.One {
			return -1
		} else if a.One > b.One {
			return 1
		} else {
			return 0
		}
	})

	return pares
}

func teams(vet []int) []Pair {
	if len(vet) == 0 {
		return nil
	}

	pares := make([]Pair, 0)
	currentVal := vet[0]
	currentCount := 1

	for i := 1; i < len(vet); i++ {
		if vet[i] == currentVal {
			currentCount++
		} else {
			pares = append(pares, Pair{One: currentVal, Two: currentCount})
			currentVal = vet[i]
			currentCount = 1
		}
	}
	pares = append(pares, Pair{One: currentVal, Two: currentCount})

	return pares
}

func mnext(vet []int) []int {
	res := make([]int, len(vet))
	for i, val := range vet {
		if val > 0 { // Se for homem
			hasWomanNeighbor := false
			// Verifica vizinho da esquerda
			if i > 0 && vet[i-1] < 0 {
				hasWomanNeighbor = true
			}
			// Verifica vizinho da direita
			if i < len(vet)-1 && vet[i+1] < 0 {
				hasWomanNeighbor = true
			}
			if hasWomanNeighbor {
				res[i] = 1
			}
		}
	}
	return res
}

func alone(vet []int) []int {
	res := make([]int, len(vet))
	for i, val := range vet {
		if val > 0 { // Se for homem
			hasWomanNeighbor := false
			if i > 0 && vet[i-1] < 0 {
				hasWomanNeighbor = true
			}
			if i < len(vet)-1 && vet[i+1] < 0 {
				hasWomanNeighbor = true
			}
			if !hasWomanNeighbor {
				res[i] = 1
			}
		}
	}
	return res
}

func couple(vet []int) int {
	men := make(map[int]int)
	women := make(map[int]int)

	for _, val := range vet {
		if val > 0 {
			men[val]++
		} else if val < 0 {
			women[-val]++
		}
	}

	couples := 0
	for stress, countMen := range men {
		if countWomen, exists := women[stress]; exists {
			if countMen < countWomen {
				couples += countMen
			} else {
				couples += countWomen
			}
		}
	}
	return couples
}

func hasSubseq(vet []int, seq []int, pos int) bool {
	if pos+len(seq) > len(vet) {
		return false
	}
	for i := 0; i < len(seq); i++ {
		if vet[pos+i] != seq[i] {
			return false
		}
	}
	return true
}

func subseq(vet []int, seq []int) int {
	if len(seq) == 0 {
		return 0
	}
	for i := 0; i <= len(vet)-len(seq); i++ {
		if hasSubseq(vet, seq, i) {
			return i
		}
	}
	return -1
}

func erase(vet []int, posList []int) []int {
	toRemove := make(map[int]bool)
	for _, pos := range posList {
		toRemove[pos] = true
	}

	res := make([]int, 0)
	for i, val := range vet {
		if !toRemove[i] {
			res = append(res, val)
		}
	}
	return res
}

func clear(vet []int, value int) []int {
	res := make([]int, 0)
	for _, val := range vet {
		if val != value {
			res = append(res, val)
		}
	}
	return res
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	for {
		fmt.Print("$")
		if !scanner.Scan() {
			break
		}
		line := scanner.Text()
		args := strings.Split(line, " ")
		fmt.Println(line)

		switch args[0] {
		case "occurr":
			printSlice(occurr(str2vet(args[1])))
		case "teams":
			printSlice(teams(str2vet(args[1])))
		case "mnext":
			printSlice(mnext(str2vet(args[1])))
		case "alone":
			printSlice(alone(str2vet(args[1])))
		case "erase":
			printSlice(erase(str2vet(args[1]), str2vet(args[2])))
		case "clear":
			val, _ := strconv.Atoi(args[2])
			printSlice(clear(str2vet(args[1]), val))
		case "subseq":
			fmt.Println(subseq(str2vet(args[1]), str2vet(args[2])))
		case "couple":
			fmt.Println(couple(str2vet(args[1])))
		case "end":
			return
		default:
			fmt.Println("Invalid command")
		}
	}
}

// Funções auxiliares

func str2vet(str string) []int {
	if str == "[]" {
		return nil
	}
	str = str[1 : len(str)-1]
	parts := strings.Split(str, ",")
	var vet []int
	for _, part := range parts {
		num, _ := strconv.Atoi(strings.TrimSpace(part))
		vet = append(vet, num)
	}
	return vet
}

func printSlice[T any](vet []T) {
	fmt.Print("[")
	for i, x := range vet {
		if i > 0 {
			fmt.Print(", ")
		}
		fmt.Print(x)
	}
	fmt.Println("]")
}

func (p Pair) String() string {
	return fmt.Sprintf("(%v, %v)", p.One, p.Two)
}

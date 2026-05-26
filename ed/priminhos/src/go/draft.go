package main

import (
	"fmt"
	"strings"
)

func Join(slice []int, sep string) string {
	if len(slice) == 0 {
		return ""
	}
	var result strings.Builder
	fmt.Fprintf(&result, "%d", slice[0])
	for _, value := range slice[1:] {
		fmt.Fprintf(&result, "%s%d", sep, value)
	}
	return result.String()
}

func ehPrimo(numAtual int) bool{

    if numAtual < 2{
        return false
    }

    for i := 2; i < numAtual; i++{
        if numAtual % i == 0{
            return false
        }
    }

    return true
}

func priminhos(n int, vet []int, numAtual int) []int {

	if n < 1 {
		return nil
	}

	if vet == nil{
        return nil
    }

    if ehPrimo(numAtual){
        vet = append(vet, numAtual)
    }

    if len(vet) == n{
        return vet
    }

	return priminhos(n, vet, numAtual+1)
}

func main() {

	var n int

	fmt.Scan(&n)

	vetBase := make([]int, 0)

	fmt.Println("[" + Join(priminhos(n, vetBase, 2), ", ") + "]")
}

package main

import "fmt"

func diagonal(s string, k int) {
	if len(s) == 0 {
		return
	}

	// imprima k caracteres
	for i := 0; i < k; i++ {
		fmt.Print(" ")
	}

	// imprima o primeiro caractere de s e pule a linha
	fmt.Println(string(s[0]))
	// faça a chamada
	diagonal(s[1:], k+1)
}

func main() {

	var palavra string
	fmt.Scan(&palavra)

	i := 0
	diagonal(palavra, i)
}

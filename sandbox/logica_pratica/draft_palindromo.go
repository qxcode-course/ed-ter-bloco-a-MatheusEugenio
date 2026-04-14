package main

import (
	"fmt"
	"strings"
	"unicode"
)

func isPalindromo(s string) bool {
	// TODO: Limpe string (lowercase, alfanum só)
	// Compare com reverse
	limpa := strings.ToLower(s) // Início
	// Adicione: remova non-alphanum, reverse e compare
	return false
}

func main() {
	var s string
	fmt.Scan(&s)
	fmt.Println(isPalindromo(s))
}


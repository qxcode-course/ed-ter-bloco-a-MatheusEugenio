package main

import "fmt"

func printTower(s string, strAnterior string) {

	ultimoIndex := len(s) - 1

	if len(s) == 0 {
		return
	}

	strAtual := string(s[ultimoIndex]) + strAnterior
	fmt.Println(strAtual)

	printTower(s[:ultimoIndex], strAtual)

}

func main() {

	var s string
	fmt.Scan(&s)

	var str string
	printTower(s, str)
}

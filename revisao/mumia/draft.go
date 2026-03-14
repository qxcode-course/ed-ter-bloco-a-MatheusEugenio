package main

import "fmt"

func main() {

	var nome string
	var idade int

	fmt.Scan(&nome)
	fmt.Scan(&idade)

	var class string

	if idade < 12 {
		class = "crianca"
	} else if idade < 18 {
		class = "jovem"
	} else if idade < 65 {
		class = "adulto"
	} else if idade < 1000 {
		class = "idoso"
	} else {
		class = "mumia"
	}

	fmt.Printf("%s eh %s\n", nome, class)

}

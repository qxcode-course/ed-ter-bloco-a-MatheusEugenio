package main

import (
	"fmt"
	"strconv"
)

// mostra a lista com o elemento sword destacado
func ToStr(l *DList[int], sword *DNode[int]) string {

	saida := "[ "

	nodeAtual := l.root.next
	for nodeAtual != l.root {

		if nodeAtual == sword {
			saida += strconv.Itoa(nodeAtual.Value) + "> "
		} else {
			saida += strconv.Itoa(nodeAtual.Value) + " "
		}

		nodeAtual = nodeAtual.next
	}

	return saida + "]"
}

// move para frente na lista circular
func Next(l *DList[int], it *DNode[int]) *DNode[int] {

	if it.Next() != it.root {
		l.Erase(it.Next())
	} else if it == l.Back() {
		l.Erase(l.Front())
		return it
	} else {
		return it
	}

	return it.next
}

func main() {
	var qtd, chosen int
	fmt.Scan(&qtd, &chosen)

	// fmt.Println(qtd, chosen)

	l := NewDList[int]()

	for i := 1; i <= qtd; i++ {
		l.PushBack(i)
	}

	sword := l.Front()

	for range chosen - 1 {
		sword = Next(l, sword)
	}

	for range qtd - 1 {
		fmt.Println(ToStr(l, sword))
		sword = Next(l, sword)
	}

	fmt.Println(ToStr(l, sword))
}

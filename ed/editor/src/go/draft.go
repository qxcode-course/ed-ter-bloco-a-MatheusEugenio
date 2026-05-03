package main

import (
	"fmt"
	// "html/template"
)

type Node struct{
    char rune
    next *Node
    prev *Node
}

type LList struct{
    head *Node
    cursor *Node //cursor
    size int
}

func NewLList() *LList{

    root := &Node{} 
    tail := &Node{}

    root.next = tail
    root.prev = tail
    tail.next = root
    tail.prev = root
    return &LList{
        head: root,
        cursor: tail,
        size: 0,
    } 

}

func (l*LList) Insert(char rune){

    newNode := &Node{}
    newNode.char = char

    newNode.prev = l.cursor
    newNode.next = l.cursor.next

    l.cursor.prev = newNode
    l.cursor.next = l.head
    l.size++
}

func editor(entrada string) string {
	var texto []rune
	posCursor := 0

	for _, char := range entrada {

		switch char {
		case '>':
			if posCursor < len(texto) {
				posCursor++
			}
		case '<':
			if posCursor > 0 {
				posCursor--
			}
		case 'R':
			texto = append(texto[:posCursor], append([]rune{'\n'}, texto[posCursor:]...)...)
			posCursor++
		case 'B':
			if posCursor > 0 {
				texto = append(texto[:posCursor-1], texto[posCursor:]...)
				posCursor--
			}
		case 'D':
			if posCursor < len(texto) {
				texto = append(texto[:posCursor], texto[posCursor+1:]...)
			}
		default:
			texto = append(texto[:posCursor], append([]rune{char}, texto[posCursor:]...)...)
			posCursor++
		}
	}

	res := append(texto[:posCursor], append([]rune{'|'}, texto[posCursor:]...)...)

	return string(res)
}

func main() {
	var texto string
	fmt.Scan(&texto)

	fmt.Println(editor(texto))
}

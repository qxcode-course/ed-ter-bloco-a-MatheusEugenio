package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Node struct {
	value int   // Valor é público
	next  *Node // o próximo nó da lista
	prev  *Node // o nó anterior
}

type LList struct { //lista ligada
	head *Node // nó de marcação
	size int
}

func NewLList() *LList {

	head := &Node{}

	head.next = head
	head.prev = head

	return &LList{
		head: head,
		size: 0,
	}
}

func (ll *LList) String() string {

	if ll.head.next == ll.head || ll.size == 0 {
		return "[]"
	}

	saida := "["

	i := 0
	noAtual := ll.Front()
	for noAtual != ll.head {
		if i == ll.size-1 {
			saida += strconv.Itoa(noAtual.value)
			noAtual = noAtual.next
			break
		}
		saida += strconv.Itoa(noAtual.value) + ", "
		noAtual = noAtual.next
		i++
	}

	return saida + "]"
}

func (ll *LList) Front() *Node {
	return ll.head.next
}

func (ll *LList) Back() *Node {
	return ll.head.prev
}

func (ll *LList) PushBack(value int) {

	nodeBack := ll.Back()
	newNode := &Node{
		value: value,
		prev:  ll.Back(),
		next:  ll.head,
	}

	nodeBack.next = newNode
	ll.head.prev = newNode
	ll.size++
}

func (ll *LList) PushFront(value int) {

	nodeFront := ll.Front()
	newNode := &Node{
		value: value,
		next:  ll.Front(),
		prev:  ll.head,
	}

	ll.head.next = newNode
	nodeFront.prev = newNode
	ll.size++
}

func (ll *LList) PopFront() {
	nodeFront := ll.Front()
	ll.head.next = nodeFront.next
	nodeFront.next.prev = ll.head
	nodeFront.next = nil
	nodeFront.prev = nil

	if ll.size > 0 {
		ll.size--
	}
}

func (ll *LList) PopBack() {
	nodeBack := ll.Back()
	nodeBack.prev.next = ll.head
	ll.head.prev = nodeBack.prev
	nodeBack.prev = nil
	nodeBack.next = nil

	if ll.size > 0 {
		ll.size--
	}
}

func (ll *LList) Size() int {
	return ll.size
}

func (ll *LList) Clear() {
	ll.size = 0
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	ll := NewLList()

	for {
		fmt.Print("$")
		if !scanner.Scan() {
			break
		}
		line := scanner.Text()
		fmt.Println(line)
		args := strings.Fields(line)

		if len(args) == 0 {
			continue
		}

		cmd := args[0]

		switch cmd {
		case "show":
			fmt.Println(ll.String())
		case "size":
			fmt.Println(ll.Size())
		case "push_back":
			for _, v := range args[1:] {
				num, _ := strconv.Atoi(v)
				ll.PushBack(num)
			}
		case "push_front":
			for _, v := range args[1:] {
				num, _ := strconv.Atoi(v)
				ll.PushFront(num)
			}
		case "pop_back":
			ll.PopBack()
		case "pop_front":
			ll.PopFront()
		case "clear":
			ll.Clear()
		case "end":
			return
		default:
			fmt.Println("fail: comando invalido")
		}
	}
}

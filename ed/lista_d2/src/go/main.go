package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Node struct {
	value int
	next  *Node
	prev  *Node
	root  *Node
}

type LList struct {
	head *Node
	size int
}

func (n *Node) Next() *Node {

	if n.next == nil {
		return nil
	}

	return n.next
}

func (n *Node) Prev() *Node {

	if n.prev == nil {
		return nil
	}

	return n.prev
}

func (l *LList) Front() *Node {

	if l.head.next == l.head {
		return nil
	}

	return l.head.next
}

func (l *LList) Back() *Node {

	if l.head.prev == l.head {
		return nil
	}

	return l.head.prev
}

func (l *LList) End() *Node {
	return l.head
}

func (l *LList) Size() int {
	return l.size
}

func (l *LList) Clear() {

	l.head.next = l.head
	l.head.prev = l.head
	l.head.root = l.head
	l.size = 0
}

func NewLList() *LList {

	head := &Node{}
	head.next = head
	head.prev = head
	head.root = head

	newNode := &LList{
		head: head,
		size: 0,
	}

	return newNode
}

func (l *LList) Search(value int) *Node {

	if l.size == 0 || l.Front() == l.End() {
		return nil
	}

	for node := l.Front(); node != l.End(); node = node.next {
		if node.value == value {
			return node
		}
	}

	return nil
}

func (l *LList) Insert(nodePosterior *Node, value int) {

	nodeAnterior := nodePosterior.prev
	newNode := &Node{
		value: value,
		next:  nodePosterior,
		prev:  nodeAnterior,
		root:  l.End(),
	}

	nodeAnterior.next = newNode
	nodePosterior.prev = newNode
	l.size++
}

func (l *LList) Remove(nodeRemove *Node) *Node {

	if nodeRemove == nil || nodeRemove == l.head {
		return nil
	}

	nodeAnterior := nodeRemove.prev
	nodePosterior := nodeRemove.next

	nodeAnterior.next = nodePosterior
	nodePosterior.prev = nodeAnterior

	nodeRemove.next = nil
	nodeRemove.prev = nil

	l.size--
	return nodePosterior
}

func (l *LList) PushBack(value int) {
	l.Insert(l.head, value)
}

func (l *LList) PushFront(value int) {
	l.Insert(l.Front(), value)
}

func (l *LList) PopFront() {
	l.Remove(l.Front())
}

func (l *LList) PopBack() {
	l.Remove(l.Back())
}

func (l *LList) String() string {

	if l.head.next == l.head || l.size == 0 {
		return "[]"
	}

	saida := "["

	i := 0
	nodeAtual := l.Front()

	for nodeAtual != l.head {

		if i == l.size-1 {
			saida += strconv.Itoa(nodeAtual.value)
			break
		}

		saida += strconv.Itoa(nodeAtual.value) + ", "
		nodeAtual = nodeAtual.next
		i++
	}

	return saida + "]"
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	l := NewLList()

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
			fmt.Println(l.String())
		case "size":
			fmt.Println(l.Size())
		case "push_back":
			for _, v := range args[1:] {
				num, _ := strconv.Atoi(v)
				l.PushBack(num)
			}
		case "push_front":
			for _, v := range args[1:] {
				num, _ := strconv.Atoi(v)
				l.PushFront(num)
			}
		case "pop_back":
			l.PopBack()
		case "pop_front":
			l.PopFront()
		case "clear":
			l.Clear()
		case "walk":
			fmt.Print("[ ")
			for node := l.Front(); node != nil && node != l.End(); node = node.Next() {
				fmt.Printf("%v ", node.value)
			}
			fmt.Print("]\n[ ")
			for node := l.Back(); node != nil && node != l.End(); node = node.Prev() {
				fmt.Printf("%v ", node.value)
			}
			fmt.Println("]")
		case "replace":
			oldvalue, _ := strconv.Atoi(args[1])
			newvalue, _ := strconv.Atoi(args[2])
			node := l.Search(oldvalue)
			if node != nil {
				node.value = newvalue
			} else {
				fmt.Println("fail: not found")
			}
		case "insert":
			oldvalue, _ := strconv.Atoi(args[1])
			newvalue, _ := strconv.Atoi(args[2])
			node := l.Search(oldvalue)
			if node != nil {
				l.Insert(node, newvalue)
			} else {
				fmt.Println("fail: not found")
			}
		case "remove":
			oldvalue, _ := strconv.Atoi(args[1])
			node := l.Search(oldvalue)
			if node != nil {
				l.Remove(node)
			} else {
				fmt.Println("fail: not found")
			}
		case "end":
			return
		default:
			fmt.Println("fail: comando invalido")
		}
	}
}

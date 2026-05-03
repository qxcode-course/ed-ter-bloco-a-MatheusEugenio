package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type MyList struct {
	data []int
}

type Iterator struct {
	data  []int
	index int
}

type ReverseIterator struct {
	data  []int
	index int
}

type CyclicIterator struct {
	data  []int
	index int
}

func NewMyList(values []int) *MyList {
	return &MyList{data: values}
}

func (l *MyList) Iterator() *Iterator {
	return &Iterator{data: l.data, index: 0}
}

func (i *Iterator) HasNext() bool {
	return i.index < len(i.data)
}

func (i *Iterator) Next() int {
	val := i.data[i.index]
	i.index++
	return val
}

func (l *MyList) ReverseIterator() *ReverseIterator {
	return &ReverseIterator{data: l.data, index: len(l.data) - 1}
}

func (ri *ReverseIterator) HasNext() bool {
	return ri.index >= 0
}

func (ri *ReverseIterator) Next() int {

	val := ri.data[ri.index]
	ri.index--
	return val
}

func (l *MyList) CyclicIterator() *CyclicIterator {
	return &CyclicIterator{
		data:  l.data,
		index: 0,
	}
}

func (c *CyclicIterator) HasNext() bool {
	return len(c.data) > 0
}

func (c *CyclicIterator) Next() int {

	val := c.data[c.index]
	c.index = (c.index + 1) % len(c.data)
	return val
}

func main() {

	scanner := bufio.NewScanner(os.Stdin)
	mylist := NewMyList([]int{})

	for scanner.Scan() {
		line := scanner.Text()
		args := strings.Fields(line)
		fmt.Println("$" + line)

		switch args[0] {
		case "end":
			return
		case "read":
			for i := 1; i < len(args); i++ {
				slice := make([]int, len(args)-1)
				for i, value := range args[1:] {
					slice[i], _ = strconv.Atoi(value)
				}
				mylist = NewMyList(slice)
			}
		case "show":
			fmt.Print("[ ")
			for it := mylist.Iterator(); it.HasNext(); {
				fmt.Printf("%v ", it.Next())
			}
			fmt.Println("]")
		case "reverse":
			fmt.Print("[ ")
			for it := mylist.ReverseIterator(); it.HasNext(); {
				fmt.Printf("%v ", it.Next())
			}
			fmt.Println("]")
		case "cyclic":
			qtd, _ := strconv.Atoi(args[1])
			fmt.Print("[ ")
			it := mylist.CyclicIterator()
			if len(mylist.data) > 0 {
				for i := 0; i < qtd; i++ {
					fmt.Printf("%v ", it.Next())
				}
			}
			fmt.Println("]")
		}
	}

}

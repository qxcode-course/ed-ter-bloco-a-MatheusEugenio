package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Set struct {
	data     []int
	size     int
	capacity int
}

func NewSet(newCacpacity int) *Set {
	return &Set{
		data:     make([]int, newCacpacity),
		size:     0,
		capacity: newCacpacity,
	}
}

// privado
func (v *Set) reserve(newCacpacity int) {

	if newCacpacity == 0 {
		v.capacity = 1
	} else if newCacpacity <= v.capacity {
		return
	} else {
		v.capacity *= 2
	}

	dataTemp := make([]int, v.capacity)

	for i, val := range v.data {
		dataTemp[i] = val
	}

	v.data = dataTemp
}

// privado
func (v *Set) binarySearch(value int) int { //retorna o indice do numero, se existe

	return -1
}

// privado
func (v *Set) insert(value int, index int) { //usa busca binaria

	if index > v.capacity || index < 0 {
		return
	}

}

// privado
func (v *Set) erase(index int) error {

	if index < 0 || index >= v.size {
		return errors.New("value not found")
	}

	if index == v.size-1 {
		v.size--
		return nil
	}

	for i := index; i < v.size-1; i++ {
		v.data[i] = v.data[i+1]
	}
	v.size--
	return nil
}

func (v *Set) Insert(value int, index int) { //usa busca binaria para saber se o elemento já existe

}

func (v *Set) Erase(index int) bool {
	res := v.erase(index)
	if res == nil {
		return true
	} else {
		return false
	}
}

func (v *Set) Contains(value int) bool { //usa busca binaria para saber se o elemento existe

	return false
}

func (v *Set) String() string {
	return "[" + Join(v.data[:v.size], ", ") + "]"
}

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

func main() {
	var line, cmd string
	scanner := bufio.NewScanner(os.Stdin)

	v := NewSet(0)
	for scanner.Scan() {
		fmt.Print("$")
		line = scanner.Text()
		fmt.Println(line)
		parts := strings.Fields(line)
		if len(parts) == 0 {
			continue
		}
		cmd = parts[0]

		switch cmd {
		case "end":
			return
		case "init":
			value, _ := strconv.Atoi(parts[1])
			v = NewSet(value)
		case "insert":
			// for _, part := range parts[1:] {
			// 	value, _ := strconv.Atoi(part)
			// }
		case "show":
			fmt.Println(v)
		case "erase":
			value, _ := strconv.Atoi(parts[1])
			if v.Erase(value) {
				fmt.Println(v)
			} else {
				// fmt.Println(v.erase(value))
				// fmt.Println("value not found")
			}
		case "contains":
			// value, _ := strconv.Atoi(parts[1])
		case "clear":
		default:
			fmt.Println("fail: comando invalido")
		}
	}
}
